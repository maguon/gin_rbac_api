package data

import (
	"gin_rbac_api/model/app"
	"gin_rbac_api/model/common"
)

var AdminParentMenuList = []app.AdminMenu{
	{Status: 1, Level: 1, EXTEND_MODEL: common.EXTEND_MODEL{ID: 1000}, ParentId: 0, Sort: 99, MenuName: "首页", Url: "/home", Icon: "mdiMonitorDashboard", Component: "Home"},
	{Status: 1, Level: 1, EXTEND_MODEL: common.EXTEND_MODEL{ID: 1001}, ParentId: 0, Sort: 5, MenuName: "系统状态", Url: "/sys_info", Icon: "mdiHarddisk", Component: "SysInfo"},
	{Status: 1, Level: 1, EXTEND_MODEL: common.EXTEND_MODEL{ID: 1002}, ParentId: 0, Sort: 3, MenuName: "系统设置", Icon: "mdiCog"},
	{Status: 1, Level: 1, EXTEND_MODEL: common.EXTEND_MODEL{ID: 1003}, ParentId: 0, Sort: 0, MenuName: "关于", Url: "/about", Icon: "mdiMonitorDashboard", Component: "About"},
}
var AdminSubMenuList = []app.AdminMenu{
	{Status: 1, Level: 2, ParentId: 1002, Sort: 9, MenuName: "API管理", Url: "/sys_api", Icon: "mdiApi", Component: "SysApi"},
	{Status: 1, Level: 2, ParentId: 1002, Sort: 7, MenuName: "菜单管理", Url: "/sys_menu", Icon: "mdiSilverware", Component: "SysMenu"},
	{Status: 1, Level: 2, ParentId: 1002, Sort: 5, MenuName: "角色管理", Url: "/sys_role", Icon: "mdiFamilyTree", Component: "SysRole"},
	{Status: 1, Level: 2, ParentId: 1002, Sort: 3, MenuName: "系统用户", Url: "/sys_user", Icon: "mdiAccountCog", Component: "SysUser"},
}
