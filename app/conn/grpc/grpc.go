package grpc

import (
	"github.com/jtzjtz/kit/conn/grpc_pool"
	"github.com/jtzjtz/kit/conn/rpc_pool"
	"github.com/jtzjtz/ys_api/app/config"
	"time"
)

var FrameworkGRPC rpc_pool.RpcPool

func Init() {

	initFrameworkGRPCPool()
}
func initFrameworkGRPCPool() {
	var err error
	url := config.App.GrpcServiceUrl
	options := &grpc_pool.Options{
		InitTargets: []string{url},
		InitCap:     2,
		MaxCap:      100,
		DialTimeout: time.Second * 30,
		IdleTimeout: time.Second * 60 * 60,
	}
	FrameworkGRPC, err = rpc_pool.RpcPool{}.Connect(options)
	if err != nil {
		println(" rpc连接池初始化失败" + err.Error())
	}
}
