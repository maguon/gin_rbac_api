package service

import (
	"errors"
	"fmt"
	"gin_rbac_api/global"
	"gin_rbac_api/utils"
	"time"

	app "gin_rbac_api/model/app"

	"gorm.io/gorm"
)

type AdminService struct{}

func (adminService *AdminService) AdminLogin(u *app.AdminInfo) (adminRoleInfoRes *app.AdminRoleInfo, err error) {
	if nil == global.SYS_DB {
		return nil, fmt.Errorf("db not init")
	}

	var adminRoleInfo []app.AdminRoleInfo
	err = global.SYS_DB.Table("admin_info ai ").Joins(
		" left join admin_role_rel arr on arr.admin_id = ai.id ").Joins(
		" left join admin_role ar on arr.admin_role_id = ar.id ").Select(
		" ai.* ,ar.id role_id, ar.role_name,ar.default_url,ar.remark ").Where("ai.phone = ?", u.Phone).Order(" ai.id desc").Find(&adminRoleInfo).Error
	global.SYS_LOG.Info(utils.BcryptHash(u.Password))
	if err == nil {
		if len(adminRoleInfo) == 0 {
			return nil, errors.New("用户不存在")
		} else {
			if ok := utils.BcryptCheck(u.Password, adminRoleInfo[0].AdminInfo.Password); !ok {
				return nil, errors.New("密码错误")
			} else {
				return &adminRoleInfo[0], err
			}
		}

	} else {
		return nil, err
	}
}

func (adminService *AdminService) GetAdminInfo(adminId int64) (adminUser app.AdminInfo, err error) {
	var reqUser app.AdminInfo
	err = global.SYS_DB.First(&reqUser, "id = ?", adminId).Error
	if err != nil {
		return reqUser, err
	}
	fmt.Println(reqUser)
	return reqUser, err
}

func (adminService *AdminService) UpdateAdminLoginAt(adminId int64) (err error) {
	err = global.SYS_DB.Model(&app.AdminInfo{}).Where("id = ? ", adminId).Updates(map[string]interface{}{"last_login_at": time.Now()}).Error
	return err
}

