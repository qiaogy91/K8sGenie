// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.27.1
// source: apps/k8s/pb/rpc.proto

package k8s

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
	Rpc_TableCreate_FullMethodName     = "/K8sGenie.k8s.Rpc/TableCreate"
	Rpc_SyncK8SWorkload_FullMethodName = "/K8sGenie.k8s.Rpc/SyncK8sWorkload"
)

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	TableCreate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	SyncK8SWorkload(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Rpc_SyncK8SWorkloadClient, error)
}

type rpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcClient(cc grpc.ClientConnInterface) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) TableCreate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Rpc_TableCreate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) SyncK8SWorkload(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Rpc_SyncK8SWorkloadClient, error) {
	stream, err := c.cc.NewStream(ctx, &Rpc_ServiceDesc.Streams[0], Rpc_SyncK8SWorkload_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &rpcSyncK8SWorkloadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Rpc_SyncK8SWorkloadClient interface {
	Recv() (*WorkLoad, error)
	grpc.ClientStream
}

type rpcSyncK8SWorkloadClient struct {
	grpc.ClientStream
}

func (x *rpcSyncK8SWorkloadClient) Recv() (*WorkLoad, error) {
	m := new(WorkLoad)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RpcServer is the server API for Rpc service.
// All implementations must embed UnimplementedRpcServer
// for forward compatibility
type RpcServer interface {
	TableCreate(context.Context, *Empty) (*Empty, error)
	SyncK8SWorkload(*Empty, Rpc_SyncK8SWorkloadServer) error
	mustEmbedUnimplementedRpcServer()
}

// UnimplementedRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServer struct {
}

func (UnimplementedRpcServer) TableCreate(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TableCreate not implemented")
}
func (UnimplementedRpcServer) SyncK8SWorkload(*Empty, Rpc_SyncK8SWorkloadServer) error {
	return status.Errorf(codes.Unimplemented, "method SyncK8SWorkload not implemented")
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

func _Rpc_TableCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).TableCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_TableCreate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).TableCreate(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_SyncK8SWorkload_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RpcServer).SyncK8SWorkload(m, &rpcSyncK8SWorkloadServer{stream})
}

type Rpc_SyncK8SWorkloadServer interface {
	Send(*WorkLoad) error
	grpc.ServerStream
}

type rpcSyncK8SWorkloadServer struct {
	grpc.ServerStream
}

func (x *rpcSyncK8SWorkloadServer) Send(m *WorkLoad) error {
	return x.ServerStream.SendMsg(m)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "K8sGenie.k8s.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TableCreate",
			Handler:    _Rpc_TableCreate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SyncK8sWorkload",
			Handler:       _Rpc_SyncK8SWorkload_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "apps/k8s/pb/rpc.proto",
}