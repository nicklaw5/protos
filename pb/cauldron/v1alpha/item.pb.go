// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: cauldron/v1alpha/item.proto

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

type ItemState int32

const (
	ItemState_UNDAMAGED ItemState = 0
	ItemState_DAMAGED   ItemState = 1
	ItemState_BROKEN    ItemState = 2
)

// Enum value maps for ItemState.
var (
	ItemState_name = map[int32]string{
		0: "UNDAMAGED",
		1: "DAMAGED",
		2: "BROKEN",
	}
	ItemState_value = map[string]int32{
		"UNDAMAGED": 0,
		"DAMAGED":   1,
		"BROKEN":    2,
	}
)

func (x ItemState) Enum() *ItemState {
	p := new(ItemState)
	*p = x
	return p
}

func (x ItemState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ItemState) Descriptor() protoreflect.EnumDescriptor {
	return file_cauldron_v1alpha_item_proto_enumTypes[0].Descriptor()
}

func (ItemState) Type() protoreflect.EnumType {
	return &file_cauldron_v1alpha_item_proto_enumTypes[0]
}

func (x ItemState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ItemState.Descriptor instead.
func (ItemState) EnumDescriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_item_proto_rawDescGZIP(), []int{0}
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                      `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Description string                      `protobuf:"bytes,20,opt,name=description,proto3" json:"description,omitempty"`
	State       ItemState                   `protobuf:"varint,30,opt,name=state,proto3,enum=nicklaw5.cauldron.v1alpha.ItemState" json:"state,omitempty"`
	Data        map[string]*structpb.Struct `protobuf:"bytes,40,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cauldron_v1alpha_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_cauldron_v1alpha_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_item_proto_rawDescGZIP(), []int{0}
}

func (x *Item) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Item) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Item) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Item) GetState() ItemState {
	if x != nil {
		return x.State
	}
	return ItemState_UNDAMAGED
}

func (x *Item) GetData() map[string]*structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

// For example, a bow and its associated arrows.
type CompositeItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string  `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Items []*Item `protobuf:"bytes,20,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CompositeItem) Reset() {
	*x = CompositeItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cauldron_v1alpha_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompositeItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompositeItem) ProtoMessage() {}

func (x *CompositeItem) ProtoReflect() protoreflect.Message {
	mi := &file_cauldron_v1alpha_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompositeItem.ProtoReflect.Descriptor instead.
func (*CompositeItem) Descriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_item_proto_rawDescGZIP(), []int{1}
}

func (x *CompositeItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CompositeItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CompositeItem) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_cauldron_v1alpha_item_proto protoreflect.FileDescriptor

var file_cauldron_v1alpha_item_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6e,
	0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x02, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x1e,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e,
	0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x3d, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x28, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64,
	0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x50, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x6a, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77,
	0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2a, 0x33,
	0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55,
	0x4e, 0x44, 0x41, 0x4d, 0x41, 0x47, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x41,
	0x4d, 0x41, 0x47, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x52, 0x4f, 0x4b, 0x45,
	0x4e, 0x10, 0x02, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cauldron_v1alpha_item_proto_rawDescOnce sync.Once
	file_cauldron_v1alpha_item_proto_rawDescData = file_cauldron_v1alpha_item_proto_rawDesc
)

func file_cauldron_v1alpha_item_proto_rawDescGZIP() []byte {
	file_cauldron_v1alpha_item_proto_rawDescOnce.Do(func() {
		file_cauldron_v1alpha_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_cauldron_v1alpha_item_proto_rawDescData)
	})
	return file_cauldron_v1alpha_item_proto_rawDescData
}

var file_cauldron_v1alpha_item_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cauldron_v1alpha_item_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cauldron_v1alpha_item_proto_goTypes = []interface{}{
	(ItemState)(0),          // 0: nicklaw5.cauldron.v1alpha.ItemState
	(*Item)(nil),            // 1: nicklaw5.cauldron.v1alpha.Item
	(*CompositeItem)(nil),   // 2: nicklaw5.cauldron.v1alpha.CompositeItem
	nil,                     // 3: nicklaw5.cauldron.v1alpha.Item.DataEntry
	(*structpb.Struct)(nil), // 4: google.protobuf.Struct
}
var file_cauldron_v1alpha_item_proto_depIdxs = []int32{
	0, // 0: nicklaw5.cauldron.v1alpha.Item.state:type_name -> nicklaw5.cauldron.v1alpha.ItemState
	3, // 1: nicklaw5.cauldron.v1alpha.Item.data:type_name -> nicklaw5.cauldron.v1alpha.Item.DataEntry
	1, // 2: nicklaw5.cauldron.v1alpha.CompositeItem.items:type_name -> nicklaw5.cauldron.v1alpha.Item
	4, // 3: nicklaw5.cauldron.v1alpha.Item.DataEntry.value:type_name -> google.protobuf.Struct
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cauldron_v1alpha_item_proto_init() }
func file_cauldron_v1alpha_item_proto_init() {
	if File_cauldron_v1alpha_item_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cauldron_v1alpha_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_cauldron_v1alpha_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompositeItem); i {
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
			RawDescriptor: file_cauldron_v1alpha_item_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cauldron_v1alpha_item_proto_goTypes,
		DependencyIndexes: file_cauldron_v1alpha_item_proto_depIdxs,
		EnumInfos:         file_cauldron_v1alpha_item_proto_enumTypes,
		MessageInfos:      file_cauldron_v1alpha_item_proto_msgTypes,
	}.Build()
	File_cauldron_v1alpha_item_proto = out.File
	file_cauldron_v1alpha_item_proto_rawDesc = nil
	file_cauldron_v1alpha_item_proto_goTypes = nil
	file_cauldron_v1alpha_item_proto_depIdxs = nil
}
