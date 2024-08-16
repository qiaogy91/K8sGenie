// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: apps/k8s/pb/model.proto

package k8s

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

type Type int32

const (
	Type_TYPE_UNSPECIFIED  Type = 0
	Type_TYPE_DEPLOYMENT   Type = 1
	Type_TYPE_CRON_JOB     Type = 2
	Type_TYPE_STATEFUL_SET Type = 3
	Type_TYPE_DAEMON_SET   Type = 4
)

// Enum value maps for Type.
var (
	Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_DEPLOYMENT",
		2: "TYPE_CRON_JOB",
		3: "TYPE_STATEFUL_SET",
		4: "TYPE_DAEMON_SET",
	}
	Type_value = map[string]int32{
		"TYPE_UNSPECIFIED":  0,
		"TYPE_DEPLOYMENT":   1,
		"TYPE_CRON_JOB":     2,
		"TYPE_STATEFUL_SET": 3,
		"TYPE_DAEMON_SET":   4,
	}
)

func (x Type) Enum() *Type {
	p := new(Type)
	*p = x
	return p
}

func (x Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Type) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_k8s_pb_model_proto_enumTypes[0].Descriptor()
}

func (Type) Type() protoreflect.EnumType {
	return &file_apps_k8s_pb_model_proto_enumTypes[0]
}

func (x Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Type.Descriptor instead.
func (Type) EnumDescriptor() ([]byte, []int) {
	return file_apps_k8s_pb_model_proto_rawDescGZIP(), []int{0}
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_k8s_pb_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_apps_k8s_pb_model_proto_msgTypes[0]
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
	return file_apps_k8s_pb_model_proto_rawDescGZIP(), []int{0}
}

type WorkLoad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"meta" gorm:"embedded"
	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" gorm:"embedded"`
	// @gotags: json:"spec" gorm:"embedded"
	Spec *Spec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" gorm:"embedded"`
}

func (x *WorkLoad) Reset() {
	*x = WorkLoad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_k8s_pb_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkLoad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkLoad) ProtoMessage() {}

