// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: pbcommon.proto

package pbcommon

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

type DiscordType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DiscordType) Reset() {
	*x = DiscordType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbcommon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscordType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscordType) ProtoMessage() {}

func (x *DiscordType) ProtoReflect() protoreflect.Message {
	mi := &file_pbcommon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscordType.ProtoReflect.Descriptor instead.
func (*DiscordType) Descriptor() ([]byte, []int) {
	return file_pbcommon_proto_rawDescGZIP(), []int{0}
}

func (x *DiscordType) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type TelegramType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TelegramType) Reset() {
	*x = TelegramType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbcommon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TelegramType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TelegramType) ProtoMessage() {}

func (x *TelegramType) ProtoReflect() protoreflect.Message {
	mi := &file_pbcommon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TelegramType.ProtoReflect.Descriptor instead.
func (*TelegramType) Descriptor() ([]byte, []int) {
	return file_pbcommon_proto_rawDescGZIP(), []int{1}
}

func (x *TelegramType) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_pbcommon_proto protoreflect.FileDescriptor

var file_pbcommon_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x1d, 0x0a, 0x0b, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x0c, 0x54, 0x65, 0x6c, 0x65, 0x67, 0x72, 0x61,
	0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x69, 0x66, 0x74, 0x79, 0x31, 0x31, 0x2f, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x2d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbcommon_proto_rawDescOnce sync.Once
	file_pbcommon_proto_rawDescData = file_pbcommon_proto_rawDesc
)

func file_pbcommon_proto_rawDescGZIP() []byte {
	file_pbcommon_proto_rawDescOnce.Do(func() {
		file_pbcommon_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbcommon_proto_rawDescData)
	})
	return file_pbcommon_proto_rawDescData
}

var file_pbcommon_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pbcommon_proto_goTypes = []interface{}{
	(*DiscordType)(nil),  // 0: cosmos_notifier_grpc.DiscordType
	(*TelegramType)(nil), // 1: cosmos_notifier_grpc.TelegramType
}
var file_pbcommon_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbcommon_proto_init() }
func file_pbcommon_proto_init() {
	if File_pbcommon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbcommon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscordType); i {
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
		file_pbcommon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TelegramType); i {
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
			RawDescriptor: file_pbcommon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbcommon_proto_goTypes,
		DependencyIndexes: file_pbcommon_proto_depIdxs,
		MessageInfos:      file_pbcommon_proto_msgTypes,
	}.Build()
	File_pbcommon_proto = out.File
	file_pbcommon_proto_rawDesc = nil
	file_pbcommon_proto_goTypes = nil
	file_pbcommon_proto_depIdxs = nil
}
