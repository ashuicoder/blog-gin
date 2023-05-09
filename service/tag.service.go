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

type TagServiceStruct struct{}

func (tagServiceStruct *TagServiceStruct) TagdTag(tag *model.Tag, ctx *gin.Context) {

	result := global.Db.Create(&tag)

	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
		return
	}

	utils.Response.Success(ctx, &tag)
}

func (tagServiceStruct *TagServiceStruct) RemoveTag(ctx *gin.Context, id string) {
	result := global.Db.Delete(&model.Tag{}, id)
	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
	}
	utils.Response.Success(ctx, true)
}

func (tagServiceStruct *TagServiceStruct) UpdateTag(id string, tag *model.Tag, ctx *gin.Context) {

	result := global.Db.Where("id = ?", id).Updates(&tag)
	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
		return
	}
	utils.Response.Success(ctx, &tag)
}

func (tagServiceStruct *TagServiceStruct) QueryTag(ctx *gin.Context) {
	var dataList []model.Tag
	// 查询全部数据 or 查询分页数据
	pageSize, _ := strconv.Atoi(ctx.Query("size"))
	pageNum, _ := strconv.Atoi(ctx.Query("current"))
	name := ctx.DefaultQuery("name", "")

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
	result := global.Db.Model(&model.Tag{}).Where("name LIKE ?", "%"+name+"%").Count(&total).Limit(pageSize).Offset(offsetVal).Order("created_at desc").Find(&dataList)

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

func (tagServiceStruct *TagServiceStruct) DetailTag(ctx *gin.Context, id string) {
	tag := model.Tag{}
	global.Db.Where("id = ?", id).First(&tag)

	if reflect.ValueOf(tag.ID).IsZero() {
		utils.Response.BindError(ctx, errors.New("id不存在"))
		return
	}

	utils.Response.Success(ctx, tag)

}

var TagService *TagServiceStruct
