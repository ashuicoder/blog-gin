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

type AdminControllerScruct struct{}

func (adminController *AdminControllerScruct) AdminRouterGroupInit(r *gin.RouterGroup) {
	adminRouter := r.Group("/admin")
	{
		adminRouter.POST("/", adminController.AddAdmin)
		adminRouter.DELETE("/:id", adminController.RemoveAdmin)
		adminRouter.PUT("/:id", adminController.UpdateAdmin)
		adminRouter.GET("/", adminController.QueryAdmin)
		adminRouter.GET("/:id", adminController.DetailAdmin)
		adminRouter.POST("/login", adminController.LoginAdmin)
	}
}

func (adminController *AdminControllerScruct) AddAdmin(ctx *gin.Context) {
	admin := model.Admin{}

	err := ctx.ShouldBindJSON(&admin)
	if err != nil {
		utils.Response.BindError(ctx, err)
		return
	}

	validate := validator.New()
	err = validate.Struct(admin)
	if err != nil {
		utils.Response.VaidateError(ctx, err)
		return
	}

	admin.Passworld = utils.EncryMd5(admin.Passworld)

	service.AdminService.AddAdmin(&admin, ctx)
}

func (adminController *AdminControllerScruct) RemoveAdmin(ctx *gin.Context) {
	id := ctx.Param("id")
	service.AdminService.RemoveAdmin(ctx, id)
}

func (adminController *AdminControllerScruct) UpdateAdmin(ctx *gin.Context) {
	id := ctx.Param("id")
	admin := model.Admin{}
	global.Db.Select("id").Where("id = ? ", id).Find(&admin)

	if reflect.ValueOf(admin.ID).IsZero() {
		utils.Response.BindError(ctx, errors.New("id不存在"))
		return
	}

	bindErr := ctx.ShouldBindJSON(&admin)
	if bindErr != nil {
		utils.Response.BindError(ctx, bindErr)
		return
	}

	validate := validator.New()
	validateErr := validate.Struct(admin)
	if validateErr != nil {
		utils.Response.VaidateError(ctx, validateErr)
		return
	}

	service.AdminService.UpdateAdmin(id, &admin, ctx)
}

func (adminController *AdminControllerScruct) QueryAdmin(ctx *gin.Context) {
	service.AdminService.QueryAdmin(ctx)
}

func (adminController *AdminControllerScruct) DetailAdmin(ctx *gin.Context) {
	id := ctx.Param("id")
	service.AdminService.DetailAdmin(ctx, id)
}

func (adminController *AdminControllerScruct) LoginAdmin(ctx *gin.Context) {
	login := model.AdminLogin{}

	bindErr := ctx.ShouldBindJSON(&login)
	if bindErr != nil {
		utils.Response.BindError(ctx, bindErr)
		return
	}

	validate := validator.New()
	validateErr := validate.Struct(&login)
	if validateErr != nil {
		utils.Response.VaidateError(ctx, validateErr)
		return
	}

	service.AdminService.LoginAdmin(ctx, &login)
}

var AdminController *AdminControllerScruct
