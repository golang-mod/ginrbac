package app

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiniuer/goutils/response"
	"net/http"
	"sync"
)

const (
	AdminId = "admin.id"
)

type Context struct {
	//context.Context
	Request *http.Request
	// This mutex protects Keys map.
	mu sync.RWMutex
	// 用户相关数据
	User struct {
		AdminId int64
	}
	// 系统相关数据
	Client struct {
		Ip string
	}
	// 请求分页相关数据
	Query struct {
		Page     int
		PageSize int
		Export   string
	}
}

// NewContext 返回APP上下文
// 取值方法
// page := c.Value(Page).(int)
// pageSize := c.Value(PageSize).(int)
func NewContext(ctx *gin.Context) *Context {
	var ac Context
	page := response.Page(ctx.Request)
	pageSize := response.PageSize(ctx.Request)
	adminUserId := ctx.GetInt64(AdminId)

	ac.Request = ctx.Request
	// 用户相关数据
	ac.User.AdminId = adminUserId
	// 请求分页相关数据
	ac.Query.Page = page
	ac.Query.PageSize = pageSize
	// 导出相关
	ac.Query.Export = ctx.Request.URL.Query().Get("_export_")

	return &ac
}
