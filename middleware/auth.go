package middleware

import (
	"reflect"
	"server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleWare(ctx *gin.Context) {

	if ctx.Request.URL.Path == "/admin/admin/login" {
		ctx.Next()
		return
	}

	authorization := ctx.Request.Header.Get("Authorization")
	if reflect.ValueOf(authorization).IsZero() {
		utils.Response.UnAuth(ctx, "无效的token")
		ctx.Abort()
		return
	}

	parts := strings.SplitN(authorization, " ", 2)

	if !(len(parts) == 2 && parts[0] == "Bearer") {
		utils.Response.UnAuth(ctx, "无效的token")
		ctx.Abort()
		return
	}

	mc, err := utils.ParseToken(parts[1])
	if err != nil {
		utils.Response.UnAuth(ctx, "无效的token")
		ctx.Abort()
		return
	}

	ctx.Set("admin", mc)

	ctx.Next()
}
