// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: keys.proto

package api

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

type GetValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetValueRequest) Reset() {
	*x = GetValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueRequest) ProtoMessage() {}

func (x *GetValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueRequest.ProtoReflect.Descriptor instead.
func (*GetValueRequest) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{0}
}

func (x *GetValueRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetValueResponse) Reset() {
	*x = GetValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_keys_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueResponse) ProtoMessage() {}

func (x *GetValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_keys_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueResponse.ProtoReflect.Descriptor instead.
func (*GetValueResponse) Descriptor() ([]byte, []int) {
	return file_keys_proto_rawDescGZIP(), []int{1}
}

func (x *GetValueResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_keys_proto protoreflect.FileDescriptor

var file_keys_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d, 0x6b,
	0x31, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x28,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x4f, 0x0a, 0x04, 0x4b, 0x65, 0x79, 0x73,
	0x12, 0x47, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1c, 0x2e, 0x6d,
	0x6b, 0x31, 0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x6b, 0x31,
	0x2e, 0x6b, 0x65, 0x79, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6c, 0x65, 0x67, 0x6d, 0x79, 0x6d, 0x72,
	0x69, 0x6e, 0x2f, 0x6d, 0x6b, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_keys_proto_rawDescOnce sync.Once
	file_keys_proto_rawDescData = file_keys_proto_rawDesc
)

func file_keys_proto_rawDescGZIP() []byte {
	file_keys_proto_rawDescOnce.Do(func() {
		file_keys_proto_rawDescData = protoimpl.X.CompressGZIP(file_keys_proto_rawDescData)
	})
	return file_keys_proto_rawDescData
}

var file_keys_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_keys_proto_goTypes = []interface{}{
	(*GetValueRequest)(nil),  // 0: mk1.keys.v1.GetValueRequest
	(*GetValueResponse)(nil), // 1: mk1.keys.v1.GetValueResponse
}
var file_keys_proto_depIdxs = []int32{
	0, // 0: mk1.keys.v1.Keys.GetValue:input_type -> mk1.keys.v1.GetValueRequest
	1, // 1: mk1.keys.v1.Keys.GetValue:output_type -> mk1.keys.v1.GetValueResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_keys_proto_init() }
func file_keys_proto_init() {
	if File_keys_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_keys_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueRequest); i {
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
		file_keys_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueResponse); i {
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
			RawDescriptor: file_keys_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_keys_proto_goTypes,
		DependencyIndexes: file_keys_proto_depIdxs,
		MessageInfos:      file_keys_proto_msgTypes,
	}.Build()
	File_keys_proto = out.File
	file_keys_proto_rawDesc = nil
	file_keys_proto_goTypes = nil
	file_keys_proto_depIdxs = nil
}
