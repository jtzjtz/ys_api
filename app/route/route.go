package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jtzjtz/kit/middleware/access_log"
	"github.com/jtzjtz/kit/middleware/recover_catch"
	"github.com/jtzjtz/ys_api/app/controller/http_controller"
	"github.com/jtzjtz/ys_api/app/controller/sq_user_blacklist_controller"
	"github.com/jtzjtz/ys_api/app/route/middleware"
)

//默认中间件
func setupMiddleware(engin *gin.Engine) {
	env := ""
	engin.Use(access_log.AccessLogMiddleware(env, "api-framework"))
	engin.Use(recover_catch.RecoverCatchMiddleware(env, "api-framework"))
}

//设置路由
func Init(engin *gin.Engine) {

	setupMiddleware(engin) //设置默认中间件
	//验签中间件demo
	engin.GET("/check_sign", middleware.CheckSign, func(c *gin.Context) {
		c.String(200, "验证通过")
	})

	engin.GET("/panic", func(c *gin.Context) {
		tmp := 0
		fmt.Println(100 / tmp)
		c.String(200, "测试异常")
	})

	httpT := engin.Group("/http")
	{
		httpT.GET("/get", http_controller.Get)
		httpT.GET("/post", http_controller.Post)
		httpT.GET("/postjson", http_controller.PostJson)
		httpT.GET("/request", http_controller.Request)
		httpT.GET("/request_timeout", http_controller.RequestTimeout)
		httpT.Any("/req_sleep", http_controller.ReqSleep)
	}

	user := engin.Group("/user")
	{

		user.GET("/getlist", sq_user_blacklist_controller.GetUserBlackList)
		user.GET("/create", sq_user_blacklist_controller.CreateUserBlacklist)
		user.POST("/create", sq_user_blacklist_controller.CreateUserWithBind)

	}

	engin.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	engin.Any("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})
}
