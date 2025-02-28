// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: proto/resiliency/resiliency.proto

package resiliency

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ResiliencyService_UnaryResiliency_FullMethodName                  = "/resiliency.ResiliencyService/UnaryResiliency"
	ResiliencyService_ServerStreamingResiliency_FullMethodName        = "/resiliency.ResiliencyService/ServerStreamingResiliency"
	ResiliencyService_ClientStreamingResiliency_FullMethodName        = "/resiliency.ResiliencyService/ClientStreamingResiliency"
	ResiliencyService_BidirectionalStreamingResiliency_FullMethodName = "/resiliency.ResiliencyService/BidirectionalStreamingResiliency"
)

// ResiliencyServiceClient is the client API for ResiliencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResiliencyServiceClient interface {
	UnaryResiliency(ctx context.Context, in *ResiliencyRequest, opts ...grpc.CallOption) (*ResiliencyResponse, error)
	ServerStreamingResiliency(ctx context.Context, in *ResiliencyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ResiliencyResponse], error)
	ClientStreamingResiliency(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ResiliencyRequest, ResiliencyResponse], error)
	BidirectionalStreamingResiliency(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ResiliencyRequest, ResiliencyResponse], error)
}

type resiliencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResiliencyServiceClient(cc grpc.ClientConnInterface) ResiliencyServiceClient {
	return &resiliencyServiceClient{cc}
}

