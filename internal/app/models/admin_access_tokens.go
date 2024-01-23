package models

import "time"

type AdminAccessTokens struct {
	Id            int        `gorm:"column:id;primaryKey;type:int(11) unsigned;not null" json:"id"`
	TokenableType string     `gorm:"column:tokenable_type;type:varchar(100);not null" json:"tokenable_type"`
	TokenableId   int        `gorm:"column:tokenable_id;type:int(11) unsigned;not null" json:"tokenable_id"`
	Name          string     `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Token         string     `gorm:"column:token;type:varchar(100);not null" json:"token"`
	Abilities     string     `gorm:"column:abilities;type:text" json:"abilities"`
	LastUsedAt    *time.Time `gorm:"column:last_used_at;type:timestamp" json:"last_used_at"`
	CreatedAt     *time.Time `gorm:"column:created_at;type:timestamp" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`
}

func (AdminAccessTokens) TableName() string {
	return "admin_access_tokens"
}
