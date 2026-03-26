package initialize

import (
	"gin_rbac_api/global"
	"gin_rbac_api/middleware"
	"gin_rbac_api/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.LoadHTMLGlob("templates/*")
	Router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:  []string{"Content-Type", "auth-token"},
		ExposeHeaders: []string{"*"},
	}))
	publicRouter := router.RouterGroupApp.Public
	adminRouter := router.RouterGroupApp.Admin
	PublicGroup := Router.Group("/api")
	{
		publicRouter.InitPublicRouter(PublicGroup)
	}
	AdminGroup := Router.Group("/api")
	AdminGroup.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "GET", "POST", "PATCH"},
		AllowHeaders:  []string{"Content-Type", "auth-token"},
		ExposeHeaders: []string{"*"},
	}))
	AdminGroup.Use(middleware.JWTAdminAuth()).Use(middleware.RoleCtrlHandler())
	{
		adminRouter.InitAdminRouter(AdminGroup)
	}

	if global.SYS_CONFIG.System.Mode == "dev" {
		docs.SwaggerInfo.BasePath = "/api"
		Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	return Router
}
