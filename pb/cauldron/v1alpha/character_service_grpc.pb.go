// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cauldron

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CharacterServiceClient is the client API for CharacterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CharacterServiceClient interface {
	// Character
	CreateCharacter(ctx context.Context, in *CreateCharacterRequest, opts ...grpc.CallOption) (*CreateCharacterResponse, error)
	GetCharacterByID(ctx context.Context, in *GetCharacterByIDRequest, opts ...grpc.CallOption) (*GetCharacterByIDResponse, error)
	// Character Race
	CreateCharacterRace(ctx context.Context, in *CreateCharacterRaceRequest, opts ...grpc.CallOption) (*CreateCharacterRaceResponse, error)
	// Character Class
	CreateCharacterClass(ctx context.Context, in *CreateCharacterClassRequest, opts ...grpc.CallOption) (*CreateCharacterClassResponse, error)
}

type characterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCharacterServiceClient(cc grpc.ClientConnInterface) CharacterServiceClient {
	return &characterServiceClient{cc}
}

func (c *characterServiceClient) CreateCharacter(ctx context.Context, in *CreateCharacterRequest, opts ...grpc.CallOption) (*CreateCharacterResponse, error) {
	out := new(CreateCharacterResponse)
	err := c.cc.Invoke(ctx, "/nicklaw5.cauldron.v1alpha.CharacterService/CreateCharacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterServiceClient) GetCharacterByID(ctx context.Context, in *GetCharacterByIDRequest, opts ...grpc.CallOption) (*GetCharacterByIDResponse, error) {
	out := new(GetCharacterByIDResponse)
	err := c.cc.Invoke(ctx, "/nicklaw5.cauldron.v1alpha.CharacterService/GetCharacterByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterServiceClient) CreateCharacterRace(ctx context.Context, in *CreateCharacterRaceRequest, opts ...grpc.CallOption) (*CreateCharacterRaceResponse, error) {
	out := new(CreateCharacterRaceResponse)
	err := c.cc.Invoke(ctx, "/nicklaw5.cauldron.v1alpha.CharacterService/CreateCharacterRace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterServiceClient) CreateCharacterClass(ctx context.Context, in *CreateCharacterClassRequest, opts ...grpc.CallOption) (*CreateCharacterClassResponse, error) {
	out := new(CreateCharacterClassResponse)
	err := c.cc.Invoke(ctx, "/nicklaw5.cauldron.v1alpha.CharacterService/CreateCharacterClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CharacterServiceServer is the server API for CharacterService service.
// All implementations must embed UnimplementedCharacterServiceServer
// for forward compatibility
type CharacterServiceServer interface {
	// Character
	CreateCharacter(context.Context, *CreateCharacterRequest) (*CreateCharacterResponse, error)
	GetCharacterByID(context.Context, *GetCharacterByIDRequest) (*GetCharacterByIDResponse, error)
	// Character Race
	CreateCharacterRace(context.Context, *CreateCharacterRaceRequest) (*CreateCharacterRaceResponse, error)
	// Character Class
	CreateCharacterClass(context.Context, *CreateCharacterClassRequest) (*CreateCharacterClassResponse, error)
	mustEmbedUnimplementedCharacterServiceServer()
}

// UnimplementedCharacterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCharacterServiceServer struct {
}

func (UnimplementedCharacterServiceServer) CreateCharacter(context.Context, *CreateCharacterRequest) (*CreateCharacterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCharacter not implemented")
}
func (UnimplementedCharacterServiceServer) GetCharacterByID(context.Context, *GetCharacterByIDRequest) (*GetCharacterByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCharacterByID not implemented")
}
func (UnimplementedCharacterServiceServer) CreateCharacterRace(context.Context, *CreateCharacterRaceRequest) (*CreateCharacterRaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCharacterRace not implemented")
}
func (UnimplementedCharacterServiceServer) CreateCharacterClass(context.Context, *CreateCharacterClassRequest) (*CreateCharacterClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCharacterClass not implemented")
}
func (UnimplementedCharacterServiceServer) mustEmbedUnimplementedCharacterServiceServer() {}

// UnsafeCharacterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CharacterServiceServer will
// result in compilation errors.
type UnsafeCharacterServiceServer interface {
	mustEmbedUnimplementedCharacterServiceServer()
}

func RegisterCharacterServiceServer(s grpc.ServiceRegistrar, srv CharacterServiceServer) {
	s.RegisterService(&CharacterService_ServiceDesc, srv)
}

func _CharacterService_CreateCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCharacterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).CreateCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nicklaw5.cauldron.v1alpha.CharacterService/CreateCharacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).CreateCharacter(ctx, req.(*CreateCharacterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterService_GetCharacterByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCharacterByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).GetCharacterByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nicklaw5.cauldron.v1alpha.CharacterService/GetCharacterByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).GetCharacterByID(ctx, req.(*GetCharacterByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterService_CreateCharacterRace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCharacterRaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).CreateCharacterRace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nicklaw5.cauldron.v1alpha.CharacterService/CreateCharacterRace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).CreateCharacterRace(ctx, req.(*CreateCharacterRaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterService_CreateCharacterClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCharacterClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).CreateCharacterClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nicklaw5.cauldron.v1alpha.CharacterService/CreateCharacterClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).CreateCharacterClass(ctx, req.(*CreateCharacterClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CharacterService_ServiceDesc is the grpc.ServiceDesc for CharacterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CharacterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nicklaw5.cauldron.v1alpha.CharacterService",
	HandlerType: (*CharacterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCharacter",
			Handler:    _CharacterService_CreateCharacter_Handler,
		},
		{
			MethodName: "GetCharacterByID",
			Handler:    _CharacterService_GetCharacterByID_Handler,
		},
		{
			MethodName: "CreateCharacterRace",
			Handler:    _CharacterService_CreateCharacterRace_Handler,
		},
		{
			MethodName: "CreateCharacterClass",
			Handler:    _CharacterService_CreateCharacterClass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cauldron/v1alpha/character_service.proto",
}
