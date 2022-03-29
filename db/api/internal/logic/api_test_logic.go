package logic

import (
	"call_center/db/rpc/pb"
	"context"
	"encoding/json"
	"log"

	"call_center/db/api/internal/svc"
	"call_center/db/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiTestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) ApiTestLogic {
	return ApiTestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiTestLogic) ApiTest(req types.Request) (*types.Response, error) {
	msgReq := new(pb.DbMsgReq)
	cmd := new(pb.DbCommandMsg)
	cmd.CmdType = pb.EDbCommand_E_DB_COMMAND_GET_CONFIG
	msgReq.Cmd = cmd

	res, err := l.svcCtx.DbRpc.DbCall(l.ctx, msgReq)
	if err != nil {
		return nil, err
	}
	confList := res.GetCmd().GetArrayConfig().GetDataList()
	for _, conf := range confList {
		var confValueMap map[string]interface{}
		confName := conf.ConfName
		confKey := conf.ConfKey
		err = json.Unmarshal([]byte(conf.ConfValue), &confValueMap)
		if err != nil {
			continue
		}
		log.Println(confName, confKey, confValueMap)
	}
	log.Println("<ApiTest>")
	return &types.Response{
		Res: 1,
	}, nil
}
