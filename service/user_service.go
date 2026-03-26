package service

import (
	"errors"
	"fmt"
	"gin_rbac_api/global"
	app "gin_rbac_api/model/app"
	"gin_rbac_api/utils"
	"time"

	"gorm.io/gorm"
)

type UserService struct{}

//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (userService *UserService) Login(u *app.UserInfo) (userInter *app.UserInfo, err error) {
	if nil == global.SYS_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user app.UserInfo
	err = global.SYS_DB.Where("phone = ?", u.Phone).First(&user).Error
	global.SYS_LOG.Info(user.Password)
	global.SYS_LOG.Info(utils.BcryptHash(u.Password))
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}
	return &user, err
}

func (userService *UserService) UpdateLoginAt(userId int64) (err error) {
	err = global.SYS_DB.Where("id = ? ", userId).First(&app.UserInfo{}).Updates(map[string]interface{}{"last_login_at": time.Now(), "last_login_id": utils.GetDateId()}).Error
	return err
}

func (userService *UserService) UpdateUserOpenId(userId int64, openId string, referId int64) (err error) {
	if referId == 0 {
		err = global.SYS_DB.Where("open_id is null and id = ? ", userId).First(&app.UserInfo{}).Updates(map[string]interface{}{"open_id": openId}).Error
	} else {
		err = global.SYS_DB.Where("open_id is null and refer_id= 0  and id = ? ", userId).First(&app.UserInfo{}).Updates(map[string]interface{}{"open_id": openId, "refer_id": referId}).Error
	}
	return err
}

//@function: Register
//@description: 用户注册
//@param: u *req.Login
//@return: err error, userId int64

func (userService *UserService) Register(userInfo *app.UserInfo) (userRes *app.UserInfo, err error) {
	if nil == global.SYS_DB {
		return nil, fmt.Errorf("db not init")
	}
	userInfo.Password = utils.BcryptHash(userInfo.Password)
	err = global.SYS_DB.Create(&userInfo).Error
	return userInfo, err
}

//@function: GetUserInfo
//@description: 获取用户信息
//@return: err error, user system.SysUser

func (userService *UserService) GetUserInfo(userId int64) (user app.UserInfo, err error) {
	var reqUser app.UserInfo
	err = global.SYS_DB.First(&reqUser, "id = ?", userId).Error
	if err != nil {
		return reqUser, err
	}
	return reqUser, err
}

