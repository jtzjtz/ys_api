package main

import (
	"flag"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/jtzjtz/ys_api/app/config"
	grpcPool "github.com/jtzjtz/ys_api/app/conn/grpc"
	"github.com/jtzjtz/ys_api/app/route"
)

var ENV string

//初始化环境变量
func init() {
	env := flag.String("env", "", "请输入运行环境:\n fat:测试环境\n uat:预上线环境\n product:正式环境\n")
	flag.Parse()
	ENV = *env
	if ENV == "" {
		println("env 不能为空")
		//os.Exit(1)
	}
	config.Init(ENV) //需要第一个加载
	grpcPool.Init()  // 初始化 grpc 连接池
	//redisPool.Init() // 初始化 redis 连接池
}

//启动服务
func main() {
	var port = ":8080"
	var engin *gin.Engine
	gin.SetMode(gin.DebugMode)
	engin = gin.New()
	pprof.Register(engin) // 性能分析

	route.Init(engin) //路由初始化
	engin.Run(port)
}
