// Code generated by goctl. DO NOT EDIT!
// Source: flowsrv.proto

package flowsrv

import (
	"context"

	"ylink/flowsrv/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CommandReq  = pb.CommandReq
	CommandResp = pb.CommandResp

	Flowsrv interface {
		Connect(ctx context.Context, in *CommandReq, opts ...grpc.CallOption) (pb.Flowsrv_ConnectClient, error)
		Disconnect(ctx context.Context, in *CommandReq, opts ...grpc.CallOption) (*CommandResp, error)
	}

	defaultFlowsrv struct {
		cli zrpc.Client
	}
)

func NewFlowsrv(cli zrpc.Client) Flowsrv {
	return &defaultFlowsrv{
		cli: cli,
	}
}

func (m *defaultFlowsrv) Connect(ctx context.Context, in *CommandReq, opts ...grpc.CallOption) (pb.Flowsrv_ConnectClient, error) {
	client := pb.NewFlowsrvClient(m.cli.Conn())
	return client.Connect(ctx, in, opts...)
}

func (m *defaultFlowsrv) Disconnect(ctx context.Context, in *CommandReq, opts ...grpc.CallOption) (*CommandResp, error) {
	client := pb.NewFlowsrvClient(m.cli.Conn())
	return client.Disconnect(ctx, in, opts...)
}
