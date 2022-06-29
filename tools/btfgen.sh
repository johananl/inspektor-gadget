#!/bin/bash

set -e
set -x

BTFHUB="${BTFHUB:-$(pwd)/../btfhub/}"
INSPEKTOR_GADGET=${INSPEKTOR_GADGET:-$(pwd)}
ARCH=x86_64
OUTPUT=/tmp/btfs

if [ ! -d "${BTFHUB}" ]; then
    echo "error: btfhub folder not found"
    exit 1
fi

if [ ! -d "${INSPEKTOR_GADGET}" ]; then
    echo "error: Inspektor Gadget not found"
    exit 1
fi

# We do not need to add network eBPF programs like dns or snisnoop here because
# they do not rely on BTF (i.e., they do not include vmlinux.h).
${BTFHUB}/tools/btfgen.sh -a ${ARCH}                                                                \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/audit-seccomp/tracer/auditseccomp_bpfel_x86.o                \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/audit-seccomp/tracer/auditseccompwithfilters_bpfel_x86.o     \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/bindsnoop/tracer/core/bindsnoop_bpfel_x86.o                  \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/biolatency/tracer/core/biolatency_bpfel_x86.o                \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/biolatency/tracer/core/biolatencybefore_bpfel_x86.go         \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/biotop/tracer/biotop_bpfel_x86.o                             \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/capabilities/tracer/core/capabilities_bpfel_x86.o            \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/execsnoop/tracer/core/execsnoop_bpfel.o                      \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/filetop/tracer/filetop_bpfel_x86.o                           \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/fsslower/tracer/core/fsslower_bpfel_x86.o                    \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/mountsnoop/tracer/core/mountsnoop_bpfel.o                    \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/oomkill/tracer/oomkill_bpfel_x86.o                           \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/opensnoop/tracer/core/opensnoop_bpfel.o                      \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/process-collector/tracer/processcollector_bpfel.o            \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/process-collector/tracer/processcollectorwithfilters_bpfel.o \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/seccomp/tracer/seccomp_bpfel_x86.o                           \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/socket-collector/tracer/itertcpv4_bpfel.o                    \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/socket-collector/tracer/iterudpv4_bpfel.o                    \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/sigsnoop/tracer/core/sigsnoop_bpfel.o                        \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/snisnoop/tracer/snisnoop_bpfel.o                             \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/tcpconnect/tracer/core/tcpconnect_bpfel_x86.o                \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/tcptracer/tracer/core/tcptracer_bpfel_x86.o                  \
    -o ${INSPEKTOR_GADGET}/pkg/gadgets/tcptop/tracer/tcptop_bpfel_x86.o                             \
    #

mkdir -p ${OUTPUT}
cp -r ${BTFHUB}/custom-archive/* ${OUTPUT}
