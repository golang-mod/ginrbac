package dao

import (
	"github.com/pkg/errors"
	"github.com/zhiniuer/goadmin/internal/app/models"
	"github.com/zhiniuer/goadmin/internal/rbac"
	"gorm.io/gorm"
	"strconv"
)

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

// AuthStore 保存授权数据
func (d *UserDao) AuthStore(tx *gorm.DB, userId int, roleList, permissionList []string) (err error) {
	r, _ := rbac.New(tx)

	// 清除user和role绑定关系
	result := tx.Where("user_id", userId).Delete(&models.AdminRoleUsers{})
	if result.Error != nil {
		return err
	}
	// 清除user和permission绑定关系
	result = tx.Where("user_id", userId).Delete(&models.AdminUserPermissions{})
	if result.Error != nil {
		return err
	}
	// 清除user_id和Casbin关系
	err = r.DeleteUserGroupingPolicy(strconv.Itoa(userId))
	if err != nil {
		return err
	}
	// 重新绑定user和role关系
	var roleUsers = make([]models.AdminRoleUsers, 0, len(roleList))
	// 防止传递无效参数
	var checkRoleList []string
	tx.Model(&models.AdminRoles{}).Where("id IN ?", roleList).Pluck("id", &checkRoleList)
	if len(checkRoleList) != len(roleList) {
		return errors.New("角色不存在或已删除，请刷新重试")
	}
	for _, v := range roleList {
		pid, _ := strconv.Atoi(v)
		roleUsers = append(roleUsers, models.AdminRoleUsers{UserId: userId, RoleId: pid})
	}
	if len(roleUsers) > 0 {
		result = tx.Create(roleUsers)
		if result.Error != nil {
			return err
		}
		// 重新绑定user和role Casbin关系
		err = r.AddUserRoles(strconv.Itoa(userId), roleList)
		if err != nil {
			return err
		}
	}

	// 重新绑定user和permission关系
	var userPermissions = make([]models.AdminUserPermissions, 0, len(permissionList))
	// 防止传递无效参数
	var checkPermissionList []string
	tx.Model(&models.AdminPermissions{}).Where("id IN ?", permissionList).Pluck("id", &checkPermissionList)
	if len(checkPermissionList) != len(permissionList) {
		return errors.New("权限不存在或已删除，请刷新重试")
	}
	for _, v := range permissionList {
		pid, _ := strconv.Atoi(v)
		userPermissions = append(userPermissions, models.AdminUserPermissions{UserId: userId, PermissionId: pid})
	}
	if len(userPermissions) > 0 {
		result = tx.Create(userPermissions)
		if result.Error != nil {
			return err
		}
		// 重新绑定user和permission Casbin关系
		err = r.AddUserPermissions(strconv.Itoa(int(userId)), permissionList)
		if err != nil {
			return err
		}
	}
	return nil
}
