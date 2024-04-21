//**
// Common definition details
//*

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: scancctv/api/common/v1/scancctv-common.proto

package commonv1

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

// Status code Enum
type StatusCode int32

const (
	StatusCode_UNSPECIFIED             StatusCode = 0
	StatusCode_SUCCESS                 StatusCode = 1
	StatusCode_SUCCEEDED_WITH_WARNINGS StatusCode = 2
	StatusCode_WARNING                 StatusCode = 3
	StatusCode_FAILED                  StatusCode = 4
)

// Enum value maps for StatusCode.
var (
	StatusCode_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "SUCCESS",
		2: "SUCCEEDED_WITH_WARNINGS",
		3: "WARNING",
		4: "FAILED",
	}
	StatusCode_value = map[string]int32{
		"UNSPECIFIED":             0,
		"SUCCESS":                 1,
		"SUCCEEDED_WITH_WARNINGS": 2,
		"WARNING":                 3,
		"FAILED":                  4,
	}
)

func (x StatusCode) Enum() *StatusCode {
	p := new(StatusCode)
	*p = x
	return p
}

func (x StatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_scancctv_api_common_v1_scancctv_common_proto_enumTypes[0].Descriptor()
}

func (StatusCode) Type() protoreflect.EnumType {
	return &file_scancctv_api_common_v1_scancctv_common_proto_enumTypes[0]
}

func (x StatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusCode.Descriptor instead.
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return file_scancctv_api_common_v1_scancctv_common_proto_rawDescGZIP(), []int{0}
}

// Detailed response details
type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// response status
	Status StatusCode `protobuf:"varint,1,opt,name=status,proto3,enum=scancctv.api.common.v1.StatusCode" json:"status,omitempty"`
	// Status message
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scancctv_api_common_v1_scancctv_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scancctv_api_common_v1_scancctv_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_scancctv_api_common_v1_scancctv_common_proto_rawDescGZIP(), []int{0}
}

func (x *StatusResponse) GetStatus() StatusCode {
	if x != nil {
		return x.Status
	}
	return StatusCode_UNSPECIFIED
}

func (x *StatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Echo Message Request
type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scancctv_api_common_v1_scancctv_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scancctv_api_common_v1_scancctv_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_scancctv_api_common_v1_scancctv_common_proto_rawDescGZIP(), []int{1}
}

func (x *EchoRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Echo Message Response
type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scancctv_api_common_v1_scancctv_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scancctv_api_common_v1_scancctv_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_scancctv_api_common_v1_scancctv_common_proto_rawDescGZIP(), []int{2}
}

func (x *EchoResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_scancctv_api_common_v1_scancctv_common_proto protoreflect.FileDescriptor

var file_scancctv_api_common_v1_scancctv_common_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x63, 0x61, 0x6e, 0x63, 0x63, 0x74, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x63, 0x63, 0x74,
	0x76, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x73, 0x63, 0x61, 0x6e, 0x63, 0x63, 0x74, 0x76, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x66, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x63,
	0x63, 0x74, 0x76, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x27,
	0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x45, 0x63, 0x68, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2a, 0x60, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x1b, 0x0a,
	0x17, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f,
	0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x53, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41,
	0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x04, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x63, 0x63, 0x74, 0x76, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_scancctv_api_common_v1_scancctv_common_proto_rawDescOnce sync.Once
	file_scancctv_api_common_v1_scancctv_common_proto_rawDescData = file_scancctv_api_common_v1_scancctv_common_proto_rawDesc
)

func file_scancctv_api_common_v1_scancctv_common_proto_rawDescGZIP() []byte {
	file_scancctv_api_common_v1_scancctv_common_proto_rawDescOnce.Do(func() {
		file_scancctv_api_common_v1_scancctv_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_scancctv_api_common_v1_scancctv_common_proto_rawDescData)
	})
	return file_scancctv_api_common_v1_scancctv_common_proto_rawDescData
}

var file_scancctv_api_common_v1_scancctv_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_scancctv_api_common_v1_scancctv_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_scancctv_api_common_v1_scancctv_common_proto_goTypes = []interface{}{
	(StatusCode)(0),        // 0: scancctv.api.common.v1.StatusCode
	(*StatusResponse)(nil), // 1: scancctv.api.common.v1.StatusResponse
	(*EchoRequest)(nil),    // 2: scancctv.api.common.v1.EchoRequest
	(*EchoResponse)(nil),   // 3: scancctv.api.common.v1.EchoResponse
}
var file_scancctv_api_common_v1_scancctv_common_proto_depIdxs = []int32{
	0, // 0: scancctv.api.common.v1.StatusResponse.status:type_name -> scancctv.api.common.v1.StatusCode
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_scancctv_api_common_v1_scancctv_common_proto_init() }
func file_scancctv_api_common_v1_scancctv_common_proto_init() {
	if File_scancctv_api_common_v1_scancctv_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scancctv_api_common_v1_scancctv_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_scancctv_api_common_v1_scancctv_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
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
		file_scancctv_api_common_v1_scancctv_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResponse); i {
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
			RawDescriptor: file_scancctv_api_common_v1_scancctv_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_scancctv_api_common_v1_scancctv_common_proto_goTypes,
		DependencyIndexes: file_scancctv_api_common_v1_scancctv_common_proto_depIdxs,
		EnumInfos:         file_scancctv_api_common_v1_scancctv_common_proto_enumTypes,
		MessageInfos:      file_scancctv_api_common_v1_scancctv_common_proto_msgTypes,
	}.Build()
	File_scancctv_api_common_v1_scancctv_common_proto = out.File
	file_scancctv_api_common_v1_scancctv_common_proto_rawDesc = nil
	file_scancctv_api_common_v1_scancctv_common_proto_goTypes = nil
	file_scancctv_api_common_v1_scancctv_common_proto_depIdxs = nil
}
