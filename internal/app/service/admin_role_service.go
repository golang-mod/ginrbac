package service

import (
	"fmt"
	"github.com/zhiniuer/goadmin/internal/app/driver"
	"github.com/zhiniuer/goadmin/internal/app/errors"
	"github.com/zhiniuer/goadmin/internal/app/forms"
	"github.com/zhiniuer/goadmin/internal/app/models"
	"github.com/zhiniuer/goadmin/internal/rbac"
	"github.com/zhiniuer/goutils/gormx"
	"gorm.io/gorm"
	"strconv"
)

type AdminRoleService struct {
}

// List 获取角色列表
func (m *AdminRoleService) List(form *forms.AdminRoleListForm, page, pageSize int) (roles []forms.AdminRoleListResult, count int64, err error) {
	db := driver.GDB.Table(models.AdminRoles{}.TableName()).Select([]string{"slug", "name", "id", "created_at", "updated_at"})
	if form.Slug != "" {
		db.Where("slug LIKE ?", "%"+form.Slug+"%")
	}
	if form.Name != "" {
		db.Where("name LIKE ?", "%"+form.Name+"%")
	}
	db.Count(&count)
	if count == 0 {
		return roles, 0, nil
	}
	// 方案一
	roleUsersTableName := models.AdminRoleUsers{}.TableName()
	usersTableName := models.AdminUsers{}.TableName()
	result := db.Scopes(gormx.Paginate(page, pageSize)).Preload("UserList", func(db *gorm.DB) *gorm.DB {
		return db.Table(roleUsersTableName+" as ru").
			Select("u.name as user_name", "ru.user_id", "ru.role_id").
			Joins("left join " + usersTableName + " as u on u.id = ru.user_id")
	}).Find(&roles)
	// 方案二
	//result := db.Scopes(utils.Gorm.Paginate(ctx.Request)).Preload("UserList", func(db *gorm.DB) *gorm.DB {
	//	return db.Joins("User")
	//}).Find(&roles)
	// 方案三
	//db.Scopes(utils.Gorm.Paginate(ctx.Request)).Preload("UserList").Preload("UserList.UserName").Find(&roles)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return roles, count, nil
}

// Options 获取角色选项接口
func (m *AdminRoleService) Options() (roles []forms.AdminRoleOptionsResult, err error) {
	result := driver.GDB.Model(&models.AdminRoles{}).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}

