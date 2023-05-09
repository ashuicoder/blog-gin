package controller

import (
	"errors"
	"reflect"
	"server/global"
	"server/model"
	"server/service"
	"server/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LinkControllerScruct struct{}

func (linkController *LinkControllerScruct) LinkRouterGroupInit(r *gin.RouterGroup) {
	linkRouter := r.Group("/link")
	{
		linkRouter.POST("/", linkController.LinkdLink)
		linkRouter.DELETE("/:id", linkController.RemoveLink)
		linkRouter.PUT("/:id", linkController.UpdateLink)
		linkRouter.GET("/", linkController.QueryLink)
		linkRouter.GET("/:id", linkController.DetailLink)
	}
}

func (linkController *LinkControllerScruct) LinkdLink(ctx *gin.Context) {
	link := model.Link{}

	err := ctx.ShouldBindJSON(&link)
	if err != nil {
		utils.Response.BindError(ctx, err)
		return
	}

	service.LinkService.LinkdLink(&link, ctx)

	validate := validator.New()
	err = validate.Struct(&link)
	if err != nil {
		utils.Response.VaidateError(ctx, err)
		return
	}

}

func (linkController *LinkControllerScruct) RemoveLink(ctx *gin.Context) {
	id := ctx.Param("id")
	service.LinkService.RemoveLink(ctx, id)
}

func (linkController *LinkControllerScruct) UpdateLink(ctx *gin.Context) {
	id := ctx.Param("id")
	link := model.Link{}
	global.Db.Select("id").Where("id = ? ", id).Find(&link)

	if reflect.ValueOf(link.ID).IsZero() {
		utils.Response.BindError(ctx, errors.New("id不存在"))
		return
	}

	bindErr := ctx.ShouldBindJSON(&link)
	if bindErr != nil {
		utils.Response.BindError(ctx, bindErr)
		return
	}

	validate := validator.New()
	validateErr := validate.Struct(&link)
	if validateErr != nil {
		utils.Response.VaidateError(ctx, validateErr)
		return
	}

	service.LinkService.UpdateLink(id, &link, ctx)
}

func (linkController *LinkControllerScruct) QueryLink(ctx *gin.Context) {
	service.LinkService.QueryLink(ctx)
}

func (linkController *LinkControllerScruct) DetailLink(ctx *gin.Context) {
	id := ctx.Param("id")
	service.LinkService.DetailLink(ctx, id)
}

var LinkController *LinkControllerScruct
