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

type TagControllerScruct struct{}

func (tagController *TagControllerScruct) TagRouterGroupInit(r *gin.RouterGroup) {
	tagRouter := r.Group("/tag")
	{
		tagRouter.POST("/", tagController.TagdTag)
		tagRouter.DELETE("/:id", tagController.RemoveTag)
		tagRouter.PUT("/:id", tagController.UpdateTag)
		tagRouter.GET("/", tagController.QueryTag)
		tagRouter.GET("/:id", tagController.DetailTag)
	}
}

func (tagController *TagControllerScruct) TagdTag(ctx *gin.Context) {
	tag := model.Tag{}

	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		utils.Response.BindError(ctx, err)
		return
	}

	service.TagService.TagdTag(&tag, ctx)

	validate := validator.New()
	err = validate.Struct(&tag)
	if err != nil {
		utils.Response.VaidateError(ctx, err)
		return
	}

}

func (tagController *TagControllerScruct) RemoveTag(ctx *gin.Context) {
	id := ctx.Param("id")
	service.TagService.RemoveTag(ctx, id)
}

func (tagController *TagControllerScruct) UpdateTag(ctx *gin.Context) {
	id := ctx.Param("id")
	tag := model.Tag{}
	global.Db.Select("id").Where("id = ? ", id).Find(&tag)

	if reflect.ValueOf(tag.ID).IsZero() {
		utils.Response.BindError(ctx, errors.New("id不存在"))
		return
	}

	bindErr := ctx.ShouldBindJSON(&tag)
	if bindErr != nil {
		utils.Response.BindError(ctx, bindErr)
		return
	}

	validate := validator.New()
	validateErr := validate.Struct(&tag)
	if validateErr != nil {
		utils.Response.VaidateError(ctx, validateErr)
		return
	}

	service.TagService.UpdateTag(id, &tag, ctx)
}

func (tagController *TagControllerScruct) QueryTag(ctx *gin.Context) {
	service.TagService.QueryTag(ctx)
}

func (tagController *TagControllerScruct) DetailTag(ctx *gin.Context) {
	id := ctx.Param("id")
	service.TagService.DetailTag(ctx, id)
}

var TagController *TagControllerScruct
