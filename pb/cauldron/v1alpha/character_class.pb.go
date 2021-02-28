// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: cauldron/v1alpha/character_class.proto

package cauldron

import (
	proto "github.com/golang/protobuf/proto"
	v1alpha "github.com/nicklaw5/protos/pb/common/types/v1alpha"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CharacterClass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique character class id.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The class name (eg. Warrior, Mage, Rogue, etc).
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// The class description.
	Description string                 `protobuf:"bytes,20,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,30,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   *timestamppb.Timestamp `protobuf:"bytes,40,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *CharacterClass) Reset() {
	*x = CharacterClass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cauldron_v1alpha_character_class_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CharacterClass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CharacterClass) ProtoMessage() {}

func (x *CharacterClass) ProtoReflect() protoreflect.Message {
	mi := &file_cauldron_v1alpha_character_class_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CharacterClass.ProtoReflect.Descriptor instead.
func (*CharacterClass) Descriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_character_class_proto_rawDescGZIP(), []int{0}
}

func (x *CharacterClass) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CharacterClass) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CharacterClass) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CharacterClass) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *CharacterClass) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateCharacterClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,20,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateCharacterClassRequest) Reset() {
	*x = CreateCharacterClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cauldron_v1alpha_character_class_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCharacterClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCharacterClassRequest) ProtoMessage() {}

func (x *CreateCharacterClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cauldron_v1alpha_character_class_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCharacterClassRequest.ProtoReflect.Descriptor instead.
func (*CreateCharacterClassRequest) Descriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_character_class_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCharacterClassRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateCharacterClassRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCharacterClassRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateCharacterClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       v1alpha.Status  `protobuf:"varint,1,opt,name=status,proto3,enum=nicklaw5.common.types.v1alpha.Status" json:"status,omitempty"`
	ErrorCode    int32           `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string          `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Class        *CharacterClass `protobuf:"bytes,10,opt,name=class,proto3" json:"class,omitempty"`
}

func (x *CreateCharacterClassResponse) Reset() {
	*x = CreateCharacterClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cauldron_v1alpha_character_class_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCharacterClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCharacterClassResponse) ProtoMessage() {}

func (x *CreateCharacterClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cauldron_v1alpha_character_class_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCharacterClassResponse.ProtoReflect.Descriptor instead.
func (*CreateCharacterClassResponse) Descriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_character_class_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCharacterClassResponse) GetStatus() v1alpha.Status {
	if x != nil {
		return x.Status
	}
	return v1alpha.Status_OK
}

func (x *CreateCharacterClassResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *CreateCharacterClassResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *CreateCharacterClassResponse) GetClass() *CharacterClass {
	if x != nil {
		return x.Class
	}
	return nil
}

type GetCharacterClassByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,10,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCharacterClassByIDRequest) Reset() {
	*x = GetCharacterClassByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cauldron_v1alpha_character_class_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCharacterClassByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCharacterClassByIDRequest) ProtoMessage() {}

func (x *GetCharacterClassByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cauldron_v1alpha_character_class_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCharacterClassByIDRequest.ProtoReflect.Descriptor instead.
func (*GetCharacterClassByIDRequest) Descriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_character_class_proto_rawDescGZIP(), []int{3}
}

func (x *GetCharacterClassByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCharacterClassByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       v1alpha.Status  `protobuf:"varint,1,opt,name=status,proto3,enum=nicklaw5.common.types.v1alpha.Status" json:"status,omitempty"`
	ErrorCode    int32           `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage string          `protobuf:"bytes,3,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	Class        *CharacterClass `protobuf:"bytes,10,opt,name=class,proto3" json:"class,omitempty"`
}

func (x *GetCharacterClassByIDResponse) Reset() {
	*x = GetCharacterClassByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cauldron_v1alpha_character_class_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCharacterClassByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCharacterClassByIDResponse) ProtoMessage() {}

func (x *GetCharacterClassByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cauldron_v1alpha_character_class_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCharacterClassByIDResponse.ProtoReflect.Descriptor instead.
func (*GetCharacterClassByIDResponse) Descriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_character_class_proto_rawDescGZIP(), []int{4}
}

func (x *GetCharacterClassByIDResponse) GetStatus() v1alpha.Status {
	if x != nil {
		return x.Status
	}
	return v1alpha.Status_OK
}

func (x *GetCharacterClassByIDResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *GetCharacterClassByIDResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *GetCharacterClassByIDResponse) GetClass() *CharacterClass {
	if x != nil {
		return x.Class
	}
	return nil
}

var File_cauldron_v1alpha_character_class_proto protoreflect.FileDescriptor

