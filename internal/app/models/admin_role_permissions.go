package models

import "time"

type AdminRolePermissions struct {
	Id           int       `gorm:"column:id;primaryKey;type:int(11) unsigned;not null" json:"id"`
	RoleId       int       `gorm:"column:role_id;type:int(11) unsigned;not null" json:"role_id"`
	PermissionId int       `gorm:"column:permission_id;type:int(11) unsigned;not null" json:"permission_id"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}

func (AdminRolePermissions) TableName() string {
	return "admin_role_permissions"
}
