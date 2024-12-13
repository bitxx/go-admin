package models

import (
	"errors"
	"go-admin/core/config"
	"go-admin/core/utils/log"
	"go-admin/core/utils/strutils"
	"go-admin/core/utils/textutils"
	"gorm.io/gorm"
)

type DataPermission struct {
	DataScope string
	UserId    int
	DeptId    int
	RoleId    int
}

func (e *DataPermission) GetDataScope(tableName string, db *gorm.DB) (*gorm.DB, error) {

	if !config.ApplicationConfig.EnableDP {
		usageStr := `数据权限已经为您` + textutils.Green(`关闭`) + `，如需开启请参考配置文件字段说明`
		log.Debug("%s\n", usageStr)
		return db, nil
	}
	user := new(SysUser)
	role := new(SysRole)
	err := db.Find(user, e.UserId).Error
	if err != nil {
		return nil, errors.New("获取用户数据出错 msg:" + err.Error())
	}
	err = db.Find(role, user.RoleId).Error
	if err != nil {
		return nil, errors.New("获取用户数据出错 msg:" + err.Error())
	}
	if role.DataScope == "2" {
		db = db.Where(tableName+".create_by in (select admin_sys_user.id from admin_sys_role_dept left join admin_sys_user on admin_sys_user.dept_id=admin_sys_role_dept.dept_id where admin_sys_role_dept.role_id = ?)", user.RoleId)
	}
	if role.DataScope == "3" {
		db = db.Where(tableName+".create_by in (SELECT id from admin_sys_user where dept_id = ? )", user.DeptId)
	}
	if role.DataScope == "4" {
		db = db.Where(tableName+".create_by in (SELECT id from admin_sys_user where admin_sys_user.dept_id in(select dept_id from admin_sys_dept where dept_path like ? ))", "%"+strutils.IntToString(user.DeptId)+"%")
	}
	if role.DataScope == "5" || role.DataScope == "" {
		db = db.Where(tableName+".create_by = ?", e.UserId)
	}

	return db, nil
}

//func DataScopes(tableName string, userid int64) func(db *gorm.DB) *gorm.DB {
//	return func(db *gorm.DB) *gorm.DB {
//		user := new(SysUser)
//		role := new(SysRole)
//		user.Id = userId
//		err := db.Find(user, userId).Error
//		if err != nil {
//			db.Error = errors.New("获取用户数据出错 msg:" + err.Error())
//			return db
//		}
//		err = db.Find(role, user.Id).Error
//		if err != nil {
//			db.Error = errors.New("获取用户数据出错 msg:" + err.Error())
//			return db
//		}
//		if role.DataScope == "2" {
//			return db.Where(tableName+".create_by in (select admin_sys_user.id from admin_sys_role_dept left join admin_sys_user on admin_sys_user.dept_id=admin_sys_role_dept.dept_id where admin_sys_role_dept.role_id = ?)", user.Id)
//		}
//		if role.DataScope == "3" {
//			return db.Where(tableName+".create_by in (SELECT id from admin_sys_user where dept_id = ? )", user.Id)
//		}
//		if role.DataScope == "4" {
//			return db.Where(tableName+".create_by in (SELECT id from admin_sys_user where admin_sys_user.dept_id in(select dept_id from admin_sys_dept where dept_path like ? ))", "%"+pkg.IntToString(user.Id)+"%")
//		}
//		if role.DataScope == "5" || role.DataScope == "" {
//			return db.Where(tableName+".create_by = ?", userId)
//		}
//		return db
//	}
//}
