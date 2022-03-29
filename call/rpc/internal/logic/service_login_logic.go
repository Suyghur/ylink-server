package logic

import (
	"call_center/call/rpc/internal/core"
	"call_center/call/rpc/pb"
	"call_center/public/exception"
	"context"
	"log"
	"time"

	"call_center/call/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewServiceLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ServiceLoginLogic {
	return &ServiceLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ServiceLoginLogic) ServiceLogin(req *pb.ServiceMsgReq, stream pb.Call_ServiceLoginServer) error {
	server := core.GetServer()

	idInfo := req.IdInfo
	if idInfo == nil {
		// 传参失效
		errStr := "<ServiceLogin> req.IdInfo is nil"
		log.Println(errStr)
		return exception.MakeError(int32(pb.EErrorCode_ERR_PARAM_ERROR), errStr)
	}

	if idInfo.Id != "" {
		// 客服重登
		service := server.GetService(idInfo.Id)
		if service != nil {
			server.KickService(idInfo.Id, int32(pb.ErrorReason_SERVICE_REPEAT_LOGIN))
			log.Println("<ServiceLogin> service already conn, disconnect first, id:", idInfo.Id)
		}
	}

	// 客服stream注册
	service, err := server.OnServiceConnect(idInfo.Id, server, stream, l.svcCtx.Db)
	if service == nil {
		log.Println("<ClientLogin> OnPlayerConnect failed, err:", err)
		return exception.MakeError(int32(pb.EErrorCode_ERR_SERVICE_CONN_ERR), err.Error())
	}

	// 初始化客服信息
	stopChan := make(chan int32)
	service.Init(stopChan)

	// 心跳ticker
	duration := time.Second * time.Duration(server.HeartBeatInterval)
	ticker := time.NewTicker(duration)
	var errCode pb.ErrorReason
	exception.Try(func() {
		defer func() {
			ticker.Stop()
		}()

		for {
			select {
			case <-stream.Context().Done():
				log.Println("<ServiceLogin> heartbeat failed, id:", service.Id, " err:", stream.Context().Err())
				errCode = pb.ErrorReason_SERVICE_HEART_BEAT_FAILED
				return
			case <-ticker.C:
				break
			case stop := <-stopChan:
				errCode = pb.ErrorReason(stop)
				log.Println("<ServiceLogin> service stop connect, code:", stop)
				return
			}
		}
	}).Catch(func(ex exception.Exception) {
		log.Println("<ServiceLogin> error:", ex)
	}).Finally(func() {
		server.OnServiceDisConnect(server, service.Id, errCode)
	})
	return nil
}
