package admin

import "gin_rbac_api/service"

type AdminApi struct {
}

type ApiGroup struct {
	AdminApi
}

var (
	adminApiService  = service.ServiceGroupApp.AdminApiService
	adminMenuService = service.ServiceGroupApp.AdminMenuService
	adminService     = service.ServiceGroupApp.AdminService
	casbinService    = service.ServiceGroupApp.CasbinService
	jwtService       = service.ServiceGroupApp.JwtService
	systemService    = service.ServiceGroupApp.SystemService
)
