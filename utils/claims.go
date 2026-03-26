package utils

import (
	"fmt"
	"gin_rbac_api/global"
	"gin_rbac_api/model/common"

	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) (*common.CustomClaims, error) {
	token := c.Request.Header.Get("auth-token")
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.SYS_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

func GetAdminClaims(c *gin.Context) (*common.AdminBaseClaims, error) {
	token := c.Request.Header.Get("auth-token")
	fmt.Println(token)
	j := NewJWT()
	claims, err := j.ParseAdminToken(token)
	if err != nil {
		global.SYS_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return &claims.AdminBaseClaims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserContext(c *gin.Context) common.BaseClaims {
	if claims, exists := c.Get("claims"); !exists {
		var baseClaims common.BaseClaims
		return baseClaims
	} else {
		waitUse := claims.(common.BaseClaims)
		return waitUse
	}
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetAdminContext(c *gin.Context) common.AdminBaseClaims {
	if claims, exists := c.Get("claims"); !exists {
		var adminBaseClaims common.AdminBaseClaims
		return adminBaseClaims
	} else {
		waitUse := claims.(common.AdminBaseClaims)
		return waitUse
	}
}
