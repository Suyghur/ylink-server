package logic

import (
	"context"
	"strings"

	"call_center/db/rpc/internal/svc"
	"call_center/db/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSensitiveWordsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSensitiveWordsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSensitiveWordsLogic {
	return &GetSensitiveWordsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSensitiveWordsLogic) GetSensitiveWords(in *pb.SensReq) (*pb.SensRes, error) {
	// todo: add your logic here and delete this line
	res := new(pb.SensRes)

	// 读取db
	wordsModel, err := l.svcCtx.SensWordModel.FindAll()
	if err != nil {
		return res, err
	}
	filterModel, err := l.svcCtx.ConfigModel.FindOneByConfKey("sensitive_words_filter")
	if err != nil {
		return res, err
	}

	// 结果赋值
	for _, v := range *wordsModel {
		res.SensWords = append(res.SensWords, v.Word)
	}
	wordsFilter := strings.Split(filterModel.ConfValue, ",")
	res.FilterWord = wordsFilter

	logx.Infof("[GetSensitiveWords] sensWordLen:%s, filterWordsLen:%s", len(res.SensWords), len(res.FilterWord))
	return res, nil
}
