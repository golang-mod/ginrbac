package schema

import (
	"github.com/zhiniuer/goadmin/tree"
)

type AdminMenuListResult struct {
	Id        string `json:"id"`
	ParentId  string `json:"parent_id"`
	Order     string `json:"order"`
	Type      string `json:"type"`
	ViewPath  string `json:"view_path"`
	IsDefault string `json:"is_default"`
	KeepAlive string `json:"keep_alive"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Uri       string `json:"uri"`
	IsShow    string `json:"is_show"`
}

func (m AdminMenuListResult) GetId() string {
	return m.Id
}

func (m AdminMenuListResult) GetParentId() string {
	return m.ParentId
}

func (m AdminMenuListResult) IsRoot() bool {
	return m.ParentId == "0" || m.ParentId == "" || m.ParentId == m.Id
}

type AdminMenuListResults []AdminMenuListResult

// ConvertToNodeArray 将当前数组转换成父类 Node 接口 数组
func (s AdminMenuListResults) ConvertToNodeArray() (nodes []tree.Node) {
	for _, v := range s {
		nodes = append(nodes, v)
	}
	return
}

type AdminMenuOptionsResult struct {
	Id             string                      `json:"id"`
	ParentId       string                      `json:"parent_id"`
	Title          string                      `json:"title"`
	PermissionList []AdminMenuPermissionResult `json:"permission_list" gorm:"foreignKey:Group;references:Id"`
}

type AdminMenuPermissionResult struct {
	Id    string `json:"id"`
	Group string `json:"group"`
	Name  string `json:"name"`
}

func (m AdminMenuOptionsResult) GetId() string {
	return m.Id
}

func (m AdminMenuOptionsResult) GetParentId() string {
	return m.ParentId
}

func (m AdminMenuOptionsResult) IsRoot() bool {
	return m.ParentId == "0" || m.ParentId == "" || m.ParentId == m.Id
}

type AdminMenuOptionsResults []AdminMenuOptionsResult

// ConvertToNodeArray 将当前数组转换成父类 Node 接口 数组
func (s AdminMenuOptionsResults) ConvertToNodeArray() (nodes []tree.Node) {
	for _, v := range s {
		nodes = append(nodes, v)
	}
	return
}

type AdminMenuDeleteForm struct {
	Id int `form:"id" json:"id" binding:"required"`
}

type AdminMenuDetailForm AdminMenuDeleteForm
type AdminMenuResult struct {
	Id        string `json:"id"`
	ParentId  string `json:"parent_id"`
	Order     int    `json:"order"`
	Type      string `json:"type"`
	ViewPath  string `json:"view_path"`
	IsDefault string `json:"is_default"`
	KeepAlive string `json:"keep_alive"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Uri       string `json:"uri"`
	IsShow    string `json:"is_show"`
}

type AdminMenuStoreForm struct {
	Id        int    `form:"id" json:"id"`
	Type      string `form:"type" json:"type"`
	ViewPath  string `json:"view_path" json:"view_path"`
	IsDefault string `json:"is_default" json:"is_default"`
	KeepAlive string `json:"keep_alive" json:"keep_alive"`
	Name      string `json:"name" json:"name"`
	ParentId  string `form:"parent_id" json:"parent_id"`
	Order     int    `form:"order" json:"order"`
	Title     string `form:"title" json:"title" binding:"required"`
	Icon      string `form:"icon" json:"icon"`
	Uri       string `form:"uri" json:"uri"`
	IsShow    string `form:"is_show" json:"is_show" binding:"required"`
}
