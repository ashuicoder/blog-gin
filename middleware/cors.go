package middleware

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5174"},
		AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{"*"},
		// ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	},
	)
}
