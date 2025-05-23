// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: donation.proto

package donationfields

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
	DonationConfig_GetDonationConfig_FullMethodName = "/DonationConfig/GetDonationConfig"
)

// DonationConfigClient is the client API for DonationConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DonationConfigClient interface {
	GetDonationConfig(ctx context.Context, in *DonationConfigRequest, opts ...grpc.CallOption) (*DonationConfigResponse, error)
}

type donationConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewDonationConfigClient(cc grpc.ClientConnInterface) DonationConfigClient {
	return &donationConfigClient{cc}
}

func (c *donationConfigClient) GetDonationConfig(ctx context.Context, in *DonationConfigRequest, opts ...grpc.CallOption) (*DonationConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DonationConfigResponse)
	err := c.cc.Invoke(ctx, DonationConfig_GetDonationConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DonationConfigServer is the server API for DonationConfig service.
// All implementations must embed UnimplementedDonationConfigServer
// for forward compatibility.
type DonationConfigServer interface {
	GetDonationConfig(context.Context, *DonationConfigRequest) (*DonationConfigResponse, error)
	mustEmbedUnimplementedDonationConfigServer()
}

// UnimplementedDonationConfigServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDonationConfigServer struct{}

func (UnimplementedDonationConfigServer) GetDonationConfig(context.Context, *DonationConfigRequest) (*DonationConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDonationConfig not implemented")
}
func (UnimplementedDonationConfigServer) mustEmbedUnimplementedDonationConfigServer() {}
func (UnimplementedDonationConfigServer) testEmbeddedByValue()                        {}

// UnsafeDonationConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DonationConfigServer will
// result in compilation errors.
type UnsafeDonationConfigServer interface {
	mustEmbedUnimplementedDonationConfigServer()
}

func RegisterDonationConfigServer(s grpc.ServiceRegistrar, srv DonationConfigServer) {
	// If the following call pancis, it indicates UnimplementedDonationConfigServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DonationConfig_ServiceDesc, srv)
}

func _DonationConfig_GetDonationConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DonationConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DonationConfigServer).GetDonationConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DonationConfig_GetDonationConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DonationConfigServer).GetDonationConfig(ctx, req.(*DonationConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DonationConfig_ServiceDesc is the grpc.ServiceDesc for DonationConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DonationConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DonationConfig",
	HandlerType: (*DonationConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDonationConfig",
			Handler:    _DonationConfig_GetDonationConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "donation.proto",
}
