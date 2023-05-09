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

type LinkServiceStruct struct{}

func (linkServiceStruct *LinkServiceStruct) LinkdLink(link *model.Link, ctx *gin.Context) {

	result := global.Db.Create(&link)

	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
		return
	}

	utils.Response.Success(ctx, &link)
}

func (linkServiceStruct *LinkServiceStruct) RemoveLink(ctx *gin.Context, id string) {
	result := global.Db.Delete(&model.Link{}, id)
	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
	}
	utils.Response.Success(ctx, true)
}

func (linkServiceStruct *LinkServiceStruct) UpdateLink(id string, link *model.Link, ctx *gin.Context) {

	result := global.Db.Where("id = ?", id).Updates(&link)
	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
		return
	}
	utils.Response.Success(ctx, &link)
}

func (linkServiceStruct *LinkServiceStruct) QueryLink(ctx *gin.Context) {
	var dataList []model.Link
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
	result := global.Db.Model(&model.Link{}).Where("title LIKE ?", "%"+title+"%").Count(&total).Limit(pageSize).Offset(offsetVal).Order("created_at desc").Find(&dataList)

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

func (linkServiceStruct *LinkServiceStruct) DetailLink(ctx *gin.Context, id string) {
	link := model.Link{}
	global.Db.Where("id = ?", id).First(&link)

	if reflect.ValueOf(link.ID).IsZero() {
		utils.Response.BindError(ctx, errors.New("id不存在"))
		return
	}

	utils.Response.Success(ctx, link)

	view := link.View
	global.Db.Model(&link).Select("view").Updates(map[string]interface{}{"view": view + 1})
}

var LinkService *LinkServiceStruct
