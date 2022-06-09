package logic

import (
	"context"
	"github.com/pkg/errors"
	"ylink/comm/result"
	"ylink/core/inner/rpc/internal/ext"
	"ylink/core/inner/rpc/internal/svc"
	"ylink/core/inner/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerFetchCsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlayerFetchCsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerFetchCsInfoLogic {
	return &PlayerFetchCsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlayerFetchCsInfoLogic) PlayerFetchCsInfo(in *pb.InnerPlayerFetchCsInfoReq) (*pb.InnerPlayerFetchCsInfoResp, error) {
	if playerInfo := ext.GetConnectedPlayerInfo(in.GameId, in.PlayerId); playerInfo != nil {
		// 玩家已连接
		if csInfo := ext.GetCsInfo(playerInfo.CsId); csInfo != nil {
			return &pb.InnerPlayerFetchCsInfoResp{
				CsId:         csInfo.CsId,
				CsNickname:   csInfo.CsNickname,
				CsAvatarUrl:  csInfo.CsAvatarUrl,
				CsSignature:  csInfo.CsSignature,
				OnlineStatus: csInfo.OnlineStatus,
			}, nil
		}
		return nil, errors.Wrap(result.NewErrMsg("Customer service information does not exist"), "")
	}
	return nil, errors.Wrap(result.NewErrMsg("The player is not connected"), "")

}
