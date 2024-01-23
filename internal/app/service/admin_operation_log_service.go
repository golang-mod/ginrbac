package service

import (
	"github.com/zhiniuer/goadmin/internal/app"
	"github.com/zhiniuer/goadmin/internal/app/driver"
	"github.com/zhiniuer/goadmin/internal/app/models"
	"github.com/zhiniuer/goadmin/internal/app/schema"
	"github.com/zhiniuer/goutils/gormx"
)

type AdminOperationLogService struct {
}

// Save 保存访问日志
func (m *AdminOperationLogService) Save(ctx *app.Context, form *schema.AdminOperationLogStoreForm) (err error) {
	db := driver.GDB
	storeData := &models.AdminOperationLog{
		Method: form.Method,
		Path:   form.Path,
		Ip:     form.Ip,
		UserId: form.UserId,
		Input:  form.Input,
		Status: form.Status,
		Output: form.Output,
	}
	err = db.Create(storeData).Error
	if err != nil {
		return
	}
	return nil
}

// List 访问日志
// ctx
// items 日志数组
// count 日志总条数
// err   error
func (m *AdminOperationLogService) List(form *schema.AdminOperationLogListFrom, page, pageSize int) (items []schema.AdminOperationLogResult, count int64, err error) {
	userTableName := models.AdminUsers{}.TableName()
	logTableName := models.AdminOperationLog{}.TableName()
	db := driver.GDB.Table(logTableName + " as l").
		Joins("left join " + userTableName + " as u on u.id = l.user_id").
		Select([]string{"u.name as user_name", "l.user_id", "l.id", "l.path", "l.method", "l.ip", "l.status", "l.created_at"})
	if form.Id != "" {
		db.Where("l.id", form.Id)
	}
	if form.UserId != "" {
		db.Where("l.user_id", form.UserId)
	}
	if form.Method != "" {
		db.Where("l.method", form.Method)
	}
	if form.Status != "" {
		db.Where("l.status", form.Status)
	}
	if form.Method != "" {
		db.Where("l.method", form.Method)
	}
	if form.Ip != "" {
		db.Where("l.ip LIKE ?", "%"+form.Ip+"%")
	}
	if form.Path != "" {
		db.Where("l.path LIKE ?", "%"+form.Path+"%")
	}
	db.Count(&count)
	if count == 0 {
		return items, 0, nil
	}
	db.Scopes(gormx.Paginate(page, pageSize)).Order("id desc").Find(&items)
	return items, count, nil
}
