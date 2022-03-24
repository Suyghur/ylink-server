// Code generated by goctl. DO NOT EDIT!
// Source: db.proto

package db

import (
	"context"

	"call_center/db/rpc/pb"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	DbMsgReq = pb.DbMsgReq
	DbMsgRes = pb.DbMsgRes
	SensReq  = pb.SensReq
	SensRes  = pb.SensRes

	Db interface {
		//  客服聊天
		DbLogin(ctx context.Context, opts ...grpc.CallOption) (pb.Db_DbLoginClient, error)
		DbCall(ctx context.Context, in *DbMsgReq, opts ...grpc.CallOption) (*DbMsgRes, error)
		//  游戏聊天
		GetSensitiveWords(ctx context.Context, in *SensReq, opts ...grpc.CallOption) (*SensRes, error)
	}

	defaultDb struct {
		cli zrpc.Client
	}
)

func NewDb(cli zrpc.Client) Db {
	return &defaultDb{
		cli: cli,
	}
}

//  客服聊天
func (m *defaultDb) DbLogin(ctx context.Context, opts ...grpc.CallOption) (pb.Db_DbLoginClient, error) {
	client := pb.NewDbClient(m.cli.Conn())
	return client.DbLogin(ctx, opts...)
}

func (m *defaultDb) DbCall(ctx context.Context, in *DbMsgReq, opts ...grpc.CallOption) (*DbMsgRes, error) {
	client := pb.NewDbClient(m.cli.Conn())
	return client.DbCall(ctx, in, opts...)
}

//  游戏聊天
func (m *defaultDb) GetSensitiveWords(ctx context.Context, in *SensReq, opts ...grpc.CallOption) (*SensRes, error) {
	client := pb.NewDbClient(m.cli.Conn())
	return client.GetSensitiveWords(ctx, in, opts...)
}
