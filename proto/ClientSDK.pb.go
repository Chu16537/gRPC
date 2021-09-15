// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: ClientSDK.proto

package ClientSDK

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

type UnaryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnaryReq string `protobuf:"bytes,1,opt,name=unaryReq,proto3" json:"unaryReq,omitempty"`
}

func (x *UnaryReq) Reset() {
	*x = UnaryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientSDK_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryReq) ProtoMessage() {}

func (x *UnaryReq) ProtoReflect() protoreflect.Message {
	mi := &file_ClientSDK_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryReq.ProtoReflect.Descriptor instead.
func (*UnaryReq) Descriptor() ([]byte, []int) {
	return file_ClientSDK_proto_rawDescGZIP(), []int{0}
}

func (x *UnaryReq) GetUnaryReq() string {
	if x != nil {
		return x.UnaryReq
	}
	return ""
}

type UnaryRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnaryRes string `protobuf:"bytes,1,opt,name=unaryRes,proto3" json:"unaryRes,omitempty"`
}

func (x *UnaryRes) Reset() {
	*x = UnaryRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientSDK_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnaryRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryRes) ProtoMessage() {}

func (x *UnaryRes) ProtoReflect() protoreflect.Message {
	mi := &file_ClientSDK_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryRes.ProtoReflect.Descriptor instead.
func (*UnaryRes) Descriptor() ([]byte, []int) {
	return file_ClientSDK_proto_rawDescGZIP(), []int{1}
}

func (x *UnaryRes) GetUnaryRes() string {
	if x != nil {
		return x.UnaryRes
	}
	return ""
}

type ServerStreamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerStreamReq string `protobuf:"bytes,1,opt,name=serverStreamReq,proto3" json:"serverStreamReq,omitempty"`
}

func (x *ServerStreamReq) Reset() {
	*x = ServerStreamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientSDK_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStreamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamReq) ProtoMessage() {}

func (x *ServerStreamReq) ProtoReflect() protoreflect.Message {
	mi := &file_ClientSDK_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStreamReq.ProtoReflect.Descriptor instead.
func (*ServerStreamReq) Descriptor() ([]byte, []int) {
	return file_ClientSDK_proto_rawDescGZIP(), []int{2}
}

func (x *ServerStreamReq) GetServerStreamReq() string {
	if x != nil {
		return x.ServerStreamReq
	}
	return ""
}

type ServerStreamRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerStreamRes string `protobuf:"bytes,1,opt,name=serverStreamRes,proto3" json:"serverStreamRes,omitempty"`
}

func (x *ServerStreamRes) Reset() {
	*x = ServerStreamRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientSDK_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStreamRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStreamRes) ProtoMessage() {}

func (x *ServerStreamRes) ProtoReflect() protoreflect.Message {
	mi := &file_ClientSDK_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStreamRes.ProtoReflect.Descriptor instead.
func (*ServerStreamRes) Descriptor() ([]byte, []int) {
	return file_ClientSDK_proto_rawDescGZIP(), []int{3}
}

func (x *ServerStreamRes) GetServerStreamRes() string {
	if x != nil {
		return x.ServerStreamRes
	}
	return ""
}

type ClientStreamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientStreamReq int32 `protobuf:"varint,1,opt,name=clientStreamReq,proto3" json:"clientStreamReq,omitempty"`
}

func (x *ClientStreamReq) Reset() {
	*x = ClientStreamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientSDK_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamReq) ProtoMessage() {}

func (x *ClientStreamReq) ProtoReflect() protoreflect.Message {
	mi := &file_ClientSDK_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamReq.ProtoReflect.Descriptor instead.
func (*ClientStreamReq) Descriptor() ([]byte, []int) {
	return file_ClientSDK_proto_rawDescGZIP(), []int{4}
}

func (x *ClientStreamReq) GetClientStreamReq() int32 {
	if x != nil {
		return x.ClientStreamReq
	}
	return 0
}

type ClientStreamRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientStreamRes int32 `protobuf:"varint,1,opt,name=clientStreamRes,proto3" json:"clientStreamRes,omitempty"`
}

func (x *ClientStreamRes) Reset() {
	*x = ClientStreamRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientSDK_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientStreamRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientStreamRes) ProtoMessage() {}

func (x *ClientStreamRes) ProtoReflect() protoreflect.Message {
	mi := &file_ClientSDK_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientStreamRes.ProtoReflect.Descriptor instead.
func (*ClientStreamRes) Descriptor() ([]byte, []int) {
	return file_ClientSDK_proto_rawDescGZIP(), []int{5}
}

func (x *ClientStreamRes) GetClientStreamRes() int32 {
	if x != nil {
		return x.ClientStreamRes
	}
	return 0
}

type BidirectionalStreamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BidirectionalStreamReq string `protobuf:"bytes,1,opt,name=bidirectionalStreamReq,proto3" json:"bidirectionalStreamReq,omitempty"`
}

func (x *BidirectionalStreamReq) Reset() {
	*x = BidirectionalStreamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientSDK_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidirectionalStreamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidirectionalStreamReq) ProtoMessage() {}

func (x *BidirectionalStreamReq) ProtoReflect() protoreflect.Message {
	mi := &file_ClientSDK_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidirectionalStreamReq.ProtoReflect.Descriptor instead.
func (*BidirectionalStreamReq) Descriptor() ([]byte, []int) {
	return file_ClientSDK_proto_rawDescGZIP(), []int{6}
}

func (x *BidirectionalStreamReq) GetBidirectionalStreamReq() string {
	if x != nil {
		return x.BidirectionalStreamReq
	}
	return ""
}

type BidirectionalStreamRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BidirectionalStreamRes string `protobuf:"bytes,1,opt,name=bidirectionalStreamRes,proto3" json:"bidirectionalStreamRes,omitempty"`
}

