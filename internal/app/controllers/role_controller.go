package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiniuer/goadmin/internal/app/forms"
	"github.com/zhiniuer/goadmin/internal/app/service"
	"github.com/zhiniuer/goutils/response"
	"github.com/zhiniuer/goutils/validator_util"
	"net/http"
)

type RoleController struct {
}

func (c *RoleController) List(ctx *gin.Context) {
	form := new(forms.AdminRoleListForm)
	if err := ctx.Bind(form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	pageSize := response.PageSize(ctx.Request)
	page := response.Page(ctx.Request)
	data, count, err := new(service.AdminRoleService).List(form, page, pageSize)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10002))
		return
	}
	pageData := response.Paginate(ctx.Request, count, data)
	ctx.JSON(http.StatusOK, response.Ok(pageData, "获取成功", 10000))
}

func (c *RoleController) Options(ctx *gin.Context) {
	data, err := new(service.AdminRoleService).Options()
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10002))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(data, "获取成功", 10000))
}

func (c *RoleController) Store(ctx *gin.Context) {
	form := new(forms.AdminRoleStoreForm)
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	err := new(service.AdminRoleService).Store(form)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10002))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok("", "保存成功", 10000))
}

func (c *RoleController) Info(ctx *gin.Context) {
	form := new(forms.AdminRoleInfoForm)
	if err := ctx.Bind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	data, err := new(service.AdminRoleService).Info(form)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10002))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(data, "获取成功", 10000))
}

func (c *RoleController) Detail(ctx *gin.Context) {
	form := new(forms.AdminRoleInfoForm)
	if err := ctx.Bind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	data, err := new(service.AdminRoleService).Detail(form)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10002))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(data, "获取成功", 10000))
}

func (c *RoleController) Delete(ctx *gin.Context) {
	form := new(forms.AdminRoleDeleteForm)
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	err := new(service.AdminRoleService).Delete(form)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10002))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok("", "删除成功", 10000))
}
