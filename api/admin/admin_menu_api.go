package admin

import (
	"fmt"
	"gin_rbac_api/global"
	"gin_rbac_api/model/app"
	"gin_rbac_api/model/common"
	"gin_rbac_api/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

// CreateAdminMenu
// @Tags      Admin
// @Summary   新增用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param    data  body      res.AdminMenu  true  ""
// @Success  200   {object}  common.Response{data=res.AdminMenu,msg=string}  ""
// @Router    /admin/menu [post]
func (b *AdminApi) CreateAdminMenu(c *gin.Context) {
	var adminMenu app.AdminMenu
	err := c.ShouldBindBodyWith(&adminMenu, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("新增菜单参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}

	if adminMenuRes, err := adminMenuService.AddAdminMenu(adminMenu); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Error(err))
		common.InteralError(err.Error(), "创建失败", c)
	} else {
		global.SYS_LOG.Info("新增菜单 " + strconv.FormatInt(adminMenuRes.ID, 10))
		common.OkWithDetailed(gin.H{"model": adminMenuRes}, "创建成功", c)
	}
}

// UpdateAdminMenu
// @Tags      Admin
// @Summary   更新系统用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param roleId path int true "role ID"
// @Param    data  body      res.AdminMenu  true  ""
// @Success  200   {object}  common.Response{data=res.AdminMenu,msg=string}  ""
// @Router    /admin/menu/{menuId} [put]
func (b *AdminApi) UpdateAdminMenu(c *gin.Context) {
	var adminMenu app.AdminMenu
	err := c.ShouldBindBodyWith(&adminMenu, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("更新系统菜单参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	menuId, _ := strconv.ParseInt(c.Param("menuId"), 10, 64)
	adminMenu.ID = menuId
	adminMenuRes, err := adminMenuService.UpdateAdminMenu(adminMenu)
	if err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Error(err))
		common.InteralError(err.Error(), "更新失败", c)
		return
	} else {
		global.SYS_LOG.Info("更新系统菜单成功 " + strconv.FormatInt(menuId, 10))
		common.OkWithDetailed(gin.H{"model": adminMenuRes}, "更新成功", c)
	}
}

// GetAdminMenu
// @Tags      Admin
// @Summary   分页获取AdminMenu列表
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Param     data  query     res.AdminMenuQuery                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  "AdminMenu列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/menu [get]
func (s *AdminApi) GetAdminMenu(c *gin.Context) {
	var queryModel app.AdminMenuQuery
	err := c.ShouldBindQuery(&queryModel)
	if err != nil {
		global.SYS_LOG.Error("参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	list, total, err := adminMenuService.GetAdminMenuList(queryModel)
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

// GetAdminMenuJson
// @Tags      Admin
// @Summary   分页获取AdminMenu列表
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Param     data  query     res.AdminMenuQuery                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  "AdminMenu列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/menuJson [get]
func (s *AdminApi) GetAdminMenuJson(c *gin.Context) {
	list, total, err := adminMenuService.GetAdminMenuJson()
	if err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Error(err))
		common.InteralError(err.Error(), "获取失败", c)
		return
	} else {
		global.SYS_LOG.Info("获取成功! ")
		common.OkWithDetailed(common.QueryResult{
			List:       list,
			Total:      total,
			PageNumber: 1,
			PageSize:   1,
		}, "获取成功", c)
	}
}

// RemoveAdminMenu
// @Tags      Admin
// @Summary   删除AdminMenu
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param menuId path int true "menu ID"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  " "
// @Router    /admin/menu/{menuId} [delete]
func (s *AdminApi) RemoveAdminMenu(c *gin.Context) {
	menuId, _ := strconv.ParseInt(c.Param("menuId"), 10, 64)

	err := adminMenuService.RemoveAdminMenu(menuId)
	if err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Error(err))
		common.InteralError(err.Error(), "删除失败", c)
		return
	} else {
		global.SYS_LOG.Info("删除系统菜单成功! " + strconv.FormatInt(menuId, 10))
		adminMenuService.RemoveAdminMenuAllBtn((menuId))
		adminMenuService.RemoveAdminRolMenuIdRel((menuId))
		common.OkWithDetailed(common.QueryResult{
			Total: 1,
		}, "删除成功", c)
	}
}

