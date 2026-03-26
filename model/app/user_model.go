package app

import (
	"gin_rbac_api/model/common"
	"time"

	"github.com/jackc/pgtype"
)

type Login struct {
	Phone     string `json:"phone" form:"phone" `        // 用户名
	Password  string `json:"password" form:"password"`   // 密码
	Captcha   string `json:"captcha" form:"captcha"`     // 验证码
	CaptchaId string `json:"captchaId" form:"captchaId"` // 验证码ID
	Code      string `json:"code" form:"code"`           // 短信验证码ID
}

type UserInfo struct {
	common.EXTEND_MODEL
	Status      *int       `json:"status" `
	LastLoginAt *time.Time `json:"lastLoginAt"  form:"lastLoginAt"  gorm:"autoCreateTime;column:last_login_at"`
	LastLoginId int        `json:"lastLoginId"  form:"lastLoginId"  gorm:"column:last_login_id"`
	DateId      int        `json:"dateId" form:"dateId"  gorm:"column:date_id"`
	Phone       string     `json:"phone"  form:"phone"`
	Password    string     `json:"-"  form:"password"`
	Email       string     `json:"email"  form:"email"`
	Avatar      string     `json:"avatar" form:"avatar"`
	Name        string     `json:"name" form:"name"`
	// Gender      *int         `json:"gender" form:"gender" binding:"required"`
	Gender  *int         `json:"gender" form:"gender" `
	Birth   *pgtype.Date `json:"birth" form:"birth" gorm:"type:date"`
	WorkAt  *pgtype.Date `json:"workAt" form:"workAt" gorm:"column:work_at;type:date"`
	OpenId  string       `json:"openId"  form:"openId"  gorm:"column:open_id;default:NULL"`
	ReferId int64        `json:"referId"  form:"referId"  gorm:"column:refer_id"`
}

type UserQuery struct {
	common.EXTEND_SEARCH
	UserInfo
	BirthStart pgtype.Date `json:"birthStart" form:"birthStart"`
	BirthEnd   pgtype.Date `json:"birthEnd" form:"birthEnd"`
	LoginStart time.Time   `json:"loginStart" form:"loginStart"`
	LoginEnd   time.Time   `json:"loginEnd" form:"loginEnd"`
}

func (UserInfo) TableName() string {
	return "user_info"
}

type LoginResponse struct {
	User      UserInfo `json:"user"`
	BizId     int64    `json:"bizId"`
	BizName   string   `json:"bizName"`
	IsAdmin   bool     `json:"isAdmin"`
	Token     string   `json:"token"`
	ExpiresAt int64    `json:"expiresAt"`
}

type UserToken struct {
	Token string `json:"token"`
}

type UserAccount struct {
	Password string `json:"-"  form:"password"`
	Phone    string `json:"phone"  form:"phone"`
	Email    string `json:"email"  form:"email"`
	Avatar   string `json:"avatar" form:"avatar"`
}
