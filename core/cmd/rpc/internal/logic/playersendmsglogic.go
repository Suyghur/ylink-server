package logic

import (
	"context"
	"encoding/json"
	"time"
	"ylink/core/cmd/rpc/internal/svc"
	"ylink/core/cmd/rpc/pb"
	"ylink/ext/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerSendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlayerSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerSendMsgLogic {
	return &PlayerSendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlayerSendMsgLogic) PlayerSendMsg(in *pb.PlayerSendMsgReq) (*pb.PlayerSendMsgResp, error) {
	// 投递到自己的发件箱
	msg, _ := json.Marshal(model.ChatMessage{
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		Content:    in.Content,
		Pic:        in.Pic,
		ReceiverId: "",
		SenderId:   in.PlayerId,
	})
	_, _, err := l.svcCtx.ChatMsgProducer.SendMessage(string(msg), in.PlayerId)
	if err != nil {
		return nil, err
	}
	return &pb.PlayerSendMsgResp{}, nil
}
