package main

import (
	"call_center/db/rpc/internal/config"
	"call_center/db/rpc/internal/server"
	"call_center/db/rpc/internal/svc"
	"call_center/db/rpc/pb"
	logic "call_center/public/common"
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"log"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/db.yaml", "the config file")

func main() {
	flag.Parse()

	logic.PreCheckConfig(*configFile)

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 日志模块初始化
	logx.MustSetup(c.Log)
	logx.CollectSysLog()

	ctx := svc.NewServiceContext(c)
	srv := server.NewDbServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterDbServer(grpcServer, srv)

		switch c.Mode {
		case service.DevMode, service.TestMode:
			reflection.Register(grpcServer)
		default:
		}

	})
	defer s.Stop()

	log.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
