package api

import (
	"gin_rbac_api/api/admin"
	"gin_rbac_api/api/public"
)

type ApiGroup struct {
	AdminApiGroup  admin.ApiGroup
	PublicApiGroup public.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
