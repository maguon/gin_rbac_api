package middleware

import (
	"gin_rbac_api/model/common"
	redis_util "gin_rbac_api/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var store = redis_util.NewCaptchaRedisStore()
var smsStore = redis_util.NewSmsCodeRedisStore()

func CheckCaptcha() gin.HandlerFunc {
	return func(c *gin.Context) {
		var l common.ReqCaptcha
		err := c.ShouldBindBodyWith(&l, binding.JSON)

		if err != nil {
			common.FailWithDetailed(gin.H{"reload": true}, "验证码错误", c)
			c.Abort()
			return
		}

		store.UseWithCtx(c.Request.Context())

		if store.Verify(l.CaptchaId, l.Captcha, true) {
			c.Next()
		} else {
			common.FailWithDetailed(gin.H{"reload": true}, "验证码错误", c)
			c.Abort()
			return
		}
	}
}

func CheckSmsCode() gin.HandlerFunc {
	return func(c *gin.Context) {
		var l common.SmsCode
		err := c.ShouldBindBodyWith(&l, binding.JSON)
		if err != nil {
			common.FailWithDetailed(gin.H{"reload": true}, "验证码错误", c)
			c.Abort()
			return
		}
		smsStore.UseWithCtx(c.Request.Context())

		if smsStore.VerifyPhoneCode(l.Phone, l.Code) {
			c.Next()
		} else {
			common.FailWithDetailed(gin.H{"reload": true}, "验证码错误", c)
			c.Abort()
			return
		}
	}
}
