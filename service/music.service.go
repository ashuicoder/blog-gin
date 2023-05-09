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

type MusicServiceStruct struct{}

func (musicServiceStruct *MusicServiceStruct) MusicdMusic(music *model.Music, ctx *gin.Context) {

	result := global.Db.Create(&music)

	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
		return
	}

	utils.Response.Success(ctx, &music)
}

func (musicServiceStruct *MusicServiceStruct) RemoveMusic(ctx *gin.Context, id string) {
	result := global.Db.Delete(&model.Music{}, id)
	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
	}
	utils.Response.Success(ctx, true)
}

func (musicServiceStruct *MusicServiceStruct) UpdateMusic(id string, music *model.Music, ctx *gin.Context) {

	result := global.Db.Where("id = ?", id).Updates(&music)
	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
		return
	}
	utils.Response.Success(ctx, &music)
}

func (musicServiceStruct *MusicServiceStruct) QueryMusic(ctx *gin.Context) {
	var dataList []model.Music
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
	result := global.Db.Model(&model.Music{}).Where("title LIKE ?", "%"+title+"%").Count(&total).Limit(pageSize).Offset(offsetVal).Order("created_at desc").Find(&dataList)

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

func (musicServiceStruct *MusicServiceStruct) DetailMusic(ctx *gin.Context, id string) {
	music := model.Music{}
	global.Db.Where("id = ?", id).First(&music)

	if reflect.ValueOf(music.ID).IsZero() {
		utils.Response.BindError(ctx, errors.New("id不存在"))
		return
	}

	utils.Response.Success(ctx, music)

}

var MusicService *MusicServiceStruct
