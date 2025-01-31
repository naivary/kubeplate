// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: input/v1/input.proto

package v1

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
	Inputer_Read_FullMethodName = "/input.v1.Inputer/Read"
)

// InputerClient is the client API for Inputer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InputerClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
}

type inputerClient struct {
	cc grpc.ClientConnInterface
}

func NewInputerClient(cc grpc.ClientConnInterface) InputerClient {
	return &inputerClient{cc}
}

func (c *inputerClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, Inputer_Read_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InputerServer is the server API for Inputer service.
// All implementations must embed UnimplementedInputerServer
// for forward compatibility.
type InputerServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	mustEmbedUnimplementedInputerServer()
}

// UnimplementedInputerServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedInputerServer struct{}

func (UnimplementedInputerServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedInputerServer) mustEmbedUnimplementedInputerServer() {}
func (UnimplementedInputerServer) testEmbeddedByValue()                 {}

// UnsafeInputerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InputerServer will
// result in compilation errors.
type UnsafeInputerServer interface {
	mustEmbedUnimplementedInputerServer()
}

func RegisterInputerServer(s grpc.ServiceRegistrar, srv InputerServer) {
	// If the following call pancis, it indicates UnimplementedInputerServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Inputer_ServiceDesc, srv)
}

func _Inputer_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InputerServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Inputer_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InputerServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Inputer_ServiceDesc is the grpc.ServiceDesc for Inputer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Inputer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "input.v1.Inputer",
	HandlerType: (*InputerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Inputer_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "input/v1/input.proto",
}
