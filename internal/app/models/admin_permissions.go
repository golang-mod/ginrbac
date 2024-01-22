package models

import "time"

type AdminPermissions struct {
	Id         int       `gorm:"column:id;primaryKey;type:int(11) unsigned;not null" json:"id"`
	Name       string    `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Slug       string    `gorm:"column:slug;type:varchar(50);not null" json:"slug"`
	Group      int       `gorm:"column:group;type:int(11) unsigned;default:0;not null" json:"group"`
	HttpMethod string    `gorm:"column:http_method;type:varchar(255)" json:"http_method"`
	HttpPath   string    `gorm:"column:http_path;type:text" json:"http_path"`
	HttpAuth   string    `gorm:"column:http_auth;type:varchar(10)" json:"http_auth"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}

func (AdminPermissions) TableName() string {
	return "permissions"
}
