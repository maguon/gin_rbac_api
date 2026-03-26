package public

import "gin_rbac_api/service"

type PublicApi struct {
}

type ApiGroup struct {
	PublicApi
}

var (
	jwtService   = service.ServiceGroupApp.JwtService
	adminService = service.ServiceGroupApp.AdminService
)
