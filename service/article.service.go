package service

import (
	"errors"
	"math"
	"reflect"
	"server/global"
	"server/model"
	"server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleServiceStruct struct{}

func (articleServiceStruct *ArticleServiceStruct) AddArticle(article *model.Article, articleMany *model.ArticleMany, ctx *gin.Context) {
	if len(articleMany.TagIds) > 0 {
		var tags []model.Tag
		for i := range articleMany.TagIds {
			tag := model.Tag{}
			global.Db.Where("id = ?", articleMany.TagIds[i]).First(&tag)
			tags = append(tags, tag)
		}
		article.Tags = tags
	}

	if len(articleMany.CollectionIds) > 0 {
		var collections []model.Collection
		for i := range articleMany.CollectionIds {
			collection := model.Collection{}
			global.Db.Where("id = ?", articleMany.CollectionIds[i]).First(&collection)
			collections = append(collections, collection)
		}
		article.Collections = collections
	}

	result := global.Db.Create(&article)

	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
		return
	}

	utils.Response.Success(ctx, &article)

}

func (articleServiceStruct *ArticleServiceStruct) RemoveArticle(ctx *gin.Context, id string) {

	result := global.Db.Delete(&model.Article{}, id)
	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
	}
	utils.Response.Success(ctx, true)
}

func (articleServiceStruct *ArticleServiceStruct) UpdateArticle(id string, article *model.Article, articleMany *model.ArticleMany, ctx *gin.Context) {
	if len(articleMany.TagIds) > 0 {
		var tags []model.Tag
		for i := range articleMany.TagIds {
			tag := model.Tag{}
			global.Db.Where("id = ?", articleMany.TagIds[i]).First(&tag)
			tags = append(tags, tag)
		}
		article.Tags = tags
	}

	if len(articleMany.CollectionIds) > 0 {
		var collections []model.Collection
		for i := range articleMany.CollectionIds {
			collection := model.Collection{}
			global.Db.Where("id = ?", articleMany.CollectionIds[i]).First(&collection)
			collections = append(collections, collection)
		}
		article.Collections = collections
	}
	result := global.Db.Where("id = ?", id).Updates(&article)
	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
		return
	}
	utils.Response.Success(ctx, &article)
}

func (articleServiceStruct *ArticleServiceStruct) QueryArticle(ctx *gin.Context) {
	var dataList []model.Article
	// 查询全部数据 or 查询分页数据
	pageSize, _ := strconv.Atoi(ctx.Query("size"))
	pageNum, _ := strconv.Atoi(ctx.Query("current"))
	title := ctx.DefaultQuery("title", "")

	if reflect.ValueOf(pageSize).IsZero() || reflect.ValueOf(pageNum).IsZero() {
		utils.Response.Fail(ctx, errors.New("分页参数错误"))
		return
	}

	offsetVal := (pageNum - 1) * pageSize // 固定写法 记住就行
	if pageNum == -1 && pageSize == -1 {
		offsetVal = -1
	}

	// 返回一个总数
	var total int64

	// 查询数据库
	result := global.Db.Model(&model.Article{}).Where("title LIKE ?", "%"+title+"%").Preload("Tags").Preload("Collections").Count(&total).Limit(pageSize).Offset(offsetVal).Order("created_at desc").Find(&dataList)

	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
		return
	}

	utils.Response.Success(ctx, gin.H{
		"total":   total,
		"current": pageNum,
		"size":    pageSize,
		"pages":   math.Ceil(float64(total) / float64(pageSize)),
		"records": dataList,
	})

}

func (articleServiceStruct *ArticleServiceStruct) DetailArticle(ctx *gin.Context, id string) {
	article := model.Article{}
	global.Db.Where("id = ?", id).Preload("Tags").Preload("Collections").First(&article)

	if reflect.ValueOf(article.ID).IsZero() {
		utils.Response.BindError(ctx, errors.New("id不存在"))
		return
	}

	utils.Response.Success(ctx, article)

	view := article.View
	global.Db.Model(&article).Select("view").Updates(map[string]interface{}{"view": view + 1})
}

var ArticleService *ArticleServiceStruct
