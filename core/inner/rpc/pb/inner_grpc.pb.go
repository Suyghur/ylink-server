// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pb/inner.proto

package pb

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

// InnerClient is the client API for Inner service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InnerClient interface {
	PlayerFetchCsInfo(ctx context.Context, in *InnerPlayerFetchCsInfoReq, opts ...grpc.CallOption) (*InnerPlayerFetchCsInfoResp, error)
	PlayerDisconnect(ctx context.Context, in *InnerPlayerDisconnectReq, opts ...grpc.CallOption) (*InnerPlayerDisconnectResp, error)
	CsFetchPlayerQueue(ctx context.Context, in *InnerCsFetchPlayerQueueReq, opts ...grpc.CallOption) (*InnerCsFetchPlayerQueueResp, error)
	CsConnectPlayer(ctx context.Context, in *InnerCsConnectPlayerReq, opts ...grpc.CallOption) (*InnerCsConnectPlayerResp, error)
	NotifyUserOnline(ctx context.Context, in *NotifyUserStatusReq, opts ...grpc.CallOption) (*NotifyUserStatusResp, error)
	NotifyUserOffline(ctx context.Context, in *NotifyUserStatusReq, opts ...grpc.CallOption) (*NotifyUserStatusResp, error)
}

type innerClient struct {
	cc grpc.ClientConnInterface
}

func NewInnerClient(cc grpc.ClientConnInterface) InnerClient {
	return &innerClient{cc}
}

