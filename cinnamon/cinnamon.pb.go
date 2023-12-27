// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: cinnamon.proto

package cinnamon

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Testing adding items to a priority queue
// arrival is added for timeout from priority queue
type InterceptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Priority  int64                  `protobuf:"varint,1,opt,name=priority,proto3" json:"priority,omitempty"`
	Route     string                 `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	Arrival   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=arrival,proto3" json:"arrival,omitempty"`
	Processed string                 `protobuf:"bytes,4,opt,name=processed,proto3" json:"processed,omitempty"`
}

func (x *InterceptRequest) Reset() {
	*x = InterceptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cinnamon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterceptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterceptRequest) ProtoMessage() {}

func (x *InterceptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cinnamon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterceptRequest.ProtoReflect.Descriptor instead.
func (*InterceptRequest) Descriptor() ([]byte, []int) {
	return file_cinnamon_proto_rawDescGZIP(), []int{0}
}

func (x *InterceptRequest) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *InterceptRequest) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *InterceptRequest) GetArrival() *timestamppb.Timestamp {
	if x != nil {
		return x.Arrival
	}
	return nil
}

func (x *InterceptRequest) GetProcessed() string {
	if x != nil {
		return x.Processed
	}
	return ""
}

type InterceptResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accepted bool   `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	Message  string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *InterceptResponse) Reset() {
	*x = InterceptResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cinnamon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InterceptResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InterceptResponse) ProtoMessage() {}

func (x *InterceptResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cinnamon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InterceptResponse.ProtoReflect.Descriptor instead.
func (*InterceptResponse) Descriptor() ([]byte, []int) {
	return file_cinnamon_proto_rawDescGZIP(), []int{1}
}

func (x *InterceptResponse) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

func (x *InterceptResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_cinnamon_proto protoreflect.FileDescriptor

var file_cinnamon_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x69, 0x6e, 0x6e, 0x61, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x98, 0x01, 0x0a, 0x10, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x61, 0x72, 0x72, 0x69,
	0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x61, 0x72, 0x72, 0x69, 0x76, 0x61, 0x6c, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x22, 0x49, 0x0a, 0x11,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x3e, 0x0a, 0x08, 0x43, 0x69, 0x6e, 0x6e, 0x61,
	0x6d, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74,
	0x12, 0x11, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x65, 0x70, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x62, 0x61, 0x6e, 0x75, 0x65, 0x6c, 0x6f, 0x2f, 0x63,
	0x69, 0x6e, 0x6e, 0x61, 0x6d, 0x6f, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x6e, 0x6e, 0x61,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cinnamon_proto_rawDescOnce sync.Once
	file_cinnamon_proto_rawDescData = file_cinnamon_proto_rawDesc
)

func file_cinnamon_proto_rawDescGZIP() []byte {
	file_cinnamon_proto_rawDescOnce.Do(func() {
		file_cinnamon_proto_rawDescData = protoimpl.X.CompressGZIP(file_cinnamon_proto_rawDescData)
	})
	return file_cinnamon_proto_rawDescData
}

var file_cinnamon_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cinnamon_proto_goTypes = []interface{}{
	(*InterceptRequest)(nil),      // 0: InterceptRequest
	(*InterceptResponse)(nil),     // 1: InterceptResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_cinnamon_proto_depIdxs = []int32{
	2, // 0: InterceptRequest.arrival:type_name -> google.protobuf.Timestamp
	0, // 1: Cinnamon.Intercept:input_type -> InterceptRequest
	1, // 2: Cinnamon.Intercept:output_type -> InterceptResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cinnamon_proto_init() }
func file_cinnamon_proto_init() {
	if File_cinnamon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cinnamon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterceptRequest); i {
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
		file_cinnamon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InterceptResponse); i {
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
			RawDescriptor: file_cinnamon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cinnamon_proto_goTypes,
		DependencyIndexes: file_cinnamon_proto_depIdxs,
		MessageInfos:      file_cinnamon_proto_msgTypes,
	}.Build()
	File_cinnamon_proto = out.File
	file_cinnamon_proto_rawDesc = nil
	file_cinnamon_proto_goTypes = nil
	file_cinnamon_proto_depIdxs = nil
}
