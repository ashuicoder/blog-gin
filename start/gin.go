package start

import (
	"server/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ServerStart() {
	r := gin.Default()

	r.Use(cors.Default())
	router.InitRouterGroup(r)
	r.Run(":3000")
}
