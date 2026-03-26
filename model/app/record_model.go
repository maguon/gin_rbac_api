package app

import (
	"gin_rbac_api/model/common"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OpRecord struct {
	Id        primitive.ObjectID `json:"id" form:"id" bson:"_id,omitempty"`
	AdminId   int64              `json:"adminId" form:"adminId" bson:"admin_id"`
	AdminName string             `json:"adminName" form:"adminName" bson:"admin_name"`
	URL       string             `json:"url" form:"url" bson:"url"`
	Method    string             `json:"method" form:"method" bson:"method"`
	RoleId    int64              `json:"roleId" form:"roleId" gorm:"bson:role_id"`
	RoleName  string             `json:"roleName" form:"roleName" bson:"role_name"`
	Params    string             `json:"params" form:"params" gorm:"bson:params"`
	CreatedAt time.Time          `json:"CreatedAt" form:"CreatedAt" bson:"created_at"`
}

type OpRecordQuery struct {
	common.EXTEND_SEARCH
	OpRecord
}
