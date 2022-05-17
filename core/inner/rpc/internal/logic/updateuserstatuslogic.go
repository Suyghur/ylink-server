package logic

import (
	"context"
	"github.com/liyue201/gostl/ds/set"
	"github.com/pkg/errors"
	"ylink/comm/globalkey"
	"ylink/comm/result"
	"ylink/core/inner/rpc/internal/ext"

	"ylink/core/inner/rpc/internal/svc"
	"ylink/core/inner/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserStatusLogic {
	return &UpdateUserStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserStatusLogic) UpdateUserStatus(in *pb.UpdateUserStatusReq) (*pb.UpdateUserStatusResp, error) {
	switch in.Type {
	case globalkey.CONNECT_TYPE_PLAYER:
		// 修改玩家在线状态
		if ext.Game2PlayerStatMap.Contains(in.GameId) {
			// 有则取出玩家的set
			playerStatSet := ext.Game2PlayerStatMap.Get(in.GameId).(*set.Set)
			if playerStatSet.Contains(in.Uid) {
				// 有则清除，代表下线
				playerStatSet.Erase(in.Uid)
			} else {
				playerStatSet.Insert(in.Uid)
			}
		} else {
			playerStatSet := set.New()
			playerStatSet.Insert(in.Uid)
			ext.Game2PlayerStatMap.Insert(in.GameId, playerStatSet)
		}
	case globalkey.CONNECT_TYPE_CS:
		// 修改客服在线状态
		if ext.CsStatSet.Contains(in.Uid) {
			// 有则清除，代表下线
			ext.CsStatSet.Erase(in.Uid)
		} else {
			ext.CsStatSet.Insert(in.Uid)
		}
	default:
		return nil, errors.Wrap(result.NewErrMsg("no such user type"), "")
	}
	return &pb.UpdateUserStatusResp{}, nil
}
