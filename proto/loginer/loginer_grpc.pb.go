// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: loginer.proto

package loginer

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

// LoginerClient is the client API for Loginer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginerClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type loginerClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginerClient(cc grpc.ClientConnInterface) LoginerClient {
	return &loginerClient{cc}
}

func (c *loginerClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/Loginer/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginerServer is the server API for Loginer service.
// All implementations must embed UnimplementedLoginerServer
// for forward compatibility
type LoginerServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedLoginerServer()
}

// UnimplementedLoginerServer must be embedded to have forward compatible implementations.
type UnimplementedLoginerServer struct {
}

func (UnimplementedLoginerServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedLoginerServer) mustEmbedUnimplementedLoginerServer() {}

// UnsafeLoginerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginerServer will
// result in compilation errors.
type UnsafeLoginerServer interface {
	mustEmbedUnimplementedLoginerServer()
}

func RegisterLoginerServer(s grpc.ServiceRegistrar, srv LoginerServer) {
	s.RegisterService(&Loginer_ServiceDesc, srv)
}

func _Loginer_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginerServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Loginer/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginerServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Loginer_ServiceDesc is the grpc.ServiceDesc for Loginer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Loginer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Loginer",
	HandlerType: (*LoginerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Loginer_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loginer.proto",
}
