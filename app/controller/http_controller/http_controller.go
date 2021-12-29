package http_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jtzjtz/kit/http"
	"time"
)

var apiUrl = "http://baidu.com"

func Get(ctx *gin.Context) {
	res, code, err := http.GetRequest(ctx, apiUrl, nil, 5*time.Second, http.ExtParams{})
	fmt.Println(res, code, err)
	ctx.JSON(200, map[string]interface{}{"method": "GET", "res": res, "code": code, "err": err})
}

func Post(ctx *gin.Context) {
	res, code, err := http.PostRequest(ctx, apiUrl, map[string]interface{}{"aa": 11, "bb": 22222}, 5*time.Second, http.ExtParams{})
	fmt.Println(res, code, err)
	ctx.JSON(200, map[string]interface{}{"method": "POST", "res": res, "code": code, "err": err})
}

func PostJson(ctx *gin.Context) {
	res, code, err := http.PostJsonRequest(ctx, apiUrl, map[string]interface{}{"aa": 11, "bb": 22222}, 5*time.Second, http.ExtParams{})
	fmt.Println(res, code, err)
	ctx.JSON(200, map[string]interface{}{"method": "POST_JSON", "res": res, "code": code, "err": err})
}

func Request(ctx *gin.Context) {
	res, code, err := http.DoRequest(ctx, apiUrl, "aa=11&bb=2", "GET", 5*time.Second, http.ExtParams{})
	fmt.Println(res, code, err)
	ctx.JSON(200, map[string]interface{}{"method": "Request", "res": res, "code": code, "err": err})
}
func RequestTimeout(ctx *gin.Context) {
	res, code, err := http.GetRequest(ctx, apiUrl+"/http/req_sleep", nil, time.Second, http.ExtParams{})
	fmt.Println(res, code, err)
	ctx.JSON(200, map[string]interface{}{"method": "GET", "res": res, "code": code, "err": err})
}

func ReqSleep(ctx *gin.Context) {
	time.Sleep(time.Second * 3)
	ctx.String(200, "hello world")
}
