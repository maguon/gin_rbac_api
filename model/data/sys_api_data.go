package data

import "gin_rbac_api/model/app"

var AdminApiList = []app.AdminSysApi{
	{Status: 1, Tag: "API", Method: "GET", Url: "/api/admin/api", Remark: "获取API列表"},
	{Status: 1, Tag: "API", Method: "GET", Url: "/api/admin/apiJson", Remark: "获取API格式JSON"},
	{Status: 1, Tag: "API", Method: "PUT", Url: "/api/admin/api/:apiId", Remark: "修改API信息"},
	{Status: 1, Tag: "API", Method: "DELETE", Url: "/api/admin/api/:apiId", Remark: "删除API信息"},
	{Status: 1, Tag: "API", Method: "POST", Url: "/api/admin/api", Remark: "创建API信息"},

	{Status: 1, Tag: "角色", Method: "GET", Url: "/api/admin/role", Remark: "获取角色信息"},
	{Status: 1, Tag: "角色", Method: "PUT", Url: "/api/admin/role/:roleId", Remark: "修改角色信息"},
	{Status: 1, Tag: "角色", Method: "DELETE", Url: "/api/admin/role/:roleId", Remark: "删除角色信息"},
	{Status: 1, Tag: "角色", Method: "POST", Url: "/api/admin/role", Remark: "创建角色信息"},
	{Status: 1, Tag: "角色", Method: "POST", Url: "/api/admin/roleRel", Remark: "创建用户角色信息"},
	{Status: 1, Tag: "角色", Method: "PUT", Url: "/api/admin/roleRel/:roleRelId", Remark: "修改用户角色信息"},
	{Status: 1, Tag: "角色", Method: "DELETE", Url: "/api/admin/roleRel", Remark: "删除用户角色信息"},
	{Status: 1, Tag: "角色", Method: "POST", Url: "/api/admin/roleApi", Remark: "创建角色API信息"},
	{Status: 1, Tag: "角色", Method: "GET", Url: "/api/admin/roleApi", Remark: "获取角色API信息"},
	{Status: 1, Tag: "角色", Method: "DELETE", Url: "role/:roleId/roleApi", Remark: "删除角色API信息"},

	{Status: 1, Tag: "菜单", Method: "POST", Url: "/api/admin/menu", Remark: "创建菜单信息"},
	{Status: 1, Tag: "菜单", Method: "PUT", Url: "/api/admin/menu/:menuId", Remark: "修改菜单信息"},
	{Status: 1, Tag: "菜单", Method: "DELETE", Url: "/api/admin/menu/:menuId", Remark: "删除菜单信息"},
	{Status: 1, Tag: "菜单", Method: "GET", Url: "/api/admin/menu", Remark: "查询菜单列表"},
	{Status: 1, Tag: "菜单", Method: "GET", Url: "/api/admin/menuJson", Remark: "查询所有菜单JSON格式"},
	{Status: 1, Tag: "菜单", Method: "POST", Url: "/api/admin/menuBtn", Remark: "创建菜单按钮信息"},
	{Status: 1, Tag: "菜单", Method: "PUT", Url: "/api/admin/menuBtn/:menuBtnId", Remark: "修改菜单按钮信息"},
	{Status: 1, Tag: "菜单", Method: "DELETE", Url: "/api/admin/menuBtn/:menuBtnId", Remark: "删除菜单按钮信息"},
	{Status: 1, Tag: "菜单", Method: "GET", Url: "/api/admin/menuBtn", Remark: "查询菜单按钮列表"},

	{Status: 1, Tag: "角色菜单", Method: "GET", Url: "/api/admin/roleMenuList", Remark: "获取角色菜单列表"},
	{Status: 1, Tag: "角色菜单", Method: "PUT", Url: "/api/admin/role/:roleId/menu/:menuId", Remark: "修改角色菜单信息"},
	{Status: 1, Tag: "角色菜单", Method: "DELETE", Url: "/api/admin/role/:roleId/menu/:menuId", Remark: "删除角色信息"},
	{Status: 1, Tag: "角色菜单", Method: "POST", Url: "/api/admin/role/:roleId/menu", Remark: "创建角色菜单信息"},

	{Status: 1, Tag: "Server", Method: "GET", Url: "/api/admin/serverInfo", Remark: "获取Server状态"},
	{Status: 1, Tag: "Server", Method: "GET", Url: "/api/admin/opRecord", Remark: "获取操作记录"},
	{Status: 1, Tag: "Server", Method: "DELETE", Url: "/api/admin/opRecord/:id", Remark: "删除操作记录"},

	{Status: 1, Tag: "用户", Method: "GET", Url: "/api/admin/sysMenu", Remark: "获取用户菜单"},
	{Status: 1, Tag: "用户", Method: "POST", Url: "/api/admin/sysUser", Remark: "创建系统用户"},
	{Status: 1, Tag: "用户", Method: "GET", Url: "/api/admin/sysUser", Remark: "获取用户信息"},
	{Status: 1, Tag: "用户", Method: "GET", Url: "/api/admin/sysUser.csv", Remark: "导出用户信息"},
	{Status: 1, Tag: "用户", Method: "PUT", Url: "/api/admin/sysUser/:sysUserId", Remark: "修改用户信息"},
	{Status: 1, Tag: "用户", Method: "GET", Url: "/api/admin/sysUserList", Remark: "获取系统用户列表"},
	{Status: 1, Tag: "用户", Method: "PUT", Url: "/api/admin/password", Remark: "修改用户密码"},
	{Status: 1, Tag: "用户", Method: "PUT", Url: "/api/admin/avatar", Remark: "修改用户头像"},
	{Status: 1, Tag: "用户", Method: "POST", Url: "/api/admin/logout", Remark: "退出登录"},
	{Status: 1, Tag: "用户", Method: "DELETE", Url: "/api/admin/sysUser/:sysUserId", Remark: "删除系统用户"},
}
