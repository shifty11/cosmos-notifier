// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: dev_service.proto

package dev_service

import (
	context "context"
	auth_service "github.com/shifty11/cosmos-notifier/services/grpc/protobuf/auth_service"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DevService_Login_FullMethodName = "/cosmos_notifier_grpc.DevService/Login"
)

// DevServiceClient is the client API for DevService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DevServiceClient interface {
	Login(ctx context.Context, in *DevLoginRequest, opts ...grpc.CallOption) (*auth_service.LoginResponse, error)
}

type devServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDevServiceClient(cc grpc.ClientConnInterface) DevServiceClient {
	return &devServiceClient{cc}
}

func (c *devServiceClient) Login(ctx context.Context, in *DevLoginRequest, opts ...grpc.CallOption) (*auth_service.LoginResponse, error) {
	out := new(auth_service.LoginResponse)
	err := c.cc.Invoke(ctx, DevService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DevServiceServer is the server API for DevService service.
// All implementations must embed UnimplementedDevServiceServer
// for forward compatibility
type DevServiceServer interface {
	Login(context.Context, *DevLoginRequest) (*auth_service.LoginResponse, error)
	mustEmbedUnimplementedDevServiceServer()
}

// UnimplementedDevServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDevServiceServer struct {
}

func (UnimplementedDevServiceServer) Login(context.Context, *DevLoginRequest) (*auth_service.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedDevServiceServer) mustEmbedUnimplementedDevServiceServer() {}

// UnsafeDevServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DevServiceServer will
// result in compilation errors.
type UnsafeDevServiceServer interface {
	mustEmbedUnimplementedDevServiceServer()
}

func RegisterDevServiceServer(s grpc.ServiceRegistrar, srv DevServiceServer) {
	s.RegisterService(&DevService_ServiceDesc, srv)
}

func _DevService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DevLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DevService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevServiceServer).Login(ctx, req.(*DevLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DevService_ServiceDesc is the grpc.ServiceDesc for DevService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DevService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos_notifier_grpc.DevService",
	HandlerType: (*DevServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _DevService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dev_service.proto",
}
