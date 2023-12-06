// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0
// source: cmd/virtual-kubelet/internal/provider/grpc/pod_provider.proto

package grpc

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

type ConfigureNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoreV1NodeJson string `protobuf:"bytes,1,opt,name=core_v1_node_json,json=coreV1NodeJson,proto3" json:"core_v1_node_json,omitempty"`
}

func (x *ConfigureNodeRequest) Reset() {
	*x = ConfigureNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureNodeRequest) ProtoMessage() {}

func (x *ConfigureNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureNodeRequest.ProtoReflect.Descriptor instead.
func (*ConfigureNodeRequest) Descriptor() ([]byte, []int) {
	return file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDescGZIP(), []int{0}
}

func (x *ConfigureNodeRequest) GetCoreV1NodeJson() string {
	if x != nil {
		return x.CoreV1NodeJson
	}
	return ""
}

type ConfigureNodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoreV1NodeJson string `protobuf:"bytes,1,opt,name=core_v1_node_json,json=coreV1NodeJson,proto3" json:"core_v1_node_json,omitempty"`
}

func (x *ConfigureNodeReply) Reset() {
	*x = ConfigureNodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigureNodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigureNodeReply) ProtoMessage() {}

func (x *ConfigureNodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigureNodeReply.ProtoReflect.Descriptor instead.
func (*ConfigureNodeReply) Descriptor() ([]byte, []int) {
	return file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigureNodeReply) GetCoreV1NodeJson() string {
	if x != nil {
		return x.CoreV1NodeJson
	}
	return ""
}

type CreatePodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoreV1PodJson string `protobuf:"bytes,1,opt,name=core_v1_pod_json,json=coreV1PodJson,proto3" json:"core_v1_pod_json,omitempty"`
}

func (x *CreatePodRequest) Reset() {
	*x = CreatePodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePodRequest) ProtoMessage() {}

func (x *CreatePodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePodRequest.ProtoReflect.Descriptor instead.
func (*CreatePodRequest) Descriptor() ([]byte, []int) {
	return file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePodRequest) GetCoreV1PodJson() string {
	if x != nil {
		return x.CoreV1PodJson
	}
	return ""
}

type CreatePodReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePodReply) Reset() {
	*x = CreatePodReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePodReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePodReply) ProtoMessage() {}

func (x *CreatePodReply) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePodReply.ProtoReflect.Descriptor instead.
func (*CreatePodReply) Descriptor() ([]byte, []int) {
	return file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDescGZIP(), []int{3}
}

var File_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto protoreflect.FileDescriptor

var file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x63, 0x6d, 0x64, 0x2f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x2d, 0x6b, 0x75,
	0x62, 0x65, 0x6c, 0x65, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x6f, 0x64,
	0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x41, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x11, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x76, 0x31, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x72, 0x65, 0x56, 0x31, 0x4e, 0x6f, 0x64, 0x65, 0x4a, 0x73,
	0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x29, 0x0a, 0x11, 0x63, 0x6f, 0x72, 0x65,
	0x5f, 0x76, 0x31, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x72, 0x65, 0x56, 0x31, 0x4e, 0x6f, 0x64, 0x65, 0x4a,
	0x73, 0x6f, 0x6e, 0x22, 0x3b, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x10, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x76, 0x31, 0x5f, 0x70, 0x6f, 0x64, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x65, 0x56, 0x31, 0x50, 0x6f, 0x64, 0x4a, 0x73, 0x6f, 0x6e,
	0x22, 0x10, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x32, 0x7b, 0x0a, 0x0b, 0x50, 0x6f, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x3b, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x12, 0x15, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x12, 0x11, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x75,
	0x69, 0x73, 0x61, 0x77, 0x65, 0x73, 0x6f, 0x6d, 0x65, 0x2f, 0x73, 0x6b, 0x79, 0x2d, 0x61, 0x74,
	0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDescOnce sync.Once
	file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDescData = file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDesc
)

func file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDescGZIP() []byte {
	file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDescOnce.Do(func() {
		file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDescData)
	})
	return file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDescData
}

var file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_goTypes = []interface{}{
	(*ConfigureNodeRequest)(nil), // 0: ConfigureNodeRequest
	(*ConfigureNodeReply)(nil),   // 1: ConfigureNodeReply
	(*CreatePodRequest)(nil),     // 2: CreatePodRequest
	(*CreatePodReply)(nil),       // 3: CreatePodReply
}
var file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_depIdxs = []int32{
	0, // 0: PodProvider.ConfigureNode:input_type -> ConfigureNodeRequest
	2, // 1: PodProvider.CreatePod:input_type -> CreatePodRequest
	1, // 2: PodProvider.ConfigureNode:output_type -> ConfigureNodeReply
	3, // 3: PodProvider.CreatePod:output_type -> CreatePodReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_init() }
func file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_init() {
	if File_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureNodeRequest); i {
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
		file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigureNodeReply); i {
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
		file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePodRequest); i {
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
		file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePodReply); i {
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
			RawDescriptor: file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_goTypes,
		DependencyIndexes: file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_depIdxs,
		MessageInfos:      file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_msgTypes,
	}.Build()
	File_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto = out.File
	file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_rawDesc = nil
	file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_goTypes = nil
	file_cmd_virtual_kubelet_internal_provider_grpc_pod_provider_proto_depIdxs = nil
}
