// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: api/v1/gopackage_name__/pb_name__.proto

package gopackage_name__

import (
	common "app_module__/pkg/proto/common"
	gopackage_name__ "app_module__/pkg/proto/gopackage_name__"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetPbName__Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` //ApiProtoParentIDs__
}

func (x *GetPbName__Request) Reset() {
	*x = GetPbName__Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPbName__Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPbName__Request) ProtoMessage() {}

func (x *GetPbName__Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPbName__Request.ProtoReflect.Descriptor instead.
func (*GetPbName__Request) Descriptor() ([]byte, []int) {
	return file_api_v1_gopackage_name___pb_name___proto_rawDescGZIP(), []int{0}
}

func (x *GetPbName__Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPbName__Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PbName__ *gopackage_name__.MessageTypeName__ `protobuf:"bytes,1,opt,name=pb_name__,json=pbName,proto3" json:"pb_name__,omitempty"`
}

func (x *GetPbName__Response) Reset() {
	*x = GetPbName__Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPbName__Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPbName__Response) ProtoMessage() {}

func (x *GetPbName__Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPbName__Response.ProtoReflect.Descriptor instead.
func (*GetPbName__Response) Descriptor() ([]byte, []int) {
	return file_api_v1_gopackage_name___pb_name___proto_rawDescGZIP(), []int{1}
}

func (x *GetPbName__Response) GetPbName__() *gopackage_name__.MessageTypeName__ {
	if x != nil {
		return x.PbName__
	}
	return nil
}

type ListPbNamePlural__Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *common.PaginationParam `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
	Ids        []string                `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"` //ApiProtoParentIDs__
}

func (x *ListPbNamePlural__Request) Reset() {
	*x = ListPbNamePlural__Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPbNamePlural__Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPbNamePlural__Request) ProtoMessage() {}

func (x *ListPbNamePlural__Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPbNamePlural__Request.ProtoReflect.Descriptor instead.
func (*ListPbNamePlural__Request) Descriptor() ([]byte, []int) {
	return file_api_v1_gopackage_name___pb_name___proto_rawDescGZIP(), []int{2}
}

func (x *ListPbNamePlural__Request) GetPagination() *common.PaginationParam {
	if x != nil {
		return x.Pagination
	}
	return nil
}

func (x *ListPbNamePlural__Request) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type ListPbNamePlural__Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PbNamePlural__ *gopackage_name__.MessageTypeNamePlural__ `protobuf:"bytes,1,opt,name=pb_name_plural__,json=pbNamePlural,proto3" json:"pb_name_plural__,omitempty"`
}

func (x *ListPbNamePlural__Response) Reset() {
	*x = ListPbNamePlural__Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPbNamePlural__Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPbNamePlural__Response) ProtoMessage() {}

func (x *ListPbNamePlural__Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPbNamePlural__Response.ProtoReflect.Descriptor instead.
func (*ListPbNamePlural__Response) Descriptor() ([]byte, []int) {
	return file_api_v1_gopackage_name___pb_name___proto_rawDescGZIP(), []int{3}
}

func (x *ListPbNamePlural__Response) GetPbNamePlural__() *gopackage_name__.MessageTypeNamePlural__ {
	if x != nil {
		return x.PbNamePlural__
	}
	return nil
}

type CreatPbName__Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PbName__ *gopackage_name__.MessageTypeName__ `protobuf:"bytes,1,opt,name=pb_name__,json=pbName,proto3" json:"pb_name__,omitempty"` //ApiProtoParentIDs__
}

func (x *CreatPbName__Request) Reset() {
	*x = CreatPbName__Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatPbName__Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatPbName__Request) ProtoMessage() {}

func (x *CreatPbName__Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatPbName__Request.ProtoReflect.Descriptor instead.
func (*CreatPbName__Request) Descriptor() ([]byte, []int) {
	return file_api_v1_gopackage_name___pb_name___proto_rawDescGZIP(), []int{4}
}

func (x *CreatPbName__Request) GetPbName__() *gopackage_name__.MessageTypeName__ {
	if x != nil {
		return x.PbName__
	}
	return nil
}

type CreatPbName__Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatPbName__Response) Reset() {
	*x = CreatPbName__Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatPbName__Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatPbName__Response) ProtoMessage() {}

