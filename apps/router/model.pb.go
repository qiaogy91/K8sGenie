// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: apps/router/pb/model.proto

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

type TYPE int32

const (
	TYPE_TYPE_UNSPECIFIED   TYPE = 0
	TYPE_TYPE_CLUSTER_SCOPE TYPE = 1
	TYPE_TYPE_PROJECT_SCOPE TYPE = 2
)

// Enum value maps for TYPE.
var (
	TYPE_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_CLUSTER_SCOPE",
		2: "TYPE_PROJECT_SCOPE",
	}
	TYPE_value = map[string]int32{
		"TYPE_UNSPECIFIED":   0,
		"TYPE_CLUSTER_SCOPE": 1,
		"TYPE_PROJECT_SCOPE": 2,
	}
)

func (x TYPE) Enum() *TYPE {
	p := new(TYPE)
	*p = x
	return p
}

func (x TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_apps_router_pb_model_proto_enumTypes[0].Descriptor()
}

func (TYPE) Type() protoreflect.EnumType {
	return &file_apps_router_pb_model_proto_enumTypes[0]
}

func (x TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TYPE.Descriptor instead.
func (TYPE) EnumDescriptor() ([]byte, []int) {
	return file_apps_router_pb_model_proto_rawDescGZIP(), []int{0}
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt int64 `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64 `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_router_pb_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_apps_router_pb_model_proto_msgTypes[0]
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
	return file_apps_router_pb_model_proto_rawDescGZIP(), []int{0}
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

	NotifyType   TYPE   `protobuf:"varint,1,opt,name=notify_type,json=notifyType,proto3,enum=K8sGenie.router.TYPE" json:"notify_type,omitempty"`
	RobotName    string `protobuf:"bytes,2,opt,name=robot_name,json=robotName,proto3" json:"robot_name,omitempty"`
	ProjectId    string `protobuf:"bytes,3,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	WebhookUrl   string `protobuf:"bytes,4,opt,name=webhook_url,json=webhookUrl,proto3" json:"webhook_url,omitempty"`
	WebhookToken string `protobuf:"bytes,5,opt,name=webhook_token,json=webhookToken,proto3" json:"webhook_token,omitempty"`
}

func (x *Spec) Reset() {
	*x = Spec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_router_pb_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Spec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Spec) ProtoMessage() {}

func (x *Spec) ProtoReflect() protoreflect.Message {
	mi := &file_apps_router_pb_model_proto_msgTypes[1]
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
	return file_apps_router_pb_model_proto_rawDescGZIP(), []int{1}
}

func (x *Spec) GetNotifyType() TYPE {
	if x != nil {
		return x.NotifyType
	}
	return TYPE_TYPE_UNSPECIFIED
}

func (x *Spec) GetRobotName() string {
	if x != nil {
		return x.RobotName
	}
	return ""
}

func (x *Spec) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Spec) GetWebhookUrl() string {
	if x != nil {
		return x.WebhookUrl
	}
	return ""
}

func (x *Spec) GetWebhookToken() string {
	if x != nil {
		return x.WebhookToken
	}
	return ""
}

type Router struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Spec *Spec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *Router) Reset() {
	*x = Router{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_router_pb_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Router) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Router) ProtoMessage() {}

func (x *Router) ProtoReflect() protoreflect.Message {
	mi := &file_apps_router_pb_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Router.ProtoReflect.Descriptor instead.
func (*Router) Descriptor() ([]byte, []int) {
	return file_apps_router_pb_model_proto_rawDescGZIP(), []int{2}
}

func (x *Router) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Router) GetSpec() *Spec {
	if x != nil {
		return x.Spec
	}
	return nil
}

var File_apps_router_pb_model_proto protoreflect.FileDescriptor

var file_apps_router_pb_model_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2f, 0x70, 0x62,
	0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x4b, 0x38,
	0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x22, 0x54, 0x0a,
	0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0xc2, 0x01, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x36, 0x0a, 0x0b,
	0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x2e, 0x54, 0x59, 0x50, 0x45, 0x52, 0x0a, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5e, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x29, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4b, 0x38,
	0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x2a, 0x4c, 0x0a, 0x04, 0x54, 0x59, 0x50, 0x45,
	0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43,
	0x4c, 0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x43, 0x4f, 0x50, 0x45, 0x10, 0x01, 0x12, 0x16,
	0x0a, 0x12, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x53,
	0x43, 0x4f, 0x50, 0x45, 0x10, 0x02, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x61, 0x6f, 0x67, 0x79, 0x39, 0x31, 0x2f, 0x4b, 0x38, 0x73,
	0x47, 0x65, 0x6e, 0x69, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_router_pb_model_proto_rawDescOnce sync.Once
	file_apps_router_pb_model_proto_rawDescData = file_apps_router_pb_model_proto_rawDesc
)

func file_apps_router_pb_model_proto_rawDescGZIP() []byte {
	file_apps_router_pb_model_proto_rawDescOnce.Do(func() {
		file_apps_router_pb_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_router_pb_model_proto_rawDescData)
	})
	return file_apps_router_pb_model_proto_rawDescData
}

var file_apps_router_pb_model_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_apps_router_pb_model_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_router_pb_model_proto_goTypes = []any{
	(TYPE)(0),      // 0: K8sGenie.router.TYPE
	(*Meta)(nil),   // 1: K8sGenie.router.Meta
	(*Spec)(nil),   // 2: K8sGenie.router.Spec
	(*Router)(nil), // 3: K8sGenie.router.Router
}
var file_apps_router_pb_model_proto_depIdxs = []int32{
	0, // 0: K8sGenie.router.Spec.notify_type:type_name -> K8sGenie.router.TYPE
	1, // 1: K8sGenie.router.Router.meta:type_name -> K8sGenie.router.Meta
	2, // 2: K8sGenie.router.Router.spec:type_name -> K8sGenie.router.Spec
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_apps_router_pb_model_proto_init() }
func file_apps_router_pb_model_proto_init() {
	if File_apps_router_pb_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_router_pb_model_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_apps_router_pb_model_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_apps_router_pb_model_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Router); i {
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
			RawDescriptor: file_apps_router_pb_model_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_router_pb_model_proto_goTypes,
		DependencyIndexes: file_apps_router_pb_model_proto_depIdxs,
		EnumInfos:         file_apps_router_pb_model_proto_enumTypes,
		MessageInfos:      file_apps_router_pb_model_proto_msgTypes,
	}.Build()
	File_apps_router_pb_model_proto = out.File
	file_apps_router_pb_model_proto_rawDesc = nil
	file_apps_router_pb_model_proto_goTypes = nil
	file_apps_router_pb_model_proto_depIdxs = nil
}
