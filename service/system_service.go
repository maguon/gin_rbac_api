package service

import (
	"context"
	"gin_rbac_api/global"
	"gin_rbac_api/model/app"
	"gin_rbac_api/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type SystemService struct{}

func (systemService *SystemService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.SYS_LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Ram, err = utils.InitRAM(); err != nil {
		global.SYS_LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.SYS_LOG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}

func (systemService *SystemService) GetRecordList(queryModel app.OpRecordQuery) (list interface{}, total int64, err error) {
	limit := queryModel.PageSize
	skip := queryModel.PageSize * (queryModel.PageNumber - 1)
	collection := global.SYS_MONGO.Collection("sys_record")
	opts := options.Find()
	var resList []app.OpRecord
	filter := bson.M{}
	if queryModel.AdminId != 0 {
		filter["admin_id"] = queryModel.AdminId
	}
	if queryModel.AdminName != "" {
		filter["admin_name"] = bson.M{"$regex": queryModel.AdminName, "$options": "i"}
	}
	if queryModel.RoleId != 0 {
		filter["role_id"] = queryModel.RoleId
	}
	if queryModel.RoleName != "" {
		filter["role_name"] = bson.M{"$regex": queryModel.RoleName, "$options": "i"}
	}
	if queryModel.Method != "" {
		filter["method"] = queryModel.Method
	}
	if queryModel.URL != "" {
		filter["url"] = bson.M{"$regex": queryModel.URL, "$options": "i"}
	}

	/* if queryModel.ActressId != "" {
		actressId, _ := primitive.ObjectIDFromHex(queryModel.ActressId)
		filter["actress._id"] = actressId
	} */

	if !queryModel.CreatedStart.IsZero() && !queryModel.CreatedEnd.IsZero() {
		filter["created_at"] = bson.M{"$gte": primitive.NewDateTimeFromTime(queryModel.CreatedStart), "$lte": primitive.NewDateTimeFromTime(queryModel.CreatedEnd)}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, filter, opts.SetSort(bson.M{"created_at": -1}).SetSkip(int64(skip)).SetLimit(int64(limit)))
	if err != nil {
		return
	}
	if err = cur.All(ctx, &resList); err != nil {
		return
	}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var item app.OpRecord
		cur.Decode(&item)
		resList = append(resList, item)
	}
	return resList, count, err
}

func (systemService *SystemService) AddRecordList(model app.OpRecord) (resModel app.OpRecord, id interface{}, err error) {
	collection := global.SYS_MONGO.Collection("sys_record")
	model.CreatedAt = time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, model)

	return model, res.InsertedID, err
}

func (systemService *SystemService) RemoveRecord(id string) (resDel *mongo.DeleteResult, err error) {
	collection := global.SYS_MONGO.Collection("sys_record")
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": objId}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := collection.DeleteOne(ctx, filter)
	return res, err
}

//todo delete service
