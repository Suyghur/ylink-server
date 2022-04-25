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

type CsAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCsAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsAuthLogic {
	return &CsAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CsAuthLogic) CsAuth(in *pb.CsAuthReq) (*pb.AuthResp, error) {
	now := time.Now().Unix()
	token, err := l.generateCsToken(now, in.CsId)
	if err != nil {
		return nil, err
	}
	data, err := structpb.NewStruct(map[string]interface{}{
		"token": token,
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
//  generateCsToken
//  @Description: 客服token签发
//  @receiver l
//  @param iat
//  @param csId
//  @return string
//  @return error
//
func (l *CsAuthLogic) generateCsToken(iat int64, csId string) (string, error) {
	secret := l.svcCtx.Config.JwtAuth.AccessSecret
	expire := l.svcCtx.Config.JwtAuth.AccessExpire
	claims := make(jwt.MapClaims)
	claims["iat"] = iat
	claims["exp"] = iat + expire
	claims[jwtdata.JwtKeyCsId] = csId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}
