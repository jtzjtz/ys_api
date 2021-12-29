package config

import "github.com/jtzjtz/kit/config"

type appConfig struct {
	GrpcServiceUrl string `json:"grpc_service_url"`
}

func (t *appConfig) Reload() error {
	println("reload app config")
	return nil
}
func (t *appConfig) IsLoad() bool {
	if len(t.GrpcServiceUrl) > 0 {
		return true
	}
	return false
}

const (
	//应用常量

	LOG_ROOT_PATH = "/data/logs/"
	LOG_EXTENSION = ".log"
	FORMAT_DATE   = `2006-01-02 15:04:05`
)

var (
	client config.Client
	App    appConfig
)

func Init(e string) (err error) {

	App = appConfig{GrpcServiceUrl: "localhost:8090"}
	//
	//nacasConfig := config.NacosConfig{
	//	Address:             "10.31.114.24",
	//	Port:                8848,
	//	Scheme:              "http",
	//	ContextPath:         "/nacos",
	//	NameSpaceId:         "fffaacb2-a7f7-45d5-83aa-dec49122a098",
	//	TimeoutMs:           5000,
	//	NotLoadCacheAtStart: true,
	//	LogDir:              "./.config/log",
	//	CacheDir:            "./.config",
	//	RotateTime:          "1h",
	//	MaxAge:              3,
	//	LogLevel:            "debug",
	//}
	//client, err = config.InitClient(nacasConfig)
	//if err != nil {
	//	return
	//}
	//client.GetDataConfig("app", "api-service", &App, true)
	return
}