func (userService *UserService) GetUserList(queryModel app.UserQuery) (list interface{}, total int64, err error) {
	limit := queryModel.PageSize
	offset := queryModel.PageSize * (queryModel.PageNumber - 1)
	db := global.SYS_DB.Table("user_info ui ")
	var userInfoList []app.UserInfo
	if queryModel.ID != 0 {
		db = db.Where("ui.id = ?", queryModel.ID)
	}
	if queryModel.Status != nil {
		db = db.Where("ui.status = ?", queryModel.Status)
	}
	if queryModel.Gender != nil {
		db = db.Where("ui.gender = ?", queryModel.Gender)
	}
	if queryModel.Name != "" {
		db = db.Where("ui.name like ?", "%"+queryModel.Name+"%")
	}
	if queryModel.Email != "" {
		db = db.Where("ui.email like ?", "%"+queryModel.Email+"%")
	}
	if queryModel.Phone != "" {
		db = db.Where("ui.phone like ?", "%"+queryModel.Phone+"%")
	}
	if queryModel.OpenId != "" {
		db = db.Where("ui.open_id = ?", queryModel.OpenId)
	}
	if queryModel.ReferId != 0 {
		db = db.Where("ui.refer_id = ?", queryModel.ReferId)
	}
	if !queryModel.LoginStart.IsZero() && !queryModel.LoginEnd.IsZero() {
		db = db.Where("ui.last_login_at >= ?", queryModel.LoginStart)
		db = db.Where("ui.last_login_at <= ?", queryModel.LoginEnd)
	}
	if !queryModel.BirthStart.Time.IsZero() && !queryModel.BirthEnd.Time.IsZero() {
		db = db.Where("ui.birth >= ?", queryModel.BirthStart)
		db = db.Where("ui.birth <= ?", queryModel.BirthEnd)
	}
	if !queryModel.CreatedStart.IsZero() && !queryModel.CreatedEnd.IsZero() {
		db = db.Where("ui.created_at >= ?", queryModel.CreatedStart)
		db = db.Where("ui.created_at <= ?", queryModel.CreatedEnd)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Debug().Limit(limit).Offset(offset).Order("id desc").Find(&userInfoList).Error
	return userInfoList, total, err
}

func (userService *UserService) GetUserAccount(userId int64) (userAccout app.UserAccount, err error) {
	var userAccount app.UserAccount
	err = global.SYS_DB.Model(&app.UserInfo{}).Where("id = ?", userId).First(&userAccount).Error
	if err != nil {
		return userAccount, err
	}
	return userAccount, err
}

func (userService *UserService) UpdateUserInfo(userId int64, userObj app.UserInfo) (user app.UserInfo, err error) {
	err = global.SYS_DB.Where("id = ? ", userId).Updates(&userObj).Error
	return userObj, err
}

func (userService *UserService) UpdateUserAccount(userId int64, userAccount interface{}) (err error) {
	err = global.SYS_DB.Model(&app.UserInfo{}).Where("id = ?", userId).Updates(userAccount).Error
	return err
}

func (userService *UserService) UpdateUserStatus(userId int64, status int) (res *gorm.DB) {
	sql := ` update user_info set status = ?  where id = ?	`
	result := global.SYS_DB.Exec(sql, status, userId)
	return result
}

func (userService *UserService) UpdateUserAvatar(userId int64, avatar string) (err error) {
	err = global.SYS_DB.Model(&app.UserInfo{}).Where("id = ?", userId).Updates(map[string]interface{}{"avatar": avatar}).Error
	return err
}

func (userService *UserService) GetUserByPhone(phone string) (userListRes []app.UserInfo, err error) {
	var userInfoList []app.UserInfo
	db := global.SYS_DB.Model(&app.UserInfo{})
	db = db.Where("phone = ?", phone)
	err = db.Find(&userInfoList).Error
	return userInfoList, err
}

func (userService *UserService) GetUserByOpenId(openId string) (userListRes []app.UserInfo, err error) {
	var userInfoList []app.UserInfo
	db := global.SYS_DB.Model(&app.UserInfo{})
	db = db.Where("open_id = ?", openId)
	err = db.Find(&userInfoList).Error
	return userInfoList, err
}

func (userService *UserService) GetUserByEmail(email string) (userListRes []app.UserInfo, err error) {
	var userInfoList []app.UserInfo
	db := global.SYS_DB.Model(&app.UserInfo{})
	db = db.Where("email = ?", email)
	err = db.Find(&userInfoList).Error
	return userInfoList, err
}

func (userService *UserService) UpdatePassword(sysUserPassword app.SysUserPassword) (userInfoRes *app.UserInfo, err error) {
	var userInfo app.UserInfo
	if err = global.SYS_DB.First(&userInfo, "id = ?", sysUserPassword.ID).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(sysUserPassword.Password, userInfo.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	userInfo.Password = utils.BcryptHash(sysUserPassword.NewPassword)
	err = global.SYS_DB.Save(&userInfo).Error
	return &userInfo, err
}

func (userService *UserService) SetPassword(login app.Login) (userInfoRes *app.UserInfo, err error) {
	var userInfo app.UserInfo
	if err = global.SYS_DB.First(&userInfo, "phone = ?", login.Phone).Error; err != nil {
		return nil, err
	}
	userInfo.Password = utils.BcryptHash(login.Password)
	err = global.SYS_DB.Save(&userInfo).Error
	return &userInfo, err
}
