package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiniuer/goadmin/internal/app/schema"
	"github.com/zhiniuer/goadmin/internal/app/service"
	"github.com/zhiniuer/goutils/response"
	"github.com/zhiniuer/goutils/validator_util"
	"net/http"
)

type OperationLogController struct {
}

func (c *OperationLogController) List(ctx *gin.Context) {
	form := new(schema.AdminOperationLogListFrom)
	if err := ctx.Bind(&form); err != nil {
		ctx.JSON(http.StatusOK, response.Fail(validator_util.ErrFirst(err).Error(), 10001))
		return
	}
	data, count, err := new(service.AdminOperationLogService).List(form, response.Page(ctx.Request), response.PageSize(ctx.Request))
	if err != nil {
		ctx.JSON(http.StatusOK, response.Fail(err.Error(), 10001))
		return
	}
	pageData := response.Paginate(ctx.Request, count, data)
	ctx.JSON(http.StatusOK, response.Ok(pageData, "获取成功", 10000))
}
