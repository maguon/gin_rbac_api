package initialize

import (
	"gin_rbac_api/global"
	"gin_rbac_api/model/app"

	"github.com/casbin/casbin/v3"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
)

func Casbin() {
	a, err := gormadapter.NewAdapterByDBWithCustomTable(global.SYS_DB, &app.AdminCasbin{}, "admin_casbin")
	if err != nil {
		global.SYS_LOG.Error("db casbin rules read failed, err:", zap.Error(err))
	}
	e, err := casbin.NewSyncedCachedEnforcer("rbac_model.conf", a)
	if err != nil {
		global.SYS_LOG.Error("system casbin config read failed, err:", zap.Error(err))
	}
	println("e")
	// Load the policy from DB.
	err = e.LoadPolicy()
	if err != nil {
		global.SYS_LOG.Error("system load casbin policy failed, err:", zap.Error(err))
	} else {
		global.SYS_CASBIN = e
		global.SYS_LOG.Info("system load casbin policy completed!")
	}
}