func (x *WorkLoad) ProtoReflect() protoreflect.Message {
	mi := &file_apps_k8s_pb_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkLoad.ProtoReflect.Descriptor instead.
func (*WorkLoad) Descriptor() ([]byte, []int) {
	return file_apps_k8s_pb_model_proto_rawDescGZIP(), []int{1}
}

func (x *WorkLoad) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *WorkLoad) GetSpec() *Spec {
	if x != nil {
		return x.Spec
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id"`
	// @gotags: gorm:"autoCreateTime"
	CreatedAt int64 `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" gorm:"autoCreateTime"`
	// @gotags: gorm:"autoUpdateTime"
	UpdatedAt int64 `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" gorm:"autoUpdateTime"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_k8s_pb_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_apps_k8s_pb_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_apps_k8s_pb_model_proto_rawDescGZIP(), []int{2}
}

func (x *Meta) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Meta) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Meta) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type Spec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"type"
	Type Type `protobuf:"varint,1,opt,name=type,proto3,enum=K8sGenie.k8s.Type" json:"type"`
	// @gotags: json:"namespace"
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace"`
	// @gotags: json:"name"
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name"`
	// @gotags: json:"replicas"
	Replicas int32 `protobuf:"varint,4,opt,name=replicas,proto3" json:"replicas"`
	// @gotags: json:"ram"
	Ram int64 `protobuf:"varint,5,opt,name=ram,proto3" json:"ram"`
	// @gotags: json:"cpu"
	Cpu int64 `protobuf:"varint,6,opt,name=cpu,proto3" json:"cpu"`
	// @gotags: json:"project_id"
	ProjectId string `protobuf:"bytes,7,opt,name=project_id,json=projectId,proto3" json:"project_id"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_k8s_pb_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_apps_k8s_pb_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Spec.ProtoReflect.Descriptor instead.
func (*Spec) Descriptor() ([]byte, []int) {
	return file_apps_k8s_pb_model_proto_rawDescGZIP(), []int{3}
}

func (x *Spec) GetType() Type {
	if x != nil {
		return x.Type
	}
	return Type_TYPE_UNSPECIFIED
}

func (x *Spec) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *Spec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Spec) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *Spec) GetRam() int64 {
	if x != nil {
		return x.Ram
	}
	return 0
}

func (x *Spec) GetCpu() int64 {
	if x != nil {
		return x.Cpu
	}
	return 0
}

func (x *Spec) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

var File_apps_k8s_pb_model_proto protoreflect.FileDescriptor

var file_apps_k8s_pb_model_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x4b, 0x38, 0x73, 0x47, 0x65,
	0x6e, 0x69, 0x65, 0x2e, 0x6b, 0x38, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x5a, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x4c, 0x6f, 0x61, 0x64, 0x12, 0x26, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4b, 0x38, 0x73,
	0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x6b, 0x38,
	0x73, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x54, 0x0a, 0x04,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0xbf, 0x01, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x26, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x4b, 0x38, 0x73, 0x47,
	0x65, 0x6e, 0x69, 0x65, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x73, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x2a, 0x70, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x45, 0x50, 0x4c, 0x4f,
	0x59, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x43, 0x52, 0x4f, 0x4e, 0x5f, 0x4a, 0x4f, 0x42, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x46, 0x55, 0x4c, 0x5f, 0x53, 0x45, 0x54, 0x10,
	0x03, 0x12, 0x13, 0x0a, 0x0f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x41, 0x45, 0x4d, 0x4f, 0x4e,
	0x5f, 0x53, 0x45, 0x54, 0x10, 0x04, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x61, 0x6f, 0x67, 0x79, 0x39, 0x31, 0x2f, 0x4b, 0x38, 0x73,
	0x47, 0x65, 0x6e, 0x69, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_k8s_pb_model_proto_rawDescOnce sync.Once
	file_apps_k8s_pb_model_proto_rawDescData = file_apps_k8s_pb_model_proto_rawDesc
)

func file_apps_k8s_pb_model_proto_rawDescGZIP() []byte {
	file_apps_k8s_pb_model_proto_rawDescOnce.Do(func() {
		file_apps_k8s_pb_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_k8s_pb_model_proto_rawDescData)
	})
	return file_apps_k8s_pb_model_proto_rawDescData
}

var file_apps_k8s_pb_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_k8s_pb_model_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_apps_k8s_pb_model_proto_goTypes = []any{
	(Type)(0),        // 0: K8sGenie.k8s.Type
	(*Empty)(nil),    // 1: K8sGenie.k8s.Empty
	(*WorkLoad)(nil), // 2: K8sGenie.k8s.WorkLoad
	(*Meta)(nil),     // 3: K8sGenie.k8s.Meta
	(*Spec)(nil),     // 4: K8sGenie.k8s.Spec
}
var file_apps_k8s_pb_model_proto_depIdxs = []int32{
	3, // 0: K8sGenie.k8s.WorkLoad.meta:type_name -> K8sGenie.k8s.Meta
	4, // 1: K8sGenie.k8s.WorkLoad.spec:type_name -> K8sGenie.k8s.Spec
	0, // 2: K8sGenie.k8s.Spec.type:type_name -> K8sGenie.k8s.Type
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_k8s_pb_model_proto_init() }
func file_apps_k8s_pb_model_proto_init() {
	if File_apps_k8s_pb_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_k8s_pb_model_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_apps_k8s_pb_model_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*WorkLoad); i {
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
		file_apps_k8s_pb_model_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Meta); i {
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
		file_apps_k8s_pb_model_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Spec); i {
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
			RawDescriptor: file_apps_k8s_pb_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_k8s_pb_model_proto_goTypes,
		DependencyIndexes: file_apps_k8s_pb_model_proto_depIdxs,
		EnumInfos:         file_apps_k8s_pb_model_proto_enumTypes,
		MessageInfos:      file_apps_k8s_pb_model_proto_msgTypes,
	}.Build()
	File_apps_k8s_pb_model_proto = out.File
	file_apps_k8s_pb_model_proto_rawDesc = nil
	file_apps_k8s_pb_model_proto_goTypes = nil
	file_apps_k8s_pb_model_proto_depIdxs = nil
}