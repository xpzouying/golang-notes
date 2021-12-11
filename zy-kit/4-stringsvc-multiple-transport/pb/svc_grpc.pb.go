// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// StringClient is the client API for String service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StringClient interface {
	Upper(ctx context.Context, in *UpperRequest, opts ...grpc.CallOption) (*UpperResponse, error)
}

type stringClient struct {
	cc grpc.ClientConnInterface
}

func NewStringClient(cc grpc.ClientConnInterface) StringClient {
	return &stringClient{cc}
}

func (c *stringClient) Upper(ctx context.Context, in *UpperRequest, opts ...grpc.CallOption) (*UpperResponse, error) {
	out := new(UpperResponse)
	err := c.cc.Invoke(ctx, "/pb.String/Upper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StringServer is the server API for String service.
// All implementations must embed UnimplementedStringServer
// for forward compatibility
type StringServer interface {
	Upper(context.Context, *UpperRequest) (*UpperResponse, error)
	mustEmbedUnimplementedStringServer()
}

// UnimplementedStringServer must be embedded to have forward compatible implementations.
type UnimplementedStringServer struct {
}

func (UnimplementedStringServer) Upper(context.Context, *UpperRequest) (*UpperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upper not implemented")
}
func (UnimplementedStringServer) mustEmbedUnimplementedStringServer() {}

// UnsafeStringServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StringServer will
// result in compilation errors.
type UnsafeStringServer interface {
	mustEmbedUnimplementedStringServer()
}

func RegisterStringServer(s grpc.ServiceRegistrar, srv StringServer) {
	s.RegisterService(&String_ServiceDesc, srv)
}

func _String_Upper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StringServer).Upper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.String/Upper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StringServer).Upper(ctx, req.(*UpperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// String_ServiceDesc is the grpc.ServiceDesc for String service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var String_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.String",
	HandlerType: (*StringServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Upper",
			Handler:    _String_Upper_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "svc.proto",
}