func (x *CreatPbName__Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatPbName__Response.ProtoReflect.Descriptor instead.
func (*CreatPbName__Response) Descriptor() ([]byte, []int) {
	return file_api_v1_gopackage_name___pb_name___proto_rawDescGZIP(), []int{5}
}

func (x *CreatPbName__Response) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdatePbName__Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string                              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PbName__ *gopackage_name__.MessageTypeName__ `protobuf:"bytes,2,opt,name=pb_name__,json=pbName,proto3" json:"pb_name__,omitempty"` //ApiProtoParentIDs__
}

func (x *UpdatePbName__Request) Reset() {
	*x = UpdatePbName__Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePbName__Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePbName__Request) ProtoMessage() {}

func (x *UpdatePbName__Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePbName__Request.ProtoReflect.Descriptor instead.
func (*UpdatePbName__Request) Descriptor() ([]byte, []int) {
	return file_api_v1_gopackage_name___pb_name___proto_rawDescGZIP(), []int{6}
}

func (x *UpdatePbName__Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePbName__Request) GetPbName__() *gopackage_name__.MessageTypeName__ {
	if x != nil {
		return x.PbName__
	}
	return nil
}

type UpdatePbName__Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Updated int64 `protobuf:"varint,1,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *UpdatePbName__Response) Reset() {
	*x = UpdatePbName__Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePbName__Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePbName__Response) ProtoMessage() {}

func (x *UpdatePbName__Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePbName__Response.ProtoReflect.Descriptor instead.
func (*UpdatePbName__Response) Descriptor() ([]byte, []int) {
	return file_api_v1_gopackage_name___pb_name___proto_rawDescGZIP(), []int{7}
}

func (x *UpdatePbName__Response) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type DeletePbName__Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Ids []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"` //ApiProtoParentIDs__
}

func (x *DeletePbName__Request) Reset() {
	*x = DeletePbName__Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePbName__Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePbName__Request) ProtoMessage() {}

func (x *DeletePbName__Request) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePbName__Request.ProtoReflect.Descriptor instead.
func (*DeletePbName__Request) Descriptor() ([]byte, []int) {
	return file_api_v1_gopackage_name___pb_name___proto_rawDescGZIP(), []int{8}
}

func (x *DeletePbName__Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeletePbName__Request) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type DeletePbName__Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deleted int64 `protobuf:"varint,1,opt,name=deleted,proto3" json:"deleted,omitempty"`
}

func (x *DeletePbName__Response) Reset() {
	*x = DeletePbName__Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePbName__Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePbName__Response) ProtoMessage() {}

func (x *DeletePbName__Response) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_gopackage_name___pb_name___proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePbName__Response.ProtoReflect.Descriptor instead.
func (*DeletePbName__Response) Descriptor() ([]byte, []int) {
	return file_api_v1_gopackage_name___pb_name___proto_rawDescGZIP(), []int{9}
}

func (x *DeletePbName__Response) GetDeleted() int64 {
	if x != nil {
		return x.Deleted
	}
	return 0
}

var File_api_v1_gopackage_name___pb_name___proto protoreflect.FileDescriptor

