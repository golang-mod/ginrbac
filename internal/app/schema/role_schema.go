package schema

import (
	"github.com/zhiniuer/goutils/gormx/datatypes"
)

// AdminRoleListForm 角色列表请求参数
type AdminRoleListForm struct {
	Name string `form:"name"`
	Slug string `form:"slug"`
}

type AdminRoleListResult struct {
	Id        string                `json:"id"`
	Name      string                `json:"name"`
	Slug      string                `json:"slug"`
	UserList  []adminRoleUserResult `json:"user_list" gorm:"foreignKey:RoleId"`
	CreatedAt datatypes.DateTime    `json:"created_at"`
	UpdatedAt datatypes.DateTime    `json:"updated_at"`
}

type adminRoleUserResult struct {
	RoleId   string `json:"role_id"`
	UserName string `json:"user_name"`
}

type AdminRoleOptionsResult struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AdminRoleDeleteForm struct {
	Id string `form:"id" binding:"required"`
}

type AdminRoleInfoForm struct {
	Id string `form:"id" binding:"required"`
}

type AdminRoleStoreForm struct {
	Id             string   `form:"id" json:"id"`
	Slug           string   `form:"slug" json:"slug" binding:"required"`
	Name           string   `form:"name" json:"name" binding:"required"`
	MenuList       []string `form:"menu_list" json:"menu_list" binding:"gt=0"`
	PermissionList []string `form:"permission_list" json:"permission_list" binding:"gt=0"`
}

type AdminRoleInfoResult struct {
	Id             string   `json:"id"`
	Slug           string   `json:"slug"`
	Name           string   `json:"name"`
	MenuList       []string `json:"menu_list"`
	PermissionList []string `json:"permission_list"`
}

type AdminRoleDetail struct {
	Id        string             `json:"id"`
	Slug      string             `json:"slug"`
	Name      string             `json:"name"`
	CreatedAt datatypes.DateTime `json:"created_at"`
	UpdatedAt datatypes.DateTime `json:"updated_at"`
}

type AdminRoleDetailResult struct {
	AdminRoleDetail
	MenuList       []string `json:"menu_list"`
	PermissionList []string `json:"permission_list"`
	UserList       []string `json:"user_list"`
}
