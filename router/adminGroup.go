package router

import (
	"server/controller"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRouterGroupInit(r *gin.Engine) {
	adminGroup := r.Group("/admin")

	adminGroup.Use(middleware.AdminAuthMiddleWare)
	{
		controller.ArticleController.ArticleRouterGroupInit(adminGroup)
		controller.AdminController.AdminRouterGroupInit(adminGroup)
		controller.CollectionController.CollectionRouterGroupInit(adminGroup)
		controller.AdController.AdRouterGroupInit(adminGroup)
		controller.LinkController.LinkRouterGroupInit(adminGroup)
		controller.TagController.TagRouterGroupInit(adminGroup)
		controller.MusicController.MusicRouterGroupInit(adminGroup)
	}
}
