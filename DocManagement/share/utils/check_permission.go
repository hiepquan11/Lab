package utils

import (
	"document-management/core/config"
	"document-management/models"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func CheckPermission(workspaceId string, currentUserId string) (bool, error) {

	adapter, err := gormadapter.NewAdapterByDB(config.ConnectDB())
	if err != nil {
		return false, err
	}
	// load model policy casbin
	enforcer, err := casbin.NewEnforcer("share/casbin/model.conf", adapter)
	if err != nil {
		return false, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return false, err
	}

	actions := []models.Permission{models.Read, models.Del, models.Edit}

	for _, action := range actions {
		ok, err := enforcer.Enforce(currentUserId, workspaceId, string(action))
		if err != nil {
			return false, err
		}
		if ok {
			return true, nil
		}
	}
	return false, nil
}
