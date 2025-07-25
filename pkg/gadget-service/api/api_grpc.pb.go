// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BuiltInGadgetManagerClient is the client API for BuiltInGadgetManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuiltInGadgetManagerClient interface {
	GetInfo(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
}

type builtInGadgetManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewBuiltInGadgetManagerClient(cc grpc.ClientConnInterface) BuiltInGadgetManagerClient {
	return &builtInGadgetManagerClient{cc}
}

func (c *builtInGadgetManagerClient) GetInfo(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/api.BuiltInGadgetManager/GetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuiltInGadgetManagerServer is the server API for BuiltInGadgetManager service.
// All implementations must embed UnimplementedBuiltInGadgetManagerServer
// for forward compatibility
type BuiltInGadgetManagerServer interface {
	GetInfo(context.Context, *InfoRequest) (*InfoResponse, error)
	mustEmbedUnimplementedBuiltInGadgetManagerServer()
}

// UnimplementedBuiltInGadgetManagerServer must be embedded to have forward compatible implementations.
type UnimplementedBuiltInGadgetManagerServer struct {
}

func (UnimplementedBuiltInGadgetManagerServer) GetInfo(context.Context, *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInfo not implemented")
}
func (UnimplementedBuiltInGadgetManagerServer) mustEmbedUnimplementedBuiltInGadgetManagerServer() {}

// UnsafeBuiltInGadgetManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuiltInGadgetManagerServer will
// result in compilation errors.
type UnsafeBuiltInGadgetManagerServer interface {
	mustEmbedUnimplementedBuiltInGadgetManagerServer()
}

func RegisterBuiltInGadgetManagerServer(s *grpc.Server, srv BuiltInGadgetManagerServer) {
	s.RegisterService(&_BuiltInGadgetManager_serviceDesc, srv)
}

func _BuiltInGadgetManager_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuiltInGadgetManagerServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.BuiltInGadgetManager/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuiltInGadgetManagerServer).GetInfo(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BuiltInGadgetManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.BuiltInGadgetManager",
	HandlerType: (*BuiltInGadgetManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _BuiltInGadgetManager_GetInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}

// GadgetManagerClient is the client API for GadgetManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GadgetManagerClient interface {
	GetGadgetInfo(ctx context.Context, in *GetGadgetInfoRequest, opts ...grpc.CallOption) (*GetGadgetInfoResponse, error)
	RunGadget(ctx context.Context, opts ...grpc.CallOption) (GadgetManager_RunGadgetClient, error)
}

type gadgetManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewGadgetManagerClient(cc grpc.ClientConnInterface) GadgetManagerClient {
	return &gadgetManagerClient{cc}
}

