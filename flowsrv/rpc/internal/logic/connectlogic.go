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
	"ylink/flowsrv/rpc/internal/model"
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
	cType, uid, gameId, err := l.checkAuth(in)
	if err != nil || cType == globalkey.ConnectTypeError {
		return stream.Send(&pb.CommandResp{
			Code: result.TokenParseError,
			Msg:  err.Error(),
			Data: nil,
		})
	}
	_, err = l.svcCtx.InnerRpc.NotifyUserOnline(l.ctx, &inner.NotifyUserStatusReq{
		Type:   cType,
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

	var flowId string
	if cType == globalkey.ConnectTypeCs {
		flowId = uid
	} else {
		flowId = gameId + "_" + uid
	}

	flow := &model.Flow{
		EndFlow:     make(chan int),
		Message:     make(chan string),
		Stream:      stream,
		RedisClient: l.svcCtx.RedisClient,
		InnerRpc:    l.svcCtx.InnerRpc,
		Type:        cType,
		Uid:         uid,
		GameId:      gameId,
		FlowId:      flowId,
	}
	defer func() {
		close(flow.EndFlow)
		flow = nil
	}()

	mgr.GetFlowMgrInstance().Register(flow)

	<-flow.EndFlow
	return nil
}

func (l *ConnectLogic) checkAuth(in *pb.CommandReq) (int32, string, string, error) {
	token, err := jwt.Parse(in.AccessToken, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(l.svcCtx.Config.JwtAuth.AccessSecret), nil
	})

	uid := ""
	gameId := ""
	var cType int32 = -1
	if token.Valid {
		//将获取的token中的Claims强转为MapClaims
		claims, _ := token.Claims.(jwt.MapClaims)
		cType = int32(claims[jwtkey.Type].(float64))
		switch cType {
		case globalkey.ConnectTypeNormalPlayer:
			uid = claims[jwtkey.PlayerId].(string)
			gameId = claims[jwtkey.GameId].(string)
		case globalkey.ConnectTypeVipPlayer:
			uid = claims[jwtkey.PlayerId].(string)
			gameId = claims[jwtkey.GameId].(string)
		case globalkey.ConnectTypeCs:
			uid = claims[jwtkey.CsId].(string)
		}
		return cType, uid, gameId, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return cType, uid, gameId, errors.Wrap(result.NewErrCode(result.TokenParseError), "")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return cType, uid, gameId, errors.Wrap(result.NewErrCode(result.TokenExpireError), "")
		} else {
			return cType, uid, gameId, errors.Wrap(result.NewErrCode(result.TokenParseError), "")
		}
	} else {
		return cType, uid, gameId, errors.Wrap(result.NewErrCode(result.TokenParseError), "")
	}
}
