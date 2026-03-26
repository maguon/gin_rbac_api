package utils

import (
	"errors"
	"gin_rbac_api/global"
	"gin_rbac_api/model/common"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.SYS_CONFIG.JWT.SigningKey),
	}
}

func ParseDuration(d string) (time.Duration, error) {
	d = strings.TrimSpace(d)
	dr, err := time.ParseDuration(d)
	if err == nil {
		return dr, nil
	}
	if strings.Contains(d, "d") {
		index := strings.Index(d, "d")

		hour, _ := strconv.Atoi(d[:index])
		dr = time.Hour * 24 * time.Duration(hour)
		ndr, err := time.ParseDuration(d[index+1:])
		if err != nil {
			return dr, nil
		}
		return dr + ndr, nil
	}

	dv, err := strconv.ParseInt(d, 10, 64)
	return time.Duration(dv), err
}

func (j *JWT) CreateClaims(baseClaims common.BaseClaims) common.CustomClaims {
	bf, _ := ParseDuration(global.SYS_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.SYS_CONFIG.JWT.ExpiresTime)
	claims := common.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,     // 签名生效时间
			ExpiresAt: time.Now().Add(ep).Unix(),    // 过期时间 7天  配置文件
			Issuer:    global.SYS_CONFIG.JWT.Issuer, // 签名的发行者
		},
	}
	return claims
}

func (j *JWT) CreateAdminClaims(adminBaseClaims common.AdminBaseClaims) common.AdminCustomClaims {
	bf, _ := ParseDuration(global.SYS_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.SYS_CONFIG.JWT.ExpiresTime)
	claims := common.AdminCustomClaims{
		AdminBaseClaims: adminBaseClaims,
		BufferTime:      int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,     // 签名生效时间
			ExpiresAt: time.Now().Add(ep).Unix(),    // 过期时间 7天  配置文件
			Issuer:    global.SYS_CONFIG.JWT.Issuer, // 签名的发行者
		},
	}
	return claims
}

// 创建一个token
func (j *JWT) CreateToken(claims common.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 创建一个token
func (j *JWT) CreateAdminToken(claims common.AdminCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func (j *JWT) CreateEmailActiveUrl(claims common.UserEmailCode) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
func (j *JWT) CreateEmailActiveToken(claims common.UserEmailCode) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
func (j *JWT) ParseEmailActiveToken(tokenString string) (*common.UserEmailCode, error) {
	token, err := jwt.ParseWithClaims(tokenString, &common.UserEmailCode{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*common.UserEmailCode); ok && token.Valid {
			return claims, nil
		} else {
			return nil, nil
		}
	} else {
		return nil, err
	}
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*common.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &common.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*common.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

// 解析 token
func (j *JWT) ParseAdminToken(tokenString string) (*common.AdminCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &common.AdminCustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*common.AdminCustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

func (j *JWT) RefreshToken(oldToken string, claims common.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}
