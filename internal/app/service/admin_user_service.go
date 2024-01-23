package service

import (
	"github.com/zhiniuer/goadmin/internal/app"
	"github.com/zhiniuer/goadmin/internal/app/dao"
	"github.com/zhiniuer/goadmin/internal/app/driver"
	"github.com/zhiniuer/goadmin/internal/app/errors"
	"github.com/zhiniuer/goadmin/internal/app/forms"
	"github.com/zhiniuer/goadmin/internal/app/models"
	"github.com/zhiniuer/goutils/gormx"
	"github.com/zhiniuer/goutils/tree"
	"gorm.io/gorm"
	"strconv"
)

type AdminUserService struct {
}

// List 获取用户列表
func (m *AdminUserService) List(page, pageSize int) (items []forms.AdminUserListResult, count int64, err error) {
	db := driver.GDB.Model(models.AdminUsers{}).Find(&items)
	roleUsersTableName := models.AdminRoleUsers{}.TableName()
	roleTableName := models.AdminRoles{}.TableName()
	db.Scopes(gormx.Paginate(page, pageSize)).Preload("RoleList", func(db *gorm.DB) *gorm.DB {
		return db.Table(roleUsersTableName+" as ru").
			Select("r.name as role_name", "ru.user_id", "ru.role_id").
			Joins("left join " + roleTableName + " as r on r.id = ru.role_id")
	}).Find(&items)

	db.Scopes(gormx.Paginate(page, pageSize)).
		Order("id DESC").Find(&items)
	return items, count, nil
}

// Menu 用户菜单列表
func (AdminUserService) Menu(userId int) (interface{}, error) {
	var roleIds []int
	driver.GDB.Model(&models.AdminRoleUsers{}).Where("user_id", userId).Pluck("role_id", &roleIds)
	if len(roleIds) == 0 {
		return roleIds, nil
	}
	var menuIds []int
	driver.GDB.Model(&models.AdminRoleMenu{}).Where("role_id IN ?", roleIds).Pluck("menu_id", &menuIds)
	if len(menuIds) == 0 {
		return menuIds, nil
	}
	var items forms.AdminMenuListResults
	_ = driver.GDB.Model(&models.AdminMenu{}).Where("id IN ?", menuIds).Order("`order` desc").Find(&items)
	if len(items) == 0 {
		return items, nil
	}
	resp := tree.GenerateTree(items.ConvertToNodeArray())
	return resp, nil
}

// AuthStore 用户授权
func (m *AdminUserService) AuthStore(form *forms.AdminUserAuthStoreForm) (err error) {
	err = driver.GDB.Transaction(func(tx *gorm.DB) (err error) {
		userId, _ := strconv.Atoi(form.Id)
		udao := &dao.UserDao{}
		return udao.AuthStore(tx, userId, form.RoleList, form.PermissionList)
	})
	if err != nil {
		return err
	}
	err = driver.Rbac.LoadPolicy()
	if err != nil {
		return err
	}
	return nil
}

// AuthInfo 授权编辑数据回显
func (m *AdminUserService) AuthInfo(form *forms.AdminUserAuthInfoForm) (res forms.AdminUserAuthInfoResult, err error) {
	db := driver.GDB
	var info models.AdminUsers
	err = db.Model(&models.AdminUsers{}).Where("id", form.Id).Select([]string{"id", "name", "username"}).First(&info).Error
	// 检查 ErrRecordNotFound 错误
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.New("没找到相关数据")
		}
		return
	}

	res.Id = form.Id
	res.Name = info.Name

	var roleList []string
	db.Model(&models.AdminRoleUsers{}).Where("user_id", info.Id).Pluck("role_id", &roleList)
	res.RoleList = roleList

	var permissionList []string
	db.Model(&models.AdminUserPermissions{}).Where("user_id", info.Id).Pluck("permission_id", &permissionList)
	res.PermissionList = permissionList
	return res, nil
}

// Check 账号检测
func (m *AdminUserService) Check(username string) (user models.AdminUsers, err error) {
	res := driver.GDB.Model(models.AdminUsers{}).Where("username", username).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
}

// Options 后台用户参数列表
func (m *AdminUserService) Options() (items []forms.AdminUserOptionResult, err error) {
	result := driver.GDB.Model(models.AdminUsers{}).Order("id DESC").Find(&items)
	if result.Error != nil {
		return items, result.Error
	}
	return items, nil
}

type UserInfoResult struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (m *AdminUserService) UserInfo(ctx *app.Context) (res UserInfoResult, err error) {
	userId := ctx.User.AdminId
	err = driver.GDB.Model(models.AdminUsers{}).Where("id = ?", userId).First(&res).Error
	if err != nil {
		err = errors.ErrRecordNotFound
		return
	}
	return
}

type RefreshForm struct {
	Token string `json:"token" form:"token"`
}
type RefreshResult struct {
	Token string `json:"token" form:"token"`
}