func (c *resiliencyServiceClient) UnaryResiliency(ctx context.Context, in *ResiliencyRequest, opts ...grpc.CallOption) (*ResiliencyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResiliencyResponse)
	err := c.cc.Invoke(ctx, ResiliencyService_UnaryResiliency_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resiliencyServiceClient) ServerStreamingResiliency(ctx context.Context, in *ResiliencyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ResiliencyResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ResiliencyService_ServiceDesc.Streams[0], ResiliencyService_ServerStreamingResiliency_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ResiliencyRequest, ResiliencyResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ResiliencyService_ServerStreamingResiliencyClient = grpc.ServerStreamingClient[ResiliencyResponse]

func (c *resiliencyServiceClient) ClientStreamingResiliency(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ResiliencyRequest, ResiliencyResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ResiliencyService_ServiceDesc.Streams[1], ResiliencyService_ClientStreamingResiliency_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ResiliencyRequest, ResiliencyResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ResiliencyService_ClientStreamingResiliencyClient = grpc.ClientStreamingClient[ResiliencyRequest, ResiliencyResponse]

func (c *resiliencyServiceClient) BidirectionalStreamingResiliency(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ResiliencyRequest, ResiliencyResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ResiliencyService_ServiceDesc.Streams[2], ResiliencyService_BidirectionalStreamingResiliency_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ResiliencyRequest, ResiliencyResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ResiliencyService_BidirectionalStreamingResiliencyClient = grpc.BidiStreamingClient[ResiliencyRequest, ResiliencyResponse]

// ResiliencyServiceServer is the server API for ResiliencyService service.
// All implementations must embed UnimplementedResiliencyServiceServer
// for forward compatibility.
type ResiliencyServiceServer interface {
	UnaryResiliency(context.Context, *ResiliencyRequest) (*ResiliencyResponse, error)
	ServerStreamingResiliency(*ResiliencyRequest, grpc.ServerStreamingServer[ResiliencyResponse]) error
	ClientStreamingResiliency(grpc.ClientStreamingServer[ResiliencyRequest, ResiliencyResponse]) error
	BidirectionalStreamingResiliency(grpc.BidiStreamingServer[ResiliencyRequest, ResiliencyResponse]) error
	mustEmbedUnimplementedResiliencyServiceServer()
}

// UnimplementedResiliencyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedResiliencyServiceServer struct{}

func (UnimplementedResiliencyServiceServer) UnaryResiliency(context.Context, *ResiliencyRequest) (*ResiliencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryResiliency not implemented")
}
func (UnimplementedResiliencyServiceServer) ServerStreamingResiliency(*ResiliencyRequest, grpc.ServerStreamingServer[ResiliencyResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamingResiliency not implemented")
}
func (UnimplementedResiliencyServiceServer) ClientStreamingResiliency(grpc.ClientStreamingServer[ResiliencyRequest, ResiliencyResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamingResiliency not implemented")
}
func (UnimplementedResiliencyServiceServer) BidirectionalStreamingResiliency(grpc.BidiStreamingServer[ResiliencyRequest, ResiliencyResponse]) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreamingResiliency not implemented")
}
func (UnimplementedResiliencyServiceServer) mustEmbedUnimplementedResiliencyServiceServer() {}
func (UnimplementedResiliencyServiceServer) testEmbeddedByValue()                           {}

// UnsafeResiliencyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResiliencyServiceServer will
// result in compilation errors.
type UnsafeResiliencyServiceServer interface {
	mustEmbedUnimplementedResiliencyServiceServer()
}

func RegisterResiliencyServiceServer(s grpc.ServiceRegistrar, srv ResiliencyServiceServer) {
	// If the following call pancis, it indicates UnimplementedResiliencyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ResiliencyService_ServiceDesc, srv)
}

func _ResiliencyService_UnaryResiliency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResiliencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResiliencyServiceServer).UnaryResiliency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResiliencyService_UnaryResiliency_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResiliencyServiceServer).UnaryResiliency(ctx, req.(*ResiliencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResiliencyService_ServerStreamingResiliency_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ResiliencyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResiliencyServiceServer).ServerStreamingResiliency(m, &grpc.GenericServerStream[ResiliencyRequest, ResiliencyResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ResiliencyService_ServerStreamingResiliencyServer = grpc.ServerStreamingServer[ResiliencyResponse]

func _ResiliencyService_ClientStreamingResiliency_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ResiliencyServiceServer).ClientStreamingResiliency(&grpc.GenericServerStream[ResiliencyRequest, ResiliencyResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ResiliencyService_ClientStreamingResiliencyServer = grpc.ClientStreamingServer[ResiliencyRequest, ResiliencyResponse]

func _ResiliencyService_BidirectionalStreamingResiliency_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ResiliencyServiceServer).BidirectionalStreamingResiliency(&grpc.GenericServerStream[ResiliencyRequest, ResiliencyResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ResiliencyService_BidirectionalStreamingResiliencyServer = grpc.BidiStreamingServer[ResiliencyRequest, ResiliencyResponse]

// ResiliencyService_ServiceDesc is the grpc.ServiceDesc for ResiliencyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResiliencyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resiliency.ResiliencyService",
	HandlerType: (*ResiliencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryResiliency",
			Handler:    _ResiliencyService_UnaryResiliency_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamingResiliency",
			Handler:       _ResiliencyService_ServerStreamingResiliency_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamingResiliency",
			Handler:       _ResiliencyService_ClientStreamingResiliency_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStreamingResiliency",
			Handler:       _ResiliencyService_BidirectionalStreamingResiliency_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/resiliency/resiliency.proto",
}

const (
	ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_FullMethodName                  = "/resiliency.ResiliencyWithMetadataService/UnaryResiliencyWithMetadata"
	ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadata_FullMethodName        = "/resiliency.ResiliencyWithMetadataService/ServerStreamingResiliencyWithMetadata"
	ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadata_FullMethodName        = "/resiliency.ResiliencyWithMetadataService/ClientStreamingResiliencyWithMetadata"
	ResiliencyWithMetadataService_BidirectionalStreamingResiliencyWithMetadata_FullMethodName = "/resiliency.ResiliencyWithMetadataService/BidirectionalStreamingResiliencyWithMetadata"
)

// ResiliencyWithMetadataServiceClient is the client API for ResiliencyWithMetadataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResiliencyWithMetadataServiceClient interface {
	UnaryResiliencyWithMetadata(ctx context.Context, in *ResiliencyRequest, opts ...grpc.CallOption) (*ResiliencyResponse, error)
	ServerStreamingResiliencyWithMetadata(ctx context.Context, in *ResiliencyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ResiliencyResponse], error)
	ClientStreamingResiliencyWithMetadata(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ResiliencyRequest, ResiliencyResponse], error)
	BidirectionalStreamingResiliencyWithMetadata(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ResiliencyRequest, ResiliencyResponse], error)
}

type resiliencyWithMetadataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResiliencyWithMetadataServiceClient(cc grpc.ClientConnInterface) ResiliencyWithMetadataServiceClient {
	return &resiliencyWithMetadataServiceClient{cc}
}

func (c *resiliencyWithMetadataServiceClient) UnaryResiliencyWithMetadata(ctx context.Context, in *ResiliencyRequest, opts ...grpc.CallOption) (*ResiliencyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResiliencyResponse)
	err := c.cc.Invoke(ctx, ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resiliencyWithMetadataServiceClient) ServerStreamingResiliencyWithMetadata(ctx context.Context, in *ResiliencyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ResiliencyResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ResiliencyWithMetadataService_ServiceDesc.Streams[0], ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadata_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ResiliencyRequest, ResiliencyResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadataClient = grpc.ServerStreamingClient[ResiliencyResponse]

func (c *resiliencyWithMetadataServiceClient) ClientStreamingResiliencyWithMetadata(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ResiliencyRequest, ResiliencyResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ResiliencyWithMetadataService_ServiceDesc.Streams[1], ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadata_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ResiliencyRequest, ResiliencyResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadataClient = grpc.ClientStreamingClient[ResiliencyRequest, ResiliencyResponse]

func (c *resiliencyWithMetadataServiceClient) BidirectionalStreamingResiliencyWithMetadata(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ResiliencyRequest, ResiliencyResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ResiliencyWithMetadataService_ServiceDesc.Streams[2], ResiliencyWithMetadataService_BidirectionalStreamingResiliencyWithMetadata_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ResiliencyRequest, ResiliencyResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ResiliencyWithMetadataService_BidirectionalStreamingResiliencyWithMetadataClient = grpc.BidiStreamingClient[ResiliencyRequest, ResiliencyResponse]

// ResiliencyWithMetadataServiceServer is the server API for ResiliencyWithMetadataService service.
// All implementations must embed UnimplementedResiliencyWithMetadataServiceServer
// for forward compatibility.
type ResiliencyWithMetadataServiceServer interface {
	UnaryResiliencyWithMetadata(context.Context, *ResiliencyRequest) (*ResiliencyResponse, error)
	ServerStreamingResiliencyWithMetadata(*ResiliencyRequest, grpc.ServerStreamingServer[ResiliencyResponse]) error
	ClientStreamingResiliencyWithMetadata(grpc.ClientStreamingServer[ResiliencyRequest, ResiliencyResponse]) error
	BidirectionalStreamingResiliencyWithMetadata(grpc.BidiStreamingServer[ResiliencyRequest, ResiliencyResponse]) error
	mustEmbedUnimplementedResiliencyWithMetadataServiceServer()
}

// UnimplementedResiliencyWithMetadataServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedResiliencyWithMetadataServiceServer struct{}

func (UnimplementedResiliencyWithMetadataServiceServer) UnaryResiliencyWithMetadata(context.Context, *ResiliencyRequest) (*ResiliencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryResiliencyWithMetadata not implemented")
}
func (UnimplementedResiliencyWithMetadataServiceServer) ServerStreamingResiliencyWithMetadata(*ResiliencyRequest, grpc.ServerStreamingServer[ResiliencyResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamingResiliencyWithMetadata not implemented")
}
func (UnimplementedResiliencyWithMetadataServiceServer) ClientStreamingResiliencyWithMetadata(grpc.ClientStreamingServer[ResiliencyRequest, ResiliencyResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamingResiliencyWithMetadata not implemented")
}
func (UnimplementedResiliencyWithMetadataServiceServer) BidirectionalStreamingResiliencyWithMetadata(grpc.BidiStreamingServer[ResiliencyRequest, ResiliencyResponse]) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreamingResiliencyWithMetadata not implemented")
}
func (UnimplementedResiliencyWithMetadataServiceServer) mustEmbedUnimplementedResiliencyWithMetadataServiceServer() {
}
func (UnimplementedResiliencyWithMetadataServiceServer) testEmbeddedByValue() {}

// UnsafeResiliencyWithMetadataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResiliencyWithMetadataServiceServer will
// result in compilation errors.
type UnsafeResiliencyWithMetadataServiceServer interface {
	mustEmbedUnimplementedResiliencyWithMetadataServiceServer()
}

func RegisterResiliencyWithMetadataServiceServer(s grpc.ServiceRegistrar, srv ResiliencyWithMetadataServiceServer) {
	// If the following call pancis, it indicates UnimplementedResiliencyWithMetadataServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ResiliencyWithMetadataService_ServiceDesc, srv)
}

func _ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResiliencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResiliencyWithMetadataServiceServer).UnaryResiliencyWithMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResiliencyWithMetadataServiceServer).UnaryResiliencyWithMetadata(ctx, req.(*ResiliencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadata_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ResiliencyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResiliencyWithMetadataServiceServer).ServerStreamingResiliencyWithMetadata(m, &grpc.GenericServerStream[ResiliencyRequest, ResiliencyResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadataServer = grpc.ServerStreamingServer[ResiliencyResponse]

func _ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadata_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ResiliencyWithMetadataServiceServer).ClientStreamingResiliencyWithMetadata(&grpc.GenericServerStream[ResiliencyRequest, ResiliencyResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadataServer = grpc.ClientStreamingServer[ResiliencyRequest, ResiliencyResponse]

func _ResiliencyWithMetadataService_BidirectionalStreamingResiliencyWithMetadata_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ResiliencyWithMetadataServiceServer).BidirectionalStreamingResiliencyWithMetadata(&grpc.GenericServerStream[ResiliencyRequest, ResiliencyResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ResiliencyWithMetadataService_BidirectionalStreamingResiliencyWithMetadataServer = grpc.BidiStreamingServer[ResiliencyRequest, ResiliencyResponse]

// ResiliencyWithMetadataService_ServiceDesc is the grpc.ServiceDesc for ResiliencyWithMetadataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResiliencyWithMetadataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resiliency.ResiliencyWithMetadataService",
	HandlerType: (*ResiliencyWithMetadataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryResiliencyWithMetadata",
			Handler:    _ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamingResiliencyWithMetadata",
			Handler:       _ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadata_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamingResiliencyWithMetadata",
			Handler:       _ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadata_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStreamingResiliencyWithMetadata",
			Handler:       _ResiliencyWithMetadataService_BidirectionalStreamingResiliencyWithMetadata_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/resiliency/resiliency.proto",
}
