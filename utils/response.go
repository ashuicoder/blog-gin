package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseStuct struct{}

func (responseStuct ResponseStuct) Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": "000000",
		"msg":  "success",
		"data": data,
	})
}

func (responseStuct ResponseStuct) Fail(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": "-1",
		"msg":  err.Error(),
	})
}

func (responseStuct ResponseStuct) BindError(ctx *gin.Context, err error) {
	responseStuct.Fail(ctx, err)
}

func (responseStuct ResponseStuct) VaidateError(ctx *gin.Context, err error) {
	responseStuct.Fail(ctx, err)
}

func (responseStuct ResponseStuct) UnAuth(ctx *gin.Context, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": "401",
		"msg":  msg,
	})
}

var Response = ResponseStuct{}
