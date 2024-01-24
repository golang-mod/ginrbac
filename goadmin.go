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
	return &Admin{
		Middleware: &_middleware{},
		Router:     &_router{},
	}, nil
}

type Admin struct {
	Middleware *_middleware
	Router     *_router
}
type _router struct {
}

// AuthRouter 没有权限控制的路由
func (*_router) AuthRouter(group *gin.RouterGroup) {
	routes.NewRouteAuth(group)
}

// RbacRouter RBAC权限控制的路由
func (*_router) RbacRouter(group *gin.RouterGroup) {
	routes.NewRouteRbac(group)
}

type _middleware struct {
}

// Rbac 中间件
func (*_middleware) Rbac() gin.HandlerFunc {
	return middleware.NewRbac()
}
