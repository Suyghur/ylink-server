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

type CsSendMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsSendMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsSendMsgLogic {
	return &CsSendMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsSendMsgLogic) CsSendMsg(in *pb.CsSendMsgReq) (*pb.CsSendMsgResp, error) {
	// 投递到自己的发件箱
	message := &model.ChatMessage{
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		Content:    in.Content,
		Pic:        in.Pic,
		ReceiverId: in.GameId + "_" + in.PlayerId,
		SenderId:   in.CsId,
		GameId:     in.GameId,
		Uid:        in.CsId,
	}
	payload, _ := sonic.MarshalString(message)
	kMsg, _ := sonic.MarshalString(&model.KqMessage{
		Opt:     model.CMD_SEND_MESSAGE,
		Payload: payload,
		Ext:     "",
	})
	_, _, err := l.svcCtx.KqMsgBoxProducer.SendMessage(l.ctx, kMsg, message.SenderId)
	if err != nil {
		return nil, err
	}
	return &pb.CsSendMsgResp{}, nil
}
