package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiniuer/goadmin/internal/app/schema"
	"github.com/zhiniuer/goadmin/internal/app/service"
	"github.com/zhiniuer/goutils/response"
	"github.com/zhiniuer/goutils/validator_util"
	"net/http"
)

type MenuController struct {
}

func (c *MenuController) List(ctx *gin.Context) {
	data, err := new(service.AdminMenuService).List()
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10001))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(data, "获取成功", 10000))
}

func (c *MenuController) Store(ctx *gin.Context) {
	form := new(schema.AdminMenuStoreForm)
	if err := ctx.Bind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	err := new(service.AdminMenuService).Store(form)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10001))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok("", "保存成功", 10000))
}

func (c *MenuController) Delete(ctx *gin.Context) {
	form := new(schema.AdminMenuDeleteForm)
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10001))
		return
	}
	err := new(service.AdminMenuService).Delete(form)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10001))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok("", "删除成功", 10000))
}

func (c *MenuController) Info(ctx *gin.Context) {
	form := new(schema.AdminMenuDetailForm)
	if err := ctx.Bind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	data, err := new(service.AdminMenuService).Info(form)
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10001))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(data, "成功", 10000))
}

// Options 表单菜单选项接口
func (c *MenuController) Options(ctx *gin.Context) {
	data, err := new(service.AdminMenuService).List()
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10001))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(data, "获取成功", 10000))
}

// PermissionOptions 附加权限列表的参数接口，用户角色编辑
func (c *MenuController) PermissionOptions(ctx *gin.Context) {
	data, err := new(service.AdminMenuService).PermissionOptions()
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10001))
		return
	}
	ctx.JSON(http.StatusOK, response.Ok(data, "获取成功", 10000))
}
