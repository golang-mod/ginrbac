package forms

import (
	"github.com/zhiniuer/goutils/datatypes"
)

type AdminPermissionsDeleteForm struct {
	Id int `form:"id" binding:"required"`
}

type AdminPermissionsListResult struct {
	Id         string             `json:"id"`
	Name       string             `json:"name"`
	Slug       string             `json:"slug"`
	Group      string             `json:"group"`
	HttpMethod datatypes.Args     `json:"http_method"`
	HttpPath   datatypes.Args     `json:"http_path"`
	HttpAuth   string             `json:"http_auth"`
	CreatedAt  datatypes.DateTime `json:"created_at"`
	UpdatedAt  datatypes.DateTime `json:"updated_at"`
}

type AdminPermissionsInfoResult struct {
	Id         string         `json:"id"`
	Name       string         `json:"name"`
	Slug       string         `json:"slug"`
	Group      string         `json:"group"`
	HttpMethod datatypes.Args `json:"http_method"`
	HttpPath   string         `json:"http_path"`
	HttpAuth   string         `json:"http_auth"`
}

type AdminPermissionsInfoForm struct {
	Id string `form:"id" json:"id" binding:"required"`
}

type AdminPermissionStoreForm struct {
	Id         string   `form:"id" json:"id"`
	Slug       string   `form:"slug" json:"slug" binding:"required"`
	Name       string   `form:"name"  json:"name" binding:"required"`
	Group      string   `form:"group" json:"group" binding:"required"`
	HttpMethod []string `form:"http_method" json:"http_method"`
	HttpPath   string   `form:"http_path" json:"http_path" binding:"required"`
	HttpAuth   string   `form:"http_auth" json:"http_auth" binding:"required"`
}

type AdminPermissionOptionsResult struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Group string `json:"group"`
}

type AdminPermissionDetail struct {
	Id         string             `json:"id"`
	Slug       string             `json:"slug"`
	Name       string             `json:"name"`
	Group      string             `json:"group"`
	HttpPath   datatypes.Args     `json:"http_path"`
	HttpMethod datatypes.Args     `json:"http_method"`
	HttpAuth   string             `json:"http_auth"`
	CreatedAt  datatypes.DateTime `json:"created_at"`
	UpdatedAt  datatypes.DateTime `json:"updated_at"`
}

type AdminPermissionDetailResult struct {
	AdminPermissionDetail
	UserList []string `json:"user_list"`
}
