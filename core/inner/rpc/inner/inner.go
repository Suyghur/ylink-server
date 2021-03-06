// Code generated by goctl. DO NOT EDIT!
// Source: inner.proto

package inner

import (
	"context"

	"ylink/core/inner/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	InnerCsConnectPlayerReq     = pb.InnerCsConnectPlayerReq
	InnerCsConnectPlayerResp    = pb.InnerCsConnectPlayerResp
	InnerCsFetchPlayerQueueReq  = pb.InnerCsFetchPlayerQueueReq
	InnerCsFetchPlayerQueueResp = pb.InnerCsFetchPlayerQueueResp
	InnerPlayerDisconnectReq    = pb.InnerPlayerDisconnectReq
	InnerPlayerDisconnectResp   = pb.InnerPlayerDisconnectResp
	InnerPlayerFetchCsInfoReq   = pb.InnerPlayerFetchCsInfoReq
	InnerPlayerFetchCsInfoResp  = pb.InnerPlayerFetchCsInfoResp
	NotifyUserStatusReq         = pb.NotifyUserStatusReq
	NotifyUserStatusResp        = pb.NotifyUserStatusResp

	Inner interface {
		PlayerFetchCsInfo(ctx context.Context, in *InnerPlayerFetchCsInfoReq, opts ...grpc.CallOption) (*InnerPlayerFetchCsInfoResp, error)
		PlayerDisconnect(ctx context.Context, in *InnerPlayerDisconnectReq, opts ...grpc.CallOption) (*InnerPlayerDisconnectResp, error)
		CsFetchPlayerQueue(ctx context.Context, in *InnerCsFetchPlayerQueueReq, opts ...grpc.CallOption) (*InnerCsFetchPlayerQueueResp, error)
		CsConnectPlayer(ctx context.Context, in *InnerCsConnectPlayerReq, opts ...grpc.CallOption) (*InnerCsConnectPlayerResp, error)
		NotifyUserOnline(ctx context.Context, in *NotifyUserStatusReq, opts ...grpc.CallOption) (*NotifyUserStatusResp, error)
		NotifyUserOffline(ctx context.Context, in *NotifyUserStatusReq, opts ...grpc.CallOption) (*NotifyUserStatusResp, error)
	}

	defaultInner struct {
		cli zrpc.Client
	}
)

func NewInner(cli zrpc.Client) Inner {
	return &defaultInner{
		cli: cli,
	}
}

func (m *defaultInner) PlayerFetchCsInfo(ctx context.Context, in *InnerPlayerFetchCsInfoReq, opts ...grpc.CallOption) (*InnerPlayerFetchCsInfoResp, error) {
	client := pb.NewInnerClient(m.cli.Conn())
	return client.PlayerFetchCsInfo(ctx, in, opts...)
}

func (m *defaultInner) PlayerDisconnect(ctx context.Context, in *InnerPlayerDisconnectReq, opts ...grpc.CallOption) (*InnerPlayerDisconnectResp, error) {
	client := pb.NewInnerClient(m.cli.Conn())
	return client.PlayerDisconnect(ctx, in, opts...)
}

func (m *defaultInner) CsFetchPlayerQueue(ctx context.Context, in *InnerCsFetchPlayerQueueReq, opts ...grpc.CallOption) (*InnerCsFetchPlayerQueueResp, error) {
	client := pb.NewInnerClient(m.cli.Conn())
	return client.CsFetchPlayerQueue(ctx, in, opts...)
}

func (m *defaultInner) CsConnectPlayer(ctx context.Context, in *InnerCsConnectPlayerReq, opts ...grpc.CallOption) (*InnerCsConnectPlayerResp, error) {
	client := pb.NewInnerClient(m.cli.Conn())
	return client.CsConnectPlayer(ctx, in, opts...)
}

func (m *defaultInner) NotifyUserOnline(ctx context.Context, in *NotifyUserStatusReq, opts ...grpc.CallOption) (*NotifyUserStatusResp, error) {
	client := pb.NewInnerClient(m.cli.Conn())
	return client.NotifyUserOnline(ctx, in, opts...)
}

func (m *defaultInner) NotifyUserOffline(ctx context.Context, in *NotifyUserStatusReq, opts ...grpc.CallOption) (*NotifyUserStatusResp, error) {
	client := pb.NewInnerClient(m.cli.Conn())
	return client.NotifyUserOffline(ctx, in, opts...)
}
