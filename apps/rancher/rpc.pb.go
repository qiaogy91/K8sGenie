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
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DESC_TYPE int32

const (
	DESC_TYPE_DESC_TYPE_UNSPECIFIED  DESC_TYPE = 0
	DESC_TYPE_DESC_TYPE_PROJECT_ID   DESC_TYPE = 1
	DESC_TYPE_DESC_TYPE_PROJECT_CODE DESC_TYPE = 2
	DESC_TYPE_DESC_TYPE_PROJECT_DESC DESC_TYPE = 3
)

// Enum value maps for DESC_TYPE.
var (
	DESC_TYPE_name = map[int32]string{
		0: "DESC_TYPE_UNSPECIFIED",
		1: "DESC_TYPE_PROJECT_ID",
		2: "DESC_TYPE_PROJECT_CODE",
		3: "DESC_TYPE_PROJECT_DESC",
	}
	DESC_TYPE_value = map[string]int32{
		"DESC_TYPE_UNSPECIFIED":  0,
		"DESC_TYPE_PROJECT_ID":   1,
		"DESC_TYPE_PROJECT_CODE": 2,
		"DESC_TYPE_PROJECT_DESC": 3,
	}
)

func (x DESC_TYPE) Enum() *DESC_TYPE {
	p := new(DESC_TYPE)
	*p = x
	return p
}

func (x DESC_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DESC_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_rancher_pb_rpc_proto_enumTypes[0].Descriptor()
}

func (DESC_TYPE) Type() protoreflect.EnumType {
	return &file_apps_rancher_pb_rpc_proto_enumTypes[0]
}