// CreateAdminMenuBtn
// @Tags      Admin
// @Summary   新增用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param    data  body      res.AdminMenuBtn  true  ""
// @Success  200   {object}  common.Response{data=res.AdminMenuBtn,msg=string}  ""
// @Router    /admin/menu [post]
func (b *AdminApi) CreateAdminMenuBtn(c *gin.Context) {
	var adminMenuBtn app.AdminMenuBtn
	err := c.ShouldBindBodyWith(&adminMenuBtn, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("新增菜单按钮参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}

	if adminMenuBtnes, err := adminMenuService.AddAdminMenuBtn(adminMenuBtn); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Error(err))
		common.InteralError(err.Error(), "创建失败", c)
	} else {
		global.SYS_LOG.Info("新增菜单按钮 " + strconv.FormatInt(adminMenuBtnes.ID, 10))
		common.OkWithDetailed(gin.H{"model": adminMenuBtnes}, "创建成功", c)
	}
}

// UpdateAdminMenuBtn
// @Tags      Admin
// @Summary   更新系统菜单按钮
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param roleId path int true "role ID"
// @Param    data  body      res.AdminMenuBtn  true  ""
// @Success  200   {object}  common.Response{data=res.AdminMenuBtn,msg=string}  ""
// @Router    /admin/menuBtn/{menuBtnId} [put]
func (b *AdminApi) UpdateAdminMenuBtn(c *gin.Context) {
	var adminMenuBtn app.AdminMenuBtn
	err := c.ShouldBindBodyWith(&adminMenuBtn, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("更新系统菜单按钮参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	menuBtnId, _ := strconv.ParseInt(c.Param("menuBtnId"), 10, 64)
	adminMenuBtn.ID = menuBtnId
	adminMenuBtnRes, err := adminMenuService.UpdateAdminMenuBtn(adminMenuBtn)
	if err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Error(err))
		common.InteralError(err.Error(), "更新失败", c)
		return
	} else {
		global.SYS_LOG.Info("更新系统菜单按钮成功 " + strconv.FormatInt(menuBtnId, 10))
		common.OkWithDetailed(gin.H{"model": adminMenuBtnRes}, "更新成功", c)
	}
}

