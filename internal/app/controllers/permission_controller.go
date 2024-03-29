package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-mod/ginrbac/internal/app/forms"
	"github.com/golang-mod/ginrbac/internal/app/service"
	"github.com/golang-mod/utils/response"
	"github.com/golang-mod/utils/validator_util"
	"net/http"
)

type PermissionController struct {
}

func (c *PermissionController) List(ctx *gin.Context) {
	form := new(forms.AdminRoleListForm)
	if err := ctx.Bind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	data, count, err := new(service.AdminPermissionsService).List(form, response.Page(ctx.Request), response.PageSize(ctx.Request))
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10001))
		return
	}
	pageData := response.Paginate(ctx.Request, count, data)
	ctx.JSON(http.StatusOK, response.Ok(pageData, "获取成功", 10000))
}

func (c *PermissionController) Options(ctx *gin.Context) {
	data, err := new(service.AdminPermissionsService).Options()
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10002))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(data, "获取成功", 10000))
}

func (c *PermissionController) Store(ctx *gin.Context) {
	form := new(forms.AdminPermissionStoreForm)
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	fmt.Println(form)
	err := new(service.AdminPermissionsService).Store(form)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10001))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok("", "保存成功", 10000))
}

func (c *PermissionController) Info(ctx *gin.Context) {
	form := new(forms.AdminPermissionsInfoForm)
	if err := ctx.Bind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	data, err := new(service.AdminPermissionsService).Info(form)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10001))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(data, "获取成功", 10000))
}

func (c *PermissionController) Detail(ctx *gin.Context) {
	form := new(forms.AdminPermissionsInfoForm)
	if err := ctx.Bind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	data, err := new(service.AdminPermissionsService).Detail(form)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10002))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(data, "获取成功", 10000))
}

func (c *PermissionController) Delete(ctx *gin.Context) {
	form := new(forms.AdminPermissionsDeleteForm)
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	err := new(service.AdminPermissionsService).Delete(form)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10001))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok("", "删除成功", 10000))
}
