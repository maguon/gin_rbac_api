package common

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type AdminCustomClaims struct {
	AdminBaseClaims
	BufferTime int64
	jwt.StandardClaims
}
type BaseClaims struct {
	ID       int64
	UserName string
	BizId    int64
	BizName  string
	IsAdmin  bool
	Phone    string
}

type AdminBaseClaims struct {
	ID       int64
	Username string
	Phone    string
	RoleId   int64
	RoleName string
	IsSuper  bool
}

type BizUserClaims struct {
	ID      int64
	BizId   int64
	BizName string
	IsAdmin bool
	Phone   string
}

type ReqCaptcha struct {
	Captcha   string `json:"captcha" form:"captcha" `    // 验证码
	CaptchaId string `json:"captchaId" form:"captchaId"` // 验证码ID
}

type SmsCode struct {
	Phone string `json:"phone" form:"phone" ` // 手机号
	Code  string `json:"code" form:"code"`    // 验证码
}

type UserEmailCode struct {
	ID    int64
	Email string
	jwt.StandardClaims
}
