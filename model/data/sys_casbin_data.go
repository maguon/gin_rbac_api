package data

import "gin_rbac_api/model/app"

var AdminCasbinList = []app.AdminCasbin{
	{Ptype: "p", V0: "10000", V1: "/api/admin/api", V2: "POST"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/api", V2: "GET"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/apiJson", V2: "GET"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/api/:apiId", V2: "PUT"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/api/:apiId", V2: "DELETE"},

	{Ptype: "p", V0: "10000", V1: "/api/admin/role", V2: "POST"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/role", V2: "GET"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/role/:roleId", V2: "PUT"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/role/:roleId", V2: "DELETE"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/roleRel", V2: "POST"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/roleRel/:roleRelId", V2: "PUT"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/roleRel", V2: "DELETE"},

	{Ptype: "p", V0: "10000", V1: "/api/admin/menu", V2: "POST"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/menu", V2: "GET"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/menuJson", V2: "GET"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/menu/:menuId", V2: "PUT"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/menu/:menuId", V2: "DELETE"},

	{Ptype: "p", V0: "10000", V1: "/api/admin/menuBtn", V2: "POST"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/menuBtn", V2: "GET"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/menuBtn/:menuBtnId", V2: "PUT"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/menuBtn/:menuBtnId", V2: "DELETE"},

	{Ptype: "p", V0: "10000", V1: "/api/admin/role/:roleId/menu", V2: "POST"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/roleMenuList", V2: "GET"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/role/:roleId/menu/:menuId", V2: "PUT"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/role/:roleId/menu/:menuId", V2: "DELETE"},

	{Ptype: "p", V0: "10000", V1: "/api/admin/serverInfo", V2: "GET"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/sysMenu", V2: "GET"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/password", V2: "PUT"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/avatar", V2: "PUT"},
	{Ptype: "p", V0: "10000", V1: "/api/admin/logout", V2: "POST"},
}
