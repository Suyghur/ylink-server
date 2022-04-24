package logic

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/structpb"
	"ylink/ext/globalkey"

	"ylink/apis/auth/internal/svc"
	"ylink/apis/auth/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAuthLogic {
	return &CheckAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckAuthLogic) CheckAuth(in *pb.CheckAuthReq) (*pb.AuthResp, error) {
	tokenKey := fmt.Sprintf(globalkey.CacheTokenKey, in.Uid)
	cacheToken, err := l.svcCtx.RedisClient.GetCtx(l.ctx, tokenKey)
	if err != nil {
		return nil, err
	}
	if cacheToken != in.Token {
		return nil, errors.New("CheckToken is invalid")
	}

	data, err := structpb.NewStruct(map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return &pb.AuthResp{
		Code: 0,
		Msg:  "success",
		Data: data,
	}, nil
}
