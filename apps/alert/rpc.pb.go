// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: apps/alert/pb/rpc.proto

package alert

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

type HandlerAlertReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Alerts []*Alert `protobuf:"bytes,2,rep,name=alerts,proto3" json:"alerts,omitempty"`
}

func (x *HandlerAlertReq) Reset() {
	*x = HandlerAlertReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_alert_pb_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandlerAlertReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandlerAlertReq) ProtoMessage() {}

func (x *HandlerAlertReq) ProtoReflect() protoreflect.Message {
	mi := &file_apps_alert_pb_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandlerAlertReq.ProtoReflect.Descriptor instead.
func (*HandlerAlertReq) Descriptor() ([]byte, []int) {
	return file_apps_alert_pb_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *HandlerAlertReq) GetAlerts() []*Alert {
	if x != nil {
		return x.Alerts
	}
	return nil
}

type Alert struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status      string            `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	StartsAt    string            `protobuf:"bytes,2,opt,name=starts_at,json=startsAt,proto3" json:"starts_at,omitempty"`
	EndsAt      string            `protobuf:"bytes,3,opt,name=ends_at,json=endsAt,proto3" json:"ends_at,omitempty"`
	Labels      map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Annotations map[string]string `protobuf:"bytes,5,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Alert) Reset() {
	*x = Alert{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_alert_pb_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alert) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alert) ProtoMessage() {}

func (x *Alert) ProtoReflect() protoreflect.Message {
	mi := &file_apps_alert_pb_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alert.ProtoReflect.Descriptor instead.
func (*Alert) Descriptor() ([]byte, []int) {
	return file_apps_alert_pb_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *Alert) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Alert) GetStartsAt() string {
	if x != nil {
		return x.StartsAt
	}
	return ""
}

func (x *Alert) GetEndsAt() string {
	if x != nil {
		return x.EndsAt
	}
	return ""
}

func (x *Alert) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Alert) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type HandlerAlertRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rsp []*Response `protobuf:"bytes,1,rep,name=rsp,proto3" json:"rsp,omitempty"`
}

func (x *HandlerAlertRsp) Reset() {
	*x = HandlerAlertRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_alert_pb_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HandlerAlertRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HandlerAlertRsp) ProtoMessage() {}

func (x *HandlerAlertRsp) ProtoReflect() protoreflect.Message {
	mi := &file_apps_alert_pb_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HandlerAlertRsp.ProtoReflect.Descriptor instead.
func (*HandlerAlertRsp) Descriptor() ([]byte, []int) {
	return file_apps_alert_pb_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *HandlerAlertRsp) GetRsp() []*Response {
	if x != nil {
		return x.Rsp
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_apps_alert_pb_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_apps_alert_pb_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_apps_alert_pb_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_apps_alert_pb_rpc_proto protoreflect.FileDescriptor

var file_apps_alert_pb_rpc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2f, 0x70, 0x62, 0x2f,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x4b, 0x38, 0x73, 0x47, 0x65,
	0x6e, 0x69, 0x65, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x22, 0x40, 0x0a, 0x0f, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2d, 0x0a, 0x06,
	0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x4b,
	0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x52, 0x06, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x73, 0x22, 0xd5, 0x02, 0x0a, 0x05,
	0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73, 0x41, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x65, 0x6e,
	0x64, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x64,
	0x73, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x48,
	0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x2e, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x3d, 0x0a, 0x0f, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x41, 0x6c,
	0x65, 0x72, 0x74, 0x52, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x03, 0x72, 0x73, 0x70, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x61,
	0x6c, 0x65, 0x72, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x03, 0x72,
	0x73, 0x70, 0x22, 0x30, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x32, 0x57, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12, 0x50, 0x0a, 0x0c, 0x48,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x12, 0x1f, 0x2e, 0x4b, 0x38,
	0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x4b,
	0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2e, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x2e, 0x48, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x72, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x52, 0x73, 0x70, 0x42, 0x28, 0x5a,
	0x26, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x71, 0x69, 0x61, 0x6f, 0x67,
	0x79, 0x39, 0x31, 0x2f, 0x4b, 0x38, 0x73, 0x47, 0x65, 0x6e, 0x69, 0x65, 0x2f, 0x61, 0x70, 0x70,
	0x73, 0x2f, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_apps_alert_pb_rpc_proto_rawDescOnce sync.Once
	file_apps_alert_pb_rpc_proto_rawDescData = file_apps_alert_pb_rpc_proto_rawDesc
)

func file_apps_alert_pb_rpc_proto_rawDescGZIP() []byte {
	file_apps_alert_pb_rpc_proto_rawDescOnce.Do(func() {
		file_apps_alert_pb_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_apps_alert_pb_rpc_proto_rawDescData)
	})
	return file_apps_alert_pb_rpc_proto_rawDescData
}

var file_apps_alert_pb_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_apps_alert_pb_rpc_proto_goTypes = []any{
	(*HandlerAlertReq)(nil), // 0: K8sGenie.alert.HandlerAlertReq
	(*Alert)(nil),           // 1: K8sGenie.alert.Alert
	(*HandlerAlertRsp)(nil), // 2: K8sGenie.alert.HandlerAlertRsp
	(*Response)(nil),        // 3: K8sGenie.alert.Response
	nil,                     // 4: K8sGenie.alert.Alert.LabelsEntry
	nil,                     // 5: K8sGenie.alert.Alert.AnnotationsEntry
}
var file_apps_alert_pb_rpc_proto_depIdxs = []int32{
	1, // 0: K8sGenie.alert.HandlerAlertReq.alerts:type_name -> K8sGenie.alert.Alert
	4, // 1: K8sGenie.alert.Alert.labels:type_name -> K8sGenie.alert.Alert.LabelsEntry
	5, // 2: K8sGenie.alert.Alert.annotations:type_name -> K8sGenie.alert.Alert.AnnotationsEntry
	3, // 3: K8sGenie.alert.HandlerAlertRsp.rsp:type_name -> K8sGenie.alert.Response
	0, // 4: K8sGenie.alert.Rpc.HandlerAlert:input_type -> K8sGenie.alert.HandlerAlertReq
	2, // 5: K8sGenie.alert.Rpc.HandlerAlert:output_type -> K8sGenie.alert.HandlerAlertRsp
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_apps_alert_pb_rpc_proto_init() }
func file_apps_alert_pb_rpc_proto_init() {
	if File_apps_alert_pb_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_apps_alert_pb_rpc_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*HandlerAlertReq); i {
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
		file_apps_alert_pb_rpc_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Alert); i {
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
		file_apps_alert_pb_rpc_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*HandlerAlertRsp); i {
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
		file_apps_alert_pb_rpc_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_apps_alert_pb_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_apps_alert_pb_rpc_proto_goTypes,
		DependencyIndexes: file_apps_alert_pb_rpc_proto_depIdxs,
		MessageInfos:      file_apps_alert_pb_rpc_proto_msgTypes,
	}.Build()
	File_apps_alert_pb_rpc_proto = out.File
	file_apps_alert_pb_rpc_proto_rawDesc = nil
	file_apps_alert_pb_rpc_proto_goTypes = nil
	file_apps_alert_pb_rpc_proto_depIdxs = nil
}
