package service

import (
	"errors"
	"math"
	"reflect"
	"server/global"
	"server/model"
	"server/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type AdminServiceStruct struct{}

func (adminServiceStruct *AdminServiceStruct) AddAdmin(admin *model.Admin, ctx *gin.Context) {

	result := global.Db.Create(&admin)

	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
		return
	}

	utils.Response.Success(ctx, &admin)
}

func (adminServiceStruct *AdminServiceStruct) RemoveAdmin(ctx *gin.Context, id string) {
	result := global.Db.Delete(&model.Admin{}, id)
	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
	}
	utils.Response.Success(ctx, true)
}

func (adminServiceStruct *AdminServiceStruct) UpdateAdmin(id string, admin *model.Admin, ctx *gin.Context) {

	result := global.Db.Where("id = ?", id).Updates(&admin)
	if result.Error != nil {
		utils.Response.Fail(ctx, result.Error)
		return
	}
	utils.Response.Success(ctx, &admin)
}

func (adminServiceStruct *AdminServiceStruct) QueryAdmin(ctx *gin.Context) {
	var dataList []model.Admin
	// 查询全部数据 or 查询分页数据
	pageSize, _ := strconv.Atoi(ctx.Query("size"))
	pageNum, _ := strconv.Atoi(ctx.Query("current"))
	username := ctx.DefaultQuery("username", "")

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
	result := global.Db.Model(&model.Admin{}).Where("username LIKE ?", "%"+username+"%").Select("username").Count(&total).Limit(pageSize).Offset(offsetVal).Order("created_at desc").Find(&dataList)

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

func (adminServiceStruct *AdminServiceStruct) DetailAdmin(ctx *gin.Context, id string) {
	admin := model.Admin{}
	global.Db.Where("id = ? ", id).Find(&admin)

	if reflect.ValueOf(admin.ID).IsZero() {
		utils.Response.BindError(ctx, errors.New("id不存在"))
		return
	}

	admin.Passworld = ""

	utils.Response.Success(ctx, admin)

}

func (adminServiceStruct *AdminServiceStruct) LoginAdmin(ctx *gin.Context, login *model.AdminLogin) {
	flag := true
	var admin model.Admin

	resul := global.Db.Where("username = ?", login.Username).First(&admin)
	if resul.Error != nil {
		utils.Response.Fail(ctx, errors.New("账户名或密码错误"))
		return
	}

	if reflect.ValueOf(admin.ID).IsZero() {
		flag = false
	} else {
		compare := strings.Compare(utils.EncryMd5(login.Passworld), admin.Passworld)

		if compare != 0 {
			flag = false
		}
	}
	if !flag {
		utils.Response.Fail(ctx, errors.New("账户名或密码错误"))
		return
	}

	admin.Passworld = ""

	tokenString, err := utils.GenToken(admin.ID, admin.Username)

	if err != nil {
		utils.Response.Fail(ctx, errors.New("token生成失败"))
	}

	utils.Response.Success(ctx, gin.H{
		"username": admin.Username,
		"token":    tokenString,
		"id":       admin.ID,
	})
}

var AdminService *AdminServiceStruct
