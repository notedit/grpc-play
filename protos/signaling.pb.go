// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: protos/signaling.proto

package protos

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

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room string `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_signaling_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_signaling_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_protos_signaling_proto_rawDescGZIP(), []int{0}
}

func (x *JoinRequest) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *JoinRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_signaling_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_signaling_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_protos_signaling_proto_rawDescGZIP(), []int{1}
}

func (x *JoinResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type LeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room string `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *LeaveRequest) Reset() {
	*x = LeaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_signaling_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRequest) ProtoMessage() {}

func (x *LeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_signaling_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRequest.ProtoReflect.Descriptor instead.
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return file_protos_signaling_proto_rawDescGZIP(), []int{2}
}

func (x *LeaveRequest) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *LeaveRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type LeaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *LeaveResponse) Reset() {
	*x = LeaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_signaling_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveResponse) ProtoMessage() {}

func (x *LeaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_signaling_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveResponse.ProtoReflect.Descriptor instead.
func (*LeaveResponse) Descriptor() ([]byte, []int) {
	return file_protos_signaling_proto_rawDescGZIP(), []int{3}
}

func (x *LeaveResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type StreamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Room string `protobuf:"bytes,1,opt,name=room,proto3" json:"room,omitempty"`
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *StreamRequest) Reset() {
	*x = StreamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_signaling_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamRequest) ProtoMessage() {}

func (x *StreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_signaling_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamRequest.ProtoReflect.Descriptor instead.
func (*StreamRequest) Descriptor() ([]byte, []int) {
	return file_protos_signaling_proto_rawDescGZIP(), []int{4}
}

func (x *StreamRequest) GetRoom() string {
	if x != nil {
		return x.Room
	}
	return ""
}

func (x *StreamRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type Login struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Login) Reset() {
	*x = Login{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_signaling_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Login) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Login) ProtoMessage() {}

func (x *Login) ProtoReflect() protoreflect.Message {
	mi := &file_protos_signaling_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Login.ProtoReflect.Descriptor instead.
func (*Login) Descriptor() ([]byte, []int) {
	return file_protos_signaling_proto_rawDescGZIP(), []int{5}
}

func (x *Login) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type Leave struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *Leave) Reset() {
	*x = Leave{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_signaling_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Leave) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Leave) ProtoMessage() {}

func (x *Leave) ProtoReflect() protoreflect.Message {
	mi := &file_protos_signaling_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Leave.ProtoReflect.Descriptor instead.
func (*Leave) Descriptor() ([]byte, []int) {
	return file_protos_signaling_proto_rawDescGZIP(), []int{6}
}

func (x *Leave) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User    string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_signaling_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_protos_signaling_proto_msgTypes[7]
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
	return file_protos_signaling_proto_rawDescGZIP(), []int{7}
}

func (x *Message) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Shutdown struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Shutdown) Reset() {
	*x = Shutdown{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_signaling_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shutdown) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shutdown) ProtoMessage() {}

func (x *Shutdown) ProtoReflect() protoreflect.Message {
	mi := &file_protos_signaling_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shutdown.ProtoReflect.Descriptor instead.
func (*Shutdown) Descriptor() ([]byte, []int) {
	return file_protos_signaling_proto_rawDescGZIP(), []int{8}
}

type StreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*StreamResponse_Login
	//	*StreamResponse_Leave
	//	*StreamResponse_Message
	//	*StreamResponse_Shutdown
	Event isStreamResponse_Event `protobuf_oneof:"event"`
}

func (x *StreamResponse) Reset() {
	*x = StreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_signaling_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamResponse) ProtoMessage() {}

func (x *StreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_signaling_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamResponse.ProtoReflect.Descriptor instead.
func (*StreamResponse) Descriptor() ([]byte, []int) {
	return file_protos_signaling_proto_rawDescGZIP(), []int{9}
}

func (m *StreamResponse) GetEvent() isStreamResponse_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *StreamResponse) GetLogin() *Login {
	if x, ok := x.GetEvent().(*StreamResponse_Login); ok {
		return x.Login
	}
	return nil
}

func (x *StreamResponse) GetLeave() *Leave {
	if x, ok := x.GetEvent().(*StreamResponse_Leave); ok {
		return x.Leave
	}
	return nil
}

func (x *StreamResponse) GetMessage() *Message {
	if x, ok := x.GetEvent().(*StreamResponse_Message); ok {
		return x.Message
	}
	return nil
}

func (x *StreamResponse) GetShutdown() *Shutdown {
	if x, ok := x.GetEvent().(*StreamResponse_Shutdown); ok {
		return x.Shutdown
	}
	return nil
}

type isStreamResponse_Event interface {
	isStreamResponse_Event()
}

type StreamResponse_Login struct {
	Login *Login `protobuf:"bytes,1,opt,name=login,proto3,oneof"`
}

type StreamResponse_Leave struct {
	Leave *Leave `protobuf:"bytes,2,opt,name=leave,proto3,oneof"`
}

type StreamResponse_Message struct {
	Message *Message `protobuf:"bytes,3,opt,name=message,proto3,oneof"`
}

type StreamResponse_Shutdown struct {
	Shutdown *Shutdown `protobuf:"bytes,4,opt,name=shutdown,proto3,oneof"`
}

func (*StreamResponse_Login) isStreamResponse_Event() {}

func (*StreamResponse_Leave) isStreamResponse_Event() {}

func (*StreamResponse_Message) isStreamResponse_Event() {}

func (*StreamResponse_Shutdown) isStreamResponse_Event() {}

var File_protos_signaling_proto protoreflect.FileDescriptor

var file_protos_signaling_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x69, 0x6e, 0x67, 0x22, 0x35, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x0c, 0x4a, 0x6f,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x36,
	0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x37, 0x0a, 0x0d, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x1b, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x1b, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x37,
	0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x53, 0x68, 0x75, 0x74, 0x64,
	0x6f, 0x77, 0x6e, 0x22, 0xd0, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x28, 0x0a, 0x05, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x65, 0x61, 0x76,
	0x65, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48,
	0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x08, 0x73, 0x68,
	0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77,
	0x6e, 0x48, 0x00, 0x52, 0x08, 0x73, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x42, 0x07, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32, 0xc9, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x16, 0x2e, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x05, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x12, 0x17, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x4c, 0x65,
	0x61, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a,
	0x06, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x18, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_signaling_proto_rawDescOnce sync.Once
	file_protos_signaling_proto_rawDescData = file_protos_signaling_proto_rawDesc
)

