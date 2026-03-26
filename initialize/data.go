package initialize

import (
	"gin_rbac_api/global"
	"gin_rbac_api/model/app"
	"gin_rbac_api/model/data"

	"go.uber.org/zap"
)

func InitDbData() {
	err := global.SYS_DB.Create(&data.AdminApiList).Error
	if err != nil {
		global.SYS_LOG.Error("Init db data api list  err:", zap.Error(err))
	} else {
		global.SYS_LOG.Info("Init db data api list success !")
	}
	err = global.SYS_DB.Create(&data.AdminRoleList).Error
	if err != nil {
		global.SYS_LOG.Error("Init db data role list  err:", zap.Error(err))
	} else {
		global.SYS_LOG.Info("Init db data role list success !")
	}
	err = global.SYS_DB.Create(&data.AdminInfoList).Error
	if err != nil {
		global.SYS_LOG.Error("Init db data user list  err:", zap.Error(err))
	} else {
		global.SYS_LOG.Info("Init db data user list success !")
	}
	err = global.SYS_DB.Create(&data.AdminRoleRelList).Error
	if err != nil {
		global.SYS_LOG.Error("Init db data user role rel list  err:", zap.Error(err))
	} else {
		global.SYS_LOG.Info("Init db data user role list success !")
	}
	err = global.SYS_DB.Create(&data.AdminCasbinList).Error
	if err != nil {
		global.SYS_LOG.Error("Init db data admin casbin list  err:", zap.Error(err))
	} else {
		global.SYS_LOG.Info("Init db data admin casbin list success !")
	}
	err = global.SYS_DB.Create(&data.AdminParentMenuList).Error
	if err != nil {
		global.SYS_LOG.Error("Init db data parent menu err:", zap.Error(err))
	} else {
		global.SYS_LOG.Info("Init db data parent menu success !")
	}
	err = global.SYS_DB.Create(&data.AdminSubMenuList).Error
	if err != nil {
		global.SYS_LOG.Error("Init db data sub menu err:", zap.Error(err))
	} else {
		global.SYS_LOG.Info("Init db data sub menu success !")
	}
	var adminRoleMenuList = []app.AdminRoleMenu{}
	for _, menu := range data.AdminParentMenuList {
		if menu.Url != "" {
			adminRoleMenuList = append(adminRoleMenuList, app.AdminRoleMenu{AdminRoleId: 10000, AdminMenuId: menu.ID})
		}
	}
	for _, menu := range data.AdminSubMenuList {
		adminRoleMenuList = append(adminRoleMenuList, app.AdminRoleMenu{AdminRoleId: 10000, AdminMenuId: menu.ID})
	}
	err = global.SYS_DB.Create(&adminRoleMenuList).Error
	if err != nil {
		global.SYS_LOG.Error("Init db data role menu rel err:", zap.Error(err))
	} else {
		global.SYS_LOG.Info("Init db data role menu rel success !")
	}
}
