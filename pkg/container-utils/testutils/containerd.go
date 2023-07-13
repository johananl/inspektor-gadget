// Copyright 2022 The Inspektor Gadget authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutils

import (
	"context"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/cio"
	"github.com/containerd/containerd/namespaces"
	"github.com/containerd/containerd/oci"
	"github.com/containerd/containerd/snapshots"
	"github.com/opencontainers/runtime-spec/specs-go"
)

const (
	// TODO containerd currently only works on k8s.io namespace
	defaultNamespace = "k8s.io"
)

func NewContainerdContainer(name, cmd string, options ...Option) Container {
	c := &ContainerdContainer{
		containerSpec: containerSpec{
			name:    name,
			cmd:     cmd,
			options: defaultContainerOptions(),
		},
	}
	for _, o := range options {
		o(c.options)
	}
	return c
}

type ContainerdContainer struct {
	containerSpec

	exitStatus <-chan containerd.ExitStatus
}

func (c *ContainerdContainer) Running() bool {
	return c.started
}

func (c *ContainerdContainer) ID() string {
	return c.id
}

func (c *ContainerdContainer) Pid() int {
	return c.pid
}

func (c *ContainerdContainer) Run(t *testing.T) {
	nsCtx := namespaces.WithNamespace(c.options.ctx, defaultNamespace)
	fullImage := getFullImage(c.options)

	client, err := containerd.New("/run/containerd/containerd.sock",
		containerd.WithTimeout(3*time.Second),
	)
	if err != nil {
		t.Fatalf("Failed to connect to containerd: %s", err)
	}

	// Download and unpack the image
	image, err := client.Pull(nsCtx, fullImage)
	if err != nil {
		t.Fatalf("Failed to pull the image %q: %s", fullImage, err)
	}

	unpacked, err := image.IsUnpacked(nsCtx, "")
	if err != nil {
		t.Fatalf("image.IsUnpacked: %v", err)
	}
	if !unpacked {
		if err := image.Unpack(nsCtx, ""); err != nil {
			t.Fatalf("image.Unpack: %v", err)
		}
	}

	// Create the container
	var specOpts []oci.SpecOpts
	specOpts = append(specOpts, oci.WithDefaultSpec())
	specOpts = append(specOpts, oci.WithDefaultUnixDevices)
	specOpts = append(specOpts, oci.WithImageConfig(image))
	if len(c.cmd) != 0 {
		specOpts = append(specOpts, oci.WithProcessArgs("/bin/sh", "-c", c.cmd))
	}
	if c.options.seccompProfile != "" {
		t.Fatalf("testutils/containerd: seccomp profiles are not supported yet")
	}

	var spec specs.Spec
	container, err := client.NewContainer(nsCtx, c.name,
		containerd.WithImage(image),
		containerd.WithImageConfigLabels(image),
		containerd.WithAdditionalContainerLabels(image.Labels()),
		containerd.WithSnapshotter(""),
		containerd.WithNewSnapshot(c.name, image, snapshots.WithLabels(map[string]string{})),
		containerd.WithImageStopSignal(image, "SIGTERM"),
		containerd.WithSpec(&spec, specOpts...),
	)
	if err != nil {
		t.Fatalf("Failed to create container %q: %s", c.name, err)
	}
	c.id = container.ID()

	containerIO := cio.NullIO
	output := &strings.Builder{}
	if c.options.logs {
		containerIO = cio.NewCreator(cio.WithStreams(nil, output, output))
	}
	// Now create and start the task
	task, err := container.NewTask(nsCtx, containerIO)
	if err != nil {
		container.Delete(nsCtx, containerd.WithSnapshotCleanup)
		t.Fatalf("Failed to create task %q: %s", c.name, err)
	}

	err = task.Start(nsCtx)
	if err != nil {
		container.Delete(nsCtx, containerd.WithSnapshotCleanup)
		t.Fatalf("Failed to start task %q: %s", c.name, err)
	}

	c.exitStatus, err = task.Wait(nsCtx)
	if err != nil {
		t.Fatalf("Failed to wait on task %q: %s", c.name, err)
	}
	c.pid = int(task.Pid())

	if c.options.wait {
		s := <-c.exitStatus
		if s.ExitCode() != 0 {
			t.Logf("Exitcode for task %q: %d", c.name, s.ExitCode())
		}
	}

	if c.options.logs {
		t.Logf("Container %q output:\n%s", c.name, output.String())
	}

	if c.options.removal {
		task.Kill(nsCtx, syscall.SIGKILL)
		// We need to wait until the task is killed before trying to delete it
		<-c.exitStatus
		_, err = task.Delete(nsCtx)
		if err != nil {
			t.Fatalf("Failed to delete task %q: %s", c.name, err)
		}
		err = container.Delete(nsCtx, containerd.WithSnapshotCleanup)
		if err != nil {
			t.Fatalf("Failed to delete container %q: %s", c.name, err)
		}
	}
}

func (c *ContainerdContainer) Start(t *testing.T) {
	if c.started {
		t.Logf("Warn(%s): trying to start already running container\n", c.name)
		return
	}
	c.start(t)
	c.started = true
}

func (c *ContainerdContainer) start(t *testing.T) {
	for _, o := range []Option{WithoutWait(), WithoutRemoval()} {
		o(c.options)
	}
	c.Run(t)
}

func (c *ContainerdContainer) Stop(t *testing.T) {
	c.stop(t)
	c.started = false
}

func (c *ContainerdContainer) stop(t *testing.T) {
	client, err := containerd.New("/run/containerd/containerd.sock",
		containerd.WithTimeout(3*time.Second),
	)
	if err != nil {
		t.Fatalf("Failed to connect to containerd: %s", err)
	}

	nsCtx := namespaces.WithNamespace(c.options.ctx, defaultNamespace)
	container, err := client.LoadContainer(nsCtx, c.name)
	if err != nil {
		t.Fatalf("Failed to get container %q: %s", c.name, err)
	}

	task, err := container.Task(nsCtx, nil)
	if err != nil {
		t.Fatalf("Failed to get task %q: %s", c.name, err)
	}

	task.Kill(nsCtx, syscall.SIGKILL)
	// We need to wait until the task is killed before trying to delete it
	<-c.exitStatus
	_, err = task.Delete(nsCtx)
	if err != nil {
		t.Fatalf("Failed to delete task %q: %s", c.name, err)
	}
	err = container.Delete(nsCtx, containerd.WithSnapshotCleanup)
	if err != nil {
		t.Fatalf("Failed to delete container %q: %s", c.name, err)
	}
}

func getFullImage(options *containerOptions) string {
	if strings.Contains(options.image, ":") {
		return options.image
	}
	return options.image + ":" + options.imageTag
}

func RunContainerdFailedContainer(ctx context.Context, t *testing.T) {
	NewContainerdContainer("test-ig-failed-container", "/none", WithoutLogs(), WithoutWait(), WithContext(ctx)).Run(t)
}
