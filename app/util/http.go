package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jtzjtz/kit/http"
	"github.com/jtzjtz/ys_api/app/config"
	"net/url"
	"time"
)

func handleException(reqUrl string, params interface{}, method string, timeOut time.Duration, headers map[string]string, res string, code int, err error) {
	if code != 200 || err != nil {
		msg := fmt.Sprintf("请求地址:%s<br/>请求参数:%v<br/>请求方式:%s<br/>请求超时时间:%v<br/>请求头:%v<br/>响应数据:%s<br/>响应状态码:%d<br/>", reqUrl, params, method, timeOut, headers, res, code)
		if err != nil {
			msg += "错误信息:" + err.Error()
		}
		_, lErr := LogError(msg)
		if lErr != nil {
			_, _ = AddLog(config.Configs.Env+"-api-zgenquiry-handleException.log", map[string]interface{}{"msg": msg}, false)
		}
	}
}

/**
timeout 超时秒数  可传0 默认30秒
headers   可传nil
*/
func GetRequest(reqUrl string, params map[string]interface{}, timeOut time.Duration, headers map[string]string) (resBody string, resErr error) {
	res, code, err := http.GetRequest(&gin.Context{}, reqUrl, params, timeOut, http.ExtParams{Headers: headers})
	handleException(reqUrl, params, "GET", timeOut, headers, res, code, err)
	return res, err
}

/**
timeout 超时秒数  可传0 默认30秒
headers   可传nil
*/
func PostRequest(reqUrl string, params map[string]interface{}, timeOut time.Duration, headers map[string]string) (resBody string, resErr error) {
	res, code, err := http.PostRequest(&gin.Context{}, reqUrl, params, timeOut, http.ExtParams{Headers: headers})
	handleException(reqUrl, params, "POST", timeOut, headers, res, code, err)
	return res, err
}

/**
timeout 超时秒数  可传0 默认30秒
headers   可传nil
*/
func PostJsonRequest(reqUrl string, params map[string]interface{}, timeOut time.Duration, headers map[string]string) (resBody string, resErr error) {
	res, code, err := http.PostJsonRequest(&gin.Context{}, reqUrl, params, timeOut, http.ExtParams{Headers: headers})
	handleException(reqUrl, params, "POST-JSON", timeOut, headers, res, code, err)
	return res, err
}

/**
timeout 超时秒数  可传0 默认30秒
headers   可传nil
*/
func DoRequest(reqUrl string, params string, method string, timeOut time.Duration, headers map[string]string) (resBody string, resErr error) {
	res, code, err := http.DoRequest(&gin.Context{}, reqUrl, params, method, timeOut, http.ExtParams{Headers: headers})
	handleException(reqUrl, params, method, timeOut, headers, res, code, err)
	return res, err
}

/**
timeout 超时秒数  可传0 默认30秒
headers   可传nil
*/
func PostUrlValuesByForm(reqUrl string, params url.Values, timeOut time.Duration, headers map[string]string) (string, error) {

	paramsStr := params.Encode()
	res, code, err := http.DoRequest(&gin.Context{}, reqUrl, paramsStr, "POST", timeOut, http.ExtParams{Headers: headers})
	handleException(reqUrl, params, "POST", timeOut, headers, res, code, err)
	return res, err
}
