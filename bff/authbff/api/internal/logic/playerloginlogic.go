package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"ylink/bff/authbff/api/internal/svc"
	"ylink/bff/authbff/api/internal/types"
	"ylink/comm/jwtkey"
	"ylink/comm/result"
)

type PlayerLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlayerLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerLoginLogic {
	return &PlayerLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlayerLoginLogic) PlayerLogin(req *types.PlayerAuthReq) (resp *types.AuthResp, err error) {
	now := time.Now().Unix()
	token, err := l.generatePlayerToken(now, req.Type, req.GameId, req.PlayerId)
	if err != nil {
		return nil, errors.Wrap(result.NewErrCode(result.TokenGenerateError), "")
	}
	return &types.AuthResp{
		AccessToken: token,
	}, nil
}

//
//  generatePlayerToken
//  @Description: 玩家token签发
//  @receiver l
//  @param iat
//  @param playerId
//  @param gameId
//  @return string
//  @return error
//
func (l *PlayerLoginLogic) generatePlayerToken(iat int64, cType int32, gameId string, playerId string) (string, error) {
	secret := l.svcCtx.Config.JwtAuth.AccessSecret
	expire := l.svcCtx.Config.JwtAuth.AccessExpire
	claims := make(jwt.MapClaims)
	claims["iat"] = iat
	claims["exp"] = iat + expire
	claims[jwtkey.GameId] = gameId
	claims[jwtkey.PlayerId] = playerId
	claims[jwtkey.Type] = cType
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}
