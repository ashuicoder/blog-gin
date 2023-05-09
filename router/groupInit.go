package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouterGroup(r *gin.Engine) {
	AdminRouterGroupInit(r)
}
