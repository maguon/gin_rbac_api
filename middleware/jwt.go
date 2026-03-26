package middleware

import (
	"encoding/json"
	"gin_rbac_api/model/common"
	"gin_rbac_api/service"

	"github.com/gin-gonic/gin"
)

var jwtService = service.ServiceGroupApp.JwtService

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("auth-token")

		if token == "" {
			common.AuthError(gin.H{"reload": true}, "未登录或非法访问", c)
			//response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		claimStr, err := jwtService.GetRedisAuth(token)
		if err != nil {
			common.AuthError(gin.H{"reload": true}, "身份验证失败", c)
			//response.FailWithDetailed(gin.H{"reload": true}, "身份验证失败", c)
			c.Abort()
			return
		}
		var baseCalims common.BaseClaims
		json.Unmarshal([]byte(claimStr), &baseCalims)

		c.Set("claims", baseCalims)
		c.Next()

		/* j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			dr, _ := utils.ParseDuration(global.SYS_CONFIG.JWT.ExpiresTime)
			claims.ExpiresAt = time.Now().Add(dr).Unix()
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if global.SYS_CONFIG.System.UseMultipoint {
				if err != nil {
					global.SYS_LOG.Error("get redis jwt failed", zap.Error(err))
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Phone)
			}
		}
		c.Set("claims", claims)
		c.Next() */
	}
}

func JWTAdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("auth-token")
		if token == "" {
			common.AuthError(gin.H{"reload": true}, "未登录或非法访问", c)
			//response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		claimStr, err := jwtService.GetRedisAdminAuth(token)
		if err != nil {
			common.AuthError(gin.H{"reload": true}, "身份验证失败", c)
			//response.FailWithDetailed(gin.H{"reload": true}, "身份验证失败", c)
			c.Abort()
			return
		}
		var adminCalims common.AdminBaseClaims
		json.Unmarshal([]byte(claimStr), &adminCalims)

		c.Set("claims", adminCalims)
		c.Next()
	}
}
