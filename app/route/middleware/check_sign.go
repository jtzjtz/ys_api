package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jtzjtz/kit/sign"
	"github.com/jtzjtz/ys_api/app"
	"net/http"
)

//签名验证
func CheckSign(ctx *gin.Context) {
	_ = ctx.Request.ParseForm()
	token := ctx.Request.FormValue("token")
	secret := "ddd"
	if token == "" || sign.CheckSn(ctx.Request.PostForm, token, secret, false) == false {
		ctx.JSON(http.StatusOK, app.Result{Code: app.CODE_ERROR, Msg: "签名未通过"})
		ctx.Abort()
	}
	ctx.Next()
}
