package models

import "time"

type AdminUsers struct {
	Id            int       `gorm:"column:id;primaryKey;type:int(11) unsigned;not null" json:"id"`
	Username      string    `gorm:"column:username;type:varchar(190);not null" json:"username"`
	Password      string    `gorm:"column:password;type:varchar(60);not null" json:"password"`
	Name          string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Avatar        string    `gorm:"column:avatar;type:varchar(255)" json:"avatar"`
	RememberToken string    `gorm:"column:remember_token;type:varchar(100)" json:"remember_token"`
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}

func (AdminUsers) TableName() string {
	return "admin_users"
}
