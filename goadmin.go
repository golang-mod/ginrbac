package goadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiniuer/goadmin/internal/app/driver"
	"github.com/zhiniuer/goadmin/internal/app/middleware"
	"github.com/zhiniuer/goadmin/internal/app/routes"
	"gorm.io/gorm"
)

// NewAdmin
// rbacFile rbac_model.conf
func NewAdmin(db *gorm.DB, rbacFile string) (a *Admin, err error) {
	driver.GDB = db
	err = driver.InitRbac(rbacFile)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return
	}
	return &Admin{
		Middleware: &middle{},
		Router:     &router{},
	}, nil
}

type Admin struct {
	Middleware *middle
	Router     *router
}
type router struct {
}

// AuthRouter 没有权限控制的路由
func (*router) AuthRouter(group *gin.RouterGroup) {
	routes.NewRouteAuth(group)
}

// RbacRouter RBAC权限控制的路由
func (*router) RbacRouter(group *gin.RouterGroup) {
	routes.NewRouteRbac(group)
}

type middle struct {
}

// Rbac 中间件
func (*middle) Rbac() gin.HandlerFunc {
	return middleware.NewRbac()
}

// Access 访问日志中间件
func (*middle) Access() gin.HandlerFunc {
	return middleware.NewAccess()
}
