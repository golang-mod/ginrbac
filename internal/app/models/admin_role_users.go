package models

import "time"

type AdminRoleUsers struct {
	Id        int       `gorm:"column:id;primaryKey;type:int(11) unsigned;not null" json:"id"`
	RoleId    int       `gorm:"column:role_id;type:int(11) unsigned;not null" json:"role_id"`
	UserId    int       `gorm:"column:user_id;type:int(11) unsigned;not null" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}

func (AdminRoleUsers) TableName() string {
	return "admin_role_users"
}
