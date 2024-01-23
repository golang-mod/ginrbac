package models

import "time"

type AdminOperationLog struct {
	Id        int64     `gorm:"column:id;primaryKey;type:bigint(20) unsigned;not null" json:"id"`
	UserId    int64     `gorm:"column:user_id;type:bigint(20) unsigned;not null" json:"user_id"`
	Path      string    `gorm:"column:path;type:varchar(255);not null" json:"path"`
	Method    string    `gorm:"column:method;type:varchar(10);not null" json:"method"`
	Ip        string    `gorm:"column:ip;type:varchar(15);not null" json:"ip"`
	Status    int       `gorm:"column:status;type:smallint(3);not null" json:"status"`
	Input     string    `gorm:"column:input;type:mediumtext;not null" json:"input"`
	Output    string    `gorm:"column:output;type:mediumtext;not null" json:"output"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null" json:"updated_at"`
}

func (AdminOperationLog) TableName() string {
	return "admin_operation_log"
}
