package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiniuer/goadmin/internal/app"
	"github.com/zhiniuer/goadmin/internal/app/errors"
	"github.com/zhiniuer/goadmin/internal/app/forms"
	"github.com/zhiniuer/goadmin/internal/app/service"
	"github.com/zhiniuer/goutils/hardware"
	"github.com/zhiniuer/goutils/ip"
	"github.com/zhiniuer/goutils/response"
	"github.com/zhiniuer/goutils/time_util"
	"github.com/zhiniuer/goutils/validator_util"
	"math"
	"net/http"
	"runtime"
	"time"
)

type AdminUserController struct {
}

// Status 系统检测
func (ctl AdminUserController) Status(ctx *gin.Context) {
	data := map[string]interface{}{
		"go_version": runtime.Version(),
		"timezone":   time.Local.String(),
		"time":       time_util.GetTimeDate(time.DateTime),
		"mode":       gin.Mode(),
		"ip":         ctx.Request.RemoteAddr,
		"ip2":        ip.GetRealIp(ctx.Request),
	}
	if ctx.Query("monitor") == "monitor" {
		h := hardware.Hardware()
		// 读取超全局变量即可
		var cpuPercent = int64(math.Floor(h.CpuPercent))
		data["cpu_percent"] = cpuPercent
		data["cpu_num"] = h.CpuNum
		data["mem_total"] = h.MemTotal
		data["mem_available"] = h.MemAvailable
		data["mem_free"] = h.MemFree
		data["mem_used"] = h.MemUsed
	}
	ctx.JSON(http.StatusOK, response.Ok(data, "获取成功", 10000))
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
	form := new(forms.AdminUserAuthInfoForm)
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
	form := new(forms.AdminUserAuthStoreForm)
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