// Store 保存角色
func (m *AdminRoleService) Store(form *forms.AdminRoleStoreForm) (err error) {
	db := driver.GDB
	oldItem := &models.AdminRoles{}
	fmt.Println(form)
	storeData := &models.AdminRoles{
		Name: form.Name,
		Slug: form.Slug,
	}
	err = db.Transaction(func(tx *gorm.DB) (err error) {
		r, err := rbac.New(tx)
		if err != nil {
			return err
		}
		newId, _ := strconv.Atoi(form.Id)
		if form.Id != "" {
			err = tx.Model(&models.AdminRoles{}).Select("id").Where("id", form.Id).First(&oldItem).Error
			if err != nil {
				return errors.New("记录未找到，参数错误")
			}
			err = tx.Model(&models.AdminRoles{}).Where("id", form.Id).Updates(storeData).Error
			if err != nil {
				return err
			}
			// 删除角色菜单
			err = tx.Where("role_id", form.Id).Delete(&models.AdminRoleMenu{}).Error
			if err != nil {
				return err
			}
			// 删除角色权限
			err = tx.Where("role_id", form.Id).Delete(&models.AdminRolePermissions{}).Error
			if err != nil {
				return err
			}
			// 删除Casbin角色相关
			err = r.RemoveRole(form.Id, false)
			if err != nil {
				return err
			}
		} else {
			err = tx.Create(storeData).Error
			if err != nil {
				return err
			}
			newId = int(storeData.Id)
		}
		// 添加角色菜单
		var roleMenus []models.AdminRoleMenu
		// 防止传递无效参数
		var menuList []int
		err = tx.Model(&models.AdminMenu{}).Where("id IN ?", form.MenuList).Pluck("id", &menuList).Error
		if err != nil {
			return err
		}
		if len(menuList) != len(form.MenuList) {
			return errors.New("菜单不存在或已删除，请刷新重试")
		}
		for _, v := range menuList {
			roleMenus = append(roleMenus, models.AdminRoleMenu{RoleId: newId, MenuId: v})
		}
		err = tx.Create(roleMenus).Error
		if err != nil {
			return err
		}
		// 添加角色权限
		var rolePermissions []models.AdminRolePermissions
		// 防止传递无效参数
		var permissionList []int
		err = tx.Model(&models.AdminPermissions{}).Where("id IN ?", form.PermissionList).Pluck("id", &permissionList).Error
		if err != nil {
			return err
		}
		if len(permissionList) != len(form.PermissionList) {
			return errors.New("权限不存在或已删除，请刷新重试")
		}
		for _, v := range permissionList {
			rolePermissions = append(rolePermissions, models.AdminRolePermissions{RoleId: newId, PermissionId: v})
		}
		err = tx.Create(rolePermissions).Error
		if err != nil {
			return err
		}
		// 添加新角色Casbin相关
		err = r.AddRole(strconv.Itoa(newId), form.PermissionList)
		if err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除角色
func (m *AdminRoleService) Delete(form *forms.AdminRoleDeleteForm) (err error) {
	role := models.AdminRoles{}
	err = driver.GDB.Transaction(func(tx *gorm.DB) (err error) {
		result := tx.Model(&models.AdminRoles{}).Select("id").Where("id", form.Id).First(&role)
		// 检查 ErrRecordNotFound 错误
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("没找到相关数据")
		}
		// 删除权限及角色绑定权限
		r, _ := rbac.New(tx)
		err = r.RemoveRole(strconv.Itoa(int(role.Id)), true)
		if err != nil {
			return err
		}
		res := tx.Where("role_id", role.Id).Delete(&models.AdminRolePermissions{})
		if res.Error != nil {
			return res.Error
		}
		res = tx.Where("role_id", role.Id).Delete(&models.AdminRoleUsers{})
		if res.Error != nil {
			return res.Error
		}
		res = tx.Where("role_id", role.Id).Delete(&models.AdminRoleMenu{})
		if res.Error != nil {
			return res.Error
		}
		res = tx.Where("id", role.Id).Delete(&AdminRoleService{})
		if res.Error != nil {
			return res.Error
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

// Info 编辑角色的基础信息接口
func (m *AdminRoleService) Info(form *forms.AdminRoleInfoForm) (res forms.AdminRoleInfoResult, err error) {
	db := driver.GDB
	var role models.AdminRoles
	result := db.Model(&models.AdminRoles{}).Where("id", form.Id).First(&role)
	// 检查 ErrRecordNotFound 错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return res, errors.New("没找到相关数据")
	}
	res.Id = form.Id
	res.Name = role.Name
	res.Slug = role.Slug

	var menuList []string
	db.Model(&models.AdminRoleMenu{}).Where("role_id", role.Id).Pluck("menu_id", &menuList)
	res.MenuList = menuList

	var permissionList []string
	db.Model(&models.AdminRolePermissions{}).Where("role_id", role.Id).Pluck("permission_id", &permissionList)
	res.PermissionList = permissionList

	return res, nil
}

// Detail 获取角色详情
func (m *AdminRoleService) Detail(form *forms.AdminRoleInfoForm) (res forms.AdminRoleDetailResult, err error) {
	db := driver.GDB
	role := forms.AdminRoleDetail{}
	result := db.Model(&models.AdminRoles{}).Where("id", form.Id).First(&role)
	// 检查 ErrRecordNotFound 错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return res, errors.New("没找到相关数据")
	}
	res.Id = form.Id
	res.Name = role.Name
	res.Slug = role.Slug
	res.CreatedAt = role.CreatedAt
	res.UpdatedAt = role.UpdatedAt

	var menuList []string
	db.Model(&models.AdminRoleMenu{}).Where("role_id", role.Id).Pluck("menu_id", &menuList)
	res.MenuList = menuList

	var permissionList []string
	db.Model(&models.AdminRolePermissions{}).Where("role_id", role.Id).Pluck("permission_id", &permissionList)
	res.PermissionList = permissionList

	var userList []string
	db.Model(&models.AdminRoleUsers{}).Where("role_id", role.Id).Pluck("user_id", &userList)
	db.Model(&models.AdminUsers{}).Where("id IN ?", userList).Pluck("name", &userList)
	res.UserList = userList

	return res, nil
}