var file_api_v1_gopackage_name___pb_name___proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x2f, 0x70, 0x62, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x5f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x70, 0x6b, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x5f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x5f, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x5f, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65,
	0x5f, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x70, 0x62,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x06, 0x70, 0x62,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7a, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x62, 0x4e, 0x61,
	0x6d, 0x65, 0x50, 0x6c, 0x75, 0x72, 0x61, 0x6c, 0x5f, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x4b, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0x79, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x6c,
	0x75, 0x72, 0x61, 0x6c, 0x5f, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b,
	0x0a, 0x10, 0x70, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x75, 0x72, 0x61, 0x6c,
	0x5f, 0x5f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x75, 0x72, 0x61, 0x6c, 0x5f, 0x5f, 0x52, 0x0c, 0x70,
	0x62, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x6c, 0x75, 0x72, 0x61, 0x6c, 0x22, 0x60, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x09, 0x70, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x5f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x5f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x06, 0x70, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a,
	0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x71, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x48, 0x0a, 0x09, 0x70, 0x62, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67,
	0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x5f,
	0x5f, 0x52, 0x06, 0x70, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x16, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x39, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x32, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x32, 0xd4, 0x06, 0x0a,
	0x0f, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xa2, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x38, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x5f, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x39, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67,
	0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x2e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x5f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x13, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f,
	0x75, 0x72, 0x69, 0x5f, 0x5f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x62, 0x09, 0x70, 0x62, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x12, 0xb3, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3f,
	0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x2e, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x5f, 0x5f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65,
	0x50, 0x6c, 0x75, 0x72, 0x61, 0x6c, 0x5f, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x40, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x2e, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x5f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x62, 0x4e, 0x61, 0x6d,
	0x65, 0x50, 0x6c, 0x75, 0x72, 0x61, 0x6c, 0x5f, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x0e, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x69, 0x5f, 0x5f, 0x62, 0x10, 0x70, 0x62, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x5f, 0x70, 0x6c, 0x75, 0x72, 0x61, 0x6c, 0x5f, 0x5f, 0x12, 0x9c, 0x01, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x5f, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x5f, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3b, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67,
	0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x2e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x5f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x50,
	0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x75, 0x72, 0x69, 0x5f, 0x5f, 0x3a, 0x01, 0x2a, 0x12, 0xa3, 0x01, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3b, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x67, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x5f, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x5f, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67,
	0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x2e,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x5f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x1a, 0x13, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x5f, 0x75, 0x72, 0x69, 0x5f, 0x5f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a,
	0x12, 0xa0, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x3b, 0x2e, 0x70, 0x6b,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x5f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x2e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x5f, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x62, 0x4e, 0x61, 0x6d, 0x65, 0x5f, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x75, 0x72, 0x69, 0x5f, 0x5f, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x42, 0x2e, 0x5a, 0x2c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x5f, 0x3b, 0x67, 0x6f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x5f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_v1_gopackage_name___pb_name___proto_rawDescOnce sync.Once
	file_api_v1_gopackage_name___pb_name___proto_rawDescData = file_api_v1_gopackage_name___pb_name___proto_rawDesc
)

func file_api_v1_gopackage_name___pb_name___proto_rawDescGZIP() []byte {
	file_api_v1_gopackage_name___pb_name___proto_rawDescOnce.Do(func() {
		file_api_v1_gopackage_name___pb_name___proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_gopackage_name___pb_name___proto_rawDescData)
	})
	return file_api_v1_gopackage_name___pb_name___proto_rawDescData
}

