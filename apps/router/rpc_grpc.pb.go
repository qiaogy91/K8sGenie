// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.28.3
// source: apps/router/pb/rpc.proto

package router

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
	Rpc_CreateTable_FullMethodName  = "/K8sGenie.router.Rpc/CreateTable"
	Rpc_CreateRoute_FullMethodName  = "/K8sGenie.router.Rpc/CreateRoute"
	Rpc_DeleteRoute_FullMethodName  = "/K8sGenie.router.Rpc/DeleteRoute"
	Rpc_QueryRoute_FullMethodName   = "/K8sGenie.router.Rpc/QueryRoute"
	Rpc_UpdateRoute_FullMethodName  = "/K8sGenie.router.Rpc/UpdateRoute"
	Rpc_DescRoute_FullMethodName    = "/K8sGenie.router.Rpc/DescRoute"
	Rpc_UrgentChange_FullMethodName = "/K8sGenie.router.Rpc/UrgentChange"
	Rpc_AlertRoute_FullMethodName   = "/K8sGenie.router.Rpc/AlertRoute"
)

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	CreateTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	CreateRoute(ctx context.Context, in *Spec, opts ...grpc.CallOption) (*Router, error)
	DeleteRoute(ctx context.Context, in *DeleteRouteReq, opts ...grpc.CallOption) (*Router, error)
	QueryRoute(ctx context.Context, in *QueryRouteReq, opts ...grpc.CallOption) (*RouterSet, error)
	UpdateRoute(ctx context.Context, in *UpdateRouteReq, opts ...grpc.CallOption) (*Router, error)
	DescRoute(ctx context.Context, in *DescRouteReq, opts ...grpc.CallOption) (*Router, error)
	UrgentChange(ctx context.Context, in *UrgentChangeReq, opts ...grpc.CallOption) (*Router, error)
	// 告警路由
	AlertRoute(ctx context.Context, in *AlertRouteReq, opts ...grpc.CallOption) (*Router, error)
}

type rpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcClient(cc grpc.ClientConnInterface) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) CreateTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Rpc_CreateTable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) CreateRoute(ctx context.Context, in *Spec, opts ...grpc.CallOption) (*Router, error) {
	out := new(Router)
	err := c.cc.Invoke(ctx, Rpc_CreateRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) DeleteRoute(ctx context.Context, in *DeleteRouteReq, opts ...grpc.CallOption) (*Router, error) {
	out := new(Router)
	err := c.cc.Invoke(ctx, Rpc_DeleteRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) QueryRoute(ctx context.Context, in *QueryRouteReq, opts ...grpc.CallOption) (*RouterSet, error) {
	out := new(RouterSet)
	err := c.cc.Invoke(ctx, Rpc_QueryRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) UpdateRoute(ctx context.Context, in *UpdateRouteReq, opts ...grpc.CallOption) (*Router, error) {
	out := new(Router)
	err := c.cc.Invoke(ctx, Rpc_UpdateRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) DescRoute(ctx context.Context, in *DescRouteReq, opts ...grpc.CallOption) (*Router, error) {
	out := new(Router)
	err := c.cc.Invoke(ctx, Rpc_DescRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) UrgentChange(ctx context.Context, in *UrgentChangeReq, opts ...grpc.CallOption) (*Router, error) {
	out := new(Router)
	err := c.cc.Invoke(ctx, Rpc_UrgentChange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) AlertRoute(ctx context.Context, in *AlertRouteReq, opts ...grpc.CallOption) (*Router, error) {
	out := new(Router)
	err := c.cc.Invoke(ctx, Rpc_AlertRoute_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServer is the server API for Rpc service.
// All implementations must embed UnimplementedRpcServer
// for forward compatibility
type RpcServer interface {
	CreateTable(context.Context, *Empty) (*Empty, error)
	CreateRoute(context.Context, *Spec) (*Router, error)
	DeleteRoute(context.Context, *DeleteRouteReq) (*Router, error)
	QueryRoute(context.Context, *QueryRouteReq) (*RouterSet, error)
	UpdateRoute(context.Context, *UpdateRouteReq) (*Router, error)
	DescRoute(context.Context, *DescRouteReq) (*Router, error)
	UrgentChange(context.Context, *UrgentChangeReq) (*Router, error)
	// 告警路由
	AlertRoute(context.Context, *AlertRouteReq) (*Router, error)
	mustEmbedUnimplementedRpcServer()
}

// UnimplementedRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServer struct {
}

func (UnimplementedRpcServer) CreateTable(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTable not implemented")
}
func (UnimplementedRpcServer) CreateRoute(context.Context, *Spec) (*Router, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoute not implemented")
}
func (UnimplementedRpcServer) DeleteRoute(context.Context, *DeleteRouteReq) (*Router, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRoute not implemented")
}
func (UnimplementedRpcServer) QueryRoute(context.Context, *QueryRouteReq) (*RouterSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryRoute not implemented")
}
func (UnimplementedRpcServer) UpdateRoute(context.Context, *UpdateRouteReq) (*Router, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoute not implemented")
}
func (UnimplementedRpcServer) DescRoute(context.Context, *DescRouteReq) (*Router, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescRoute not implemented")
}
func (UnimplementedRpcServer) UrgentChange(context.Context, *UrgentChangeReq) (*Router, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UrgentChange not implemented")
}
func (UnimplementedRpcServer) AlertRoute(context.Context, *AlertRouteReq) (*Router, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AlertRoute not implemented")
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

func _Rpc_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_CreateTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).CreateTable(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_CreateRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Spec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).CreateRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_CreateRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).CreateRoute(ctx, req.(*Spec))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_DeleteRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRouteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).DeleteRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_DeleteRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).DeleteRoute(ctx, req.(*DeleteRouteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_QueryRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRouteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).QueryRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_QueryRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).QueryRoute(ctx, req.(*QueryRouteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_UpdateRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRouteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).UpdateRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_UpdateRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).UpdateRoute(ctx, req.(*UpdateRouteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_DescRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescRouteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).DescRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_DescRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).DescRoute(ctx, req.(*DescRouteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_UrgentChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrgentChangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).UrgentChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_UrgentChange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).UrgentChange(ctx, req.(*UrgentChangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_AlertRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertRouteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).AlertRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_AlertRoute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).AlertRoute(ctx, req.(*AlertRouteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "K8sGenie.router.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTable",
			Handler:    _Rpc_CreateTable_Handler,
		},
		{
			MethodName: "CreateRoute",
			Handler:    _Rpc_CreateRoute_Handler,
		},
		{
			MethodName: "DeleteRoute",
			Handler:    _Rpc_DeleteRoute_Handler,
		},
		{
			MethodName: "QueryRoute",
			Handler:    _Rpc_QueryRoute_Handler,
		},
		{
			MethodName: "UpdateRoute",
			Handler:    _Rpc_UpdateRoute_Handler,
		},
		{
			MethodName: "DescRoute",
			Handler:    _Rpc_DescRoute_Handler,
		},
		{
			MethodName: "UrgentChange",
			Handler:    _Rpc_UrgentChange_Handler,
		},
		{
			MethodName: "AlertRoute",
			Handler:    _Rpc_AlertRoute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/router/pb/rpc.proto",
}
