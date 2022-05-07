package logic

import (
	"context"
	"encoding/json"
	"time"
	"ylink/core/cmd/rpc/internal/svc"
	"ylink/core/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerSendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

type message struct {
	CreateTime string `json:"create_time"`
	Content    string `json:"content"`
	Pic        string `json:"pic"`
	ReceiverId string `json:"receiver_id"`
	SenderId   string `json:"sender_id"`
}

func NewPlayerSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerSendMsgLogic {
	return &PlayerSendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlayerSendMsgLogic) PlayerSendMsg(in *pb.PlayerSendMsgReq) (*pb.PlayerSendMsgResp, error) {
	// todo 投递到对应客服的收件箱
	// todo 写入db
	msg, _ := json.Marshal(message{
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		Content:    in.Content,
		Pic:        in.Pic,
		ReceiverId: "",
		SenderId:   in.PlayerId,
	})

	//if err := l.svcCtx.ChatMsgProducerClient.Push(string(msg)); err != nil {
	//	return nil, err
	//}
	pid, offset, err := l.svcCtx.ChatMsgProducer.SendMessage(string(msg), in.PlayerId)
	if err != nil {
		return nil, err
	}
	l.Logger.Infof("pid: %d", pid)
	l.Logger.Infof("offset: %d", offset)
	return &pb.PlayerSendMsgResp{}, nil
}