func (x DESC_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DESC_TYPE.Descriptor instead.
func (DESC_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_rancher_pb_rpc_proto_rawDescGZIP(), []int{0}
}

type QUERY_TYPE int32

const (
	QUERY_TYPE_QUERY_TYPE_UNSPECIFIED  QUERY_TYPE = 0
	QUERY_TYPE_QUERY_TYPE_ALL          QUERY_TYPE = 1
	QUERY_TYPE_QUERY_TYPE_CLUSTER_CODE QUERY_TYPE = 2
	QUERY_TYPE_QUERY_TYPE_PROJECT_LINE QUERY_TYPE = 3
	QUERY_TYPE_QUERY_TYPE_PROJECT_CODE QUERY_TYPE = 4
	QUERY_TYPE_QUERY_TYPE_PROJECT_DESC QUERY_TYPE = 5
)

// Enum value maps for QUERY_TYPE.
var (
	QUERY_TYPE_name = map[int32]string{
		0: "QUERY_TYPE_UNSPECIFIED",
		1: "QUERY_TYPE_ALL",
		2: "QUERY_TYPE_CLUSTER_CODE",
		3: "QUERY_TYPE_PROJECT_LINE",
		4: "QUERY_TYPE_PROJECT_CODE",
		5: "QUERY_TYPE_PROJECT_DESC",
	}
	QUERY_TYPE_value = map[string]int32{
		"QUERY_TYPE_UNSPECIFIED":  0,
		"QUERY_TYPE_ALL":          1,
		"QUERY_TYPE_CLUSTER_CODE": 2,
		"QUERY_TYPE_PROJECT_LINE": 3,
		"QUERY_TYPE_PROJECT_CODE": 4,
		"QUERY_TYPE_PROJECT_DESC": 5,
	}
)

func (x QUERY_TYPE) Enum() *QUERY_TYPE {
	p := new(QUERY_TYPE)
	*p = x
	return p
}

func (x QUERY_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QUERY_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_rancher_pb_rpc_proto_enumTypes[1].Descriptor()
}

func (QUERY_TYPE) Type() protoreflect.EnumType {
	return &file_apps_rancher_pb_rpc_proto_enumTypes[1]
}

func (x QUERY_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QUERY_TYPE.Descriptor instead.
func (QUERY_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_rancher_pb_rpc_proto_rawDescGZIP(), []int{1}
}

type DescProjectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DescType DESC_TYPE `protobuf:"varint,1,opt,name=desc_type,json=descType,proto3,enum=K8sGenie.resourcer.DESC_TYPE" json:"desc_type,omitempty"`
	KeyWord  string    `protobuf:"bytes,2,opt,name=key_word,json=keyWord,proto3" json:"key_word,omitempty"`
}

func (x *DescProjectReq) Reset() {
	*x = DescProjectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_rancher_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescProjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescProjectReq) ProtoMessage() {}

func (x *DescProjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_rancher_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescProjectReq.ProtoReflect.Descriptor instead.
func (*DescProjectReq) Descriptor() ([]byte, []int) {
	return file_apps_rancher_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *DescProjectReq) GetDescType() DESC_TYPE {
	if x != nil {
		return x.DescType
	}
	return DESC_TYPE_DESC_TYPE_UNSPECIFIED
}

func (x *DescProjectReq) GetKeyWord() string {
	if x != nil {
		return x.KeyWord
	}
	return ""
}

type QueryProjectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueryType QUERY_TYPE `protobuf:"varint,1,opt,name=query_type,json=queryType,proto3,enum=K8sGenie.resourcer.QUERY_TYPE" json:"query_type,omitempty"`
	KeyWord   string     `protobuf:"bytes,2,opt,name=key_word,json=keyWord,proto3" json:"key_word,omitempty"`
}

func (x *QueryProjectReq) Reset() {
	*x = QueryProjectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_rancher_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryProjectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryProjectReq) ProtoMessage() {}

func (x *QueryProjectReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_rancher_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryProjectReq.ProtoReflect.Descriptor instead.
func (*QueryProjectReq) Descriptor() ([]byte, []int) {
	return file_apps_rancher_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *QueryProjectReq) GetQueryType() QUERY_TYPE {
	if x != nil {
		return x.QueryType
	}
	return QUERY_TYPE_QUERY_TYPE_UNSPECIFIED
}

func (x *QueryProjectReq) GetKeyWord() string {
	if x != nil {
		return x.KeyWord
	}
	return ""
}

var File_apps_rancher_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_rancher_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x70,
	0x62, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x4b, 0x38, 0x73,
	0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x1a,
	0x1b, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x62,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a, 0x0e,
	0x44, 0x65, 0x73, 0x63, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x3a,
	0x0a, 0x09, 0x64, 0x65, 0x73, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1d, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x44, 0x45, 0x53, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x52, 0x08, 0x64, 0x65, 0x73, 0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65,
	0x79, 0x5f, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65,
	0x79, 0x57, 0x6f, 0x72, 0x64, 0x22, 0x6b, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x3d, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x4b,
	0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x72, 0x2e, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x09, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x57, 0x6f,
	0x72, 0x64, 0x2a, 0x78, 0x0a, 0x09, 0x44, 0x45, 0x53, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12,
	0x19, 0x0a, 0x15, 0x44, 0x45, 0x53, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x44, 0x45,
	0x53, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f,
	0x49, 0x44, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x44, 0x45, 0x53, 0x43, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x02,
	0x12, 0x1a, 0x0a, 0x16, 0x44, 0x45, 0x53, 0x43, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52,
	0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x03, 0x2a, 0xb0, 0x01, 0x0a,
	0x0a, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x1a, 0x0a, 0x16, 0x51,
	0x55, 0x45, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x51, 0x55, 0x45, 0x52, 0x59,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x51,
	0x55, 0x45, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x55, 0x53, 0x54, 0x45,
	0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x51, 0x55, 0x45, 0x52,
	0x59, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4c,
	0x49, 0x4e, 0x45, 0x10, 0x03, 0x12, 0x1b, 0x0a, 0x17, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x10, 0x04, 0x12, 0x1b, 0x0a, 0x17, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x05, 0x32,
	0xb8, 0x02, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12, 0x43, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x19, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69,
	0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x19, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x47, 0x0a, 0x0b,
	0x53, 0x79, 0x6e, 0x63, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x19, 0x2e, 0x4b, 0x38,
	0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1b, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69,
	0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65,
	0x6e, 0x69, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x53, 0x0a, 0x0c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x23, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x4b, 0x38, 0x73,
	0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69,
	0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x61, 0x6f, 0x67, 0x79, 0x39, 0x31,
	0x2f, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_rancher_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_rancher_pb_rpc_proto_rawDescData = file_apps_rancher_pb_rpc_proto_rawDesc
)

func file_apps_rancher_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_rancher_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_rancher_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_rancher_pb_rpc_proto_rawDescData)
	})
	return file_apps_rancher_pb_rpc_proto_rawDescData
}

