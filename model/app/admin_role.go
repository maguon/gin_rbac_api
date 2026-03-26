package app

import (
	"gin_rbac_api/model/common"
	"time"

	"github.com/jackc/pgtype"
)

type AdminRoleQuery struct {
	common.EXTEND_SEARCH
	AdminRole
}

type AdminRole struct {
	common.EXTEND_MODEL
	Status     int    `json:"status" form:"status" gorm:"comment:状态"`
	RoleName   string `json:"roleName" form:"roleName" gorm:"column:role_name;comment:角色名称"`
	DefaultUrl string `json:"defaultUrl" form:"defaultUrl" gorm:"column:default_url;comment:角色名称"`
	Remark     string `json:"remark" form:"remark" gorm:"comment:备注"`
}

func (AdminRole) TableName() string {
	return "admin_role"
}

type AdminInfoBase struct {
	common.EXTEND_MODEL
	LastLoginAt time.Time `json:"lastLoginAt"  form:"-"  gorm:"autoCreateTime;column:last_login_at"`
	Username    string    `json:"username" form:"username"  gorm:"column:username;comment:用户登录名"` // 用户登录名
	Gender      int8      `json:"gender"  form:"gender" gorm:"comment:管理员性别"`
	Avatar      string    `json:"avatar" form:"avatar" gorm:"comment:用户头像"`
	Phone       string    `json:"phone" form:"phone" gorm:"comment:用户手机号"` // 用户手机号
	Email       string    `json:"email" form:"email" gorm:"comment:用户邮箱"`
	Status      *int8     `json:"status" form:"status" gorm:"comment:状态"`
}
type AdminInfo struct {
	common.EXTEND_MODEL
	LastLoginAt time.Time `json:"lastLoginAt"  form:"-"  gorm:"autoCreateTime;column:last_login_at"`
	Username    string    `json:"username" form:"username"  gorm:"column:username;comment:用户登录名"` // 用户登录名
	Gender      int8      `json:"gender"  form:"gender" gorm:"comment:管理员性别"`
	Avatar      string    `json:"avatar" form:"avatar" gorm:"comment:用户头像"`
	Phone       string    `json:"phone" form:"phone" gorm:"comment:用户手机号"` // 用户手机号
	Email       string    `json:"email" form:"email" gorm:"comment:用户邮箱"`
	Status      *int8     `json:"status" form:"status" gorm:"comment:状态"`
	Password    string    `json:"password" form:"password" gorm:"comment:用户登录密码"`
}

func (AdminInfo) TableName() string {
	return "admin_info"
}

type AdminRoleInfo struct {
	AdminInfo
	RoleId     int64  `json:"roleId" form:"roleId" gorm:"column:role_id"`
	RoleName   string `json:"roleName" form:"roleName" gorm:"column:role_name"`
	DefaultUrl string `json:"defaultUrl" form:"defaultUrl" gorm:"column:default_url"`
	Remark     string `json:"remark" form:"remark" gorm:"column:remark"`
}
type AdminToken struct {
	ID    int64  `json:"id" form:"id" gorm:"column:id"`
	Token string `json:"token" form:"token" gorm:"column:token"`
}

func (AdminToken) TableName() string {
	return "admin_token"
}

type AdminLoginResponse struct {
	AdminRoleInfo AdminRoleInfo `json:"adminRoleInfo"`
	Token         string        `json:"token"`
	ExpiresAt     int64         `json:"expiresAt"`
}

type AdminUserQuery struct {
	common.EXTEND_SEARCH
	AdminInfo
	AdminRoleId int64 `json:"adminRoleId"  form:"adminRoleId" gorm:"column:admin_role_id;comment:管理员角色ID"`
}

type AdminRoleRel struct {
	common.EXTEND_MODEL
	AdminId     int64 `json:"adminId"  form:"adminId" gorm:"column:admin_id;comment:管理员ID"`
	AdminRoleId int64 `json:"adminRoleId"  form:"adminRoleId" gorm:"column:admin_role_id;comment:管理员角色ID"`
}

func (AdminRoleRel) TableName() string {
	return "admin_role_rel"
}

type AdminRoleRelQuery struct {
	common.EXTEND_SEARCH
	AdminRoleRel
}

type AdminRoleMenu struct {
	common.EXTEND_MODEL
	AdminMenuId int64         `json:"adminMenuId"  form:"adminMenuId" gorm:"column:admin_menu_id;comment:菜单ID"`
	AdminRoleId int64         `json:"adminRoleId"  form:"adminRoleId" gorm:"column:admin_role_id;comment:管理员角色ID"`
	MenuBtn     *pgtype.JSONB `json:"menuBtn" form:"menuBtn" gorm:"column:menu_btn;comment:菜单项的按钮列表"`
}

func (AdminRoleMenu) TableName() string {
	return "admin_role_menu"
}

type AdminRoleMenuQuery struct {
	common.EXTEND_SEARCH
	AdminRoleMenu
}
