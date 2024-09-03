// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: apps/router/pb/model.proto

package router

import (
	rancher "gitee.com/qiaogy91/K8sGenie/apps/rancher"
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

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @gotags: gorm:"autoCreateTime"
	CreatedAt int64 `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" gorm:"autoCreateTime"`
	// @gotags: gorm:"autoUpdateTime"
	UpdatedAt int64 `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" gorm:"autoUpdateTime"`
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

	// @gotags: validate:"required" gorm:"unique"
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty" validate:"required" gorm:"unique"`
	// @gotags: validate:"required"
	WebhookUrl string `protobuf:"bytes,2,opt,name=webhook_url,json=webhookUrl,proto3" json:"webhook_url,omitempty" validate:"required"`
	// @gotags: validate:"required"
	WebhookToken string `protobuf:"bytes,3,opt,name=webhook_token,json=webhookToken,proto3" json:"webhook_token,omitempty" validate:"required"`
	// @gotags: gorm:"-"
	ProjectSpec *rancher.Spec `protobuf:"bytes,4,opt,name=project_spec,json=projectSpec,proto3" json:"project_spec,omitempty" gorm:"-"`
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

func (x *Spec) GetIdentity() string {
	if x != nil {
		return x.Identity
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

func (x *Spec) GetProjectSpec() *rancher.Spec {
	if x != nil {
		return x.ProjectSpec
	}
	return nil
}

type Router struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: json:"meta" gorm:"embedded"
	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta" gorm:"embedded"`
	// @gotags: json:"spec" gorm:"embedded"
	Spec *Spec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec" gorm:"embedded"`
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
	0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x1a, 0x1b, 0x61,
	0x70, 0x70, 0x73, 0x2f, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x04, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0xa5, 0x01, 0x0a, 0x04, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x3b, 0x0a, 0x0c, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x53, 0x70, 0x65, 0x63, 0x22, 0x5e, 0x0a, 0x06, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x29, 0x0a,
	0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4b, 0x38,
	0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x65,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x61, 0x6f, 0x67, 0x79, 0x39, 0x31, 0x2f, 0x4b,
	0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_apps_router_pb_model_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_apps_router_pb_model_proto_goTypes = []any{
	(*Meta)(nil),         // 0: K8sGenie.router.Meta
	(*Spec)(nil),         // 1: K8sGenie.router.Spec
	(*Router)(nil),       // 2: K8sGenie.router.Router
	(*rancher.Spec)(nil), // 3: K8sGenie.resourcer.Spec
}
var file_apps_router_pb_model_proto_depIdxs = []int32{
	3, // 0: K8sGenie.router.Spec.project_spec:type_name -> K8sGenie.resourcer.Spec
	0, // 1: K8sGenie.router.Router.meta:type_name -> K8sGenie.router.Meta
	1, // 2: K8sGenie.router.Router.spec:type_name -> K8sGenie.router.Spec
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
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_apps_router_pb_model_proto_goTypes,
		DependencyIndexes: file_apps_router_pb_model_proto_depIdxs,
		MessageInfos:      file_apps_router_pb_model_proto_msgTypes,
	}.Build()
	File_apps_router_pb_model_proto = out.File
	file_apps_router_pb_model_proto_rawDesc = nil
	file_apps_router_pb_model_proto_goTypes = nil
	file_apps_router_pb_model_proto_depIdxs = nil
}