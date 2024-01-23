package service

import (
	"github.com/zhiniuer/goadmin/internal/app/driver"
	"github.com/zhiniuer/goadmin/internal/app/errors"
	"github.com/zhiniuer/goadmin/internal/app/models"
	"github.com/zhiniuer/goadmin/internal/app/schema"
	"github.com/zhiniuer/goadmin/internal/rbac"
	"github.com/zhiniuer/goutils/gormx"
	"github.com/zhiniuer/goutils/str"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type AdminPermissionsService struct {
	Model models.AdminPermissions
}

// List 获取权限列表
func (m *AdminPermissionsService) List(form *schema.AdminRoleListForm, page, pageSize int) (items []schema.AdminPermissionsListResult, count int64, err error) {
	menuTableName := models.AdminMenu{}.TableName()
	db := driver.GDB.Table(m.Model.TableName() + " as p").
		Joins("left join " + menuTableName + " as m on m.id = p.group").
		Select([]string{"m.title as `group`, p.id, p.name, p.slug, p.http_path, p.http_auth, p.created_at, p.http_method"})
	if form.Slug != "" {
		db.Where("p.slug LIKE ?", "%"+form.Slug+"%")
	}
	if form.Name != "" {
		db.Where("p.name LIKE ?", "%"+form.Name+"%")
	}
	db.Count(&count)
	if count == 0 {
		return items, 0, nil
	}
	db.Scopes(gormx.Paginate(page, pageSize)).Find(&items)
	return items, count, nil
}

// Options 获取权限选项接口
func (m *AdminPermissionsService) Options() (items []schema.AdminPermissionOptionsResult, err error) {
	menuTableName := models.AdminMenu{}.TableName()
	result := driver.GDB.Table(m.Model.TableName() + " as p").
		Joins("left join " + menuTableName + " as m on m.id = p.group").
		Select([]string{"m.title as `group`, p.id, p.name"}).
		Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, nil
}

func (m *AdminPermissionsService) Store(form *schema.AdminPermissionStoreForm) (err error) {
	db := driver.GDB
	oldItem := models.AdminPermissions{}
	storeData := &models.AdminPermissions{
		Name:       form.Name,
		Slug:       form.Slug,
		Group:      str.S(form.Group).DefaultInt(0),
		HttpMethod: strings.Join(form.HttpMethod, ","),
		HttpPath:   strings.ReplaceAll(form.HttpPath, "\n", ","),
		HttpAuth:   form.HttpAuth,
	}
	err = db.Transaction(func(tx *gorm.DB) (err error) {
		r, _ := rbac.New(tx)
		newId := form.Id
		if form.Id != "" {
			err = tx.Model(&models.AdminPermissions{}).Where("id", form.Id).First(&oldItem).Error
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return errors.New("记录未找到，参数错误")
				}
				return err
			}
			err = tx.Model(&models.AdminPermissions{}).Where("id", form.Id).Select([]string{
				"name",
				"slug",
				"group",
				"http_method",
				"http_path",
				"http_auth",
			}).Updates(storeData).Error
			if err != nil {
				return err
			}
			err = r.RemovePermission(strconv.Itoa(int(oldItem.Id)), false)
			if err != nil {
				return err
			}
		} else {
			err = tx.Create(storeData).Error
			if err != nil {
				return err
			}
			newId = strconv.Itoa(int(storeData.Id))
		}
		paths := strings.Split(storeData.HttpPath, ",")
		err = r.AddPermission(newId, paths, form.HttpMethod)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// Info 用户编辑的权限接口
func (m *AdminPermissionsService) Info(form *schema.AdminPermissionsInfoForm) (item schema.AdminPermissionsInfoResult, err error) {
	result := driver.GDB.Model(&models.AdminPermissions{}).Where("id", form.Id).First(&item)
	// 检查 ErrRecordNotFound 错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return item, errors.New("没找到相关数据")
	}
	item.HttpPath = strings.ReplaceAll(item.HttpPath, ",", "\n")
	return item, nil
}

// Detail 获取权限详情
func (m *AdminPermissionsService) Detail(form *schema.AdminPermissionsInfoForm) (res schema.AdminPermissionDetailResult, err error) {
	db := driver.GDB
	permission := schema.AdminPermissionDetail{}
	result := db.Model(&models.AdminPermissions{}).Where("id", form.Id).First(&permission)
	// 检查 ErrRecordNotFound 错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return res, errors.New("没找到相关数据")
	}
	res.Id = form.Id
	res.Name = permission.Name
	res.Slug = permission.Slug
	res.Group = permission.Group
	res.HttpMethod = permission.HttpMethod
	res.HttpPath = permission.HttpPath
	res.HttpAuth = permission.HttpAuth
	res.CreatedAt = permission.CreatedAt
	res.UpdatedAt = permission.UpdatedAt

	menu := models.AdminMenu{}
	err = db.Select("title").Where("id", permission.Group).Find(&menu).Error
	if err != nil {
		return res, err
	}
	res.Group = menu.Title
	var userList []string
	db.Model(&models.AdminUserPermissions{}).Where("permission_id", permission.Id).Pluck("user_id", &userList)
	db.Model(&models.AdminUsers{}).Where("id IN ?", userList).Pluck("name", &userList)
	res.UserList = userList

	return res, nil
}

// Delete 删除权限
func (m *AdminPermissionsService) Delete(form *schema.AdminPermissionsDeleteForm) (err error) {
	p := models.AdminPermissions{}
	err = driver.GDB.Transaction(func(tx *gorm.DB) (err error) {
		err = tx.Model(&models.AdminPermissions{}).Where("id", form.Id).First(&p).Error
		// 检查 ErrRecordNotFound 错误
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errors.New("没找到相关数据")
			}
			return err
		}
		// 删除权限及角色绑定权限
		r, _ := rbac.New(tx)
		err = r.RemovePermission(strconv.Itoa(int(p.Id)), true)
		if err != nil {
			return err
		}
		err = tx.Where("permission_id", p.Id).Delete(&models.AdminRolePermissions{}).Error
		if err != nil {
			return err
		}
		err = tx.Where("permission_id", p.Id).Delete(&models.AdminUserPermissions{}).Error
		if err != nil {
			return err
		}
		err = tx.Where("id", p.Id).Delete(&models.AdminPermissions{}).Error
		if err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	err = driver.Rbac.LoadPolicy()
	if err != nil {
		return err
	}
	return err
}
