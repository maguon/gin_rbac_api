package app

type AdminCasbin struct {
	ID    int64  `json:"id" form:"id" gorm:"primarykey"`
	Ptype string `json:"ptype" form:"ptype" gorm:"column:ptype;comment:策略类型"`
	V0    string `json:"v0" form:"v0" gorm:"column:v0;comment:角色ID"`
	V1    string `json:"v1" form:"v1" gorm:"column:v1;comment:资源路径"`
	V2    string `json:"v2" form:"v2" gorm:"column:v2;comment:请求方法"`
	V3    string `json:"v3" form:"v3" gorm:"column:v3;"`
	V4    string `json:"v4" form:"v4" gorm:"column:v4;"`
	V5    string `json:"v5" form:"v5" gorm:"column:v5;"`
}

func (AdminCasbin) TableName() string {
	return "admin_casbin"
}
