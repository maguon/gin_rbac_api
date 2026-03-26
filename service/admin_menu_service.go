package service

import (
	"gin_rbac_api/global"
	app "gin_rbac_api/model/app"
)

type AdminMenuService struct{}

func (adminMenuService *AdminMenuService) AddAdminMenu(adminMenu app.AdminMenu) (app.AdminMenu, error) {
	err := global.SYS_DB.Create(&adminMenu).Error
	return adminMenu, err
}

func (adminMenuService *AdminMenuService) UpdateAdminMenu(adminMenu app.AdminMenu) (adminMenuRes app.AdminMenu, err error) {
	err = global.SYS_DB.Where("id = ? ", adminMenu.ID).Updates(&adminMenu).Error
	return adminMenu, err
}

func (adminMenuService *AdminMenuService) RemoveAdminMenu(menuId int64) (err error) {
	var adminMenu app.AdminMenu
	err = global.SYS_DB.Where("id = ? ", menuId).Delete(&adminMenu).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (adminMenuService *AdminMenuService) GetAdminMenuList(queryModel app.AdminMenuQuery) (list interface{}, total int64, err error) {
	db := global.SYS_DB.Model(&app.AdminMenu{})
	var resList []app.AdminMenu

	if queryModel.ID != 0 {
		db = db.Where(" id = ? ", queryModel.ID)
	}
	if queryModel.Level != 0 {
		db = db.Where(" level = ? ", queryModel.Level)
	}
	if queryModel.ParentId != 0 {
		db = db.Where(" parent_id = ? ", queryModel.ParentId)
	}
	if queryModel.MenuName != "" {
		db = db.Where(" menu_name like ? ", "%"+queryModel.MenuName+"%")
	}
	if queryModel.Status != 0 {
		db = db.Where(" status = ? ", queryModel.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Debug().Order("sort desc ").Find(&resList).Error
	return resList, total, err
}

func (adminMenuService *AdminMenuService) GetAdminMenuJson() (list interface{}, total int64, err error) {
	var resultList *[]app.AdminMenuJson
	query := ` select am1.id,am1.parent_id,am1.url,am1.menu_name,am1.icon,am1.sort,am1.component,
		jsonb_agg(jsonb_strip_nulls(json_build_object('id',amb1.id,'btn_name',btn_name,'remark',amb1.remark)::jsonb))  as menu_btn,
		jsonb_agg(jsonb_strip_nulls(json_build_object('id',asm.id,'parent_id',asm.parent_id,'url',asm.url,
		'menu_name',asm.menu_name ,'icon',asm.icon,'component',asm.component,'sort',asm.sort,'menu_btn',asm.menu_btn  )::jsonb))  as children
		from admin_menu am1 left join admin_menu_btn amb1 on am1.id = amb1.admin_menu_id 
		left join (select am2.id,am2.parent_id,am2.url,am2.menu_name,am2.icon,am2.sort,am2.component,
		jsonb_agg(jsonb_strip_nulls(json_build_object('id',amb2.id,'btn_name',btn_name,'remark',amb2.remark)::jsonb))  as menu_btn
		from admin_menu am2 left join admin_menu_btn amb2 on am2.id = amb2.admin_menu_id  
		where am2.level=2 group by am2.id order by am2.sort desc) as asm on asm.parent_id = am1.id
		where am1.level=1 group by am1.id  order by am1.sort desc `

	queryParamArray := []interface{}{}

	global.SYS_DB.Raw(query, queryParamArray...).Scan(&resultList)
	return resultList, total, nil
}

func (adminMenuService *AdminMenuService) GetAdminRoleMenuJson(adminId int64) ([]app.AdminMenuJson, error) {
	var resultList []app.AdminMenuJson
	query := ` (select  distinct(amp.*) ,null::jsonb as menu_btn,ammp.children::jsonb from  admin_role_rel arr  
		left join admin_role_menu arm on arr.admin_role_id  = arm.admin_role_id 
		left join admin_menu am on arm.admin_menu_id  = am.id
		left join admin_menu amp on am.parent_id  = amp.id
		left join (select am.parent_id , jsonb_agg(json_build_object('menu_btn',arm.menu_btn ,
		'admin_menu_id',arm.admin_menu_id ,'level', am.level,'icon',am.icon ,'component',am.component ,'menu_name',am.menu_name ,
		'parent_id',am.parent_id ,'url',am.url,'sort',am.sort)) as children  
		from admin_role_rel arr 
		left join admin_role_menu arm on arr.admin_role_id  = arm.admin_role_id 
		left join admin_menu am on arm.admin_menu_id  = am.id
		where arr.admin_id  = ? and parent_id >0 

		group by am.parent_id ) as ammp on ammp.parent_id  = amp.id
		where arr.admin_id = ? and am.parent_id>0)
		union
		(select am.*,arm.menu_btn,null::jsonb as children from admin_role_rel arr   
		left join admin_role_menu arm on arr.admin_role_id  = arm.admin_role_id 
		left join admin_menu am on arm.admin_menu_id  = am.id
		where arr.admin_id  = ? and am.parent_id =0 )

		order by sort desc `

	queryParamArray := []interface{}{}
	queryParamArray = append(queryParamArray, adminId)
	queryParamArray = append(queryParamArray, adminId)
	queryParamArray = append(queryParamArray, adminId)
	global.SYS_DB.Raw(query, queryParamArray...).Find(&resultList)
	return resultList, nil
}

/* func (adminMenuService *AdminMenuService) GetAdminRoleMenuList(adminId int64) (list interface{}, total int64, err error) {
	var resultList *[]app.AdminMenuJson
	db := global.SYS_DB.Table("admin_info ai ").Joins(
		" left join admin_role_rel arr on arr.admin_id = ai.id ").Joins(
		" left join admin_role_menu arm on arm.admin_role_id = arr.admin_role_id ").Joins(
		" left join admin_menu am on arm.admin_menu_id = am.id ").Select(
		" am.* ,arm.menu_btn ")

	db = db.Where("arr.admin_id = ?", adminId)

	err = db.Debug().Order("bjts.id ").Find(&resultList).Error
	return resultList, total, nil
} */

func (adminMenuService *AdminMenuService) AddAdminMenuBtn(adminMenuBtn app.AdminMenuBtn) (app.AdminMenuBtn, error) {
	err := global.SYS_DB.Create(&adminMenuBtn).Error
	return adminMenuBtn, err
}

func (adminMenuService *AdminMenuService) UpdateAdminMenuBtn(adminMenuBtn app.AdminMenuBtn) (adminMenuBtnRes app.AdminMenuBtn, err error) {
	err = global.SYS_DB.Where("id = ? ", adminMenuBtn.ID).Updates(&adminMenuBtn).Error
	return adminMenuBtn, err
}

func (adminMenuService *AdminMenuService) RemoveAdminMenuAllBtn(menuId int64) (err error) {
	var adminMenuBtn app.AdminMenuBtn
	err = global.SYS_DB.Where("admin_menu_id = ? ", menuId).Delete(&adminMenuBtn).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (adminMenuService *AdminMenuService) RemoveAdminMenuBtn(menuBtnId int64) (err error) {
	var adminMenuBtn app.AdminMenuBtn
	err = global.SYS_DB.Where("id = ? ", menuBtnId).Delete(&adminMenuBtn).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (adminMenuService *AdminMenuService) GetAdminMenuBtn(queryModel app.AdminMenuBtnQuery) (list interface{}, total int64, err error) {
	db := global.SYS_DB.Model(&app.AdminMenuBtn{})
	var resList []app.AdminMenuBtn

	if queryModel.ID != 0 {
		db = db.Where(" id = ? ", queryModel.ID)
	}
	if queryModel.AdminMenuId != 0 {
		db = db.Where(" admin_menu_id = ? ", queryModel.AdminMenuId)
	}
	if queryModel.Status != nil {
		db = db.Where(" status = ? ", queryModel.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Debug().Order("id  ").Find(&resList).Error
	return resList, total, err
}

func (adminMenuService *AdminMenuService) AddAdminRoleMenu(adminRoleMenu app.AdminRoleMenu) (app.AdminRoleMenu, error) {
	err := global.SYS_DB.Create(&adminRoleMenu).Error
	return adminRoleMenu, err
}

func (adminMenuService *AdminMenuService) UpdateAdminRoleMenu(adminRoleMenu app.AdminRoleMenu) (adminRoleMenuRes app.AdminRoleMenu, err error) {
	err = global.SYS_DB.Where("admin_role_id = ? AND admin_menu_id = ?", adminRoleMenu.AdminRoleId, adminRoleMenu.AdminMenuId).Updates(&adminRoleMenu).Error
	return adminRoleMenu, err
}

func (adminMenuService *AdminMenuService) RemoveAdminRolMenu(menuId int64, roleId int64) (err error) {
	var adminRoleMenu app.AdminRoleMenu
	err = global.SYS_DB.Where("admin_menu_id = ? AND admin_role_id = ? ", menuId, roleId).Delete(&adminRoleMenu).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}
func (adminMenuService *AdminMenuService) RemoveAdminRolMenuIdRel(menuId int64) (err error) {
	var adminRoleMenu app.AdminRoleMenu
	err = global.SYS_DB.Where("admin_menu_id = ? ", menuId).Delete(&adminRoleMenu).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (adminMenuService *AdminMenuService) GetAdminRoleMenuList(queryModel app.AdminRoleMenuQuery) (list interface{}, total int64, err error) {
	db := global.SYS_DB.Model(&app.AdminRoleMenu{})
	var resList []app.AdminRoleMenu

	if queryModel.ID != 0 {
		db = db.Where(" id = ? ", queryModel.ID)
	}
	if queryModel.AdminMenuId != 0 {
		db = db.Where(" admin_menu_id = ? ", queryModel.AdminMenuId)
	}
	if queryModel.AdminRoleId != 0 {
		db = db.Where(" admin_role_id = ? ", queryModel.AdminRoleId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Debug().Order("id  ").Find(&resList).Error
	return resList, total, err
}
