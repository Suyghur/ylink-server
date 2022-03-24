package main

import (
	"call_center/call/rpc/internal/core"
	"call_center/call/rpc/pb"
	logic "call_center/public/common"
	"flag"
	"fmt"

	"call_center/call/rpc/internal/config"
	"call_center/call/rpc/internal/server"
	"call_center/call/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/core/service"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/call.yaml", "the config file")

func main() {

	flag.Parse()

	logic.PreCheckConfig(*configFile)

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewCallServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterCallServer(grpcServer, srv)

		switch c.Mode {
		case service.DevMode, service.TestMode:
			reflection.Register(grpcServer)
		default:
		}

	})
	defer s.Stop()

	// core server init
	core.ServerInit(c.CoreConf)

	// start db
	if ctx.Db != nil {
		go ctx.Db.Start()
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