var file_cauldron_v1alpha_character_class_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61,
	0x77, 0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x1a, 0x21, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x63, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe2, 0x01, 0x0a, 0x1c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6e,
	0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x3f, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x22, 0x2e, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xe3, 0x01, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x25, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e,
	0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52,
	0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x32, 0xb2, 0x02, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x89, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x36, 0x2e, 0x6e, 0x69, 0x63, 0x6b,
	0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x37, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61, 0x75,
	0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8c, 0x01, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x42, 0x79, 0x49, 0x44, 0x12, 0x37, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77,
	0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x38, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64,
	0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61,
	0x77, 0x35, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x75,
	0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x63, 0x61,
	0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cauldron_v1alpha_character_class_proto_rawDescOnce sync.Once
	file_cauldron_v1alpha_character_class_proto_rawDescData = file_cauldron_v1alpha_character_class_proto_rawDesc
)

func file_cauldron_v1alpha_character_class_proto_rawDescGZIP() []byte {
	file_cauldron_v1alpha_character_class_proto_rawDescOnce.Do(func() {
		file_cauldron_v1alpha_character_class_proto_rawDescData = protoimpl.X.CompressGZIP(file_cauldron_v1alpha_character_class_proto_rawDescData)
	})
	return file_cauldron_v1alpha_character_class_proto_rawDescData
}

var file_cauldron_v1alpha_character_class_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cauldron_v1alpha_character_class_proto_goTypes = []interface{}{
	(*CharacterClass)(nil),                // 0: nicklaw5.cauldron.v1alpha.CharacterClass
	(*CreateCharacterClassRequest)(nil),   // 1: nicklaw5.cauldron.v1alpha.CreateCharacterClassRequest
	(*CreateCharacterClassResponse)(nil),  // 2: nicklaw5.cauldron.v1alpha.CreateCharacterClassResponse
	(*GetCharacterClassByIDRequest)(nil),  // 3: nicklaw5.cauldron.v1alpha.GetCharacterClassByIDRequest
	(*GetCharacterClassByIDResponse)(nil), // 4: nicklaw5.cauldron.v1alpha.GetCharacterClassByIDResponse
	(*timestamppb.Timestamp)(nil),         // 5: google.protobuf.Timestamp
	(v1alpha.Status)(0),                   // 6: nicklaw5.common.types.v1alpha.Status
}
var file_cauldron_v1alpha_character_class_proto_depIdxs = []int32{
	5, // 0: nicklaw5.cauldron.v1alpha.CharacterClass.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: nicklaw5.cauldron.v1alpha.CharacterClass.updated_at:type_name -> google.protobuf.Timestamp
	6, // 2: nicklaw5.cauldron.v1alpha.CreateCharacterClassResponse.status:type_name -> nicklaw5.common.types.v1alpha.Status
	0, // 3: nicklaw5.cauldron.v1alpha.CreateCharacterClassResponse.class:type_name -> nicklaw5.cauldron.v1alpha.CharacterClass
	6, // 4: nicklaw5.cauldron.v1alpha.GetCharacterClassByIDResponse.status:type_name -> nicklaw5.common.types.v1alpha.Status
	0, // 5: nicklaw5.cauldron.v1alpha.GetCharacterClassByIDResponse.class:type_name -> nicklaw5.cauldron.v1alpha.CharacterClass
	1, // 6: nicklaw5.cauldron.v1alpha.CharacterClassService.CreateCharacterClass:input_type -> nicklaw5.cauldron.v1alpha.CreateCharacterClassRequest
	3, // 7: nicklaw5.cauldron.v1alpha.CharacterClassService.GetCharacterClassByID:input_type -> nicklaw5.cauldron.v1alpha.GetCharacterClassByIDRequest
	2, // 8: nicklaw5.cauldron.v1alpha.CharacterClassService.CreateCharacterClass:output_type -> nicklaw5.cauldron.v1alpha.CreateCharacterClassResponse
	4, // 9: nicklaw5.cauldron.v1alpha.CharacterClassService.GetCharacterClassByID:output_type -> nicklaw5.cauldron.v1alpha.GetCharacterClassByIDResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cauldron_v1alpha_character_class_proto_init() }
func file_cauldron_v1alpha_character_class_proto_init() {
	if File_cauldron_v1alpha_character_class_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cauldron_v1alpha_character_class_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CharacterClass); i {
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
		file_cauldron_v1alpha_character_class_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCharacterClassRequest); i {
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
		file_cauldron_v1alpha_character_class_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCharacterClassResponse); i {
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
		file_cauldron_v1alpha_character_class_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCharacterClassByIDRequest); i {
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
		file_cauldron_v1alpha_character_class_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCharacterClassByIDResponse); i {
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
			RawDescriptor: file_cauldron_v1alpha_character_class_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cauldron_v1alpha_character_class_proto_goTypes,
		DependencyIndexes: file_cauldron_v1alpha_character_class_proto_depIdxs,
		MessageInfos:      file_cauldron_v1alpha_character_class_proto_msgTypes,
	}.Build()
	File_cauldron_v1alpha_character_class_proto = out.File
	file_cauldron_v1alpha_character_class_proto_rawDesc = nil
	file_cauldron_v1alpha_character_class_proto_goTypes = nil
	file_cauldron_v1alpha_character_class_proto_depIdxs = nil
}
