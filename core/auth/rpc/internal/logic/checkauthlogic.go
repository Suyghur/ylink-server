package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"ylink/ext/result"

	"ylink/core/auth/rpc/internal/svc"
	"ylink/core/auth/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAuthLogic {
	return &CheckAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckAuthLogic) CheckAuth(in *pb.CheckAuthReq) (*pb.CheckAuthResp, error) {
	token, err := jwt.Parse(in.AccessToken, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(l.svcCtx.Config.JwtAuth.AccessSecret), nil
	})

	if token.Valid {
		return &pb.CheckAuthResp{}, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.Wrap(result.NewErrCode(result.TokenParseError), "")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return nil, errors.Wrap(result.NewErrCode(result.TokenExpireError), "")
		} else {
			return nil, errors.Wrap(result.NewErrCode(result.TokenParseError), "")
		}
	} else {
		return nil, errors.Wrap(result.NewErrCode(result.TokenParseError), "")
	}
}
