package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiniuer/goadmin/internal/app"
	"github.com/zhiniuer/goadmin/internal/app/errors"
	"github.com/zhiniuer/goadmin/internal/app/schema"
	"github.com/zhiniuer/goadmin/internal/app/service"
	"github.com/zhiniuer/goutils/response"
	"github.com/zhiniuer/goutils/validator_util"
	"net/http"
)

type AdminUserController struct {
}

func (ctl AdminUserController) List(ctx *gin.Context) {
	data, count, err := new(service.AdminUserService).List(response.Page(ctx.Request), response.PageSize(ctx.Request))
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10002))
		return
	}
	pageData := response.Paginate(ctx.Request, count, data)
	ctx.JSON(http.StatusOK, response.Ok(pageData, "成功", 10000))
}
func (ctl AdminUserController) Options(ctx *gin.Context) {
	data, err := new(service.AdminUserService).Options()
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10002))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(data, "成功", 10000))
}

func (ctl AdminUserController) AuthInfo(ctx *gin.Context) {
	form := new(schema.AdminUserAuthInfoForm)
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	res, err := new(service.AdminUserService).AuthInfo(form)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10002))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(res, "成功", 10000))
}

func (ctl AdminUserController) AuthStore(ctx *gin.Context) {
	form := new(schema.AdminUserAuthStoreForm)
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	err := new(service.AdminUserService).AuthStore(form)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10002))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok("", "分配成功", 10000))
}

func (ctl AdminUserController) Menu(ctx *gin.Context) {
	kc := app.NewContext(ctx)
	userId := kc.User.AdminId
	data, err := new(service.AdminUserService).Menu(int(userId))
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10002))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(data, "获取成功", 10000))
}

type UserLoginForm struct {
	Username string `form:"username" json:"username" binding:"required" name:"用户名"`
	Password string `form:"password" json:"password" binding:"required" name:"密码"`
}

type UserPasswordForm struct {
	CheckPassword string `form:"checkPassword" json:"checkPassword" binding:"required"`
	Password      string `form:"password" json:"password" binding:"required"`
}

func (ctl AdminUserController) UserInfo(ctx *gin.Context) {
	//var user models.AdminUsers

	ac := app.NewContext(ctx)

	//_ = driver.GDB.Where("id = ?", ac.User.UserId).First(&user).Error
	adminUserService := service.AdminUserService{}
	user, err := adminUserService.UserInfo(ac)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(errors.E(err).MessageCode()))
	}
	menus, _ := adminUserService.Menu(int(ac.User.AdminId))
	ctx.JSON(http.StatusOK, response.Ok(gin.H{
		"id":    user.Id,
		"name":  user.Name,
		"menus": menus,
	}, "成功", 10000))
}
