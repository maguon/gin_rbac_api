package app

import (
	"gin_rbac_api/model/common"

	"github.com/jackc/pgtype"
)

// Modify password structure
type SysUserPassword struct {
	ID          int64  `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

type UserAvatar struct {
	ID     int64  `json:"-"`
	Avatar string `json:"avatar"`
}
type AdminLogin struct {
	Phone     string `json:"Phone"`     // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type AdminSysApi struct {
	common.EXTEND_MODEL
	Status int    `json:"status" form:"status" gorm:"column:status"`
	Method string `json:"method" form:"method" gorm:"column:method"`
	Url    string `json:"url" form:"url" gorm:"comment:column:url"`
	Tag    string `json:"tag" form:"tag" gorm:"comment:column:tag"`
	Remark string `json:"remark" form:"remark" gorm:"column:remark;comment:备注"`
}

func (AdminSysApi) TableName() string {
	return "admin_sys_api"
}

type AdminSysApiQuery struct {
	common.EXTEND_SEARCH
	AdminSysApi
}

type AdminSysApiTagGroup struct {
	Tag     string        `json:"tag" form:"tag" gorm:"comment:column:tag"`
	ApiList *pgtype.JSONB `json:"apiList" form:"apiList" gorm:"column:api_list;comment:api表"`
}

type AdminMenu struct {
	common.EXTEND_MODEL
	Status    int    `json:"status" form:"status" gorm:"column:status"`
	Level     int    `json:"level" form:"level" gorm:"column:level"`
	ParentId  int64  `json:"parentId" form:"parentId" gorm:"column:parent_id"`
	Sort      int    `json:"sort" form:"sort" gorm:"column:sort"`
	MenuName  string `json:"menuName" form:"menuName" gorm:"column:menu_name"`
	Url       string `json:"url" form:"url" gorm:"comment:column:url"`
	Icon      string `json:"icon" form:"icon" gorm:"comment:column:icon"`
	Component string `json:"component" form:"component" gorm:"comment:column:component"`
}

func (AdminMenu) TableName() string {
	return "admin_menu"
}

type AdminMenuQuery struct {
	common.EXTEND_SEARCH
	AdminMenu
}
type AdminMenuBtn struct {
	common.EXTEND_MODEL
	Status      *int   `json:"status" form:"status" gorm:"column:status"`
	AdminMenuId int64  `json:"adminMenuId" form:"adminMenuId" gorm:"column:admin_menu_id"`
	BtnName     string `json:"btnName" form:"btnName" gorm:"column:btn_name"`
	Remark      string `json:"remark" form:"remark" gorm:"comment:column:remark"`
}

func (AdminMenuBtn) TableName() string {
	return "admin_menu_btn"
}

type AdminMenuBtnQuery struct {
	common.EXTEND_SEARCH
	AdminMenuBtn
}

type AdminMenuJson struct {
	common.EXTEND_MODEL
	Status    int           `json:"status" form:"status" gorm:"column:status"`
	ParentId  int64         `json:"parent_id" form:"parent_id" gorm:"column:parent_id"`
	Sort      int           `json:"sort" form:"sort" gorm:"column:sort"`
	MenuName  string        `json:"menu_name" form:"menu_name" gorm:"column:menu_name"`
	Url       string        `json:"url" form:"url" gorm:"comment:column:url"`
	Icon      string        `json:"icon" form:"icon" gorm:"comment:column:icon"`
	Component string        `json:"component" form:"component" gorm:"comment:column:component"`
	Children  *pgtype.JSONB `json:"children" form:"children" gorm:"column:children;type:jsonb;comment:菜单项的二级选项"`
	MenuBtn   *pgtype.JSONB `json:"menu_btn" form:"menu_btn" gorm:"type:jsonb;column:menu_btn;comment:菜单项的按钮列表"`
}
