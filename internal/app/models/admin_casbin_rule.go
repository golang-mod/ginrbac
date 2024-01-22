package models

type AdminCasbinRule struct {
	Id    int    `gorm:"column:id;primaryKey;type:int(11) unsigned;not null" json:"id"`
	Ptype string `gorm:"column:ptype;type:varchar(100)" json:"ptype"`
	V0    string `gorm:"column:v0;type:varchar(100)" json:"v0"`
	V1    string `gorm:"column:v1;type:varchar(100)" json:"v1"`
	V2    string `gorm:"column:v2;type:varchar(100)" json:"v2"`
	V3    string `gorm:"column:v3;type:varchar(100)" json:"v3"`
	V4    string `gorm:"column:v4;type:varchar(100)" json:"v4"`
	V5    string `gorm:"column:v5;type:varchar(100)" json:"v5"`
	V6    string `gorm:"column:v6;type:varchar(25)" json:"v6"`
	V7    string `gorm:"column:v7;type:varchar(25)" json:"v7"`
}

func (AdminCasbinRule) TableName() string {
	return "casbin_rule"
}