func (adminService *AdminService) GetSysUserList(queryModel app.AdminUserQuery) (list interface{}, total int64, err error) {
	limit := queryModel.PageSize
	offset := queryModel.PageSize * (queryModel.PageNumber - 1)
	fmt.Println("limit", limit)
	fmt.Println("offset", queryModel.PageNumber)
	db := global.SYS_DB.Table("admin_info ai ").
		Joins("left join admin_role_rel arr on arr.admin_id = ai.id").
		Joins("left join admin_role ar on arr.admin_role_id = ar.id").
		Select("ai.id,ai.created_at,ai.updated_at,ai.last_login_at,ai.status,ai.username,ai.avatar,ai.phone,ai.email,ai.gender,ar.id role_id,ar.role_name,ar.default_url,ar.remark")
	var adminInfoList []app.AdminRoleInfo
	if queryModel.ID != 0 {
		db = db.Where("ai.id = ?", queryModel.ID)
	}
	if queryModel.Gender != 0 {
		db = db.Where("ai.gender = ?", queryModel.Gender)
	}
	if queryModel.Email != "" {
		db = db.Where("ai.email = ?", queryModel.Email)
	}
	if queryModel.Phone != "" {
		db = db.Where("ai.phone = ?", queryModel.Phone)
	}
	if queryModel.Username != "" {
		db = db.Where("ai.user_name = ?", queryModel.Username)
	}
	if queryModel.AdminRoleId != 0 {
		db = db.Where("arr.admin_role_id = ?", queryModel.AdminRoleId)
	}
	if queryModel.Status != nil {
		db = db.Where("ai.status = ?", queryModel.Status)
	}
	fmt.Println(queryModel.CreatedStart)
	if !queryModel.CreatedStart.IsZero() && !queryModel.CreatedEnd.IsZero() {
		db = db.Where("ai.created_at >= ?", queryModel.CreatedStart)
		db = db.Where("ai.created_at <= ?", queryModel.CreatedEnd)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Debug().Limit(limit).Offset(offset).Order("id desc").Find(&adminInfoList).Error
	return adminInfoList, total, err
}

func (adminService *AdminService) AddAdminInfo(adminInfo app.AdminInfo) (adminInfoRes app.AdminInfo, err error) {
	err = global.SYS_DB.Create(&adminInfo).Error
	return adminInfo, err
}

func (adminService *AdminService) UpdateAdminInfo(adminInfo app.AdminInfo) (adminInfoRes app.AdminInfo, err error) {
	err = global.SYS_DB.Where("id = ? ", adminInfo.ID).Updates(&adminInfo).Error
	return adminInfo, err
}

func (adminService *AdminService) RemoveAdminInfo(adminId int64) (err error) {
	var adminInfo app.AdminInfo
	err = global.SYS_DB.Where(" id = ?  ", adminId).Delete(&adminInfo).Error
	return err
}

func (adminService *AdminService) GetAdminToken() (list interface{}, total int64, err error) {
	db := global.SYS_DB.Model(&app.AdminToken{})
	var resList []app.AdminToken

	err = db.Debug().Order("id").Find(&resList).Error
	return resList, 1, err
}

func (adminService *AdminService) UpdateAdminToken(adminToken app.AdminToken) (res *gorm.DB) {
	result := global.SYS_DB.Where("id = ? ", adminToken.ID).Attrs(adminToken).Assign(adminToken).FirstOrCreate(&adminToken)
	return result
}

func (adminService *AdminService) AddAdminRole(adminRole app.AdminRole) (adminRoleRes app.AdminRole, err error) {
	err = global.SYS_DB.Create(&adminRole).Error
	return adminRole, err
}

func (adminService *AdminService) UpdateAdminRole(adminRole app.AdminRole) (adminRoleRes app.AdminRole, err error) {
	err = global.SYS_DB.Where("id = ? ", adminRole.ID).Updates(&adminRole).Error
	return adminRole, err
}

func (adminService *AdminService) UpdateAdminPassword(sysUserPassword app.SysUserPassword) (adminUser *app.AdminInfo, err error) {
	var adminInfo app.AdminInfo
	if err = global.SYS_DB.First(&adminInfo, "id = ?", sysUserPassword.ID).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(sysUserPassword.Password, adminInfo.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	adminInfo.Password = utils.BcryptHash(sysUserPassword.NewPassword)
	err = global.SYS_DB.Save(&adminInfo).Error
	return &adminInfo, err
}

func (adminService *AdminService) UpdateAdminAvatar(adminId int64, avatar string) (err error) {
	err = global.SYS_DB.Model(&app.AdminInfo{}).Where("id = ?", adminId).Updates(map[string]interface{}{"avatar": avatar}).Error
	return err
}

func (adminService *AdminService) GetAdminRole(adminRoleQuery app.AdminRoleQuery) (list interface{}, total int64, err error) {
	limit := adminRoleQuery.PageSize
	offset := adminRoleQuery.PageSize * (adminRoleQuery.PageNumber - 1)
	db := global.SYS_DB.Model(&app.AdminRole{})
	var adminRoleList []app.AdminRole
	if adminRoleQuery.ID != 0 {
		db = db.Where("id = ?", adminRoleQuery.ID)
	}
	if adminRoleQuery.RoleName != "" {
		db = db.Where("role_name = ?", adminRoleQuery.RoleName)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit > 0 {
		db.Limit(limit)
	}
	err = db.Offset(offset).Order("id desc").Find(&adminRoleList).Error
	return adminRoleList, total, err
}

func (adminService *AdminService) RemoveAdminRole(adminRoleId int64) (total int64, err error) {
	db := global.SYS_DB.Table("admin_role_rel ").Select("id")
	db = db.Where("admin_role_id = ?", adminRoleId)

	err = db.Count(&total).Error
	if err != nil {
		return total, err
	}
	if total == 0 {
		var adminRole app.AdminRole
		err = global.SYS_DB.Where("id = ? ", adminRoleId).Delete(&adminRole).Error
		if err != nil {
			return total, err
		}
	}
	return total, err
}

func (adminService *AdminService) AddAdminRoleRel(adminRoleRel app.AdminRoleRel) (adminRoleRelRes app.AdminRoleRel, err error) {
	err = global.SYS_DB.Create(&adminRoleRel).Error
	return adminRoleRelRes, err
}

func (adminService *AdminService) UpdateAdminRoleRel(adminRoleRel app.AdminRoleRel) (adminRoleRelRes app.AdminRoleRel, err error) {
	err = global.SYS_DB.Where("id = ? ", adminRoleRel.ID).Updates(&adminRoleRel).Error
	return adminRoleRel, err
}

func (adminService *AdminService) UpdateAdminUserRoleRel(adminRoleRel app.AdminRoleRel) (adminRoleRelRes app.AdminRoleRel, err error) {
	err = global.SYS_DB.Where("admin_id = ? ", adminRoleRel.AdminId).Updates(&adminRoleRel).Error
	return adminRoleRel, err
}
func (adminService *AdminService) GetAdminRoleRel(adminRoleRelQuery app.AdminRoleRelQuery) (list interface{}, total int64, err error) {
	db := global.SYS_DB.Model(&app.AdminRoleRel{})
	var adminRoleRelList []app.AdminRoleRel
	if adminRoleRelQuery.ID != 0 {
		db = db.Where("id = ?", adminRoleRelQuery.ID)
	}
	if adminRoleRelQuery.AdminRoleId != 0 {
		db = db.Where("admin_role_id = ?", adminRoleRelQuery.AdminRoleId)
	}
	if adminRoleRelQuery.AdminId != 0 {
		db = db.Where("admin_id = ?", adminRoleRelQuery.AdminId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("id desc").Find(&adminRoleRelList).Error
	return adminRoleRelList, total, err
}

func (adminService *AdminService) RemoveAdminRoleRel(adminId int64, adminRoleId int64) (err error) {
	var adminRoleRel app.AdminRoleRel
	err = global.SYS_DB.Where("admin_role_id = ? and admin_id = ? ", adminRoleId, adminId).Delete(&adminRoleRel).Error
	return err
}
