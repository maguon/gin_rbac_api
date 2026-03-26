package admin

import (
	"gin_rbac_api/global"
	"gin_rbac_api/model/app"
	"gin_rbac_api/model/common"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetServerInfo
// @Tags      System
// @Summary   获取服务器信息
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  common.Response{data=map[string]interface{},msg=string}  "获取服务器信息"
// @Router    /admin/serverInfo [get]
func (s *AdminApi) GetServerInfo(c *gin.Context) {
	server, err := systemService.GetServerInfo()
	if err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Error(err))
		common.InteralError(err.Error(), "获取失败", c)
		return
	} else {
		global.SYS_LOG.Info("获取服务器信息成功! ")
		common.OkWithDetailed(gin.H{"server": server}, "获取成功", c)
	}
}

// GetOpRecord
// @Tags      System
// @Summary   获取服务器信息
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  common.Response{data=map[string]interface{},msg=string}  "获取服务器信息"
// @Router    /admin/opRecord [get]
func (s *AdminApi) GetOpRecord(c *gin.Context) {
	var queryModel app.OpRecordQuery
	err := c.ShouldBindQuery(&queryModel)
	if err != nil {
		global.SYS_LOG.Error("参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	list, total, err := systemService.GetRecordList(queryModel)
	if err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Error(err))
		common.InteralError(err.Error(), "获取失败", c)
		return
	} else {
		global.SYS_LOG.Info("获取成功! ")
		common.OkWithDetailed(common.QueryResult{
			List:       list,
			Total:      total,
			PageNumber: queryModel.PageNumber,
			PageSize:   queryModel.PageSize,
		}, "获取成功", c)
	}
}

// DelOpRecord
// @Tags      System
// @Summary   获取服务器信息
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  common.Response{data=map[string]interface{},msg=string}  "获取服务器信息"
// @Router    /admin/opRecord/{id} [delete]
func (s *AdminApi) DeleteOpRecord(c *gin.Context) {
	id := c.Param("id")
	res, err := systemService.RemoveRecord(id)
	if err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Error(err))
		common.InteralError(err.Error(), "获取失败", c)
		return
	} else {
		global.SYS_LOG.Info("删除成功! ")
		common.OkWithDetailed(gin.H{"res": res}, "删除成功", c)
	}
}
