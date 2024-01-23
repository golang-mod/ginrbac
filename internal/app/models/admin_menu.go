package models

import "time"

type AdminMenu struct {
	Id        int       `gorm:"column:id;primaryKey;type:int(11) unsigned;not null" json:"id"`
	ParentId  int       `gorm:"column:parent_id;type:int(11) unsigned;default:0;not null" json:"parent_id"` // 父级ID
	Order     int       `gorm:"column:order;type:int(11);default:0;not null" json:"order"`                  // 排序
	Title     string    `gorm:"column:title;type:varchar(50);not null" json:"title"`                        // 菜单标题
	Type      int       `gorm:"column:type;type:tinyint(2);default:0;not null" json:"type"`                 // 类型 1分组 2 菜单 3 按钮
	ViewPath  string    `gorm:"column:view_path;type:varchar(100);default:" json:"view_path"`               // 组件地址
	KeepAlive int       `gorm:"column:keep_alive;type:tinyint(2);default:0;not null" json:"keep_alive"`     // 前端页面缓存
	Icon      string    `gorm:"column:icon;type:varchar(50);default:" json:"icon"`                          // 图标
	Name      string    `gorm:"column:name;type:varchar(50);default:;not null" json:"name"`                 // 前端路由名称
	Uri       string    `gorm:"column:uri;type:varchar(50);default:" json:"uri"`                            // 页面地址
	IsShow    int       `gorm:"column:is_show;type:tinyint(2);default:0;not null" json:"is_show"`           // 是否在菜单显示
	IsDefault int       `gorm:"column:is_default;type:tinyint(2);default:0;not null" json:"is_default"`     // 默认页面
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}

func (AdminMenu) TableName() string {
	return "admin_menu"
}
