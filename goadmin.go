package goadmin

import (
	"github.com/gin-gonic/gin"
	"github.com/zhiniuer/goadmin/internal/app/driver"
	"github.com/zhiniuer/goadmin/internal/app/middleware"
	"github.com/zhiniuer/goadmin/internal/app/routes"
	"gorm.io/gorm"
)

// New configs/rbac_model.conf
func New(db *gorm.DB, rbacFile string) (a *Admin, err error) {
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
	return &Admin{}, nil
}

type Admin struct {
}

// NewRouteAuth 没有权限控制的路由
func (*Admin) NewRouteAuth(group *gin.RouterGroup) {
	routes.NewRouteRbac(group)
}

// NewRouteRbac RBAC权限控制的路由
func (*Admin) NewRouteRbac(group *gin.RouterGroup) {
	routes.NewRouteRbac(group)
}

func RbacMiddleware() gin.HandlerFunc {
	return middleware.NewRbac()
}
