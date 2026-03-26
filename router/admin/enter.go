package admin

import (
	"gin_rbac_api/api"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	AdminRouter
}

type AdminRouter struct{}

func (s *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	adminRouter := Router.Group("admin")
	adminApi := api.ApiGroupApp.AdminApiGroup.AdminApi
	{
		adminRouter.GET("sysUser", adminApi.GetAdminUserInfo)
		adminRouter.GET("sysUserList", adminApi.GetAdminUserList)
		adminRouter.POST("sysUser", adminApi.CreateAdminInfo)
		adminRouter.PUT("sysUser/:sysUserId", adminApi.UpdateAdmiInfo)
		adminRouter.DELETE("sysUser/:sysUserId", adminApi.RemoveAdminInfo)
		adminRouter.GET("sysUser.csv", adminApi.ExportAdminInfo)

		adminRouter.POST("role", adminApi.CreateAdminRole)
		adminRouter.PUT("role/:roleId", adminApi.UpdateAdminRole)
		adminRouter.DELETE("role/:roleId", adminApi.RemoveAdminRole)
		adminRouter.GET("role", adminApi.GetAdminRole)
		adminRouter.POST("roleRel", adminApi.CreateAdminRoleRel)
		adminRouter.POST("roleRel/:roleRelId", adminApi.UpdateAdminRoleRel)
		adminRouter.DELETE("roleRel", adminApi.RemoveAdminRoleRel)
		adminRouter.GET("roleApi", adminApi.GetAdminRoleApi)
		adminRouter.POST("roleApi", adminApi.CreateAdminRoleApi)
		adminRouter.DELETE("role/:roleId/roleApi", adminApi.RemoveAdminRoleApi)

		adminRouter.POST("menu", adminApi.CreateAdminMenu)
		adminRouter.PUT("menu/:menuId", adminApi.UpdateAdminMenu)
		adminRouter.DELETE("menu/:menuId", adminApi.RemoveAdminMenu)
		adminRouter.GET("menu", adminApi.GetAdminMenu)
		adminRouter.GET("menuJson", adminApi.GetAdminMenuJson)
		adminRouter.POST("menuBtn", adminApi.CreateAdminMenuBtn)
		adminRouter.PUT("menuBtn/:menuBtnId", adminApi.UpdateAdminMenuBtn)
		adminRouter.DELETE("menuBtn/:menuBtnId", adminApi.RemoveAdminMenuBtn)
		adminRouter.GET("menuBtn", adminApi.GetAdminMenuBtn)
		adminRouter.POST("api", adminApi.CreateAdminSysApi)
		adminRouter.PUT("api/:apiId", adminApi.UpdateAdminSysApi)
		adminRouter.DELETE("api/:apiId", adminApi.RemoveAdminSysApi)
		adminRouter.GET("api", adminApi.GetAdminSysApi)
		adminRouter.GET("apiJson", adminApi.GetAdminSysApiJson)
		adminRouter.GET("sysMenu", adminApi.GetAdminRoleSysMenuJson)
		adminRouter.GET("roleMenuList", adminApi.GetAdminRoleMenuList)
		adminRouter.POST("role/:roleId/menu", adminApi.CreateAdminRoleMenu)
		adminRouter.PUT("role/:roleId/menu/:menuId", adminApi.UpdateAdminRoleMenu)
		adminRouter.DELETE("role/:roleId/menu/:menuId", adminApi.RemoveAdminRoleMenu)

		adminRouter.GET("serverInfo", adminApi.GetServerInfo)
		adminRouter.GET("opRecord", adminApi.GetOpRecord)
		adminRouter.DELETE("opRecord/:id", adminApi.DeleteOpRecord)

		adminRouter.PUT("password", adminApi.UpdateAdminPassword)
		adminRouter.POST("logout", adminApi.AdminLogout)
		adminRouter.PUT("avatar", adminApi.UpdateAdminAvatar)

	}
	return adminRouter
}
