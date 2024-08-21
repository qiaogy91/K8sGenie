// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: apps/router/pb/rpc.proto

package router

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_router_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_apps_router_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_apps_router_pb_rpc_proto_rawDescGZIP(), []int{0}
}

type DeleteRouteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validate:"required"`
}

func (x *DeleteRouteReq) Reset() {
	*x = DeleteRouteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_router_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRouteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRouteReq) ProtoMessage() {}

func (x *DeleteRouteReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_router_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRouteReq.ProtoReflect.Descriptor instead.
func (*DeleteRouteReq) Descriptor() ([]byte, []int) {
	return file_apps_router_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteRouteReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AlertRouteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RobotName     string `protobuf:"bytes,1,opt,name=robot_name,json=robotName,proto3" json:"robot_name,omitempty"`
	NamespaceName string `protobuf:"bytes,2,opt,name=namespace_name,json=namespaceName,proto3" json:"namespace_name,omitempty"`
}

func (x *AlertRouteReq) Reset() {
	*x = AlertRouteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_router_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlertRouteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlertRouteReq) ProtoMessage() {}

func (x *AlertRouteReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_router_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlertRouteReq.ProtoReflect.Descriptor instead.
func (*AlertRouteReq) Descriptor() ([]byte, []int) {
	return file_apps_router_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *AlertRouteReq) GetRobotName() string {
	if x != nil {
		return x.RobotName
	}
	return ""
}

func (x *AlertRouteReq) GetNamespaceName() string {
	if x != nil {
		return x.NamespaceName
	}
	return ""
}

var File_apps_router_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_router_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x62,
	0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x4b, 0x38, 0x73, 0x47,
	0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x1a, 0x1a, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x55, 0x0a, 0x0d, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x32, 0x93, 0x02, 0x0a, 0x03, 0x52, 0x70,
	0x63, 0x12, 0x3d, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x12, 0x16, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65,
	0x6e, 0x69, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x3d, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12,
	0x15, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x17, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69,
	0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12,
	0x47, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1f,
	0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a,
	0x17, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x0a, 0x41, 0x6c, 0x65, 0x72,
	0x74, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69,
	0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69,
	0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x42,
	0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x61,
	0x6f, 0x67, 0x79, 0x39, 0x31, 0x2f, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2f, 0x61,
	0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_apps_router_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_router_pb_rpc_proto_rawDescData = file_apps_router_pb_rpc_proto_rawDesc
)

func file_apps_router_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_router_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_router_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_router_pb_rpc_proto_rawDescData)
	})
	return file_apps_router_pb_rpc_proto_rawDescData
}

var file_apps_router_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_router_pb_rpc_proto_goTypes = []any{
	(*Empty)(nil),          // 0: K8sGenie.router.Empty
	(*DeleteRouteReq)(nil), // 1: K8sGenie.router.DeleteRouteReq
	(*AlertRouteReq)(nil),  // 2: K8sGenie.router.AlertRouteReq
	(*Spec)(nil),           // 3: K8sGenie.router.Spec
	(*Router)(nil),         // 4: K8sGenie.router.Router
}
var file_apps_router_pb_rpc_proto_depIdxs = []int32{
	0, // 0: K8sGenie.router.Rpc.CreateTable:input_type -> K8sGenie.router.Empty
	3, // 1: K8sGenie.router.Rpc.CreateRoute:input_type -> K8sGenie.router.Spec
	1, // 2: K8sGenie.router.Rpc.DeleteRoute:input_type -> K8sGenie.router.DeleteRouteReq
	2, // 3: K8sGenie.router.Rpc.AlertRoute:input_type -> K8sGenie.router.AlertRouteReq
	0, // 4: K8sGenie.router.Rpc.CreateTable:output_type -> K8sGenie.router.Empty
	4, // 5: K8sGenie.router.Rpc.CreateRoute:output_type -> K8sGenie.router.Router
	4, // 6: K8sGenie.router.Rpc.DeleteRoute:output_type -> K8sGenie.router.Router
	4, // 7: K8sGenie.router.Rpc.AlertRoute:output_type -> K8sGenie.router.Router
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apps_router_pb_rpc_proto_init() }
func file_apps_router_pb_rpc_proto_init() {
	if File_apps_router_pb_rpc_proto != nil {
		return
	}
	file_apps_router_pb_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_router_pb_rpc_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_router_pb_rpc_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRouteReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_apps_router_pb_rpc_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AlertRouteReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_router_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_router_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_router_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_router_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_router_pb_rpc_proto = out.File
	file_apps_router_pb_rpc_proto_rawDesc = nil
	file_apps_router_pb_rpc_proto_goTypes = nil
	file_apps_router_pb_rpc_proto_depIdxs = nil
}
