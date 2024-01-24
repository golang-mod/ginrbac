package driver

import (
	"github.com/golang-mod/ginrbac/internal/rbac"
)

var Rbac *rbac.Synced

func InitRbac(filePath string) (err error) {
	Rbac, err = rbac.NewSynced(GDB, filePath)
	if err != nil {
		return err
	}
	return nil
}
