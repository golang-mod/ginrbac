package service

import (
	"errors"
	"github.com/zhiniuer/goadmin/internal/app/driver"
	"github.com/zhiniuer/goadmin/internal/app/models"
	"github.com/zhiniuer/goadmin/internal/app/schema"
	"github.com/zhiniuer/goadmin/str"
	"github.com/zhiniuer/goadmin/tree"
	"gorm.io/gorm"
)

type AdminMenuService struct {
}

// Store 保存菜单
func (m *AdminMenuService) Store(form *schema.AdminMenuStoreForm) (err error) {
	db := driver.GDB
	oldItem := &models.AdminMenu{}
	storeData := map[string]interface{}{
		"ParentId":  str.S(form.ParentId).DefaultInt(0),
		"Order":     form.Order,
		"Title":     form.Title,
		"Icon":      form.Icon,
		"Uri":       form.Uri,
		"Type":      form.Type,
		"ViewPath":  form.ViewPath,
		"Name":      form.Name,
		"KeepAlive": form.KeepAlive,
		"IsShow":    form.IsShow,
	}
	if form.Id != 0 {
		err = db.Model(&models.AdminMenu{}).Where("id", form.Id).First(&oldItem).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("没找到相关数据")
			}
			return err
		}
		err = db.Model(&models.AdminMenu{}).Where("id", form.Id).Updates(storeData).Error
		if err != nil {
			return err
		}
	} else {
		//result := db.Model(&oldItem).Create(form)
		err = db.Model(&models.AdminMenu{}).Create(&models.AdminMenu{
			ParentId:  str.S(form.ParentId).DefaultInt(0),
			Order:     form.Order,
			Title:     form.Title,
			Type:      str.S(form.Type).DefaultInt(0),
			ViewPath:  form.ViewPath,
			KeepAlive: str.S(form.KeepAlive).DefaultInt(0),
			Icon:      form.Icon,
			Name:      form.Name,
			Uri:       form.Uri,
			IsShow:    str.S(form.IsShow).DefaultInt(0),
			IsDefault: 0,
		}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// List 菜单列表
func (m *AdminMenuService) List() (interface{}, error) {
	var items schema.AdminMenuListResults
	db := driver.GDB.Model(&models.AdminMenu{})
	db.Order("`order` desc, `id` asc")
	db.Find(&items)
	resp := tree.GenerateTree(items.ConvertToNodeArray())
	return resp, nil
}

// PermissionOptions 菜单列表
func (m *AdminMenuService) PermissionOptions() (interface{}, error) {
	var items schema.AdminMenuOptionsResults
	db := driver.GDB.Table(models.AdminMenu{}.TableName())
	db.Order("`order` desc")

	db.Preload("PermissionList", func(db *gorm.DB) *gorm.DB {
		return db.Model(models.AdminPermissions{}).
			Select("name", "id", "group")
	}).Find(&items)

	resp := tree.GenerateTree(items.ConvertToNodeArray())
	return resp, nil
}

// Delete 删除菜单
func (m *AdminMenuService) Delete(form *schema.AdminMenuDeleteForm) (err error) {
	err = driver.GDB.Transaction(func(tx *gorm.DB) (err error) {
		var count int64
		tx.Model(&models.AdminMenu{}).Where("parent_id", form.Id).Count(&count)
		if count > 0 {
			return errors.New("请先删除子菜单")
		}
		// 删除菜单本身
		result := tx.Where("id", form.Id).Delete(&models.AdminMenu{})
		if result.RowsAffected == 0 {
			return result.Error
		}
		// 删除角色绑定的菜单
		result = tx.Where("menu_id = ?", form.Id).Delete(&models.AdminRoleMenu{})
		if result.Error != nil {
			return result.Error
		}
		var permissions []int
		err = tx.Model(&models.AdminPermissions{}).Where("`group` = ?", form.Id).Pluck("id", &permissions).Error
		if err != nil {
			return err
		}
		if len(permissions) > 0 {
			// 删除绑定菜单的权限
			err = tx.Where("`group` = ?", form.Id).Delete(&models.AdminPermissions{}).Error
			if err != nil {
				return
			}
			err = tx.Where("permissions_id IN", permissions).Delete(&models.AdminRolePermissions{}).Error
			if err != nil {
				return
			}
			err = tx.Where("permissions_id IN", permissions).Delete(&models.AdminUserPermissions{}).Error
			if err != nil {
				return
			}
		}
		// 返回 nil 提交事务
		return nil
	})
	return err
}

// Info 用户编辑的基础信息接口
func (m *AdminMenuService) Info(form *schema.AdminMenuDetailForm) (item schema.AdminMenuResult, err error) {
	result := driver.GDB.Model(&models.AdminMenu{}).Where("id", form.Id).First(&item)
	// 检查 ErrRecordNotFound 错误
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return item, errors.New("没找到相关数据")
		}
		return item, result.Error
	}
	return item, nil
}
