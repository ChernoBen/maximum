// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: src/maximum_proto/maximum.proto

package maximum_proto

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

type FindMaximumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *FindMaximumRequest) Reset() {
	*x = FindMaximumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_maximum_proto_maximum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaximumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaximumRequest) ProtoMessage() {}

func (x *FindMaximumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_maximum_proto_maximum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaximumRequest.ProtoReflect.Descriptor instead.
func (*FindMaximumRequest) Descriptor() ([]byte, []int) {
	return file_src_maximum_proto_maximum_proto_rawDescGZIP(), []int{0}
}

func (x *FindMaximumRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type FindMaximumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Maximum int32 `protobuf:"varint,1,opt,name=maximum,proto3" json:"maximum,omitempty"`
}

func (x *FindMaximumResponse) Reset() {
	*x = FindMaximumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_maximum_proto_maximum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMaximumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMaximumResponse) ProtoMessage() {}

func (x *FindMaximumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_maximum_proto_maximum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMaximumResponse.ProtoReflect.Descriptor instead.
func (*FindMaximumResponse) Descriptor() ([]byte, []int) {
	return file_src_maximum_proto_maximum_proto_rawDescGZIP(), []int{1}
}

func (x *FindMaximumResponse) GetMaximum() int32 {
	if x != nil {
		return x.Maximum
	}
	return 0
}

var File_src_maximum_proto_maximum_proto protoreflect.FileDescriptor

var file_src_maximum_proto_maximum_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x73, 0x72, 0x63, 0x22, 0x2c, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61,
	0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x61,
	0x78, 0x69, 0x6d, 0x75, 0x6d, 0x32, 0x5b, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x46, 0x69,
	0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x12, 0x17, 0x2e, 0x73, 0x72, 0x63, 0x2e,
	0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x72, 0x63, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x61, 0x78,
	0x69, 0x6d, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61, 0x78, 0x69,
	0x6d, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_src_maximum_proto_maximum_proto_rawDescOnce sync.Once
	file_src_maximum_proto_maximum_proto_rawDescData = file_src_maximum_proto_maximum_proto_rawDesc
)

func file_src_maximum_proto_maximum_proto_rawDescGZIP() []byte {
	file_src_maximum_proto_maximum_proto_rawDescOnce.Do(func() {
		file_src_maximum_proto_maximum_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_maximum_proto_maximum_proto_rawDescData)
	})
	return file_src_maximum_proto_maximum_proto_rawDescData
}

var file_src_maximum_proto_maximum_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_src_maximum_proto_maximum_proto_goTypes = []interface{}{
	(*FindMaximumRequest)(nil),  // 0: src.FindMaximumRequest
	(*FindMaximumResponse)(nil), // 1: src.FindMaximumResponse
}
var file_src_maximum_proto_maximum_proto_depIdxs = []int32{
	0, // 0: src.CalculatorService.FindMaximum:input_type -> src.FindMaximumRequest
	1, // 1: src.CalculatorService.FindMaximum:output_type -> src.FindMaximumResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_maximum_proto_maximum_proto_init() }
func file_src_maximum_proto_maximum_proto_init() {
	if File_src_maximum_proto_maximum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_maximum_proto_maximum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaximumRequest); i {
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
		file_src_maximum_proto_maximum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMaximumResponse); i {
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
			RawDescriptor: file_src_maximum_proto_maximum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_maximum_proto_maximum_proto_goTypes,
		DependencyIndexes: file_src_maximum_proto_maximum_proto_depIdxs,
		MessageInfos:      file_src_maximum_proto_maximum_proto_msgTypes,
	}.Build()
	File_src_maximum_proto_maximum_proto = out.File
	file_src_maximum_proto_maximum_proto_rawDesc = nil
	file_src_maximum_proto_maximum_proto_goTypes = nil
	file_src_maximum_proto_maximum_proto_depIdxs = nil
}
