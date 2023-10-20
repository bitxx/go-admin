package middleware

import (
	"errors"
	"github.com/bitxx/logger/logbase"
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/constant"
	"go-admin/core/config"
	"go-admin/core/dto/response"
	"go-admin/core/middleware/auth"
	"go-admin/core/utils/ginutils"
	"go-admin/core/utils/strutils"
	"gorm.io/gorm"
)

const (
	PermissionKey = "dataPermission"
)

type DataPermission struct {
	DataScope string
	UserId    int64
	DeptId    int64
	RoleId    int64
}

func PermissionAction() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := ginutils.GetOrm(c)
		if err != nil {
			logbase.Error(err)
			return
		}

		msgID := strutils.GenerateMsgIDFromContext(c)
		var p = new(DataPermission)
		userId, _, _ := auth.Auth.GetUserId(c)
		if userId > 0 {
			p, err = newDataPermission(db, userId)
			if err != nil {
				logbase.Errorf("MsgID[%s] PermissionAction error: %s", msgID, err)
				response.Error(c, 500, "权限范围鉴定错误")
				c.Abort()
				return
			}
		}
		c.Set(PermissionKey, p)
		c.Next()
	}
}

func newDataPermission(tx *gorm.DB, userId interface{}) (*DataPermission, error) {
	var err error
	p := &DataPermission{}

	err = tx.Table("sys_user").
		Select("sys_user.id", "sys_role.id", "sys_user.dept_id", "sys_role.data_scope").
		Joins("left join sys_role on sys_role.id = sys_user.role_id").
		Where("sys_user.id = ?", userId).
		Scan(p).Error
	if err != nil {
		err = errors.New("获取用户数据出错 msg:" + err.Error())
		return nil, err
	}
	return p, nil
}

func Permission(tableName string, p *DataPermission) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !config.ApplicationConfig.EnableDP || p == nil {
			return db
		}
		switch p.DataScope {
		case constant.DataScope2:
			//自定数据权限
			return db.Where(tableName+".create_by in (select sys_user.id from sys_role_dept left join sys_user on sys_user.dept_id=sys_role_dept.dept_id where sys_role_dept.role_id = ?)", p.RoleId)
		case constant.DataScope3:
			//本部门数据权限
			return db.Where(tableName+".create_by in (SELECT id from sys_user where dept_id = ? )", p.DeptId)
		case constant.DataScope4:
			//本部门及以下数据权限
			return db.Where(tableName+".create_by in (SELECT id from sys_user where sys_user.dept_id in(select dept_id from sys_dept where dept_path like ? ))", "%/"+strutils.Int64ToString(p.DeptId)+"/%")
		case constant.DataScope5:
			//仅本人数据权限
			return db.Where(tableName+".create_by = ?", p.UserId)
		default:
			return db
		}
	}
}

func getPermissionFromContext(c *gin.Context) *DataPermission {
	p := new(DataPermission)
	if pm, ok := c.Get(PermissionKey); ok {
		switch pm.(type) {
		case *DataPermission:
			p = pm.(*DataPermission)
		}
	}
	return p
}

// GetPermissionFromContext 提供非action写法数据范围约束
func GetPermissionFromContext(c *gin.Context) *DataPermission {
	return getPermissionFromContext(c)
}
