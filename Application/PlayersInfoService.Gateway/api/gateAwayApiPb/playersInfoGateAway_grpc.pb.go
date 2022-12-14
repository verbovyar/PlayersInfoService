// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: playersInfoGateAway.proto

package gateAwayApiPb

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

// PlayersInfoGateAwayClient is the client API for PlayersInfoGateAway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PlayersInfoGateAwayClient interface {
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	Drop(ctx context.Context, in *DropRequest, opts ...grpc.CallOption) (*DropResponse, error)
}

type playersInfoGateAwayClient struct {
	cc grpc.ClientConnInterface
}

func NewPlayersInfoGateAwayClient(cc grpc.ClientConnInterface) PlayersInfoGateAwayClient {
	return &playersInfoGateAwayClient{cc}
}

func (c *playersInfoGateAwayClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/PlayersInfoGateAway/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playersInfoGateAwayClient) Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/PlayersInfoGateAway/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playersInfoGateAwayClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/PlayersInfoGateAway/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *playersInfoGateAwayClient) Drop(ctx context.Context, in *DropRequest, opts ...grpc.CallOption) (*DropResponse, error) {
	out := new(DropResponse)
	err := c.cc.Invoke(ctx, "/PlayersInfoGateAway/Drop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PlayersInfoGateAwayServer is the server API for PlayersInfoGateAway service.
// All implementations must embed UnimplementedPlayersInfoGateAwayServer
// for forward compatibility
type PlayersInfoGateAwayServer interface {
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	Post(context.Context, *PostRequest) (*PostResponse, error)
	Put(context.Context, *PutRequest) (*PutResponse, error)
	Drop(context.Context, *DropRequest) (*DropResponse, error)
	mustEmbedUnimplementedPlayersInfoGateAwayServer()
}

// UnimplementedPlayersInfoGateAwayServer must be embedded to have forward compatible implementations.
type UnimplementedPlayersInfoGateAwayServer struct {
}

func (UnimplementedPlayersInfoGateAwayServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedPlayersInfoGateAwayServer) Post(context.Context, *PostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedPlayersInfoGateAwayServer) Put(context.Context, *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedPlayersInfoGateAwayServer) Drop(context.Context, *DropRequest) (*DropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Drop not implemented")
}
func (UnimplementedPlayersInfoGateAwayServer) mustEmbedUnimplementedPlayersInfoGateAwayServer() {}

// UnsafePlayersInfoGateAwayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PlayersInfoGateAwayServer will
// result in compilation errors.
type UnsafePlayersInfoGateAwayServer interface {
	mustEmbedUnimplementedPlayersInfoGateAwayServer()
}

func RegisterPlayersInfoGateAwayServer(s grpc.ServiceRegistrar, srv PlayersInfoGateAwayServer) {
	s.RegisterService(&PlayersInfoGateAway_ServiceDesc, srv)
}

func _PlayersInfoGateAway_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayersInfoGateAwayServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlayersInfoGateAway/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayersInfoGateAwayServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayersInfoGateAway_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayersInfoGateAwayServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlayersInfoGateAway/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayersInfoGateAwayServer).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayersInfoGateAway_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayersInfoGateAwayServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlayersInfoGateAway/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayersInfoGateAwayServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PlayersInfoGateAway_Drop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DropRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PlayersInfoGateAwayServer).Drop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PlayersInfoGateAway/Drop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PlayersInfoGateAwayServer).Drop(ctx, req.(*DropRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PlayersInfoGateAway_ServiceDesc is the grpc.ServiceDesc for PlayersInfoGateAway service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PlayersInfoGateAway_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "PlayersInfoGateAway",
	HandlerType: (*PlayersInfoGateAwayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _PlayersInfoGateAway_GetAll_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _PlayersInfoGateAway_Post_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _PlayersInfoGateAway_Put_Handler,
		},
		{
			MethodName: "Drop",
			Handler:    _PlayersInfoGateAway_Drop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "playersInfoGateAway.proto",
}
