package util

import (
	"fmt"
	"github.com/jtzjtz/kit/log"
	"github.com/jtzjtz/ys_api/app/config"
	"runtime"
	"time"
)

func LogError(msg string) {
	emailContent := fmt.Sprintf("当前时间:%s<br/>%s", time.Now().Format(config.FORMAT_DATE), msg)
	_, file, line, _ := runtime.Caller(1)
	emailContent += fmt.Sprintf("<br>%s %v", file, line)
	m := map[string]interface{}{}
	m["msg"] = emailContent
	log.AddLog("./.log/", "error.log", m, false)
}

//添加本地日志  isAsync：是否使用异步记录
func AddLog(fileName string, data map[string]interface{}, isAsync bool) (bool, error) {
	data["time"] = time.Now().Format(config.FORMAT_DATE)
	return log.AddLog(config.LOG_ROOT_PATH, fileName, data, isAsync)
}
