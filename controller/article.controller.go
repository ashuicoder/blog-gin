package controller

import (
	"errors"
	"reflect"
	"server/global"
	"server/model"
	"server/service"
	"server/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ArticleControllerScruct struct{}

func (articleController *ArticleControllerScruct) ArticleRouterGroupInit(r *gin.RouterGroup) {
	articleRouter := r.Group("/article")
	{
		articleRouter.POST("/", articleController.AddArticle)
		articleRouter.DELETE("/:id", articleController.RemoveArticle)
		articleRouter.PUT("/:id", articleController.UpdateArticle)
		articleRouter.GET("/", articleController.QueryArticle)
		articleRouter.GET("/:id", articleController.DetailArticle)
	}
}

func (articleController *ArticleControllerScruct) AddArticle(ctx *gin.Context) {
	articleMany := model.ArticleMany{}
	err := ctx.ShouldBindBodyWith(&articleMany, binding.JSON)

	if err != nil {
		utils.Response.BindError(ctx, err)
		return
	}

	article := model.Article{}
	err = ctx.ShouldBindBodyWith(&article, binding.JSON)
	if err != nil {
		utils.Response.BindError(ctx, err)
		return
	}

	if err != nil {
		utils.Response.BindError(ctx, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(&article)
	if err != nil {
		utils.Response.VaidateError(ctx, err)
		return
	}

	service.ArticleService.AddArticle(&article, &articleMany, ctx)
}

func (articleController *ArticleControllerScruct) RemoveArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	service.ArticleService.RemoveArticle(ctx, id)
}

func (articleController *ArticleControllerScruct) UpdateArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	article := model.Article{}
	global.Db.Select("id").Where("id = ? ", id).Find(&article)

	if reflect.ValueOf(article.ID).IsZero() {
		utils.Response.BindError(ctx, errors.New("id不存在"))
		return
	}

	articleMany := model.ArticleMany{}
	bindErr := ctx.ShouldBindBodyWith(&articleMany, binding.JSON)
	if bindErr != nil {
		utils.Response.BindError(ctx, bindErr)
		return
	}

	bindErr = ctx.ShouldBindBodyWith(&article, binding.JSON)
	if bindErr != nil {
		utils.Response.BindError(ctx, bindErr)
		return
	}

	validate := validator.New()
	validateErr := validate.Struct(&article)
	if validateErr != nil {
		utils.Response.VaidateError(ctx, validateErr)
		return
	}

	service.ArticleService.UpdateArticle(id, &article, &articleMany, ctx)
}

func (articleController *ArticleControllerScruct) QueryArticle(ctx *gin.Context) {
	service.ArticleService.QueryArticle(ctx)
}

func (articleController *ArticleControllerScruct) DetailArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	service.ArticleService.DetailArticle(ctx, id)
}

var ArticleController *ArticleControllerScruct
