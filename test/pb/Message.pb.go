// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: Message.proto

package Message

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Pb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GamerLoginC2S      *Pb_GamerLoginC2S      `protobuf:"bytes,1,opt,name=gamerLoginC2S" json:"gamerLoginC2S,omitempty"`
	GamerLoginS2C      *Pb_GamerLoginS2C      `protobuf:"bytes,2,opt,name=gamerLoginS2C" json:"gamerLoginS2C,omitempty"`
	GamerGlobalChatC2S *Pb_GamerGlobalChatC2S `protobuf:"bytes,3,opt,name=gamerGlobalChatC2S" json:"gamerGlobalChatC2S,omitempty"`
	GamerGlobalChatS2C *Pb_GamerGlobalChatS2C `protobuf:"bytes,4,opt,name=gamerGlobalChatS2C" json:"gamerGlobalChatS2C,omitempty"`
	GamerChatC2S       *Pb_GamerChatC2S       `protobuf:"bytes,5,opt,name=gamerChatC2S" json:"gamerChatC2S,omitempty"`
	GamerChatS2C       *Pb_GamerChatS2C       `protobuf:"bytes,6,opt,name=gamerChatS2C" json:"gamerChatS2C,omitempty"`
}

func (x *Pb) Reset() {
	*x = Pb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb) ProtoMessage() {}

func (x *Pb) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb.ProtoReflect.Descriptor instead.
func (*Pb) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0}
}

func (x *Pb) GetGamerLoginC2S() *Pb_GamerLoginC2S {
	if x != nil {
		return x.GamerLoginC2S
	}
	return nil
}

func (x *Pb) GetGamerLoginS2C() *Pb_GamerLoginS2C {
	if x != nil {
		return x.GamerLoginS2C
	}
	return nil
}

func (x *Pb) GetGamerGlobalChatC2S() *Pb_GamerGlobalChatC2S {
	if x != nil {
		return x.GamerGlobalChatC2S
	}
	return nil
}

func (x *Pb) GetGamerGlobalChatS2C() *Pb_GamerGlobalChatS2C {
	if x != nil {
		return x.GamerGlobalChatS2C
	}
	return nil
}

func (x *Pb) GetGamerChatC2S() *Pb_GamerChatC2S {
	if x != nil {
		return x.GamerChatC2S
	}
	return nil
}

func (x *Pb) GetGamerChatS2C() *Pb_GamerChatS2C {
	if x != nil {
		return x.GamerChatS2C
	}
	return nil
}

type Pb_GamerLoginC2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"` //Íæ¼Òid
}

func (x *Pb_GamerLoginC2S) Reset() {
	*x = Pb_GamerLoginC2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pb_GamerLoginC2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb_GamerLoginC2S) ProtoMessage() {}

func (x *Pb_GamerLoginC2S) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb_GamerLoginC2S.ProtoReflect.Descriptor instead.
func (*Pb_GamerLoginC2S) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Pb_GamerLoginC2S) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

type Pb_GamerLoginS2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json *string `protobuf:"bytes,1,opt,name=json" json:"json,omitempty"` //Êý¾Ý
}

func (x *Pb_GamerLoginS2C) Reset() {
	*x = Pb_GamerLoginS2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pb_GamerLoginS2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb_GamerLoginS2C) ProtoMessage() {}

func (x *Pb_GamerLoginS2C) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb_GamerLoginS2C.ProtoReflect.Descriptor instead.
func (*Pb_GamerLoginS2C) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Pb_GamerLoginS2C) GetJson() string {
	if x != nil && x.Json != nil {
		return *x.Json
	}
	return ""
}

type Pb_GamerGlobalChatC2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"` //
}

func (x *Pb_GamerGlobalChatC2S) Reset() {
	*x = Pb_GamerGlobalChatC2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pb_GamerGlobalChatC2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb_GamerGlobalChatC2S) ProtoMessage() {}

func (x *Pb_GamerGlobalChatC2S) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb_GamerGlobalChatC2S.ProtoReflect.Descriptor instead.
func (*Pb_GamerGlobalChatC2S) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Pb_GamerGlobalChatC2S) GetData() string {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return ""
}

type Pb_GamerGlobalChatS2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"` //
}

func (x *Pb_GamerGlobalChatS2C) Reset() {
	*x = Pb_GamerGlobalChatS2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pb_GamerGlobalChatS2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb_GamerGlobalChatS2C) ProtoMessage() {}

func (x *Pb_GamerGlobalChatS2C) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb_GamerGlobalChatS2C.ProtoReflect.Descriptor instead.
func (*Pb_GamerGlobalChatS2C) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Pb_GamerGlobalChatS2C) GetData() string {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return ""
}

type Pb_GamerChatC2S struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"` //
	Id   *string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
}

func (x *Pb_GamerChatC2S) Reset() {
	*x = Pb_GamerChatC2S{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pb_GamerChatC2S) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb_GamerChatC2S) ProtoMessage() {}

func (x *Pb_GamerChatC2S) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb_GamerChatC2S.ProtoReflect.Descriptor instead.
func (*Pb_GamerChatC2S) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Pb_GamerChatC2S) GetData() string {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return ""
}

func (x *Pb_GamerChatC2S) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

type Pb_GamerChatS2C struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   *string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"` //
	FromId *string `protobuf:"bytes,2,opt,name=fromId" json:"fromId,omitempty"`
}

func (x *Pb_GamerChatS2C) Reset() {
	*x = Pb_GamerChatS2C{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pb_GamerChatS2C) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pb_GamerChatS2C) ProtoMessage() {}

