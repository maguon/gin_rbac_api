package public

import (
	"encoding/json"
	"gin_rbac_api/global"
	"gin_rbac_api/utils"
	"strconv"

	app "gin_rbac_api/model/app"
	common "gin_rbac_api/model/common"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

// Login
// @Tags     Admin
// @Summary  用户登录
// @Produce   application/json
// @Param    data  body      req.AdminLogin                                             true  "用户名, 密码, 验证码"
// @Success  200   {object}  common.Response{data=res.AdminLoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /public/adminLogin [post]
func (b *PublicApi) AdminLogin(c *gin.Context) {
	var l app.AdminLogin
	err := c.ShouldBindBodyWith(&l, binding.JSON)
	if err != nil {
		global.SYS_LOG.Error("系统用户登录参数错误! ", zap.Error(err))
		common.RequestError(err.Error(), "参数错误", c)
		return
	}

	u := &app.AdminInfo{Phone: l.Phone, Password: l.Password}
	user, err := adminService.AdminLogin(u)
	if err != nil {
		global.SYS_LOG.Error("登陆失败! 用户名或密码错误!", zap.Error(err))
		common.InteralError(err.Error(), "用户名或密码错误", c)
		return
	} else {
		if *user.AdminInfo.Status == utils.USER_STATUS_FORBIDDEN {
			global.SYS_LOG.Warn("系统用户已被禁用! " + strconv.FormatInt(user.AdminInfo.ID, 10))
			common.FailWithMessage("系统用户已被禁用", c)
			return
		} else {
			user.AdminInfo.Password = "" //清除用户密码
			global.SYS_LOG.Info("系统用户登录成功! " + strconv.FormatInt(user.AdminInfo.ID, 10))
			b.AdminTokenNext(c, *user)
		}
	}
}

// TokenNext 登录以后签发jwt
func (b *PublicApi) AdminTokenNext(c *gin.Context, user app.AdminRoleInfo) {
	j := &utils.JWT{SigningKey: []byte(global.SYS_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateAdminClaims(common.AdminBaseClaims{
		ID:       user.AdminInfo.ID,
		Username: user.AdminInfo.Username,
		Phone:    user.AdminInfo.Phone,
		RoleId:   user.RoleId,
		RoleName: user.RoleName,
		IsSuper:  false,
	})
	token, err := j.CreateAdminToken(claims)

	if err != nil {
		global.SYS_LOG.Error("获取token失败!", zap.Error(err))
		common.InteralError(err.Error(), "获取身份失败", c)
		return
	}
	claimByte, _ := json.Marshal(claims)
	claimStr := string(claimByte)
	if err := jwtService.SetRedisAdminAuth(claimStr, token); err != nil {
		global.SYS_LOG.Error("设置登录状态失败!", zap.Error(err))
		common.InteralError(err.Error(), "设置登录状态失败", c)
		return
	} else {
		adminService.UpdateAdminLoginAt(user.AdminInfo.ID)
		common.OkWithDetailed(app.AdminLoginResponse{
			AdminRoleInfo: user,
			Token:         token,
			ExpiresAt:     claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}
