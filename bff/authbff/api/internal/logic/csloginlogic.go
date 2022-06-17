package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
	"ylink/bff/authbff/api/internal/svc"
	"ylink/bff/authbff/api/internal/types"
	"ylink/comm/globalkey"
	"ylink/comm/jwtkey"
	"ylink/comm/result"
)

type CsLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCsLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CsLoginLogic {
	return &CsLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CsLoginLogic) CsLogin(req *types.CsAuthReq) (resp *types.AuthResp, err error) {
	now := time.Now().Unix()
	token, err := l.generateCsToken(now, req.CsId)
	if err != nil {
		return nil, errors.Wrap(result.NewErrCode(result.TokenGenerateError), "")
	}
	return &types.AuthResp{
		AccessToken: token,
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
func (l *CsLoginLogic) generateCsToken(iat int64, csId string) (string, error) {
	secret := l.svcCtx.Config.JwtAuth.AccessSecret
	expire := l.svcCtx.Config.JwtAuth.AccessExpire
	claims := make(jwt.MapClaims)
	claims["iat"] = iat
	claims["exp"] = iat + expire
	claims[jwtkey.CsId] = csId
	claims[jwtkey.Type] = globalkey.ConnectTypeCs
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secret))
}
