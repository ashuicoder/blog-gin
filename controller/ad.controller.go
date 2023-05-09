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

type AdControllerScruct struct{}

func (adController *AdControllerScruct) AdRouterGroupInit(r *gin.RouterGroup) {
	adRouter := r.Group("/ad")
	{
		adRouter.POST("/", adController.AddAd)
		adRouter.DELETE("/:id", adController.RemoveAd)
		adRouter.PUT("/:id", adController.UpdateAd)
		adRouter.GET("/", adController.QueryAd)
		adRouter.GET("/:id", adController.DetailAd)
	}
}

func (adController *AdControllerScruct) AddAd(ctx *gin.Context) {
	ad := model.Ad{}

	err := ctx.ShouldBindJSON(&ad)
	if err != nil {
		utils.Response.BindError(ctx, err)
		return
	}

	service.AdService.AddAd(&ad, ctx)

	validate := validator.New()
	err = validate.Struct(&ad)
	if err != nil {
		utils.Response.VaidateError(ctx, err)
		return
	}

}

func (adController *AdControllerScruct) RemoveAd(ctx *gin.Context) {
	id := ctx.Param("id")
	service.AdService.RemoveAd(ctx, id)
}

func (adController *AdControllerScruct) UpdateAd(ctx *gin.Context) {
	id := ctx.Param("id")
	ad := model.Ad{}
	global.Db.Select("id").Where("id = ? ", id).Find(&ad)

	if reflect.ValueOf(ad.ID).IsZero() {
		utils.Response.BindError(ctx, errors.New("id不存在"))
		return
	}

	bindErr := ctx.ShouldBindJSON(&ad)
	if bindErr != nil {
		utils.Response.BindError(ctx, bindErr)
		return
	}

	validate := validator.New()
	validateErr := validate.Struct(&ad)
	if validateErr != nil {
		utils.Response.VaidateError(ctx, validateErr)
		return
	}

	service.AdService.UpdateAd(id, &ad, ctx)
}

func (adController *AdControllerScruct) QueryAd(ctx *gin.Context) {
	service.AdService.QueryAd(ctx)
}

func (adController *AdControllerScruct) DetailAd(ctx *gin.Context) {
	id := ctx.Param("id")
	service.AdService.DetailAd(ctx, id)
}

var AdController *AdControllerScruct
