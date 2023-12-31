// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/v1/message/message.proto

package messagev1

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

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string  `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	FileUrl *string `protobuf:"bytes,2,opt,name=file_url,json=fileUrl,proto3,oneof" json:"file_url,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_message_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_message_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_proto_v1_message_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetFileUrl() string {
	if x != nil && x.FileUrl != nil {
		return *x.FileUrl
	}
	return ""
}

type SendDirectMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	FromMemberId string   `protobuf:"bytes,2,opt,name=from_member_id,json=fromMemberId,proto3" json:"from_member_id,omitempty"`
	ToMemberId   string   `protobuf:"bytes,3,opt,name=to_member_id,json=toMemberId,proto3" json:"to_member_id,omitempty"`
}

func (x *SendDirectMessageRequest) Reset() {
	*x = SendDirectMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_message_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDirectMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDirectMessageRequest) ProtoMessage() {}

func (x *SendDirectMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_message_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDirectMessageRequest.ProtoReflect.Descriptor instead.
func (*SendDirectMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_message_message_proto_rawDescGZIP(), []int{1}
}

func (x *SendDirectMessageRequest) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *SendDirectMessageRequest) GetFromMemberId() string {
	if x != nil {
		return x.FromMemberId
	}
	return ""
}

func (x *SendDirectMessageRequest) GetToMemberId() string {
	if x != nil {
		return x.ToMemberId
	}
	return ""
}

type SendDirectMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	SentAt  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=sent_at,json=sentAt,proto3" json:"sent_at,omitempty"`
}

func (x *SendDirectMessageResponse) Reset() {
	*x = SendDirectMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_message_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendDirectMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendDirectMessageResponse) ProtoMessage() {}

func (x *SendDirectMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_message_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendDirectMessageResponse.ProtoReflect.Descriptor instead.
func (*SendDirectMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_message_message_proto_rawDescGZIP(), []int{2}
}

func (x *SendDirectMessageResponse) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *SendDirectMessageResponse) GetSentAt() *timestamppb.Timestamp {
	if x != nil {
		return x.SentAt
	}
	return nil
}

var File_proto_v1_message_message_proto protoreflect.FileDescriptor

var file_proto_v1_message_message_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x66, 0x69,
	0x6c, 0x65, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x97, 0x01, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x72, 0x6f, 0x6d, 0x5f,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0c, 0x74, 0x6f, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x6f, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x85, 0x01, 0x0a, 0x19, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x32, 0x7f, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0xc6, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69,
	0x6c, 0x77, 0x6f, 0x6e, 0x67, 0x30, 0x30, 0x2f, 0x67, 0x6f, 0x2d, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x72, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x50, 0x4d, 0x58, 0xaa, 0x02, 0x10, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_v1_message_message_proto_rawDescOnce sync.Once
	file_proto_v1_message_message_proto_rawDescData = file_proto_v1_message_message_proto_rawDesc
)

func file_proto_v1_message_message_proto_rawDescGZIP() []byte {
	file_proto_v1_message_message_proto_rawDescOnce.Do(func() {
		file_proto_v1_message_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_message_message_proto_rawDescData)
	})
	return file_proto_v1_message_message_proto_rawDescData
}

var file_proto_v1_message_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_v1_message_message_proto_goTypes = []interface{}{
	(*Message)(nil),                   // 0: proto.message.v1.Message
	(*SendDirectMessageRequest)(nil),  // 1: proto.message.v1.SendDirectMessageRequest
	(*SendDirectMessageResponse)(nil), // 2: proto.message.v1.SendDirectMessageResponse
	(*timestamppb.Timestamp)(nil),     // 3: google.protobuf.Timestamp
}
var file_proto_v1_message_message_proto_depIdxs = []int32{
	0, // 0: proto.message.v1.SendDirectMessageRequest.message:type_name -> proto.message.v1.Message
	0, // 1: proto.message.v1.SendDirectMessageResponse.message:type_name -> proto.message.v1.Message
	3, // 2: proto.message.v1.SendDirectMessageResponse.sent_at:type_name -> google.protobuf.Timestamp
	1, // 3: proto.message.v1.ServerService.SendDirectMessage:input_type -> proto.message.v1.SendDirectMessageRequest
	2, // 4: proto.message.v1.ServerService.SendDirectMessage:output_type -> proto.message.v1.SendDirectMessageResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_v1_message_message_proto_init() }
func file_proto_v1_message_message_proto_init() {
	if File_proto_v1_message_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_message_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_proto_v1_message_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDirectMessageRequest); i {
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
		file_proto_v1_message_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendDirectMessageResponse); i {
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
	file_proto_v1_message_message_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_v1_message_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_message_message_proto_goTypes,
		DependencyIndexes: file_proto_v1_message_message_proto_depIdxs,
		MessageInfos:      file_proto_v1_message_message_proto_msgTypes,
	}.Build()
	File_proto_v1_message_message_proto = out.File
	file_proto_v1_message_message_proto_rawDesc = nil
	file_proto_v1_message_message_proto_goTypes = nil
	file_proto_v1_message_message_proto_depIdxs = nil
}
