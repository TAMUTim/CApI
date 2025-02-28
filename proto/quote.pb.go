// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: proto/quote.proto

package proto

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

type QuoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Temp string `protobuf:"bytes,1,opt,name=temp,proto3" json:"temp,omitempty"`
}

func (x *QuoteRequest) Reset() {
	*x = QuoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quote_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteRequest) ProtoMessage() {}

func (x *QuoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quote_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteRequest.ProtoReflect.Descriptor instead.
func (*QuoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_quote_proto_rawDescGZIP(), []int{0}
}

func (x *QuoteRequest) GetTemp() string {
	if x != nil {
		return x.Temp
	}
	return ""
}

type QuoteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quote string                 `protobuf:"bytes,1,opt,name=quote,proto3" json:"quote,omitempty"`
	Time  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *QuoteReply) Reset() {
	*x = QuoteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_quote_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuoteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuoteReply) ProtoMessage() {}

func (x *QuoteReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_quote_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuoteReply.ProtoReflect.Descriptor instead.
func (*QuoteReply) Descriptor() ([]byte, []int) {
	return file_proto_quote_proto_rawDescGZIP(), []int{1}
}

func (x *QuoteReply) GetQuote() string {
	if x != nil {
		return x.Quote
	}
	return ""
}

func (x *QuoteReply) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

var File_proto_quote_proto protoreflect.FileDescriptor

var file_proto_quote_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x6d, 0x70, 0x22,
	0x52, 0x0a, 0x0a, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x6f, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74,
	0x69, 0x6d, 0x65, 0x32, 0x45, 0x0a, 0x0c, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x61, 0x62,
	0x62, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x09, 0x47, 0x72, 0x61, 0x62, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x51, 0x75,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x43, 0x41,
	0x70, 0x49, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_quote_proto_rawDescOnce sync.Once
	file_proto_quote_proto_rawDescData = file_proto_quote_proto_rawDesc
)

func file_proto_quote_proto_rawDescGZIP() []byte {
	file_proto_quote_proto_rawDescOnce.Do(func() {
		file_proto_quote_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_quote_proto_rawDescData)
	})
	return file_proto_quote_proto_rawDescData
}

var file_proto_quote_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_quote_proto_goTypes = []interface{}{
	(*QuoteRequest)(nil),          // 0: quote.QuoteRequest
	(*QuoteReply)(nil),            // 1: quote.QuoteReply
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_proto_quote_proto_depIdxs = []int32{
	2, // 0: quote.QuoteReply.time:type_name -> google.protobuf.Timestamp
	0, // 1: quote.QuoteGrabber.GrabQuote:input_type -> quote.QuoteRequest
	1, // 2: quote.QuoteGrabber.GrabQuote:output_type -> quote.QuoteReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_quote_proto_init() }
func file_proto_quote_proto_init() {
	if File_proto_quote_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_quote_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteRequest); i {
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
		file_proto_quote_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuoteReply); i {
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
			RawDescriptor: file_proto_quote_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_quote_proto_goTypes,
		DependencyIndexes: file_proto_quote_proto_depIdxs,
		MessageInfos:      file_proto_quote_proto_msgTypes,
	}.Build()
	File_proto_quote_proto = out.File
	file_proto_quote_proto_rawDesc = nil
	file_proto_quote_proto_goTypes = nil
	file_proto_quote_proto_depIdxs = nil
}
