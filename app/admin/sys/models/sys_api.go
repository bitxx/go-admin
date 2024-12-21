package models

import (
	"errors"
	"fmt"
	"go-admin/config/base/constant"
	"go-admin/core/runtime"
	"gorm.io/gorm"
	"strings"
	"time"
)

type SysApi struct {
	Id          int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Description string     `json:"description" gorm:"size:256;comment:功能描述"`
	Path        string     `json:"path" gorm:"size:128;comment:地址"`
	Method      string     `json:"method" gorm:"size:16;comment:请求类型"`
	ApiType     string     `json:"apiType" gorm:"size:16;comment:接口类型"`
	Remark      string     `json:"remark" gorm:"size:128;comment:备注"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	CreateBy    int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy    int64      `json:"updateBy" gorm:"index;comment:更新者"`
	SysMenu     []SysMenu  `json:"sysMenu" gorm:"many2many:admin_sys_menu_api_rule;foreignKey:id;joinForeignKey:admin_sys_api_id;references:Id;joinReferences:admin_sys_menu_menu_id;"`
}

func (SysApi) TableName() string {
	return "admin_sys_api"
}

func SaveSysApi(db *gorm.DB, routers []runtime.Router) (err error) {
	tx := db.Begin()
	var dbApiCacheMap = make(map[string]bool)
	var handlerApiCacheMap = make(map[string]*runtime.Router) //缓存路由中实际包含的地址，用于对比数据库，删除库中已经实效的路由地址
	defer func() {
		dbApiCacheMap = nil
		handlerApiCacheMap = nil
		if err != nil {
			tx.Rollback()
			return
		} else {
			tx.Commit()
		}
	}()

	//读取库中所有接口并加入map缓存
	var dbApilist []SysApi
	err = db.Model(&SysApi{}).Find(&dbApilist).Error
	if err != nil {
		err = errors.New(fmt.Sprintf("get Api dbApilist error: %s \r\n ", err.Error()))
		return
	}
	for _, item := range dbApilist {
		dbApiCacheMap[item.Path+"-"+item.Method] = true
	}
	for _, v := range routers {
		handlerApiCacheMap[v.RelativePath+"-"+v.HttpMethod] = &v
	}

	//根据实际路由对比库中路由，将新路由加入库中
	var newSysApis []SysApi
	for k, v := range handlerApiCacheMap {

		if v.HttpMethod == "HEAD" {
			continue
		}
		//缓存
		if dbApiCacheMap[k] {
			continue
		}
		paths := strings.Split(v.RelativePath, "/")
		apiType := ""
		if len(paths) >= 4 {
			if strings.HasPrefix(paths[3], "admin") {
				apiType = constant.ApiTypeSys
			} else if strings.HasPrefix(paths[3], "plugins") {
				apiType = constant.ApiTypePlugin
			} else if strings.HasPrefix(paths[3], "app") {
				apiType = constant.ApiTypeApp
			}
		}
		newSysApi := SysApi{Path: v.RelativePath, Method: v.HttpMethod}
		if apiType != "" {
			newSysApi.ApiType = apiType
		}
		if ApiDescMap[v.Handler] != "" {
			newSysApi.Description = ApiDescMap[v.Handler]
		}
		newSysApis = append(newSysApis, newSysApi)
	}
	if len(newSysApis) > 0 {
		//事务批量插入，提高效率
		if err = tx.Debug().Model(&SysApi{}).Create(&newSysApis).Error; err != nil {
			return err
		}
		if err != nil {
			err = errors.New(fmt.Sprintf("Models SaveSysApi error: %s \r\n ", err.Error()))
			return
		}
		for _, item := range newSysApis {
			dbApiCacheMap[item.Path+"-"+item.Method] = true
		}
	}

	// 删除库中无效接口
	var delIds []int64
	for _, item := range dbApilist {
		if handlerApiCacheMap[item.Path+"-"+item.Method] == nil {
			delIds = append(delIds, item.Id)
		}
	}
	if len(delIds) > 0 {
		if err = tx.Table("admin_sys_menu_api_rule").Where("admin_sys_api_id in (?)", delIds).Delete(nil).Error; err != nil {
			return
		}
		// 删除主表数据
		if err = tx.Model(&SysApi{}).Delete(&SysApi{}, delIds).Error; err != nil {
			return
		}
		if err != nil {
			err = errors.New(fmt.Sprintf("sync delete api error: %s \r\n ", err.Error()))
			return
		}
	}
	return
}
