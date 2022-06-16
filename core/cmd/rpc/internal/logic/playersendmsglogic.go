package logic

import (
	"context"
	"github.com/bytedance/sonic"
	"time"
	"ylink/comm/model"
	"ylink/core/cmd/rpc/internal/svc"
	"ylink/core/cmd/rpc/pb"

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
	uniqueId := in.GameId + "_" + in.PlayerId
	t := time.Now()
	payload, _ := sonic.MarshalString(&model.ChatMessage{
		CreateTime: t.Format("2006-01-02 15:04:05"),
		Content:    in.Content,
		Pic:        in.Pic,
	})
	kMsg, _ := sonic.MarshalString(&model.KqMessage{
		Opt:      model.CMD_SEND_MESSAGE,
		CreateTs: t.Unix(),
		Payload:  payload,
		SenderId: uniqueId,
		GameId:   in.GameId,
		Uid:      in.PlayerId,
		Ext:      "",
	})
	_, _, err := l.svcCtx.KqMsgBoxProducer.SendMessage(l.ctx, kMsg, uniqueId)
	if err != nil {
		return nil, err
	}
	return &pb.PlayerSendMsgResp{}, nil
}
