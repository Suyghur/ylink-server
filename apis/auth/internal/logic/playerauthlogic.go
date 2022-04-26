package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/protobuf/types/known/structpb"
	"time"
	"ylink/ext/jwtdata"

	"ylink/apis/auth/internal/svc"
	"ylink/apis/auth/pb"

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
		return nil, err
	}

	data, err := structpb.NewStruct(map[string]interface{}{
		"access_token": token,
	})
	if err != nil {
		return nil, err
	}

	return &pb.AuthResp{
		Code: 0,
		Msg:  "success",
		Data: data,
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
	claims[jwtdata.JwtKeyPlayerId] = playerId
	claims[jwtdata.JwtKeyGameId] = gameId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}
