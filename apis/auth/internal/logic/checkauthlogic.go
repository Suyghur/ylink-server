package logic

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/structpb"
	"ylink/apis/auth/internal/svc"
	"ylink/apis/auth/pb"
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

func (l *CheckAuthLogic) CheckAuth(in *pb.CheckAuthReq) (*pb.AuthResp, error) {

	// 解析传入的token
	// 第二个参数是一个回调函数，作用是判断生成token所用的签名算法是否和传入token的签名算法是否一致。
	// 算法匹配就返回密钥，用来解析token.
	token, err := jwt.Parse(in.Token, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(l.svcCtx.Config.JwtAuth.AccessSecret), nil
	})

	// err不为空，说明token已过期
	if err != nil {
		return nil, err
	}

	// 将获取的token中的Claims强转为MapClaims
	_, ok := token.Claims.(jwt.MapClaims)
	// 判断token是否有效
	if !(ok && token.Valid) {
		return nil, errors.New("cannot convert claim to mapClaim")
	}

	data, err := structpb.NewStruct(map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return &pb.AuthResp{
		Code: 0,
		Msg:  "success",
		Data: data,
	}, nil
}