var file_apps_rancher_pb_rpc_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_apps_rancher_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_apps_rancher_pb_rpc_proto_goTypes = []any{
	(DESC_TYPE)(0),          // 0: K8sGenie.resourcer.DESC_TYPE
	(QUERY_TYPE)(0),         // 1: K8sGenie.resourcer.QUERY_TYPE
	(*DescProjectReq)(nil),  // 2: K8sGenie.resourcer.DescProjectReq
	(*QueryProjectReq)(nil), // 3: K8sGenie.resourcer.QueryProjectReq
	(*Empty)(nil),           // 4: K8sGenie.resourcer.Empty
	(*Project)(nil),         // 5: K8sGenie.resourcer.Project
	(*ProjectSet)(nil),      // 6: K8sGenie.resourcer.ProjectSet
}
var file_apps_rancher_pb_rpc_proto_depIdxs = []int32{
	0, // 0: K8sGenie.resourcer.DescProjectReq.desc_type:type_name -> K8sGenie.resourcer.DESC_TYPE
	1, // 1: K8sGenie.resourcer.QueryProjectReq.query_type:type_name -> K8sGenie.resourcer.QUERY_TYPE
	4, // 2: K8sGenie.resourcer.Rpc.CreateTable:input_type -> K8sGenie.resourcer.Empty
	4, // 3: K8sGenie.resourcer.Rpc.SyncProject:input_type -> K8sGenie.resourcer.Empty
	2, // 4: K8sGenie.resourcer.Rpc.DescProject:input_type -> K8sGenie.resourcer.DescProjectReq
	3, // 5: K8sGenie.resourcer.Rpc.QueryProject:input_type -> K8sGenie.resourcer.QueryProjectReq
	4, // 6: K8sGenie.resourcer.Rpc.CreateTable:output_type -> K8sGenie.resourcer.Empty
	5, // 7: K8sGenie.resourcer.Rpc.SyncProject:output_type -> K8sGenie.resourcer.Project
	5, // 8: K8sGenie.resourcer.Rpc.DescProject:output_type -> K8sGenie.resourcer.Project
	6, // 9: K8sGenie.resourcer.Rpc.QueryProject:output_type -> K8sGenie.resourcer.ProjectSet
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_apps_rancher_pb_rpc_proto_init() }
func file_apps_rancher_pb_rpc_proto_init() {
	if File_apps_rancher_pb_rpc_proto != nil {
		return
	}
	file_apps_rancher_pb_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_apps_rancher_pb_rpc_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DescProjectReq); i {
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
		file_apps_rancher_pb_rpc_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*QueryProjectReq); i {
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
			RawDescriptor: file_apps_rancher_pb_rpc_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_rancher_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_rancher_pb_rpc_proto_depIdxs,
		EnumInfos:         file_apps_rancher_pb_rpc_proto_enumTypes,
		MessageInfos:      file_apps_rancher_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_rancher_pb_rpc_proto = out.File
	file_apps_rancher_pb_rpc_proto_rawDesc = nil
	file_apps_rancher_pb_rpc_proto_goTypes = nil
	file_apps_rancher_pb_rpc_proto_depIdxs = nil
}