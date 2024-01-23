package forms

import "github.com/zhiniuer/goutils/datatypes"

type AdminUserMenuForm struct {
	UserId int `form:"user_id" binding:"required"`
}

type AdminUserRolesResult struct {
	UserId   string `json:"user_id"`
	RoleName string `json:"role_name"`
}

type AdminUserListResult struct {
	Id        int64                  `json:"id"`
	Username  string                 `json:"username"`
	Name      string                 `json:"name"`
	CreatedAt datatypes.DateTime     `json:"created_at"`
	UpdatedAt datatypes.DateTime     `json:"updated_at"`
	RoleList  []AdminUserRolesResult `json:"role_list" gorm:"foreignKey:UserId;references:Id"`
}

type AdminUserOptionResult struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AdminUserAuthStoreForm struct {
	Id             string   `form:"id" json:"id" binding:"required"`
	RoleList       []string `form:"role_list" json:"role_list"`
	PermissionList []string `form:"permission_list" json:"permission_list"`
}

type AdminUserAuthInfoForm struct {
	Id string `form:"id" binding:"required"`
}

type AdminUserAuthInfoResult struct {
	Id             string   `json:"id"`
	Name           string   `json:"name"`
	RoleList       []string `json:"role_list"`
	PermissionList []string `json:"permission_list"`
}
