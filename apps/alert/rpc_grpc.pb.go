// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: apps/alert/pb/rpc.proto

package alert

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

const (
	Rpc_HandlerAlert_FullMethodName    = "/K8sGenie.alert.Rpc/HandlerAlert"
	Rpc_UrgentAlertCall_FullMethodName = "/K8sGenie.alert.Rpc/UrgentAlertCall"
)

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	HandlerAlert(ctx context.Context, in *HandlerAlertReq, opts ...grpc.CallOption) (*HandlerAlertRsp, error)
	UrgentAlertCall(ctx context.Context, in *UrgentAlertCallRequest, opts ...grpc.CallOption) (*UrgentAlert, error)
}

type rpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcClient(cc grpc.ClientConnInterface) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) HandlerAlert(ctx context.Context, in *HandlerAlertReq, opts ...grpc.CallOption) (*HandlerAlertRsp, error) {
	out := new(HandlerAlertRsp)
	err := c.cc.Invoke(ctx, Rpc_HandlerAlert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) UrgentAlertCall(ctx context.Context, in *UrgentAlertCallRequest, opts ...grpc.CallOption) (*UrgentAlert, error) {
	out := new(UrgentAlert)
	err := c.cc.Invoke(ctx, Rpc_UrgentAlertCall_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServer is the server API for Rpc service.
// All implementations must embed UnimplementedRpcServer
// for forward compatibility
type RpcServer interface {
	HandlerAlert(context.Context, *HandlerAlertReq) (*HandlerAlertRsp, error)
	UrgentAlertCall(context.Context, *UrgentAlertCallRequest) (*UrgentAlert, error)
	mustEmbedUnimplementedRpcServer()
}

// UnimplementedRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServer struct {
}

func (UnimplementedRpcServer) HandlerAlert(context.Context, *HandlerAlertReq) (*HandlerAlertRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandlerAlert not implemented")
}
func (UnimplementedRpcServer) UrgentAlertCall(context.Context, *UrgentAlertCallRequest) (*UrgentAlert, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UrgentAlertCall not implemented")
}
func (UnimplementedRpcServer) mustEmbedUnimplementedRpcServer() {}

// UnsafeRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcServer will
// result in compilation errors.
type UnsafeRpcServer interface {
	mustEmbedUnimplementedRpcServer()
}

func RegisterRpcServer(s grpc.ServiceRegistrar, srv RpcServer) {
	s.RegisterService(&Rpc_ServiceDesc, srv)
}

func _Rpc_HandlerAlert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandlerAlertReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).HandlerAlert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_HandlerAlert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).HandlerAlert(ctx, req.(*HandlerAlertReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_UrgentAlertCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrgentAlertCallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).UrgentAlertCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_UrgentAlertCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).UrgentAlertCall(ctx, req.(*UrgentAlertCallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "K8sGenie.alert.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandlerAlert",
			Handler:    _Rpc_HandlerAlert_Handler,
		},
		{
			MethodName: "UrgentAlertCall",
			Handler:    _Rpc_UrgentAlertCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/alert/pb/rpc.proto",
}
