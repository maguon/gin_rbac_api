package public

import (
	api "gin_rbac_api/api"

	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	PublicRouter
}

type PublicRouter struct{}

func (s *PublicRouter) InitPublicRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	publicRouter := Router.Group("public")
	publicApi := api.ApiGroupApp.PublicApiGroup.PublicApi
	{
		publicRouter.POST("adminLogin", publicApi.AdminLogin)
		publicRouter.GET("captcha", publicApi.Captcha)
	}
	return publicRouter
}
