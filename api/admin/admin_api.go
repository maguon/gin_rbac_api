package admin

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
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

// GetAdminInfo
// @Tags      Admin
// @Summary   获取用户信息
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  common.Response{data=map[string]interface{},msg=string}  "获取用户信息"
// @Router    /admin/sysUser [get]
func (b *AdminApi) GetAdminUserInfo(c *gin.Context) {
	adminCalims := utils.GetAdminContext(c)

	var adminUserQuery app.AdminUserQuery
	adminUserQuery.PageSize = 10
	adminUserQuery.PageNumber = 1
	adminUserQuery.ID = adminCalims.ID
	adminUserQuery.Status = utils.GetInt8Pointer(utils.USER_STATUS_NORMAL)
	adminRole, total, err := adminService.GetSysUserList(adminUserQuery)
	adminMenu, err := adminMenuService.GetAdminRoleMenuJson(adminCalims.ID)
	if err != nil {
		global.SYS_LOG.Error("获取当前用户信息失败!", zap.Error(err))
		common.FailWithMessage("获取当前用户信息失败", c)
		return
	} else {
		if total == 0 {
			global.SYS_LOG.Warn("获取当前用户不存在或被禁用! " + strconv.FormatInt(adminCalims.ID, 10))
			common.FailWithMessage("获取当前用户不存在或被禁用", c)
			return
		} else {
			global.SYS_LOG.Info("获取当前用户信息成功! " + strconv.FormatInt(adminCalims.ID, 10))
			common.OkWithDetailed(gin.H{"adminInfo": adminRole, "adminMenu": adminMenu}, "获取成功", c)
		}

	}
}

