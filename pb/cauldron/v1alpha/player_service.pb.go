// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: cauldron/v1alpha/player_service.proto

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

type NewPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email       string           `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Usernname   string           `protobuf:"bytes,3,opt,name=usernname,proto3" json:"usernname,omitempty"`
	DisplayName string           `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Attributes  *structpb.Struct `protobuf:"bytes,10,opt,name=attributes,proto3" json:"attributes,omitempty"`
}

func (x *NewPlayerRequest) Reset() {
	*x = NewPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cauldron_v1alpha_player_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPlayerRequest) ProtoMessage() {}

func (x *NewPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cauldron_v1alpha_player_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPlayerRequest.ProtoReflect.Descriptor instead.
func (*NewPlayerRequest) Descriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_player_service_proto_rawDescGZIP(), []int{0}
}

func (x *NewPlayerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NewPlayerRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *NewPlayerRequest) GetUsernname() string {
	if x != nil {
		return x.Usernname
	}
	return ""
}

func (x *NewPlayerRequest) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *NewPlayerRequest) GetAttributes() *structpb.Struct {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type NewPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *NewPlayerResponse) Reset() {
	*x = NewPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cauldron_v1alpha_player_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPlayerResponse) ProtoMessage() {}

func (x *NewPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cauldron_v1alpha_player_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPlayerResponse.ProtoReflect.Descriptor instead.
func (*NewPlayerResponse) Descriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_player_service_proto_rawDescGZIP(), []int{1}
}

func (x *NewPlayerResponse) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

var File_cauldron_v1alpha_player_service_proto protoreflect.FileDescriptor

var file_cauldron_v1alpha_player_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77,
	0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb2, 0x01, 0x0a, 0x10, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x22, 0x4e, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x6e, 0x69, 0x63, 0x6b,
	0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x32, 0x79, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x68, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x2b, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61,
	0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4e,
	0x65, 0x77, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2c, 0x2e, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64,
	0x72, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4e, 0x65, 0x77, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69,
	0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62,
	0x2f, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x3b, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cauldron_v1alpha_player_service_proto_rawDescOnce sync.Once
	file_cauldron_v1alpha_player_service_proto_rawDescData = file_cauldron_v1alpha_player_service_proto_rawDesc
)

func file_cauldron_v1alpha_player_service_proto_rawDescGZIP() []byte {
	file_cauldron_v1alpha_player_service_proto_rawDescOnce.Do(func() {
		file_cauldron_v1alpha_player_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cauldron_v1alpha_player_service_proto_rawDescData)
	})
	return file_cauldron_v1alpha_player_service_proto_rawDescData
}

var file_cauldron_v1alpha_player_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cauldron_v1alpha_player_service_proto_goTypes = []interface{}{
	(*NewPlayerRequest)(nil),  // 0: nicklaw5.cauldron.v1alpha.NewPlayerRequest
	(*NewPlayerResponse)(nil), // 1: nicklaw5.cauldron.v1alpha.NewPlayerResponse
	(*structpb.Struct)(nil),   // 2: google.protobuf.Struct
	(*Player)(nil),            // 3: nicklaw5.cauldron.v1alpha.Player
}
var file_cauldron_v1alpha_player_service_proto_depIdxs = []int32{
	2, // 0: nicklaw5.cauldron.v1alpha.NewPlayerRequest.attributes:type_name -> google.protobuf.Struct
	3, // 1: nicklaw5.cauldron.v1alpha.NewPlayerResponse.player:type_name -> nicklaw5.cauldron.v1alpha.Player
	0, // 2: nicklaw5.cauldron.v1alpha.PlayerService.NewPlayer:input_type -> nicklaw5.cauldron.v1alpha.NewPlayerRequest
	1, // 3: nicklaw5.cauldron.v1alpha.PlayerService.NewPlayer:output_type -> nicklaw5.cauldron.v1alpha.NewPlayerResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cauldron_v1alpha_player_service_proto_init() }
func file_cauldron_v1alpha_player_service_proto_init() {
	if File_cauldron_v1alpha_player_service_proto != nil {
		return
	}
	file_cauldron_v1alpha_player_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cauldron_v1alpha_player_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPlayerRequest); i {
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
		file_cauldron_v1alpha_player_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPlayerResponse); i {
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
			RawDescriptor: file_cauldron_v1alpha_player_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cauldron_v1alpha_player_service_proto_goTypes,
		DependencyIndexes: file_cauldron_v1alpha_player_service_proto_depIdxs,
		MessageInfos:      file_cauldron_v1alpha_player_service_proto_msgTypes,
	}.Build()
	File_cauldron_v1alpha_player_service_proto = out.File
	file_cauldron_v1alpha_player_service_proto_rawDesc = nil
	file_cauldron_v1alpha_player_service_proto_goTypes = nil
	file_cauldron_v1alpha_player_service_proto_depIdxs = nil
}