// GetAdminMenuBtn
// @Tags      Admin
// @Summary   分页获取AdminMenuBtn列表
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Param     data  query     res.AdminMenuBtnQuery                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  "AdminMenuBtn列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/menuBtn [get]
func (s *AdminApi) GetAdminMenuBtn(c *gin.Context) {
	var queryModel app.AdminMenuBtnQuery
	err := c.ShouldBindQuery(&queryModel)
	if err != nil {
		global.SYS_LOG.Error("参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	list, total, err := adminMenuService.GetAdminMenuBtn(queryModel)
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

// RemoveAdminMenuBtn
// @Tags      Admin
// @Summary   删除AdminMenuBtn
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param menuBtnId path int true "menu btn ID"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  " "
// @Router    /admin/menuBtn/{menuBtnId} [delete]
func (s *AdminApi) RemoveAdminMenuBtn(c *gin.Context) {
	menuBtnId, _ := strconv.ParseInt(c.Param("menuBtnId"), 10, 64)

	err := adminMenuService.RemoveAdminMenuBtn(menuBtnId)
	if err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Error(err))
		common.InteralError(err.Error(), "删除失败", c)
		return
	} else {
		global.SYS_LOG.Info("删除系统菜单按钮成功! " + strconv.FormatInt(menuBtnId, 10))
		common.OkWithDetailed(common.QueryResult{
			Total: 1,
		}, "删除成功", c)
	}
}

// CreateAdminSysApi
// @Tags      Admin
// @Summary   新增用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param    data  body      res.AdminSysApi  true  ""
// @Success  200   {object}  common.Response{data=res.AdminSysApi,msg=string}  ""
// @Router    /admin/api [post]
func (b *AdminApi) CreateAdminSysApi(c *gin.Context) {
	var adminSysApi app.AdminSysApi
	err := c.ShouldBindBodyWith(&adminSysApi, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("新增API参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}

	if adminSysApies, err := adminApiService.AddAdminApi(adminSysApi); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Error(err))
		common.InteralError(err.Error(), "创建失败", c)
	} else {
		global.SYS_LOG.Info("新增系统API " + strconv.FormatInt(adminSysApies.ID, 10))
		common.OkWithDetailed(gin.H{"model": adminSysApies}, "创建成功", c)
	}
}

// UpdateAdminSysApi
// @Tags      Admin
// @Summary   更新系统用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param roleId path int true "role ID"
// @Param    data  body      res.AdminSysApi  true  ""
// @Success  200   {object}  common.Response{data=res.AdminSysApi,msg=string}  ""
// @Router    /admin/api/{apiId} [put]
func (b *AdminApi) UpdateAdminSysApi(c *gin.Context) {
	var adminSysApi app.AdminSysApi
	err := c.ShouldBindBodyWith(&adminSysApi, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("更新系统API参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	apiId, _ := strconv.ParseInt(c.Param("apiId"), 10, 64)

	var adminSysApiQuery app.AdminSysApiQuery
	adminSysApiQuery.Method = adminSysApi.Method
	adminSysApiQuery.Url = adminSysApi.Url
	_, total, err := adminApiService.GetAdminApiList(adminSysApiQuery)
	if err != nil {
		global.SYS_LOG.Error("更新系统API查询错误! ", zap.Error(err))
		common.RequestError(err.Error(), "查询错误", c)
		return
	}
	if total > 0 {
		global.SYS_LOG.Warn("更新系统API失败 " + strconv.FormatInt(apiId, 10))
		common.FailWithMessage("更新失败已存在相同的URL和METHOD", c)
		return
	}
	var adminSysApiQuery2 app.AdminSysApiQuery
	adminSysApiQuery2.ID = apiId
	adminSysApiQuery2.PageNumber = 1
	adminSysApiQuery2.PageSize = 10
	list, total, err := adminApiService.GetAdminApiList(adminSysApiQuery2)

	adminSysApi.ID = apiId
	adminSysApiRes, err := adminApiService.UpdateAdminApi(adminSysApi)
	if err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Error(err))
		common.InteralError(err.Error(), "更新失败", c)
		return
	} else {

		if list[0].Method == adminSysApi.Method && list[0].Url == adminSysApi.Url {
			global.SYS_LOG.Info("更新系统API成功 " + strconv.FormatInt(apiId, 10))
			common.OkWithDetailed(gin.H{"model": adminSysApiRes}, "更新成功", c)
			return
		} else {
			err = casbinService.UpdateCasbinApiPath(list[0].Url, adminSysApi.Url, list[0].Method, adminSysApi.Method)
			if err != nil {
				global.SYS_LOG.Error("同步更新RBAC失败!", zap.Error(err))
				common.InteralError(err.Error(), "同步更新RBAC失败", c)
				return
			} else {
				global.SYS_LOG.Info("更新系统API并同步casbin成功 " + strconv.FormatInt(apiId, 10))
				common.OkWithDetailed(gin.H{"model": adminSysApiRes}, "更新成功", c)
				return
			}
		}
	}
}

// GetAdminSysApi
// @Tags      Admin
// @Summary   分页获取AdminSysApi列表
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Param     data  query     res.AdminSysApiQuery                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  "AdminSysApi列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/api [get]
func (s *AdminApi) GetAdminSysApi(c *gin.Context) {
	var queryModel app.AdminSysApiQuery
	err := c.ShouldBindQuery(&queryModel)
	if err != nil {
		global.SYS_LOG.Error("参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	list, total, err := adminApiService.GetAdminApiList(queryModel)
	fmt.Println(len(list))
	fmt.Println(queryModel)
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

// GetAdminSysApiJson
// @Tags      Admin
// @Summary   分页获取AdminSysApi列表
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Param     data  query     res.AdminSysApiQuery                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  "AdminSysApi列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/apiJson[get]
func (s *AdminApi) GetAdminSysApiJson(c *gin.Context) {

	list, total, err := adminApiService.GetAdminApiJson()
	if err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Error(err))
		common.InteralError(err.Error(), "获取失败", c)
		return
	} else {
		global.SYS_LOG.Info("获取成功! ")
		common.OkWithDetailed(common.QueryResult{
			List:       list,
			Total:      total,
			PageNumber: 1,
			PageSize:   1,
		}, "获取成功", c)
	}
}

// RemoveAdminSysApi
// @Tags      Admin
// @Summary   删除AdminSysApi
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param apiId path int true "api ID"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  " "
// @Router    /admin/api/{apiId} [delete]
func (s *AdminApi) RemoveAdminSysApi(c *gin.Context) {
	apiId, _ := strconv.ParseInt(c.Param("apiId"), 10, 64)
	var adminSysApi app.AdminSysApi
	err := c.ShouldBindBodyWith(&adminSysApi, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("删除角色API参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}

	err = adminApiService.RemoveAdminApi(apiId)

	if err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Error(err))
		common.InteralError(err.Error(), "删除失败", c)
		return
	} else {
		global.SYS_LOG.Info("删除系统Api成功! " + strconv.FormatInt(apiId, 10))
		err = casbinService.RemoveAdminCasbinApi(adminSysApi.Url, adminSysApi.Method)
		if err != nil {
			global.SYS_LOG.Error("同步删除RBAC失败!", zap.Error(err))
			common.InteralError(err.Error(), "同步删除RBAC失败", c)
			return
		} else {
			common.OkWithDetailed(common.QueryResult{
				Total: 1,
			}, "删除成功", c)
			return
		}
	}
}

// GetAdminRoleMenuList
// @Tags      Admin
// @Summary   分页获取AdminRoleMenu列表
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Param     data  query     res.GetAdminRoleMenuList                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  "AdminRoleMenu列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/roleMenuList[get]
func (s *AdminApi) GetAdminRoleMenuList(c *gin.Context) {
	var queryModel app.AdminRoleMenuQuery
	err := c.ShouldBindQuery(&queryModel)
	if err != nil {
		global.SYS_LOG.Error("参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	list, total, err := adminMenuService.GetAdminRoleMenuList(queryModel)
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

// GetAdminRoleSysMenuJson
// @Tags      Admin
// @Summary   分页获取AdminRoleMenu列表
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Param     data  query     res.GetAdminRoleMenuList                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  "AdminRoleMenu列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/sysMenu[get]
func (s *AdminApi) GetAdminRoleSysMenuJson(c *gin.Context) {
	adminCalims := utils.GetAdminContext(c)
	list, err := adminMenuService.GetAdminRoleMenuJson(adminCalims.ID)
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

// CreateAdminRoleMenu
// @Tags      Admin
// @Summary   新增用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param    data  body      res.AdminRoleMenu  true  ""
// @Success  200   {object}  common.Response{data=res.AdminRoleMenu,msg=string}  ""
// @Router    /admin/role/{roleId}/menu [post]
func (b *AdminApi) CreateAdminRoleMenu(c *gin.Context) {
	var adminRoleMenu app.AdminRoleMenu
	err := c.ShouldBindBodyWith(&adminRoleMenu, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("新增菜单参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}

	if adminRoleMenuRes, err := adminMenuService.AddAdminRoleMenu(adminRoleMenu); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Error(err))
		common.InteralError(err.Error(), "创建失败", c)
	} else {
		global.SYS_LOG.Info("新增角色菜单 " + strconv.FormatInt(adminRoleMenuRes.ID, 10))
		common.OkWithDetailed(gin.H{"model": adminRoleMenuRes}, "创建成功", c)
	}
}

// UpdateAdminRoleMenu
// @Tags      Admin
// @Summary   更新系统用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param roleId path int true "role ID"
// @Param menuId path int true "menu ID"
// @Param    data  body      res.AdminRoleMenu  true  ""
// @Success  200   {object}  common.Response{data=res.AdminRoleMenu,msg=string}  ""
// @Router    /admin/roleId/{roleId}/menu/:menuId [put]
func (b *AdminApi) UpdateAdminRoleMenu(c *gin.Context) {
	var adminRoleMenu app.AdminRoleMenu
	err := c.ShouldBindBodyWith(&adminRoleMenu, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("更新系统菜单参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	roleId, _ := strconv.ParseInt(c.Param("roleId"), 10, 64)
	menuId, _ := strconv.ParseInt(c.Param("menuId"), 10, 64)
	adminRoleMenu.AdminRoleId = roleId
	adminRoleMenu.AdminMenuId = menuId
	adminRoleMenuRes, err := adminMenuService.UpdateAdminRoleMenu(adminRoleMenu)
	if err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Error(err))
		common.InteralError(err.Error(), "更新失败", c)
		return
	} else {
		global.SYS_LOG.Info("更新角色菜单成功 " + strconv.FormatInt(roleId, 10) + "-" + strconv.FormatInt(menuId, 10))
		common.OkWithDetailed(gin.H{"model": adminRoleMenuRes}, "更新成功", c)
	}
}

// RemoveAdminRoleMenu
// @Tags      Admin
// @Summary   删除AdminRoleMenu
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param roleId path int true "role ID"
// @Param menuId path int true "menu ID"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  " "
// @Router    /admin/role/{roleId}/menu/{menuId} [delete]
func (s *AdminApi) RemoveAdminRoleMenu(c *gin.Context) {
	roleId, _ := strconv.ParseInt(c.Param("roleId"), 10, 64)
	menuId, _ := strconv.ParseInt(c.Param("menuId"), 10, 64)

	err := adminMenuService.RemoveAdminRolMenu(menuId, roleId)
	if err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Error(err))
		common.InteralError(err.Error(), "删除失败", c)
		return
	} else {
		global.SYS_LOG.Info("删除角色菜单成功! " + strconv.FormatInt(roleId, 10) + "-" + strconv.FormatInt(menuId, 10))
		common.OkWithDetailed(common.QueryResult{
			Total: 1,
		}, "删除成功", c)
	}
}