// GetAdminUserList
// @Tags      Admin
// @Summary   分页获取SysUser列表
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Param     data  query     res.AdminUserQuery                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  "SysUser列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/sysUserList [get]
func (s *AdminApi) GetAdminUserList(c *gin.Context) {
	var queryModel app.AdminUserQuery

	/* if claims, exists := c.Get("claims"); !exists {
		global.SYS_LOG.Error("获取用户上下文失败!", zap.Error())
		//fmt.Println(exists)
	}  */

	err := c.ShouldBindQuery(&queryModel)
	if err != nil {
		global.SYS_LOG.Error("获取系统用户参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	list, total, err := adminService.GetSysUserList(queryModel)
	if err != nil {

		common.InteralError(err.Error(), "获取失败", c)
		return
	} else {
		global.SYS_LOG.Info("获取系统用户成功 ")
		common.OkWithDetailed(common.QueryResult{
			List:       list,
			Total:      total,
			PageNumber: queryModel.PageNumber,
			PageSize:   queryModel.PageSize,
		}, "获取成功", c)
	}

}

// UpdateAdminPassword
// @Tags      Admin
// @Summary   更改系统用户密码
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param    data  body      req.SysUserPassword  true  " 原密码, 新密码"
// @Success  200   {object}  common.Response{msg=string}  "返回更新结果信息"
// @Router    /admin/password [put]
func (b *AdminApi) UpdateAdminPassword(c *gin.Context) {
	var sysUserPassword app.SysUserPassword
	err := c.ShouldBindBodyWith(&sysUserPassword, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("更新密码参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	adminCalims := utils.GetAdminContext(c)
	sysUserPassword.ID = adminCalims.ID
	_, err = adminService.UpdateAdminPassword(sysUserPassword)
	if err != nil {
		global.SYS_LOG.Error("更新密码失败!", zap.Error(err))
		common.InteralError(err.Error(), "更新密码失败", c)
		return
	} else {
		global.SYS_LOG.Info("更新密码成功 " + strconv.FormatInt(adminCalims.ID, 10))
		common.OkWithMessage("修改成功", c)
	}
}

// UpdateAdminAvatar
// @Tags      Admin
// @Summary   更换系统用户头衔
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param    data  body      req.UserAvatar  true  " avatar url "
// @Success  200   {object}  common.Response{msg=string}  "返回更新结果信息"
// @Router    /admin/avatar [put]
func (b *AdminApi) UpdateAdminAvatar(c *gin.Context) {
	var model app.UserAvatar
	err := c.ShouldBindBodyWith(&model, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("更新头像错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	adminCalims := utils.GetAdminContext(c)
	model.ID = adminCalims.ID
	err = adminService.UpdateAdminAvatar(model.ID, model.Avatar)
	if err != nil {
		global.SYS_LOG.Error("更新头像失败!", zap.Error(err))
		common.InteralError(err.Error(), "更新头像失败", c)
		return
	} else {
		global.SYS_LOG.Info("更新头像成功 " + strconv.FormatInt(adminCalims.ID, 10))
		common.OkWithMessage("修改成功", c)
	}
}

// CreateAdminInfo
// @Tags      Admin
// @Summary   新增系统用户
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param    data  body      res.AdminRoleInfo  true  "用户信息"
// @Success  200   {object}  common.Response{data=res.AdminInfo,msg=string}  "返回包括用户组信息"
// @Router    /admin/sysUser [post]
func (b *AdminApi) CreateAdminInfo(c *gin.Context) {
	var adminRoleInfo app.AdminRoleInfo
	err := c.ShouldBindBodyWith(&adminRoleInfo, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("创建系统用户参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}

	fmt.Println(adminRoleInfo.AdminInfo.Password)
	adminRoleInfo.AdminInfo.Password = utils.BcryptHash(adminRoleInfo.AdminInfo.Password)
	fmt.Println(adminRoleInfo.AdminInfo.Password)
	if adminInfoRes, err := adminService.AddAdminInfo(adminRoleInfo.AdminInfo); err != nil {
		global.SYS_LOG.Error("创建系统用户失败!", zap.Error(err))
		common.InteralError(err.Error(), "创建失败", c)
	} else {
		global.SYS_LOG.Info("创建系统用户成功 " + strconv.FormatInt(adminInfoRes.ID, 10))
		var adminRoleRel app.AdminRoleRel
		adminRoleRel.AdminId = adminInfoRes.ID
		adminRoleRel.AdminRoleId = adminRoleInfo.RoleId

		adminRoleRes, err := adminService.AddAdminRoleRel(adminRoleRel)
		if err != nil {
			global.SYS_LOG.Error("创建用户角色关联!", zap.Error(err))
			common.InteralError(err.Error(), "创建失败", c)
			return
		} else {
			global.SYS_LOG.Info("创建系统用户成功 " + strconv.FormatInt(adminRoleRes.AdminRoleId, 10))
			common.OkWithDetailed(gin.H{"model": adminRoleInfo}, "创建成功", c)
		}
	}
}

// UpdateAdminInfo
// @Tags      Admin
// @Summary   更新系统用户信息
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param typeId path int true "sys user ID"
// @Param    data  body      res.AdminInfo  true  "用户名, 密码, 验证码"
// @Success  200   {object}  common.Response{data=res.AdminInfo,msg=string}  "返回包括用户组信息"
// @Router    /admin/sysUser/{sysUserId} [put]
func (b *AdminApi) UpdateAdmiInfo(c *gin.Context) {
	var adminRoleInfo app.AdminRoleInfo
	err := c.ShouldBindBodyWith(&adminRoleInfo, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("更新系统用户参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}

	sysUserId, _ := strconv.ParseInt(c.Param("sysUserId"), 10, 64)
	adminRoleInfo.AdminInfo.ID = sysUserId

	_, err = adminService.UpdateAdminInfo(adminRoleInfo.AdminInfo)
	if err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Error(err))
		common.InteralError(err.Error(), "更新失败", c)
		return
	} else {
		global.SYS_LOG.Info("更新系统用户成功 " + strconv.FormatInt(sysUserId, 10))
		var adminRoleRel app.AdminRoleRel
		adminRoleRel.AdminId = sysUserId
		adminRoleRel.AdminRoleId = adminRoleInfo.RoleId
		adminRoleRes, err := adminService.UpdateAdminUserRoleRel(adminRoleRel)
		if err != nil {
			global.SYS_LOG.Error("更新失败!", zap.Error(err))
			common.InteralError(err.Error(), "更新失败", c)
			return
		} else {
			global.SYS_LOG.Info("更新系统用户角色成功 " + strconv.FormatInt(adminRoleRes.AdminRoleId, 10))
			common.OkWithDetailed(gin.H{"model": adminRoleInfo}, "更新成功", c)
		}
	}

}

// RemoveAdminInfo
// @Tags      Admin
// @Summary   删除AdminInfo
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param adminId path int true "admin ID"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  " "
// @Router    /admin/sysUser/{sysUserId} [delete]
func (s *AdminApi) RemoveAdminInfo(c *gin.Context) {
	sysUserId, _ := strconv.ParseInt(c.Param("sysUserId"), 10, 64)

	err := adminService.RemoveAdminInfo(sysUserId)
	if err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Error(err))
		common.InteralError(err.Error(), "删除失败", c)
		return
	} else {
		global.SYS_LOG.Info("删除系统用户组成功! " + strconv.FormatInt(sysUserId, 10))
		common.OkWithDetailed(gin.H{"model": sysUserId}, "删除成功", c)
	}
}

// CreateAdminRole
// @Tags      Admin
// @Summary   新增用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param    data  body      res.AdminRole  true  "用户名, 密码, 验证码"
// @Success  200   {object}  common.Response{data=res.LoginResponse,msg=string}  "返回包括用户组信息"
// @Router    /admin/role [post]
func (b *AdminApi) CreateAdminRole(c *gin.Context) {
	var adminRole app.AdminRole
	err := c.ShouldBindBodyWith(&adminRole, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("新增用户组参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}

	if adminRoleRes, err := adminService.AddAdminRole(adminRole); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Error(err))
		common.InteralError(err.Error(), "创建失败", c)
	} else {
		global.SYS_LOG.Info("新增用户组成功 " + strconv.FormatInt(adminRole.ID, 10))
		common.OkWithDetailed(gin.H{"model": adminRoleRes}, "创建成功", c)
	}
}

// UpdateAdminRole
// @Tags      Admin
// @Summary   更新系统用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param roleId path int true "role ID"
// @Param    data  body      res.AdminRole  true  "用户名, 密码, 验证码"
// @Success  200   {object}  common.Response{data=res.LoginResponse,msg=string}  "返回包括用户组信息"
// @Router    /admin/role/{roleId} [put]
func (b *AdminApi) UpdateAdminRole(c *gin.Context) {
	var adminRole app.AdminRole
	err := c.ShouldBindBodyWith(&adminRole, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("更新系统用户组参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	roleId, _ := strconv.ParseInt(c.Param("roleId"), 10, 64)
	adminRole.ID = roleId
	adminRoleRes, err := adminService.UpdateAdminRole(adminRole)
	if err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Error(err))
		common.InteralError(err.Error(), "更新失败", c)
		return
	} else {
		global.SYS_LOG.Info("更新系统用户组成功 " + strconv.FormatInt(roleId, 10))
		common.OkWithDetailed(gin.H{"model": adminRoleRes}, "更新成功", c)
	}
}

// GetAdminRole
// @Tags      Admin
// @Summary   分页获取AdminRole列表
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Param     data  query     res.AdminRoleQuery                       true  "页码, 每页大小, 搜索条件"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  "AdminRole列表,返回包括列表,总数,页码,每页数量"
// @Router    /admin/role [get]
func (s *AdminApi) GetAdminRole(c *gin.Context) {
	var queryModel app.AdminRoleQuery
	err := c.ShouldBindQuery(&queryModel)
	if err != nil {
		global.SYS_LOG.Error("参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	list, total, err := adminService.GetAdminRole(queryModel)
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

// RemoveAdminRole
// @Tags      Admin
// @Summary   删除AdminRole
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param roleId path int true "role ID"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  " "
// @Router    /admin/role/{roleId} [delete]
func (s *AdminApi) RemoveAdminRole(c *gin.Context) {
	roleId, _ := strconv.ParseInt(c.Param("roleId"), 10, 64)

	total, err := adminService.RemoveAdminRole(roleId)
	if err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Error(err))
		common.InteralError(err.Error(), "删除失败", c)
		return
	} else {
		if total > 0 {
			global.SYS_LOG.Info("删除用户组失败! " + strconv.FormatInt(roleId, 10))
			common.FailWithMessage("获取当前角色下存在用户", c)
		} else {
			global.SYS_LOG.Info("删除用户组成功! " + strconv.FormatInt(roleId, 10))
			common.OkWithDetailed(common.QueryResult{
				Total: total,
			}, "删除成功", c)
		}

	}
}

// CreateAdminRoleRel
// @Tags      Admin
// @Summary   新增用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param    data  body      res.AdminRoleRel  true  "用户名, 密码, 验证码"
// @Success  200   {object}  common.Response{data=res.LoginResponse,msg=string}  "返回包括用户组信息"
// @Router    /admin/roleRel [post]
func (b *AdminApi) CreateAdminRoleRel(c *gin.Context) {
	var adminRoleRel app.AdminRoleRel
	err := c.ShouldBindBodyWith(&adminRoleRel, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("新增用户组参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}

	if adminRoleRes, err := adminService.AddAdminRoleRel(adminRoleRel); err != nil {
		global.SYS_LOG.Error("创建失败!", zap.Error(err))
		common.InteralError(err.Error(), "创建失败", c)
	} else {
		global.SYS_LOG.Info("新增用户组成功 " + strconv.FormatInt(adminRoleRel.ID, 10))
		common.OkWithDetailed(gin.H{"model": adminRoleRes}, "创建成功", c)
	}
}

// UpdateAdminRoleRel
// @Tags      Admin
// @Summary   更新系统用户组
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param roleRelId path int true "roleRel ID"
// @Param    data  body      res.AdminRoleRel  true  "用户名, 密码, 验证码"
// @Success  200   {object}  common.Response{data=res.LoginResponse,msg=string}  "返回包括用户组信息"
// @Router    /admin/roleRel/{roleRelId} [put]
func (b *AdminApi) UpdateAdminRoleRel(c *gin.Context) {
	var adminRoleRel app.AdminRoleRel
	err := c.ShouldBindBodyWith(&adminRoleRel, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("更新系统用户角色参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	roleRelId, _ := strconv.ParseInt(c.Param("roleRelId"), 10, 64)
	adminRoleRel.ID = roleRelId
	adminRoleRelRes, err := adminService.UpdateAdminRoleRel(adminRoleRel)
	if err != nil {
		global.SYS_LOG.Error("更新失败!", zap.Error(err))
		common.InteralError(err.Error(), "更新失败", c)
		return
	} else {
		global.SYS_LOG.Info("更新系统用户角色成功 " + strconv.FormatInt(roleRelId, 10))
		common.OkWithDetailed(gin.H{"model": adminRoleRelRes}, "更新成功", c)
	}
}

// RemoveAdminRoleRel
// @Tags      Admin
// @Summary   删除AdminRoleRel
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Param roleId path int true "admin role rel ID"
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  " "
// @Router    /admin/roleRel[delete]
func (s *AdminApi) RemoveAdminRoleRel(c *gin.Context) {
	var adminRoleRel app.AdminRoleRel
	err := c.ShouldBindBodyWith(&adminRoleRel, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("删除角色API参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	err = adminService.RemoveAdminRoleRel(adminRoleRel.AdminId, adminRoleRel.AdminRoleId)
	if err != nil {
		global.SYS_LOG.Error("删除失败!", zap.Error(err))
		common.InteralError(err.Error(), "删除失败", c)
		return
	} else {
		global.SYS_LOG.Info("删除用户组成功! " + strconv.FormatInt(adminRoleRel.AdminId, 10))
		common.OkWithDetailed(common.QueryResult{
			Total: 1,
		}, "删除成功", c)
	}
}

// ExportAdminInfoCsv
// @Tags      Admin
// @Summary   导出Sys AdminInfo CSV
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @Produce   application/json
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  " "
// @Router    /admin/sysUser.csv [get]
func (s *AdminApi) ExportAdminInfo(c *gin.Context) {
	var queryModel app.AdminUserQuery
	err := c.ShouldBindQuery(&queryModel)
	if err != nil {
		global.SYS_LOG.Error("参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}
	list, _, err := adminService.GetSysUserList(queryModel)
	if err != nil {
		global.SYS_LOG.Error("获取失败!", zap.Error(err))
		common.InteralError(err.Error(), "获取失败", c)
		return
	}
	var byteList []byte
	byteList, _ = json.Marshal(list)
	var adminInfo []app.AdminRoleInfo
	json.Unmarshal(byteList, &adminInfo)
	record := []string{"ID", "用户名称", "用户群组", "手机", "邮箱", "性别", "创建时间", "状态"} // just some test data to use for the wr.Writer() method below.
	fmt.Println(len(adminInfo))
	b := &bytes.Buffer{}   // creates IO Writer
	wr := csv.NewWriter(b) // creates a csv writer that uses the io buffer.
	wr.Write(record)
	for i := 0; i < len(adminInfo); i++ { // make a loop for 100 rows just for testing purposes
		recordTemp := []string{strconv.FormatInt(adminInfo[i].AdminInfo.ID, 10), adminInfo[i].AdminInfo.Username, adminInfo[i].AdminInfo.Username, adminInfo[i].AdminInfo.Phone,
			adminInfo[i].AdminInfo.Email, utils.GetGenderStr(adminInfo[i].AdminInfo.Gender), adminInfo[i].AdminInfo.CreatedAt.Format("2006-01-02 15:04:05"), utils.GetStatusStr(*adminInfo[i].AdminInfo.Status)}
		wr.Write(recordTemp) // converts array of string to comma seperated values for 1 row.
	}
	wr.Flush() // writes the csv writer data to  the buffered data io writer(b(bytes.buffer))
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Write(b.Bytes())
	global.SYS_LOG.Info("获取成功! ")
}

// AdminLogout
// @Tags      Admin
// @Summary   管理员退出登录
// @Security  ApiKeyAuth
// @param auth-token header string true "auth-token"
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  common.Response{data=common.QueryResult,msg=string}  " "
// @Router    /admin/logout [post]
func (s *AdminApi) AdminLogout(c *gin.Context) {
	token := c.Request.Header.Get("auth-token")
	if token == "" {
		global.SYS_LOG.Error("未登录或非法访问! ")
		common.AuthError(gin.H{"reload": true}, "未登录或非法访问", c)
		common.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
		return
	}
	if err := jwtService.DelRedisAdminAuth(token); err != nil {
		global.SYS_LOG.Error("退出登录状态失败!", zap.Error(err))
		common.InteralError(err.Error(), "设置登录状态失败", c)
		return
	} else {
		global.SYS_LOG.Error("退出登录状态成功!")
		common.OkWithMessage("退出登录成功", c)
	}
}
