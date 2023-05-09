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

type MusicControllerScruct struct{}

func (musicController *MusicControllerScruct) MusicRouterGroupInit(r *gin.RouterGroup) {
	musicRouter := r.Group("/music")
	{
		musicRouter.POST("/", musicController.MusicdMusic)
		musicRouter.DELETE("/:id", musicController.RemoveMusic)
		musicRouter.PUT("/:id", musicController.UpdateMusic)
		musicRouter.GET("/", musicController.QueryMusic)
		musicRouter.GET("/:id", musicController.DetailMusic)
	}
}

func (musicController *MusicControllerScruct) MusicdMusic(ctx *gin.Context) {
	music := model.Music{}

	err := ctx.ShouldBindJSON(&music)
	if err != nil {
		utils.Response.BindError(ctx, err)
		return
	}

	service.MusicService.MusicdMusic(&music, ctx)

	validate := validator.New()
	err = validate.Struct(&music)
	if err != nil {
		utils.Response.VaidateError(ctx, err)
		return
	}

}

func (musicController *MusicControllerScruct) RemoveMusic(ctx *gin.Context) {
	id := ctx.Param("id")
	service.MusicService.RemoveMusic(ctx, id)
}

func (musicController *MusicControllerScruct) UpdateMusic(ctx *gin.Context) {
	id := ctx.Param("id")
	music := model.Music{}
	global.Db.Select("id").Where("id = ? ", id).Find(&music)

	if reflect.ValueOf(music.ID).IsZero() {
		utils.Response.BindError(ctx, errors.New("id不存在"))
		return
	}

	bindErr := ctx.ShouldBindJSON(&music)
	if bindErr != nil {
		utils.Response.BindError(ctx, bindErr)
		return
	}

	validate := validator.New()
	validateErr := validate.Struct(&music)
	if validateErr != nil {
		utils.Response.VaidateError(ctx, validateErr)
		return
	}

	service.MusicService.UpdateMusic(id, &music, ctx)
}

func (musicController *MusicControllerScruct) QueryMusic(ctx *gin.Context) {
	service.MusicService.QueryMusic(ctx)
}

func (musicController *MusicControllerScruct) DetailMusic(ctx *gin.Context) {
	id := ctx.Param("id")
	service.MusicService.DetailMusic(ctx, id)
}

var MusicController *MusicControllerScruct
