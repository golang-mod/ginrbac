package driver

import (
	"github.com/zhiniuer/goadmin/internal/rbac"
)

var Rbac *rbac.Synced

func InitRbac(filePath string) (err error) {
	Rbac, err = rbac.NewSynced(GDB, filePath)
	if err != nil {
		return err
	}
	return nil
}
