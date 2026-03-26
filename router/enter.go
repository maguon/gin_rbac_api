package router

import (
	"gin_rbac_api/router/admin"
	"gin_rbac_api/router/public"
)

type RouterGroup struct {
	Admin  admin.RouterGroup
	Public public.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
