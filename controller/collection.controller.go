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

type CollectionControllerScruct struct{}

func (collectionController *CollectionControllerScruct) CollectionRouterGroupInit(r *gin.RouterGroup) {
	collectionRouter := r.Group("/collection")
	{
		collectionRouter.POST("/", collectionController.AddCollection)
		collectionRouter.DELETE("/:id", collectionController.RemoveCollection)
		collectionRouter.PUT("/:id", collectionController.UpdateCollection)
		collectionRouter.GET("/", collectionController.QueryCollection)
		collectionRouter.GET("/:id", collectionController.DetailCollection)
	}
}

func (collectionController *CollectionControllerScruct) AddCollection(ctx *gin.Context) {
	collection := model.Collection{}

	err := ctx.ShouldBindJSON(&collection)
	if err != nil {
		utils.Response.BindError(ctx, err)
		return
	}

	service.CollectionService.AddCollection(&collection, ctx)

	validate := validator.New()
	err = validate.Struct(&collection)
	if err != nil {
		utils.Response.VaidateError(ctx, err)
		return
	}

}

func (collectionController *CollectionControllerScruct) RemoveCollection(ctx *gin.Context) {
	id := ctx.Param("id")
	service.CollectionService.RemoveCollection(ctx, id)
}

func (collectionController *CollectionControllerScruct) UpdateCollection(ctx *gin.Context) {
	id := ctx.Param("id")
	collection := model.Collection{}
	global.Db.Select("id").Where("id = ? ", id).Find(&collection)

	if reflect.ValueOf(collection.ID).IsZero() {
		utils.Response.BindError(ctx, errors.New("id不存在"))
		return
	}

	bindErr := ctx.ShouldBindJSON(&collection)
	if bindErr != nil {
		utils.Response.BindError(ctx, bindErr)
		return
	}

	validate := validator.New()
	validateErr := validate.Struct(&collection)
	if validateErr != nil {
		utils.Response.VaidateError(ctx, validateErr)
		return
	}

	service.CollectionService.UpdateCollection(id, &collection, ctx)
}

func (collectionController *CollectionControllerScruct) QueryCollection(ctx *gin.Context) {
	service.CollectionService.QueryCollection(ctx)
}

func (collectionController *CollectionControllerScruct) DetailCollection(ctx *gin.Context) {
	id := ctx.Param("id")
	service.CollectionService.DetailCollection(ctx, id)
}

var CollectionController *CollectionControllerScruct