func (x *BidirectionalStreamRes) Reset() {
	*x = BidirectionalStreamRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClientSDK_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidirectionalStreamRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidirectionalStreamRes) ProtoMessage() {}

func (x *BidirectionalStreamRes) ProtoReflect() protoreflect.Message {
	mi := &file_ClientSDK_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidirectionalStreamRes.ProtoReflect.Descriptor instead.
func (*BidirectionalStreamRes) Descriptor() ([]byte, []int) {
	return file_ClientSDK_proto_rawDescGZIP(), []int{7}
}

func (x *BidirectionalStreamRes) GetBidirectionalStreamRes() string {
	if x != nil {
		return x.BidirectionalStreamRes
	}
	return ""
}

var File_ClientSDK_proto protoreflect.FileDescriptor

var file_ClientSDK_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x44, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x44, 0x4b, 0x22, 0x26, 0x0a, 0x08,
	0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x61, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x22, 0x26, 0x0a, 0x08, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x0f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x12,
	0x28, 0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x22, 0x3b, 0x0a, 0x0f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x22, 0x3b, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x52, 0x65, 0x71, 0x22, 0x3b, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x22, 0x50, 0x0a, 0x16, 0x62, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x36, 0x0a, 0x16, 0x62, 0x69,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x62, 0x69, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x22, 0x50, 0x0a, 0x16, 0x62, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x16,
	0x62, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x62, 0x69,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x32, 0xc8, 0x02, 0x0a, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x05, 0x75, 0x6e, 0x61, 0x72, 0x79,
	0x12, 0x13, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x44, 0x4b, 0x2e, 0x55, 0x6e, 0x61,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x44,
	0x4b, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x10,
	0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x1a, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x44, 0x4b, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x44, 0x4b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4e, 0x0a, 0x10,
	0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x1a, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x44, 0x4b, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x44, 0x4b, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x22, 0x00, 0x28, 0x01, 0x12, 0x61, 0x0a, 0x13,
	0x62, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x21, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x44, 0x4b, 0x2e,
	0x62, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53,
	0x44, 0x4b, 0x2e, 0x62, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x44, 0x4b, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ClientSDK_proto_rawDescOnce sync.Once
	file_ClientSDK_proto_rawDescData = file_ClientSDK_proto_rawDesc
)

func file_ClientSDK_proto_rawDescGZIP() []byte {
	file_ClientSDK_proto_rawDescOnce.Do(func() {
		file_ClientSDK_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClientSDK_proto_rawDescData)
	})
	return file_ClientSDK_proto_rawDescData
}

var file_ClientSDK_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ClientSDK_proto_goTypes = []interface{}{
	(*UnaryReq)(nil),               // 0: ClientSDK.UnaryReq
	(*UnaryRes)(nil),               // 1: ClientSDK.UnaryRes
	(*ServerStreamReq)(nil),        // 2: ClientSDK.serverStreamReq
	(*ServerStreamRes)(nil),        // 3: ClientSDK.serverStreamRes
	(*ClientStreamReq)(nil),        // 4: ClientSDK.clientStreamReq
	(*ClientStreamRes)(nil),        // 5: ClientSDK.clientStreamRes
	(*BidirectionalStreamReq)(nil), // 6: ClientSDK.bidirectionalStreamReq
	(*BidirectionalStreamRes)(nil), // 7: ClientSDK.bidirectionalStreamRes
}
var file_ClientSDK_proto_depIdxs = []int32{
	0, // 0: ClientSDK.GameController.unary:input_type -> ClientSDK.UnaryReq
	2, // 1: ClientSDK.GameController.testStreamServer:input_type -> ClientSDK.serverStreamReq
	4, // 2: ClientSDK.GameController.testStreamClient:input_type -> ClientSDK.clientStreamReq
	6, // 3: ClientSDK.GameController.bidirectionalStream:input_type -> ClientSDK.bidirectionalStreamReq
	1, // 4: ClientSDK.GameController.unary:output_type -> ClientSDK.UnaryRes
	3, // 5: ClientSDK.GameController.testStreamServer:output_type -> ClientSDK.serverStreamRes
	5, // 6: ClientSDK.GameController.testStreamClient:output_type -> ClientSDK.clientStreamRes
	7, // 7: ClientSDK.GameController.bidirectionalStream:output_type -> ClientSDK.bidirectionalStreamRes
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ClientSDK_proto_init() }
func file_ClientSDK_proto_init() {
	if File_ClientSDK_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ClientSDK_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryReq); i {
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
		file_ClientSDK_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnaryRes); i {
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
		file_ClientSDK_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStreamReq); i {
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
		file_ClientSDK_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStreamRes); i {
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
		file_ClientSDK_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamReq); i {
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
		file_ClientSDK_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientStreamRes); i {
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
		file_ClientSDK_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidirectionalStreamReq); i {
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
		file_ClientSDK_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidirectionalStreamRes); i {
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
			RawDescriptor: file_ClientSDK_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ClientSDK_proto_goTypes,
		DependencyIndexes: file_ClientSDK_proto_depIdxs,
		MessageInfos:      file_ClientSDK_proto_msgTypes,
	}.Build()
	File_ClientSDK_proto = out.File
	file_ClientSDK_proto_rawDesc = nil
	file_ClientSDK_proto_goTypes = nil
	file_ClientSDK_proto_depIdxs = nil
}