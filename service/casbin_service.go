package service

import (
	"gin_rbac_api/global"
	"gin_rbac_api/model/app"
)

type CasbinService struct{}

func (casbinService *CasbinService) AddPolicy(adminCasbin app.AdminCasbin) (res app.AdminCasbin, err error) {
	err = global.SYS_DB.Create(&adminCasbin).Error
	global.SYS_CASBIN.LoadPolicy()
	return adminCasbin, err
}

// casbin table v0 roleid ,v1 path ,v2 method
func (casbinService *CasbinService) UpdateCasbinApiPath(oldPath string, newPath string, oldMethed string, newMethod string) error {
	err := global.SYS_DB.Model(&app.AdminCasbin{}).Where("v1 = ? AND v2 = ? ", oldPath, oldMethed).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	if err != nil {
		return err
	}
	err = global.SYS_CASBIN.LoadPolicy()
	return err
}

func (casbinService *CasbinService) RemoveAdminCasbinApi(path string, method string) (err error) {
	var adminCasbin app.AdminCasbin
	err = global.SYS_DB.Where("v1 = ? AND v2 = ? ", path, method).Delete(&adminCasbin).Error
	if err != nil {
		return err
	} else {
		global.SYS_CASBIN.LoadPolicy()
		return nil
	}
}

// 根据角色ID获取对应API信息
func (casbinService *CasbinService) GetPolicyByRole(roleId string) ([][]string, error) {
	list, err := global.SYS_CASBIN.GetFilteredPolicy(0, roleId)
	return list, err
}

// 清除匹配的权限
// @param:v int policy index, p ...string value
func (casbinService *CasbinService) ClearCasbin(v int, p ...string) (casbinRes bool, err error) {
	res, err := global.SYS_CASBIN.RemoveFilteredPolicy(v, p...)
	return res, err
}
