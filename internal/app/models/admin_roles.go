package models

import "time"

type AdminRoles struct {
	Id        int       `gorm:"column:id;primaryKey;type:int(11) unsigned;not null" json:"id"`
	Name      string    `gorm:"column:name;type:varchar(50);not null" json:"name"`
	Slug      string    `gorm:"column:slug;type:varchar(50);not null" json:"slug"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}

func (AdminRoles) TableName() string {
	return "admin_roles"
}
