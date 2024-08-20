// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.1
// source: apps/rancher/pb/rpc.proto

package rancher

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
	Rpc_CreateTable_FullMethodName  = "/K8sGenie.resourcer.Rpc/CreateTable"
	Rpc_SyncProject_FullMethodName  = "/K8sGenie.resourcer.Rpc/SyncProject"
	Rpc_QueryProject_FullMethodName = "/K8sGenie.resourcer.Rpc/QueryProject"
)

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	// for rancher
	CreateTable(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	SyncProject(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Rpc_SyncProjectClient, error)
	QueryProject(ctx context.Context, in *QueryProjectReq, opts ...grpc.CallOption) (*ProjectSet, error)
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

func (c *rpcClient) SyncProject(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Rpc_SyncProjectClient, error) {
	stream, err := c.cc.NewStream(ctx, &Rpc_ServiceDesc.Streams[0], Rpc_SyncProject_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcSyncProjectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Rpc_SyncProjectClient interface {
	Recv() (*Project, error)
	grpc.ClientStream
}

type rpcSyncProjectClient struct {
	grpc.ClientStream
}

func (x *rpcSyncProjectClient) Recv() (*Project, error) {
	m := new(Project)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rpcClient) QueryProject(ctx context.Context, in *QueryProjectReq, opts ...grpc.CallOption) (*ProjectSet, error) {
	out := new(ProjectSet)
	err := c.cc.Invoke(ctx, Rpc_QueryProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServer is the server API for Rpc service.
// All implementations must embed UnimplementedRpcServer
// for forward compatibility
type RpcServer interface {
	// for rancher
	CreateTable(context.Context, *Empty) (*Empty, error)
	SyncProject(*Empty, Rpc_SyncProjectServer) error
	QueryProject(context.Context, *QueryProjectReq) (*ProjectSet, error)
	mustEmbedUnimplementedRpcServer()
}

// UnimplementedRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServer struct {
}

func (UnimplementedRpcServer) CreateTable(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTable not implemented")
}
func (UnimplementedRpcServer) SyncProject(*Empty, Rpc_SyncProjectServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncProject not implemented")
}
func (UnimplementedRpcServer) QueryProject(context.Context, *QueryProjectReq) (*ProjectSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryProject not implemented")
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

func _Rpc_SyncProject_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RpcServer).SyncProject(m, &rpcSyncProjectServer{stream})
}

type Rpc_SyncProjectServer interface {
	Send(*Project) error
	grpc.ServerStream
}

type rpcSyncProjectServer struct {
	grpc.ServerStream
}

func (x *rpcSyncProjectServer) Send(m *Project) error {
	return x.ServerStream.SendMsg(m)
}

func _Rpc_QueryProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).QueryProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_QueryProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).QueryProject(ctx, req.(*QueryProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "K8sGenie.resourcer.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTable",
			Handler:    _Rpc_CreateTable_Handler,
		},
		{
			MethodName: "QueryProject",
			Handler:    _Rpc_QueryProject_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SyncProject",
			Handler:       _Rpc_SyncProject_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "apps/rancher/pb/rpc.proto",
}