func (c *gadgetManagerClient) GetGadgetInfo(ctx context.Context, in *GetGadgetInfoRequest, opts ...grpc.CallOption) (*GetGadgetInfoResponse, error) {
	out := new(GetGadgetInfoResponse)
	err := c.cc.Invoke(ctx, "/api.GadgetManager/GetGadgetInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gadgetManagerClient) RunGadget(ctx context.Context, opts ...grpc.CallOption) (GadgetManager_RunGadgetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GadgetManager_serviceDesc.Streams[0], "/api.GadgetManager/RunGadget", opts...)
	if err != nil {
		return nil, err
	}
	x := &gadgetManagerRunGadgetClient{stream}
	return x, nil
}

type GadgetManager_RunGadgetClient interface {
	Send(*GadgetControlRequest) error
	Recv() (*GadgetEvent, error)
	grpc.ClientStream
}

type gadgetManagerRunGadgetClient struct {
	grpc.ClientStream
}

func (x *gadgetManagerRunGadgetClient) Send(m *GadgetControlRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gadgetManagerRunGadgetClient) Recv() (*GadgetEvent, error) {
	m := new(GadgetEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GadgetManagerServer is the server API for GadgetManager service.
// All implementations must embed UnimplementedGadgetManagerServer
// for forward compatibility
type GadgetManagerServer interface {
	GetGadgetInfo(context.Context, *GetGadgetInfoRequest) (*GetGadgetInfoResponse, error)
	RunGadget(GadgetManager_RunGadgetServer) error
	mustEmbedUnimplementedGadgetManagerServer()
}

// UnimplementedGadgetManagerServer must be embedded to have forward compatible implementations.
type UnimplementedGadgetManagerServer struct {
}

func (UnimplementedGadgetManagerServer) GetGadgetInfo(context.Context, *GetGadgetInfoRequest) (*GetGadgetInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGadgetInfo not implemented")
}
func (UnimplementedGadgetManagerServer) RunGadget(GadgetManager_RunGadgetServer) error {
	return status.Errorf(codes.Unimplemented, "method RunGadget not implemented")
}
func (UnimplementedGadgetManagerServer) mustEmbedUnimplementedGadgetManagerServer() {}

// UnsafeGadgetManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GadgetManagerServer will
// result in compilation errors.
type UnsafeGadgetManagerServer interface {
	mustEmbedUnimplementedGadgetManagerServer()
}

func RegisterGadgetManagerServer(s *grpc.Server, srv GadgetManagerServer) {
	s.RegisterService(&_GadgetManager_serviceDesc, srv)
}

func _GadgetManager_GetGadgetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGadgetInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GadgetManagerServer).GetGadgetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GadgetManager/GetGadgetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GadgetManagerServer).GetGadgetInfo(ctx, req.(*GetGadgetInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GadgetManager_RunGadget_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GadgetManagerServer).RunGadget(&gadgetManagerRunGadgetServer{stream})
}

type GadgetManager_RunGadgetServer interface {
	Send(*GadgetEvent) error
	Recv() (*GadgetControlRequest, error)
	grpc.ServerStream
}

type gadgetManagerRunGadgetServer struct {
	grpc.ServerStream
}

func (x *gadgetManagerRunGadgetServer) Send(m *GadgetEvent) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gadgetManagerRunGadgetServer) Recv() (*GadgetControlRequest, error) {
	m := new(GadgetControlRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GadgetManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.GadgetManager",
	HandlerType: (*GadgetManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGadgetInfo",
			Handler:    _GadgetManager_GetGadgetInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunGadget",
			Handler:       _GadgetManager_RunGadget_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/api.proto",
}

// GadgetInstanceManagerClient is the client API for GadgetInstanceManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GadgetInstanceManagerClient interface {
	CreateGadgetInstance(ctx context.Context, in *CreateGadgetInstanceRequest, opts ...grpc.CallOption) (*CreateGadgetInstanceResponse, error)
	ListGadgetInstances(ctx context.Context, in *ListGadgetInstancesRequest, opts ...grpc.CallOption) (*ListGadgetInstanceResponse, error)
	GetGadgetInstance(ctx context.Context, in *GadgetInstanceId, opts ...grpc.CallOption) (*GadgetInstance, error)
	RemoveGadgetInstance(ctx context.Context, in *GadgetInstanceId, opts ...grpc.CallOption) (*StatusResponse, error)
}

type gadgetInstanceManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewGadgetInstanceManagerClient(cc grpc.ClientConnInterface) GadgetInstanceManagerClient {
	return &gadgetInstanceManagerClient{cc}
}

func (c *gadgetInstanceManagerClient) CreateGadgetInstance(ctx context.Context, in *CreateGadgetInstanceRequest, opts ...grpc.CallOption) (*CreateGadgetInstanceResponse, error) {
	out := new(CreateGadgetInstanceResponse)
	err := c.cc.Invoke(ctx, "/api.GadgetInstanceManager/CreateGadgetInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gadgetInstanceManagerClient) ListGadgetInstances(ctx context.Context, in *ListGadgetInstancesRequest, opts ...grpc.CallOption) (*ListGadgetInstanceResponse, error) {
	out := new(ListGadgetInstanceResponse)
	err := c.cc.Invoke(ctx, "/api.GadgetInstanceManager/ListGadgetInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gadgetInstanceManagerClient) GetGadgetInstance(ctx context.Context, in *GadgetInstanceId, opts ...grpc.CallOption) (*GadgetInstance, error) {
	out := new(GadgetInstance)
	err := c.cc.Invoke(ctx, "/api.GadgetInstanceManager/GetGadgetInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gadgetInstanceManagerClient) RemoveGadgetInstance(ctx context.Context, in *GadgetInstanceId, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/api.GadgetInstanceManager/RemoveGadgetInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GadgetInstanceManagerServer is the server API for GadgetInstanceManager service.
// All implementations must embed UnimplementedGadgetInstanceManagerServer
// for forward compatibility
type GadgetInstanceManagerServer interface {
	CreateGadgetInstance(context.Context, *CreateGadgetInstanceRequest) (*CreateGadgetInstanceResponse, error)
	ListGadgetInstances(context.Context, *ListGadgetInstancesRequest) (*ListGadgetInstanceResponse, error)
	GetGadgetInstance(context.Context, *GadgetInstanceId) (*GadgetInstance, error)
	RemoveGadgetInstance(context.Context, *GadgetInstanceId) (*StatusResponse, error)
	mustEmbedUnimplementedGadgetInstanceManagerServer()
}

// UnimplementedGadgetInstanceManagerServer must be embedded to have forward compatible implementations.
type UnimplementedGadgetInstanceManagerServer struct {
}

func (UnimplementedGadgetInstanceManagerServer) CreateGadgetInstance(context.Context, *CreateGadgetInstanceRequest) (*CreateGadgetInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGadgetInstance not implemented")
}
func (UnimplementedGadgetInstanceManagerServer) ListGadgetInstances(context.Context, *ListGadgetInstancesRequest) (*ListGadgetInstanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListGadgetInstances not implemented")
}
func (UnimplementedGadgetInstanceManagerServer) GetGadgetInstance(context.Context, *GadgetInstanceId) (*GadgetInstance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGadgetInstance not implemented")
}
func (UnimplementedGadgetInstanceManagerServer) RemoveGadgetInstance(context.Context, *GadgetInstanceId) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGadgetInstance not implemented")
}
func (UnimplementedGadgetInstanceManagerServer) mustEmbedUnimplementedGadgetInstanceManagerServer() {}

// UnsafeGadgetInstanceManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GadgetInstanceManagerServer will
// result in compilation errors.
type UnsafeGadgetInstanceManagerServer interface {
	mustEmbedUnimplementedGadgetInstanceManagerServer()
}

func RegisterGadgetInstanceManagerServer(s *grpc.Server, srv GadgetInstanceManagerServer) {
	s.RegisterService(&_GadgetInstanceManager_serviceDesc, srv)
}

func _GadgetInstanceManager_CreateGadgetInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGadgetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GadgetInstanceManagerServer).CreateGadgetInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GadgetInstanceManager/CreateGadgetInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GadgetInstanceManagerServer).CreateGadgetInstance(ctx, req.(*CreateGadgetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GadgetInstanceManager_ListGadgetInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGadgetInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GadgetInstanceManagerServer).ListGadgetInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GadgetInstanceManager/ListGadgetInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GadgetInstanceManagerServer).ListGadgetInstances(ctx, req.(*ListGadgetInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GadgetInstanceManager_GetGadgetInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GadgetInstanceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GadgetInstanceManagerServer).GetGadgetInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GadgetInstanceManager/GetGadgetInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GadgetInstanceManagerServer).GetGadgetInstance(ctx, req.(*GadgetInstanceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _GadgetInstanceManager_RemoveGadgetInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GadgetInstanceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GadgetInstanceManagerServer).RemoveGadgetInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.GadgetInstanceManager/RemoveGadgetInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GadgetInstanceManagerServer).RemoveGadgetInstance(ctx, req.(*GadgetInstanceId))
	}
	return interceptor(ctx, in, info, handler)
}

var _GadgetInstanceManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.GadgetInstanceManager",
	HandlerType: (*GadgetInstanceManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGadgetInstance",
			Handler:    _GadgetInstanceManager_CreateGadgetInstance_Handler,
		},
		{
			MethodName: "ListGadgetInstances",
			Handler:    _GadgetInstanceManager_ListGadgetInstances_Handler,
		},
		{
			MethodName: "GetGadgetInstance",
			Handler:    _GadgetInstanceManager_GetGadgetInstance_Handler,
		},
		{
			MethodName: "RemoveGadgetInstance",
			Handler:    _GadgetInstanceManager_RemoveGadgetInstance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}
