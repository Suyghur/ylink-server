package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"ylink/comm/globalkey"
	"ylink/comm/jwtkey"
	"ylink/comm/result"
	"ylink/core/inner/rpc/inner"
	"ylink/flowsrv/rpc/internal/mgr"

	"ylink/flowsrv/rpc/internal/svc"
	"ylink/flowsrv/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ConnectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewConnectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConnectLogic {
	return &ConnectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ConnectLogic) Connect(in *pb.CommandReq, stream pb.Flowsrv_ConnectServer) error {
	uid, gameId, err := l.checkAuth(in)
	if err != nil {
		return stream.Send(&pb.CommandResp{
			Code: result.TokenParseError,
			Msg:  err.Error(),
			Data: nil,
		})
	}
	_, err = l.svcCtx.InnerRpc.NotifyUserOnline(l.ctx, &inner.NotifyUserStatusReq{
		Type:   in.Type,
		Uid:    uid,
		GameId: gameId,
	})
	if err != nil {
		return stream.Send(&pb.CommandResp{
			Code: result.ServerCommonError,
			Msg:  err.Error(),
			Data: nil,
		})
	}

	mgr.GetFlowMgrInstance().SetFlow(uid, stream)

	return stream.Send(&pb.CommandResp{
		Code: result.Ok,
		Msg:  "success",
		Data: nil,
	})
}

func (l *ConnectLogic) checkAuth(in *pb.CommandReq) (string, string, error) {
	token, err := jwt.Parse(in.AccessToken, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(l.svcCtx.Config.JwtAuth.AccessSecret), nil
	})

	uid := ""
	gameId := ""
	if token.Valid {
		//将获取的token中的Claims强转为MapClaims
		claims, _ := token.Claims.(jwt.MapClaims)
		if in.Type == globalkey.CONNECT_TYPE_PLAYER {
			uid = claims[jwtkey.PlayerId].(string)
			gameId = claims[jwtkey.GameId].(string)
		} else {
			uid = claims[jwtkey.CsId].(string)
		}
		return uid, gameId, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return uid, gameId, errors.Wrap(result.NewErrCode(result.TokenParseError), "")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return uid, gameId, errors.Wrap(result.NewErrCode(result.TokenExpireError), "")
		} else {
			return uid, gameId, errors.Wrap(result.NewErrCode(result.TokenParseError), "")
		}
	} else {
		return uid, gameId, errors.Wrap(result.NewErrCode(result.TokenParseError), "")
	}
}
