// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: cauldron/v1alpha/player.proto

package cauldron

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email          string `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	Usernname      string `protobuf:"bytes,20,opt,name=usernname,proto3" json:"usernname,omitempty"`
	DisplayName    string `protobuf:"bytes,30,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	PasswordHashed string `protobuf:"bytes,40,opt,name=password_hashed,json=passwordHashed,proto3" json:"password_hashed,omitempty"`
	FirstName      string `protobuf:"bytes,50,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName       string `protobuf:"bytes,60,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// The number of characters a single player can own.
	// A negative value represents unlimited characters.
	//
	// Options:
	//  - default_character_limit: The default limit of characters a player can own.
	CharacterLimit int32                  `protobuf:"varint,70,opt,name=character_limit,json=characterLimit,proto3" json:"character_limit,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,80,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	CustomData     *structpb.Struct       `protobuf:"bytes,90,opt,name=custom_data,json=customData,proto3" json:"custom_data,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cauldron_v1alpha_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_cauldron_v1alpha_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_cauldron_v1alpha_player_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Player) GetUsernname() string {
	if x != nil {
		return x.Usernname
	}
	return ""
}

func (x *Player) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Player) GetPasswordHashed() string {
	if x != nil {
		return x.PasswordHashed
	}
	return ""
}

func (x *Player) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Player) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Player) GetCharacterLimit() int32 {
	if x != nil {
		return x.CharacterLimit
	}
	return 0
}

func (x *Player) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Player) GetCustomData() *structpb.Struct {
	if x != nil {
		return x.CustomData
	}
	return nil
}

var file_cauldron_v1alpha_player_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         5020,
		Name:          "nicklaw5.cauldron.v1alpha.Player.default_character_limit",
		Tag:           "varint,5020,opt,name=default_character_limit",
		Filename:      "cauldron/v1alpha/player.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional int32 default_character_limit = 5020;
	E_Player_DefaultCharacterLimit = &file_cauldron_v1alpha_player_proto_extTypes[0]
)

var File_cauldron_v1alpha_player_proto protoreflect.FileDescriptor

var file_cauldron_v1alpha_player_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77, 0x35, 0x2e, 0x63, 0x61, 0x75, 0x6c, 0x64, 0x72,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x03, 0x0a, 0x06,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a,
	0x0f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x65, 0x64,
	0x18, 0x28, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x48, 0x61, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x36, 0x0a, 0x0f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x46, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0d, 0xe0, 0xb9, 0x02,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x72,
	0x61, 0x63, 0x74, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x50, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x5a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x61, 0x74, 0x61, 0x32,
	0x59, 0x0a, 0x17, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x63, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x9c, 0x27, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x15, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74,
	0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x69, 0x63, 0x6b, 0x6c, 0x61, 0x77,
	0x35, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x75, 0x6c,
	0x64, 0x72, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x63, 0x61, 0x75,
	0x6c, 0x64, 0x72, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cauldron_v1alpha_player_proto_rawDescOnce sync.Once
	file_cauldron_v1alpha_player_proto_rawDescData = file_cauldron_v1alpha_player_proto_rawDesc
)

func file_cauldron_v1alpha_player_proto_rawDescGZIP() []byte {
	file_cauldron_v1alpha_player_proto_rawDescOnce.Do(func() {
		file_cauldron_v1alpha_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_cauldron_v1alpha_player_proto_rawDescData)
	})
	return file_cauldron_v1alpha_player_proto_rawDescData
}

var file_cauldron_v1alpha_player_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cauldron_v1alpha_player_proto_goTypes = []interface{}{
	(*Player)(nil),                    // 0: nicklaw5.cauldron.v1alpha.Player
	(*timestamppb.Timestamp)(nil),     // 1: google.protobuf.Timestamp
	(*structpb.Struct)(nil),           // 2: google.protobuf.Struct
	(*descriptorpb.FieldOptions)(nil), // 3: google.protobuf.FieldOptions
}
var file_cauldron_v1alpha_player_proto_depIdxs = []int32{
	1, // 0: nicklaw5.cauldron.v1alpha.Player.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: nicklaw5.cauldron.v1alpha.Player.custom_data:type_name -> google.protobuf.Struct
	3, // 2: nicklaw5.cauldron.v1alpha.Player.default_character_limit:extendee -> google.protobuf.FieldOptions
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	2, // [2:3] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cauldron_v1alpha_player_proto_init() }
func file_cauldron_v1alpha_player_proto_init() {
	if File_cauldron_v1alpha_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cauldron_v1alpha_player_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
			RawDescriptor: file_cauldron_v1alpha_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_cauldron_v1alpha_player_proto_goTypes,
		DependencyIndexes: file_cauldron_v1alpha_player_proto_depIdxs,
		MessageInfos:      file_cauldron_v1alpha_player_proto_msgTypes,
		ExtensionInfos:    file_cauldron_v1alpha_player_proto_extTypes,
	}.Build()
	File_cauldron_v1alpha_player_proto = out.File
	file_cauldron_v1alpha_player_proto_rawDesc = nil
	file_cauldron_v1alpha_player_proto_goTypes = nil
	file_cauldron_v1alpha_player_proto_depIdxs = nil
}
