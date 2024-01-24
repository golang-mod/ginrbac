package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-mod/ginrbac/internal/app/controllers"
)

var (
	adminRole       = new(controllers.RoleController)
	adminUser       = new(controllers.UserController)
	adminMenu       = new(controllers.MenuController)
	adminPermission = new(controllers.PermissionController)
)

// 登录 退出 token 生成验证不应在本包范围内
//func NewRouteDefault(rg *gin.RouterGroup) {
//	rg.POST("login", adminUser.UserLogin)
//	rg.GET("refresh", adminUser.Refresh)
//}

// NewRouteAuth 没有权限控制的路由
func NewRouteAuth(rg *gin.RouterGroup) {
	//rg.POST("auth/logout", adminUser.Logout)
	rg.GET("auth/user/menu", adminUser.Menu)
	rg.GET("status", adminUser.Status)
	rg.GET("auth/user/info", adminUser.UserInfo)
	//rg.POST("auth/user/password", adminUser.UserPassword)
}

// NewRouteRbac RBAC权限控制的路由
func NewRouteRbac(rbac *gin.RouterGroup) {
	rbac.GET("auth/user/list", adminUser.List)
	rbac.GET("auth/user/options", adminUser.Options)
	rbac.POST("auth/user/auth-store", adminUser.AuthStore)
	rbac.GET("auth/user/auth-info", adminUser.AuthInfo)

	rbac.GET("auth/role/list", adminRole.List)
	rbac.GET("auth/role/options", adminRole.Options)
	rbac.GET("auth/role/info", adminRole.Info)
	rbac.GET("auth/role/detail", adminRole.Detail)
	rbac.POST("auth/role/store", adminRole.Store)
	rbac.POST("auth/role/delete", adminRole.Delete)

	rbac.GET("auth/menu/list", adminMenu.List)
	rbac.GET("auth/menu/info", adminMenu.Info)
	rbac.POST("auth/menu/store", adminMenu.Store)
	rbac.POST("auth/menu/delete", adminMenu.Delete)
	rbac.GET("auth/menu/options", adminMenu.List)
	rbac.GET("auth/menu/permission-options", adminMenu.PermissionOptions)

	rbac.GET("auth/permissions/list", adminPermission.List)
	rbac.GET("auth/permissions/options", adminPermission.Options)
	rbac.GET("auth/permissions/info", adminPermission.Info)
	rbac.GET("auth/permissions/detail", adminPermission.Detail)
	rbac.POST("auth/permissions/store", adminPermission.Store)
	rbac.POST("auth/permissions/delete", adminPermission.Delete)
}
