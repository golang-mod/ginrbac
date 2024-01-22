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
	//err = utils.InitTrans("zh")
	driver.GDB = db
	rbacFile = "configs/rbac_model.conf"
	err = driver.InitRbac(rbacFile)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return
	}
	return &Admin{
		Middleware: &Middleware{},
	}, nil
}

type Admin struct {
	Middleware *Middleware
}

// NewRouteAuth 没有权限控制的路由
func (*Admin) NewRouteAuth(group *gin.RouterGroup) {
	routes.NewRouteAuth(group)
}

// NewRouteRbac RBAC权限控制的路由
func (*Admin) NewRouteRbac(group *gin.RouterGroup) {
	routes.NewRouteRbac(group)
}

type Middleware struct {
}

// Rbac 中间件
func (*Middleware) Rbac() gin.HandlerFunc {
	return middleware.NewRbac()
}

// Access 访问日志中间件
func (*Middleware) Access() gin.HandlerFunc {
	return middleware.NewAccess()
}