var file_api_v1_gopackage_name___pb_name___proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_v1_gopackage_name___pb_name___proto_goTypes = []interface{}{
	(*GetPbName__Request)(nil),                       // 0: pkg.proto.gopackage_name__.version__.GetPbName__Request
	(*GetPbName__Response)(nil),                      // 1: pkg.proto.gopackage_name__.version__.GetPbName__Response
	(*ListPbNamePlural__Request)(nil),                // 2: pkg.proto.gopackage_name__.version__.ListPbNamePlural__Request
	(*ListPbNamePlural__Response)(nil),               // 3: pkg.proto.gopackage_name__.version__.ListPbNamePlural__Response
	(*CreatPbName__Request)(nil),                     // 4: pkg.proto.gopackage_name__.version__.CreatPbName__Request
	(*CreatPbName__Response)(nil),                    // 5: pkg.proto.gopackage_name__.version__.CreatPbName__Response
	(*UpdatePbName__Request)(nil),                    // 6: pkg.proto.gopackage_name__.version__.UpdatePbName__Request
	(*UpdatePbName__Response)(nil),                   // 7: pkg.proto.gopackage_name__.version__.UpdatePbName__Response
	(*DeletePbName__Request)(nil),                    // 8: pkg.proto.gopackage_name__.version__.DeletePbName__Request
	(*DeletePbName__Response)(nil),                   // 9: pkg.proto.gopackage_name__.version__.DeletePbName__Response
	(*gopackage_name__.MessageTypeName__)(nil),       // 10: pkg.proto.gopackage_name__.MessageTypeName__
	(*common.PaginationParam)(nil),                   // 11: pkg.proto.common.PaginationParam
	(*gopackage_name__.MessageTypeNamePlural__)(nil), // 12: pkg.proto.gopackage_name__.MessageTypeNamePlural__
}
var file_api_v1_gopackage_name___pb_name___proto_depIdxs = []int32{
	10, // 0: pkg.proto.gopackage_name__.version__.GetPbName__Response.pb_name__:type_name -> pkg.proto.gopackage_name__.MessageTypeName__
	11, // 1: pkg.proto.gopackage_name__.version__.ListPbNamePlural__Request.pagination:type_name -> pkg.proto.common.PaginationParam
	12, // 2: pkg.proto.gopackage_name__.version__.ListPbNamePlural__Response.pb_name_plural__:type_name -> pkg.proto.gopackage_name__.MessageTypeNamePlural__
	10, // 3: pkg.proto.gopackage_name__.version__.CreatPbName__Request.pb_name__:type_name -> pkg.proto.gopackage_name__.MessageTypeName__
	10, // 4: pkg.proto.gopackage_name__.version__.UpdatePbName__Request.pb_name__:type_name -> pkg.proto.gopackage_name__.MessageTypeName__
	0,  // 5: pkg.proto.gopackage_name__.version__.PbName__Service.Get:input_type -> pkg.proto.gopackage_name__.version__.GetPbName__Request
	2,  // 6: pkg.proto.gopackage_name__.version__.PbName__Service.List:input_type -> pkg.proto.gopackage_name__.version__.ListPbNamePlural__Request
	4,  // 7: pkg.proto.gopackage_name__.version__.PbName__Service.Create:input_type -> pkg.proto.gopackage_name__.version__.CreatPbName__Request
	6,  // 8: pkg.proto.gopackage_name__.version__.PbName__Service.Update:input_type -> pkg.proto.gopackage_name__.version__.UpdatePbName__Request
	8,  // 9: pkg.proto.gopackage_name__.version__.PbName__Service.Delete:input_type -> pkg.proto.gopackage_name__.version__.DeletePbName__Request
	1,  // 10: pkg.proto.gopackage_name__.version__.PbName__Service.Get:output_type -> pkg.proto.gopackage_name__.version__.GetPbName__Response
	3,  // 11: pkg.proto.gopackage_name__.version__.PbName__Service.List:output_type -> pkg.proto.gopackage_name__.version__.ListPbNamePlural__Response
	5,  // 12: pkg.proto.gopackage_name__.version__.PbName__Service.Create:output_type -> pkg.proto.gopackage_name__.version__.CreatPbName__Response
	7,  // 13: pkg.proto.gopackage_name__.version__.PbName__Service.Update:output_type -> pkg.proto.gopackage_name__.version__.UpdatePbName__Response
	9,  // 14: pkg.proto.gopackage_name__.version__.PbName__Service.Delete:output_type -> pkg.proto.gopackage_name__.version__.DeletePbName__Response
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_v1_gopackage_name___pb_name___proto_init() }
func file_api_v1_gopackage_name___pb_name___proto_init() {
	if File_api_v1_gopackage_name___pb_name___proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_gopackage_name___pb_name___proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPbName__Request); i {
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
		file_api_v1_gopackage_name___pb_name___proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPbName__Response); i {
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
		file_api_v1_gopackage_name___pb_name___proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPbNamePlural__Request); i {
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
		file_api_v1_gopackage_name___pb_name___proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPbNamePlural__Response); i {
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
		file_api_v1_gopackage_name___pb_name___proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatPbName__Request); i {
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
		file_api_v1_gopackage_name___pb_name___proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatPbName__Response); i {
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
		file_api_v1_gopackage_name___pb_name___proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePbName__Request); i {
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
		file_api_v1_gopackage_name___pb_name___proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePbName__Response); i {
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
		file_api_v1_gopackage_name___pb_name___proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePbName__Request); i {
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
		file_api_v1_gopackage_name___pb_name___proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePbName__Response); i {
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
			RawDescriptor: file_api_v1_gopackage_name___pb_name___proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_gopackage_name___pb_name___proto_goTypes,
		DependencyIndexes: file_api_v1_gopackage_name___pb_name___proto_depIdxs,
		MessageInfos:      file_api_v1_gopackage_name___pb_name___proto_msgTypes,
	}.Build()
	File_api_v1_gopackage_name___pb_name___proto = out.File
	file_api_v1_gopackage_name___pb_name___proto_rawDesc = nil
	file_api_v1_gopackage_name___pb_name___proto_goTypes = nil
	file_api_v1_gopackage_name___pb_name___proto_depIdxs = nil
}
