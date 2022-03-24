//@Author   : KaiShin
//@Time     : 2021/10/28

package handler

import (
	"call_center/db/rpc/internal/svc"
	"call_center/db/rpc/pb"
	"log"
)

func GetConfig(svcCtx *svc.ServiceContext) *pb.DbCommandMsg {
	dataList := new(pb.ArrayConfig)
	configs, err := svcCtx.ConfigModel.FindAll()
	if err != nil {
		log.Println("<handler.GetConfig> [ERROR] err:", err)
	}
	for _, conf := range configs {
		p := new(pb.DbConfig)
		p.ConfName = conf.ConfName
		p.ConfValue = conf.ConfValue
		p.ConfKey = conf.ConfKey
		dataList.DataList = append(dataList.DataList, p)
	}

	return &pb.DbCommandMsg{Data: &pb.DbCommandMsg_ArrayConfig{ArrayConfig: dataList}}
}
