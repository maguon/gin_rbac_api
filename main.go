package main

import (
	"fmt"
	"gin_rbac_api/global"
	"gin_rbac_api/initialize"
	"strconv"
)

func main() {
	global.SYS_VP = initialize.Viper()

	fmt.Println(global.SYS_CONFIG.Zap.OutputPaths)
	global.SYS_LOG = initialize.Zap()
	initialize.Redis()
	initialize.Mongo()
	global.SYS_DB = initialize.Gorm()
	initialize.Casbin()
	Router := initialize.Routers()
	global.SYS_LOG.Info("server start ")
	// initialize.InitDbData() //第一次登陆需要初始化数据
	Router.Run(":" + strconv.Itoa(global.SYS_CONFIG.System.Port))
}
