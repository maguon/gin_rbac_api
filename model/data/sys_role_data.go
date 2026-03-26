package data

import (
	"gin_rbac_api/model/app"
	"gin_rbac_api/model/common"
	"gin_rbac_api/utils"
)

var AdminInfoList = []app.AdminInfo{
	{EXTEND_MODEL: common.EXTEND_MODEL{ID: 10000}, Status: utils.GetInt8Pointer(1), Username: "admin", Phone: "18888888888", Gender: 1, Password: utils.BcryptHash("123456")},
}

var AdminRoleList = []app.AdminRole{
	{EXTEND_MODEL: common.EXTEND_MODEL{ID: 10000}, Status: 1, RoleName: "super", DefaultUrl: "/home", Remark: "超级管理员"},
}

var AdminRoleRelList = []app.AdminRoleRel{
	{AdminId: 10000, AdminRoleId: 10000},
}
