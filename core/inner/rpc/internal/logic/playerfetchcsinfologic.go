package logic

import (
	"context"
	treemap "github.com/liyue201/gostl/ds/map"
	"github.com/pkg/errors"
	"ylink/comm/model"
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
	if ext.GameConnectedMap.Contains(in.GameId) {
		playerConnMap := ext.GameConnectedMap.Get(in.GameId).(*treemap.Map)
		csId := playerConnMap.Get(in.PlayerId).(string)
		if ext.CsInfoMap.Contains(csId) {
			csInfo := ext.CsInfoMap.Get(csId).(model.CsInfo)
			return &pb.InnerPlayerFetchCsInfoResp{
				CsId:         csInfo.CsId,
				CsNickname:   csInfo.CsNickname,
				CsAvatarUrl:  csInfo.CsAvatarUrl,
				CsSignature:  csInfo.CsSignature,
				OnlineStatus: csInfo.OnlineStatus,
			}, nil
		}
	}
	return nil, errors.Wrap(result.NewErrMsg("Customer service information does not exist"), "")
}
