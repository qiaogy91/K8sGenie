// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: apps/rancher/pb/rpc.proto

package rancher

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_apps_rancher_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_rancher_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x4b, 0x38, 0x73,
	0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x1a,
	0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x62,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9a, 0x01, 0x0a,
	0x03, 0x52, 0x70, 0x63, 0x12, 0x43, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19,
	0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4e, 0x0a, 0x12, 0x53, 0x79, 0x6e,
	0x63, 0x52, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12,
	0x19, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x4b, 0x38, 0x73,
	0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x30, 0x01, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x61, 0x6f, 0x67, 0x79, 0x39, 0x31, 0x2f,
	0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_apps_rancher_pb_rpc_proto_goTypes = []any{
	(*Empty)(nil),   // 0: K8sGenie.resourcer.Empty
	(*Project)(nil), // 1: K8sGenie.resourcer.Project
}
var file_apps_rancher_pb_rpc_proto_depIdxs = []int32{
	0, // 0: K8sGenie.resourcer.Rpc.CreateTable:input_type -> K8sGenie.resourcer.Empty
	0, // 1: K8sGenie.resourcer.Rpc.SyncRancherProject:input_type -> K8sGenie.resourcer.Empty
	0, // 2: K8sGenie.resourcer.Rpc.CreateTable:output_type -> K8sGenie.resourcer.Empty
	1, // 3: K8sGenie.resourcer.Rpc.SyncRancherProject:output_type -> K8sGenie.resourcer.Project
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_apps_rancher_pb_rpc_proto_init() }
func file_apps_rancher_pb_rpc_proto_init() {
	if File_apps_rancher_pb_rpc_proto != nil {
		return
	}
	file_apps_rancher_pb_model_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_apps_rancher_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_rancher_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_rancher_pb_rpc_proto_depIdxs,
	}.Build()
	File_apps_rancher_pb_rpc_proto = out.File
	file_apps_rancher_pb_rpc_proto_rawDesc = nil
	file_apps_rancher_pb_rpc_proto_goTypes = nil
	file_apps_rancher_pb_rpc_proto_depIdxs = nil
}
