package admin

import (
	"gin_rbac_api/global"
	"gin_rbac_api/model/app"
	"gin_rbac_api/model/common"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

// CreateAdminRoleApi
// @Tags      Admin
// @Summary   新增用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param    data  body      res.AdminCasbin  true  ""
// @Success  200   {object}  common.Response{data=res.AdminCasbin,msg=string}  ""
// @Router    /admin/roleApi [post]
func (b *AdminApi) CreateAdminRoleApi(c *gin.Context) {
	var adminCasbin app.AdminCasbin
	err := c.ShouldBindBodyWith(&adminCasbin, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("新增角色API参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}

	if adminCasbinRes, err := casbinService.AddPolicy(adminCasbin); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Error(err))
		common.InteralError(err.Error(), "创建失败", c)
	} else {
		global.SYS_LOG.Info("新增角色API " + strconv.FormatInt(adminCasbinRes.ID, 10))
		common.OkWithDetailed(gin.H{"model": adminCasbinRes}, "创建成功", c)
	}
}

// GetAdminRoleApi
// @Tags      Admin
// @Summary   分页获取AdminCasbin列表
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Param     data  query     res.AdminCasbin                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  "AdminCasbin列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/roleApi [get]
func (s *AdminApi) GetAdminRoleApi(c *gin.Context) {
	var queryModel app.AdminCasbin
	err := c.ShouldBindQuery(&queryModel)
	if err != nil {
		global.SYS_LOG.Error("参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	list, err := casbinService.GetPolicyByRole(queryModel.V0)
	if err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Error(err))
		common.InteralError(err.Error(), "获取失败", c)
		return
	} else {
		global.SYS_LOG.Info("获取成功! ")
		common.OkWithDetailed(common.QueryResult{
			List:       list,
			Total:      1,
			PageNumber: 1,
			PageSize:   1,
		}, "获取成功", c)
	}
}

// RemoveAdminRoleApi
// @Tags      Admin
// @Summary   删除AdminRoleApi
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param roleId path int true "role ID"
// @Param     data  query     res.AdminCasbin                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  " "
// @Router    /admin/role/:roleId/roleApi [delete]
func (s *AdminApi) RemoveAdminRoleApi(c *gin.Context) {
	var adminCasbin app.AdminCasbin
	err := c.ShouldBindBodyWith(&adminCasbin, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("删除角色API参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	resFlag, err := casbinService.ClearCasbin(0, adminCasbin.V0, adminCasbin.V1, adminCasbin.V2)
	if err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Error(err))
		common.InteralError(err.Error(), "删除失败", c)
		return
	} else {
		if resFlag {
			global.SYS_LOG.Info("删除系统菜单成功! " + adminCasbin.V0 + "-" + adminCasbin.V1 + "-" + adminCasbin.V2)

		} else {
			global.SYS_LOG.Warn("删除系统菜单失败! " + adminCasbin.V0 + "-" + adminCasbin.V1 + "-" + adminCasbin.V2)

		}
		common.OkWithDetailed(common.QueryResult{
			Total: 1,
		}, "删除成功", c)
	}
}
