// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: admin_service.proto

package admin_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AdminService_BroadcastMessage_FullMethodName = "/cosmos_notifier_grpc.AdminService/BroadcastMessage"
	AdminService_GetStats_FullMethodName         = "/cosmos_notifier_grpc.AdminService/GetStats"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	BroadcastMessage(ctx context.Context, in *BroadcastMessageRequest, opts ...grpc.CallOption) (AdminService_BroadcastMessageClient, error)
	GetStats(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetStatsResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) BroadcastMessage(ctx context.Context, in *BroadcastMessageRequest, opts ...grpc.CallOption) (AdminService_BroadcastMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &AdminService_ServiceDesc.Streams[0], AdminService_BroadcastMessage_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &adminServiceBroadcastMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AdminService_BroadcastMessageClient interface {
	Recv() (*BroadcastMessageResponse, error)
	grpc.ClientStream
}

type adminServiceBroadcastMessageClient struct {
	grpc.ClientStream
}

func (x *adminServiceBroadcastMessageClient) Recv() (*BroadcastMessageResponse, error) {
	m := new(BroadcastMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *adminServiceClient) GetStats(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, AdminService_GetStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	BroadcastMessage(*BroadcastMessageRequest, AdminService_BroadcastMessageServer) error
	GetStats(context.Context, *empty.Empty) (*GetStatsResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) BroadcastMessage(*BroadcastMessageRequest, AdminService_BroadcastMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method BroadcastMessage not implemented")
}
func (UnimplementedAdminServiceServer) GetStats(context.Context, *empty.Empty) (*GetStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStats not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_BroadcastMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BroadcastMessageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AdminServiceServer).BroadcastMessage(m, &adminServiceBroadcastMessageServer{stream})
}

type AdminService_BroadcastMessageServer interface {
	Send(*BroadcastMessageResponse) error
	grpc.ServerStream
}

type adminServiceBroadcastMessageServer struct {
	grpc.ServerStream
}

func (x *adminServiceBroadcastMessageServer) Send(m *BroadcastMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _AdminService_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_GetStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetStats(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos_notifier_grpc.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStats",
			Handler:    _AdminService_GetStats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BroadcastMessage",
			Handler:       _AdminService_BroadcastMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "admin_service.proto",
}