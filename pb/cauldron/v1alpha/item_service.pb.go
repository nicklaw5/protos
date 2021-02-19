// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: cauldron/v1alpha/item_service.proto

package cauldron

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type Status int32

const (
	Status_OK              Status = 0
	Status_ERROR           Status = 1
	Status_UNAUTHENTICATED Status = 2
	Status_UNAUTHORIZED    Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "OK",
		1: "ERROR",
		2: "UNAUTHENTICATED",
		3: "UNAUTHORIZED",
	}
	Status_value = map[string]int32{
		"OK":              0,
		"ERROR":           1,
		"UNAUTHENTICATED": 2,
		"UNAUTHORIZED":    3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_cauldron_v1alpha_item_service_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_cauldron_v1alpha_item_service_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_item_service_proto_rawDescGZIP(), []int{0}
}

type GetItemByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
}

func (x *GetItemByIdRequest) Reset() {
	*x = GetItemByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cauldron_v1alpha_item_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemByIdRequest) ProtoMessage() {}

func (x *GetItemByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cauldron_v1alpha_item_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemByIdRequest.ProtoReflect.Descriptor instead.
func (*GetItemByIdRequest) Descriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_item_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetItemByIdRequest) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

type GetItemByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId string `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Item    *Item  `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *GetItemByIdResponse) Reset() {
	*x = GetItemByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cauldron_v1alpha_item_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemByIdResponse) ProtoMessage() {}

func (x *GetItemByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cauldron_v1alpha_item_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemByIdResponse.ProtoReflect.Descriptor instead.
func (*GetItemByIdResponse) Descriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_item_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetItemByIdResponse) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *GetItemByIdResponse) GetItem() *Item {
	if x != nil {
		return x.Item
	}
	return nil
}

type NewItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CatalogueId string           `protobuf:"bytes,1,opt,name=catalogue_id,json=catalogueId,proto3" json:"catalogue_id,omitempty"`
	Name        string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Attributes  *structpb.Struct `protobuf:"bytes,10,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *NewItemRequest) Reset() {
	*x = NewItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cauldron_v1alpha_item_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewItemRequest) ProtoMessage() {}

func (x *NewItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cauldron_v1alpha_item_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewItemRequest.ProtoReflect.Descriptor instead.
func (*NewItemRequest) Descriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_item_service_proto_rawDescGZIP(), []int{2}
}

func (x *NewItemRequest) GetCatalogueId() string {
	if x != nil {
		return x.CatalogueId
	}
	return ""
}

func (x *NewItemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewItemRequest) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type NewItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId      string `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	Status       Status `protobuf:"varint,2,opt,name=status,proto3,enum=nicklaw5.cauldron.v1alpha.Status" json:"status,omitempty"` // 0 = error || 1 = success
	ErrorCode    int32  `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`                // only if ok = 0
	ErrorMessage string `protobuf:"bytes,4,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`        // only if ok = 0
	Data         *Item  `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *NewItemResponse) Reset() {
	*x = NewItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cauldron_v1alpha_item_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewItemResponse) ProtoMessage() {}

func (x *NewItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cauldron_v1alpha_item_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewItemResponse.ProtoReflect.Descriptor instead.
func (*NewItemResponse) Descriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_item_service_proto_rawDescGZIP(), []int{3}
}

func (x *NewItemResponse) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *NewItemResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_OK
}

func (x *NewItemResponse) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *NewItemResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *NewItemResponse) GetData() *Item {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_cauldron_v1alpha_item_service_proto protoreflect.FileDescriptor

var file_cauldron_v1alpha_item_service_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e,
	0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x04,
	0x69, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x69, 0x63,
	0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65,
	0x6d, 0x22, 0x80, 0x01, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x75, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x22, 0xe0, 0x01, 0x0a, 0x0f, 0x4e, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63,
	0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c,
	0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x42, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x41, 0x55, 0x54, 0x48, 0x45, 0x4e,
	0x54, 0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x41,
	0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x45, 0x44, 0x10, 0x03, 0x32, 0xe1, 0x01, 0x0a, 0x0b,
	0x49, 0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49, 0x44, 0x12, 0x2d, 0x2e, 0x6e, 0x69, 0x63,
	0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6e, 0x69, 0x63, 0x6b,
	0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x07, 0x4e,
	0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x29, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77,
	0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x4e, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2a, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61, 0x75,
	0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4e, 0x65,
	0x77, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69,
	0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62,
	0x2f, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x3b, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cauldron_v1alpha_item_service_proto_rawDescOnce sync.Once
	file_cauldron_v1alpha_item_service_proto_rawDescData = file_cauldron_v1alpha_item_service_proto_rawDesc
)

func file_cauldron_v1alpha_item_service_proto_rawDescGZIP() []byte {
	file_cauldron_v1alpha_item_service_proto_rawDescOnce.Do(func() {
		file_cauldron_v1alpha_item_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cauldron_v1alpha_item_service_proto_rawDescData)
	})
	return file_cauldron_v1alpha_item_service_proto_rawDescData
}

var file_cauldron_v1alpha_item_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cauldron_v1alpha_item_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cauldron_v1alpha_item_service_proto_goTypes = []interface{}{
	(Status)(0),                 // 0: nicklaw5.cauldron.v1alpha.Status
	(*GetItemByIdRequest)(nil),  // 1: nicklaw5.cauldron.v1alpha.GetItemByIdRequest
	(*GetItemByIdResponse)(nil), // 2: nicklaw5.cauldron.v1alpha.GetItemByIdResponse
	(*NewItemRequest)(nil),      // 3: nicklaw5.cauldron.v1alpha.NewItemRequest
	(*NewItemResponse)(nil),     // 4: nicklaw5.cauldron.v1alpha.NewItemResponse
	(*Item)(nil),                // 5: nicklaw5.cauldron.v1alpha.Item
	(*structpb.Struct)(nil),     // 6: google.protobuf.Struct
}
var file_cauldron_v1alpha_item_service_proto_depIdxs = []int32{
	5, // 0: nicklaw5.cauldron.v1alpha.GetItemByIdResponse.item:type_name -> nicklaw5.cauldron.v1alpha.Item
	6, // 1: nicklaw5.cauldron.v1alpha.NewItemRequest.attributes:type_name -> google.protobuf.Struct
	0, // 2: nicklaw5.cauldron.v1alpha.NewItemResponse.status:type_name -> nicklaw5.cauldron.v1alpha.Status
	5, // 3: nicklaw5.cauldron.v1alpha.NewItemResponse.data:type_name -> nicklaw5.cauldron.v1alpha.Item
	1, // 4: nicklaw5.cauldron.v1alpha.ItemService.GetItemByID:input_type -> nicklaw5.cauldron.v1alpha.GetItemByIdRequest
	3, // 5: nicklaw5.cauldron.v1alpha.ItemService.NewItem:input_type -> nicklaw5.cauldron.v1alpha.NewItemRequest
	2, // 6: nicklaw5.cauldron.v1alpha.ItemService.GetItemByID:output_type -> nicklaw5.cauldron.v1alpha.GetItemByIdResponse
	4, // 7: nicklaw5.cauldron.v1alpha.ItemService.NewItem:output_type -> nicklaw5.cauldron.v1alpha.NewItemResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cauldron_v1alpha_item_service_proto_init() }
func file_cauldron_v1alpha_item_service_proto_init() {
	if File_cauldron_v1alpha_item_service_proto != nil {
		return
	}
	file_cauldron_v1alpha_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cauldron_v1alpha_item_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemByIdRequest); i {
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
		file_cauldron_v1alpha_item_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemByIdResponse); i {
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
		file_cauldron_v1alpha_item_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewItemRequest); i {
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
		file_cauldron_v1alpha_item_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewItemResponse); i {
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
			RawDescriptor: file_cauldron_v1alpha_item_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cauldron_v1alpha_item_service_proto_goTypes,
		DependencyIndexes: file_cauldron_v1alpha_item_service_proto_depIdxs,
		EnumInfos:         file_cauldron_v1alpha_item_service_proto_enumTypes,
		MessageInfos:      file_cauldron_v1alpha_item_service_proto_msgTypes,
	}.Build()
	File_cauldron_v1alpha_item_service_proto = out.File
	file_cauldron_v1alpha_item_service_proto_rawDesc = nil
	file_cauldron_v1alpha_item_service_proto_goTypes = nil
	file_cauldron_v1alpha_item_service_proto_depIdxs = nil
}
