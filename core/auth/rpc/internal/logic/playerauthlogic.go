package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"
	"ylink/comm/jwtkey"
	"ylink/comm/result"

	"ylink/core/auth/rpc/internal/svc"
	"ylink/core/auth/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlayerAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlayerAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlayerAuthLogic {
	return &PlayerAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlayerAuthLogic) PlayerAuth(in *pb.PlayerAuthReq) (*pb.AuthResp, error) {
	now := time.Now().Unix()
	token, err := l.generatePlayerToken(now, in.PlayerId, in.GameId)
	if err != nil {
		return nil, errors.Wrap(result.NewErrCode(result.TokenGenerateError), "")
	}
	return &pb.AuthResp{
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
func (l *PlayerAuthLogic) generatePlayerToken(iat int64, playerId string, gameId string) (string, error) {
	secret := l.svcCtx.Config.JwtAuth.AccessSecret
	expire := l.svcCtx.Config.JwtAuth.AccessExpire
	claims := make(jwt.MapClaims)
	claims["iat"] = iat
	claims["exp"] = iat + expire
	claims[jwtkey.PlayerId] = playerId
	claims[jwtkey.GameId] = gameId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}
