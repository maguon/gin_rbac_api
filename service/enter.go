package service

type ServiceGroup struct {
	AdminApiService
	AdminMenuService
	AdminService
	CasbinService
	JwtService
	SystemService
	UserService
}

var ServiceGroupApp = new(ServiceGroup)