func file_protos_signaling_proto_rawDescGZIP() []byte {
	file_protos_signaling_proto_rawDescOnce.Do(func() {
		file_protos_signaling_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_signaling_proto_rawDescData)
	})
	return file_protos_signaling_proto_rawDescData
}

var file_protos_signaling_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_protos_signaling_proto_goTypes = []interface{}{
	(*JoinRequest)(nil),    // 0: signaling.JoinRequest
	(*JoinResponse)(nil),   // 1: signaling.JoinResponse
	(*LeaveRequest)(nil),   // 2: signaling.LeaveRequest
	(*LeaveResponse)(nil),  // 3: signaling.LeaveResponse
	(*StreamRequest)(nil),  // 4: signaling.StreamRequest
	(*Login)(nil),          // 5: signaling.Login
	(*Leave)(nil),          // 6: signaling.Leave
	(*Message)(nil),        // 7: signaling.Message
	(*Shutdown)(nil),       // 8: signaling.Shutdown
	(*StreamResponse)(nil), // 9: signaling.StreamResponse
}
var file_protos_signaling_proto_depIdxs = []int32{
	5, // 0: signaling.StreamResponse.login:type_name -> signaling.Login
	6, // 1: signaling.StreamResponse.leave:type_name -> signaling.Leave
	7, // 2: signaling.StreamResponse.message:type_name -> signaling.Message
	8, // 3: signaling.StreamResponse.shutdown:type_name -> signaling.Shutdown
	0, // 4: signaling.Signaling.Join:input_type -> signaling.JoinRequest
	2, // 5: signaling.Signaling.Leave:input_type -> signaling.LeaveRequest
	4, // 6: signaling.Signaling.Stream:input_type -> signaling.StreamRequest
	1, // 7: signaling.Signaling.Join:output_type -> signaling.JoinResponse
	3, // 8: signaling.Signaling.Leave:output_type -> signaling.LeaveResponse
	9, // 9: signaling.Signaling.Stream:output_type -> signaling.StreamResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_signaling_proto_init() }
func file_protos_signaling_proto_init() {
	if File_protos_signaling_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_signaling_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_protos_signaling_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinResponse); i {
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
		file_protos_signaling_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRequest); i {
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
		file_protos_signaling_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveResponse); i {
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
		file_protos_signaling_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamRequest); i {
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
		file_protos_signaling_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Login); i {
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
		file_protos_signaling_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Leave); i {
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
		file_protos_signaling_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_protos_signaling_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shutdown); i {
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
		file_protos_signaling_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamResponse); i {
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
	file_protos_signaling_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*StreamResponse_Login)(nil),
		(*StreamResponse_Leave)(nil),
		(*StreamResponse_Message)(nil),
		(*StreamResponse_Shutdown)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_signaling_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_signaling_proto_goTypes,
		DependencyIndexes: file_protos_signaling_proto_depIdxs,
		MessageInfos:      file_protos_signaling_proto_msgTypes,
	}.Build()
	File_protos_signaling_proto = out.File
	file_protos_signaling_proto_rawDesc = nil
	file_protos_signaling_proto_goTypes = nil
	file_protos_signaling_proto_depIdxs = nil
}