package models

import "time"

type AdminUserPermissions struct {
	Id           int       `gorm:"column:id;primaryKey;type:int(11) unsigned;not null" json:"id"`
	UserId       int       `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	PermissionId int       `gorm:"column:permission_id;type:int(11);not null" json:"permission_id"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}

func (AdminUserPermissions) TableName() string {
	return "admin_user_permissions"
}
