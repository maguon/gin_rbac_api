package middleware

import (
	"bytes"
	"gin_rbac_api/global"
	"gin_rbac_api/model/app"
	"gin_rbac_api/model/common"
	"gin_rbac_api/service"
	"gin_rbac_api/utils"
	"io/ioutil"
	"strconv"

	"github.com/gin-gonic/gin"
)

var systemService = service.ServiceGroupApp.SystemService

func RoleCtrlHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		adminCalims := utils.GetAdminContext(c)
		obj := c.Request.URL.Path
		//obj := strings.TrimPrefix(c.Request.URL.Path, "/api/admin")
		act := c.Request.Method
		sub := strconv.FormatInt(adminCalims.RoleId, 10)
		success, _ := global.SYS_CASBIN.Enforce(sub, obj, act)
		if !success {
			common.ForbbidnError(gin.H{"reload": true}, "权限不足", c)
			//response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		//create oprecord obj
		var opRecord app.OpRecord
		opRecord.AdminId = adminCalims.ID
		opRecord.RoleId = adminCalims.RoleId
		opRecord.AdminName = adminCalims.Username
		opRecord.RoleName = adminCalims.RoleName
		opRecord.Method = c.Request.Method
		opRecord.URL = c.Request.URL.Path
		if c.Request.Method != "GET" {
			body, err := ioutil.ReadAll(c.Request.Body)
			if err == nil {
				c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			}
			opRecord.Params = string(body)
		}

		c.Next()

		if c.Writer.Status() == 200 {
			if c.Request.Method != "GET" {
				systemService.AddRecordList(opRecord)
			}
		}
	}
}
