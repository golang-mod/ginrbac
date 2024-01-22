package schema

import (
	"github.com/zhiniuer/goadmin/gormx/datatypes"
)

type AdminOperationLogListFrom struct {
	Id     string `form:"id"`
	Method string `form:"method"`
	UserId string `form:"user_id"`
	Path   string `form:"path"`
	Status string `form:"status"`
	Ip     string `form:"ip"`
}

type AdminOperationLogResult struct {
	Id        uint               `json:"id"`
	UserId    string             `json:"user_id"`
	UserName  string             `json:"user_name"`
	Path      string             `json:"path"`
	Method    string             `json:"method"`
	Ip        string             `json:"ip"`
	Status    string             `json:"status"`
	CreatedAt datatypes.DateTime `json:"created_at"`
}

type AdminOperationLogStoreForm struct {
	Method string
	Path   string
	Ip     string
	UserId int
	Input  string
}
