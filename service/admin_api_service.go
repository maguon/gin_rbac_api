package service

import (
	"gin_rbac_api/global"
	app "gin_rbac_api/model/app"
)

type AdminApiService struct{}

func (adminApiService *AdminApiService) AddAdminApi(adminSysApi app.AdminSysApi) (app.AdminSysApi, error) {
	err := global.SYS_DB.Create(&adminSysApi).Error
	return adminSysApi, err
}

func (adminApiService *AdminApiService) UpdateAdminApi(adminSysApi app.AdminSysApi) (adminInfoRes app.AdminSysApi, err error) {
	err = global.SYS_DB.Where("id = ? ", adminSysApi.ID).Updates(&adminSysApi).Error
	return adminSysApi, err
}

func (adminApiService *AdminApiService) RemoveAdminApi(apiId int64) (err error) {
	var adminSysApi app.AdminSysApi
	err = global.SYS_DB.Where("id = ? ", apiId).Delete(&adminSysApi).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (adminApiService *AdminApiService) GetAdminApiAll(queryModel app.AdminSysApiQuery) (list interface{}, total int64, err error) {
	var resList []app.AdminSysApi
	err = global.SYS_DB.Model(&app.AdminSysApi{}).Find(&resList).Error
	total = int64(len(resList))
	return
}

func (adminApiService *AdminApiService) GetAdminApiList(queryModel app.AdminSysApiQuery) (list []app.AdminSysApi, total int64, err error) {
	limit := queryModel.PageSize
	offset := queryModel.PageSize * (queryModel.PageNumber - 1)

	db := global.SYS_DB.Model(&app.AdminSysApi{})
	var resList []app.AdminSysApi

	if queryModel.ID != 0 {
		db = db.Where(" id = ? ", queryModel.ID)
	}
	if queryModel.Tag != "" {
		db = db.Where(" tag like ? ", "%"+queryModel.Tag+"%")
	}
	if queryModel.Method != "" {
		db = db.Where(" method = ? ", queryModel.Method)
	}
	if queryModel.Remark != "" {
		db = db.Where(" remark like ? ", "%"+queryModel.Remark+"%")
	}
	if queryModel.Url != "" {
		db = db.Where(" url = ?", queryModel.Url)
	}
	if queryModel.Status != 0 {
		db = db.Where(" status = ? ", queryModel.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Debug().Order("id desc ").Limit(limit).Offset(offset).Find(&resList).Error
	return resList, total, err
}

func (adminApiService *AdminApiService) GetAdminApiJson() (list interface{}, total int64, err error) {

	var resultList *[]app.AdminSysApiTagGroup
	query := ` SELECT  tag, JSONB_AGG(json_build_object('id',id,'method',method,'url',url,'remark',remark) ) 
				as api_list
				FROM admin_sys_api
				group by tag `

	queryParamArray := []interface{}{}

	global.SYS_DB.Raw(query, queryParamArray...).Scan(&resultList)
	return resultList, total, nil
}