func (c *innerClient) PlayerFetchCsInfo(ctx context.Context, in *InnerPlayerFetchCsInfoReq, opts ...grpc.CallOption) (*InnerPlayerFetchCsInfoResp, error) {
	out := new(InnerPlayerFetchCsInfoResp)
	err := c.cc.Invoke(ctx, "/pb.Inner/playerFetchCsInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerClient) PlayerDisconnect(ctx context.Context, in *InnerPlayerDisconnectReq, opts ...grpc.CallOption) (*InnerPlayerDisconnectResp, error) {
	out := new(InnerPlayerDisconnectResp)
	err := c.cc.Invoke(ctx, "/pb.Inner/playerDisconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerClient) CsFetchPlayerQueue(ctx context.Context, in *InnerCsFetchPlayerQueueReq, opts ...grpc.CallOption) (*InnerCsFetchPlayerQueueResp, error) {
	out := new(InnerCsFetchPlayerQueueResp)
	err := c.cc.Invoke(ctx, "/pb.Inner/csFetchPlayerQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerClient) CsConnectPlayer(ctx context.Context, in *InnerCsConnectPlayerReq, opts ...grpc.CallOption) (*InnerCsConnectPlayerResp, error) {
	out := new(InnerCsConnectPlayerResp)
	err := c.cc.Invoke(ctx, "/pb.Inner/csConnectPlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerClient) NotifyUserOnline(ctx context.Context, in *NotifyUserStatusReq, opts ...grpc.CallOption) (*NotifyUserStatusResp, error) {
	out := new(NotifyUserStatusResp)
	err := c.cc.Invoke(ctx, "/pb.Inner/notifyUserOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *innerClient) NotifyUserOffline(ctx context.Context, in *NotifyUserStatusReq, opts ...grpc.CallOption) (*NotifyUserStatusResp, error) {
	out := new(NotifyUserStatusResp)
	err := c.cc.Invoke(ctx, "/pb.Inner/notifyUserOffline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InnerServer is the server API for Inner service.
// All implementations must embed UnimplementedInnerServer
// for forward compatibility
type InnerServer interface {
	PlayerFetchCsInfo(context.Context, *InnerPlayerFetchCsInfoReq) (*InnerPlayerFetchCsInfoResp, error)
	PlayerDisconnect(context.Context, *InnerPlayerDisconnectReq) (*InnerPlayerDisconnectResp, error)
	CsFetchPlayerQueue(context.Context, *InnerCsFetchPlayerQueueReq) (*InnerCsFetchPlayerQueueResp, error)
	CsConnectPlayer(context.Context, *InnerCsConnectPlayerReq) (*InnerCsConnectPlayerResp, error)
	NotifyUserOnline(context.Context, *NotifyUserStatusReq) (*NotifyUserStatusResp, error)
	NotifyUserOffline(context.Context, *NotifyUserStatusReq) (*NotifyUserStatusResp, error)
	mustEmbedUnimplementedInnerServer()
}

// UnimplementedInnerServer must be embedded to have forward compatible implementations.
type UnimplementedInnerServer struct {
}

func (UnimplementedInnerServer) PlayerFetchCsInfo(context.Context, *InnerPlayerFetchCsInfoReq) (*InnerPlayerFetchCsInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerFetchCsInfo not implemented")
}
func (UnimplementedInnerServer) PlayerDisconnect(context.Context, *InnerPlayerDisconnectReq) (*InnerPlayerDisconnectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayerDisconnect not implemented")
}
func (UnimplementedInnerServer) CsFetchPlayerQueue(context.Context, *InnerCsFetchPlayerQueueReq) (*InnerCsFetchPlayerQueueResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CsFetchPlayerQueue not implemented")
}
func (UnimplementedInnerServer) CsConnectPlayer(context.Context, *InnerCsConnectPlayerReq) (*InnerCsConnectPlayerResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CsConnectPlayer not implemented")
}
func (UnimplementedInnerServer) NotifyUserOnline(context.Context, *NotifyUserStatusReq) (*NotifyUserStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyUserOnline not implemented")
}
func (UnimplementedInnerServer) NotifyUserOffline(context.Context, *NotifyUserStatusReq) (*NotifyUserStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyUserOffline not implemented")
}
func (UnimplementedInnerServer) mustEmbedUnimplementedInnerServer() {}

// UnsafeInnerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InnerServer will
// result in compilation errors.
type UnsafeInnerServer interface {
	mustEmbedUnimplementedInnerServer()
}

func RegisterInnerServer(s grpc.ServiceRegistrar, srv InnerServer) {
	s.RegisterService(&Inner_ServiceDesc, srv)
}

func _Inner_PlayerFetchCsInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InnerPlayerFetchCsInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerServer).PlayerFetchCsInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Inner/playerFetchCsInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerServer).PlayerFetchCsInfo(ctx, req.(*InnerPlayerFetchCsInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inner_PlayerDisconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InnerPlayerDisconnectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerServer).PlayerDisconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Inner/playerDisconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerServer).PlayerDisconnect(ctx, req.(*InnerPlayerDisconnectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inner_CsFetchPlayerQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InnerCsFetchPlayerQueueReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerServer).CsFetchPlayerQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Inner/csFetchPlayerQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerServer).CsFetchPlayerQueue(ctx, req.(*InnerCsFetchPlayerQueueReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inner_CsConnectPlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InnerCsConnectPlayerReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerServer).CsConnectPlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Inner/csConnectPlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerServer).CsConnectPlayer(ctx, req.(*InnerCsConnectPlayerReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inner_NotifyUserOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyUserStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerServer).NotifyUserOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Inner/notifyUserOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerServer).NotifyUserOnline(ctx, req.(*NotifyUserStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Inner_NotifyUserOffline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyUserStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InnerServer).NotifyUserOffline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Inner/notifyUserOffline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InnerServer).NotifyUserOffline(ctx, req.(*NotifyUserStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Inner_ServiceDesc is the grpc.ServiceDesc for Inner service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Inner_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Inner",
	HandlerType: (*InnerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "playerFetchCsInfo",
			Handler:    _Inner_PlayerFetchCsInfo_Handler,
		},
		{
			MethodName: "playerDisconnect",
			Handler:    _Inner_PlayerDisconnect_Handler,
		},
		{
			MethodName: "csFetchPlayerQueue",
			Handler:    _Inner_CsFetchPlayerQueue_Handler,
		},
		{
			MethodName: "csConnectPlayer",
			Handler:    _Inner_CsConnectPlayer_Handler,
		},
		{
			MethodName: "notifyUserOnline",
			Handler:    _Inner_NotifyUserOnline_Handler,
		},
		{
			MethodName: "notifyUserOffline",
			Handler:    _Inner_NotifyUserOffline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/inner.proto",
}