func (x *Pb_GamerChatS2C) ProtoReflect() protoreflect.Message {
	mi := &file_Message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pb_GamerChatS2C.ProtoReflect.Descriptor instead.
func (*Pb_GamerChatS2C) Descriptor() ([]byte, []int) {
	return file_Message_proto_rawDescGZIP(), []int{0, 5}
}

func (x *Pb_GamerChatS2C) GetData() string {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return ""
}

func (x *Pb_GamerChatS2C) GetFromId() string {
	if x != nil && x.FromId != nil {
		return *x.FromId
	}
	return ""
}

var File_Message_proto protoreflect.FileDescriptor

var file_Message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfc, 0x04, 0x0a, 0x02, 0x70, 0x62, 0x12, 0x37, 0x0a, 0x0d, 0x67, 0x61, 0x6d, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x43, 0x32, 0x53, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x32, 0x53,
	0x52, 0x0d, 0x67, 0x61, 0x6d, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x32, 0x53, 0x12,
	0x37, 0x0a, 0x0d, 0x67, 0x61, 0x6d, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x32, 0x43,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x32, 0x43, 0x52, 0x0d, 0x67, 0x61, 0x6d, 0x65, 0x72,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x53, 0x32, 0x43, 0x12, 0x46, 0x0a, 0x12, 0x67, 0x61, 0x6d, 0x65,
	0x72, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x43, 0x32, 0x53, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x47,
	0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x43, 0x32, 0x53, 0x52, 0x12, 0x67, 0x61,
	0x6d, 0x65, 0x72, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x43, 0x32, 0x53,
	0x12, 0x46, 0x0a, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x72, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43,
	0x68, 0x61, 0x74, 0x53, 0x32, 0x43, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x68, 0x61,
	0x74, 0x53, 0x32, 0x43, 0x52, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x72, 0x47, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x43, 0x68, 0x61, 0x74, 0x53, 0x32, 0x43, 0x12, 0x34, 0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65,
	0x72, 0x43, 0x68, 0x61, 0x74, 0x43, 0x32, 0x53, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x43, 0x32, 0x53,
	0x52, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x43, 0x32, 0x53, 0x12, 0x34,
	0x0a, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x53, 0x32, 0x43, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x43,
	0x68, 0x61, 0x74, 0x53, 0x32, 0x43, 0x52, 0x0c, 0x67, 0x61, 0x6d, 0x65, 0x72, 0x43, 0x68, 0x61,
	0x74, 0x53, 0x32, 0x43, 0x1a, 0x1f, 0x0a, 0x0d, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x43, 0x32, 0x53, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x1a, 0x23, 0x0a, 0x0d, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x53, 0x32, 0x43, 0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x1a, 0x28, 0x0a, 0x12, 0x47, 0x61,
	0x6d, 0x65, 0x72, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x43, 0x32, 0x53,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x28, 0x0a, 0x12, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x47, 0x6c, 0x6f,
	0x62, 0x61, 0x6c, 0x43, 0x68, 0x61, 0x74, 0x53, 0x32, 0x43, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x32,
	0x0a, 0x0c, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x43, 0x32, 0x53, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x1a, 0x3a, 0x0a, 0x0c, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x53,
	0x32, 0x43, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x72, 0x6f, 0x6d, 0x49, 0x64,
}

var (
	file_Message_proto_rawDescOnce sync.Once
	file_Message_proto_rawDescData = file_Message_proto_rawDesc
)

func file_Message_proto_rawDescGZIP() []byte {
	file_Message_proto_rawDescOnce.Do(func() {
		file_Message_proto_rawDescData = protoimpl.X.CompressGZIP(file_Message_proto_rawDescData)
	})
	return file_Message_proto_rawDescData
}

var file_Message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_Message_proto_goTypes = []interface{}{
	(*Pb)(nil),                    // 0: pb
	(*Pb_GamerLoginC2S)(nil),      // 1: pb.GamerLoginC2S
	(*Pb_GamerLoginS2C)(nil),      // 2: pb.GamerLoginS2C
	(*Pb_GamerGlobalChatC2S)(nil), // 3: pb.GamerGlobalChatC2S
	(*Pb_GamerGlobalChatS2C)(nil), // 4: pb.GamerGlobalChatS2C
	(*Pb_GamerChatC2S)(nil),       // 5: pb.GamerChatC2S
	(*Pb_GamerChatS2C)(nil),       // 6: pb.GamerChatS2C
}
var file_Message_proto_depIdxs = []int32{
	1, // 0: pb.gamerLoginC2S:type_name -> pb.GamerLoginC2S
	2, // 1: pb.gamerLoginS2C:type_name -> pb.GamerLoginS2C
	3, // 2: pb.gamerGlobalChatC2S:type_name -> pb.GamerGlobalChatC2S
	4, // 3: pb.gamerGlobalChatS2C:type_name -> pb.GamerGlobalChatS2C
	5, // 4: pb.gamerChatC2S:type_name -> pb.GamerChatC2S
	6, // 5: pb.gamerChatS2C:type_name -> pb.GamerChatS2C
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_Message_proto_init() }
func file_Message_proto_init() {
	if File_Message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pb); i {
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
		file_Message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pb_GamerLoginC2S); i {
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
		file_Message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pb_GamerLoginS2C); i {
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
		file_Message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pb_GamerGlobalChatC2S); i {
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
		file_Message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pb_GamerGlobalChatS2C); i {
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
		file_Message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pb_GamerChatC2S); i {
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
		file_Message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pb_GamerChatS2C); i {
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
			RawDescriptor: file_Message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Message_proto_goTypes,
		DependencyIndexes: file_Message_proto_depIdxs,
		MessageInfos:      file_Message_proto_msgTypes,
	}.Build()
	File_Message_proto = out.File
	file_Message_proto_rawDesc = nil
	file_Message_proto_goTypes = nil
	file_Message_proto_depIdxs = nil
}
