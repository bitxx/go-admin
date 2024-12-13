package service

import (
	"go-admin/app/app/user/constant"
	"go-admin/app/app/user/models"
	"go-admin/app/app/user/service/dto"
	"go-admin/core/config"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/utils/encrypt"
	"go-admin/core/utils/strutils"
	"gorm.io/gorm"
	"time"
)

type UserConf struct {
	service.Service
}

// NewUserConfService app-实例用户配置管理记录
func NewUserConfService(s *service.Service) *UserConf {
	var srv = new(UserConf)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage app-获取用户配置管理分页列表
func (e *UserConf) GetPage(c *dto.UserConfQueryReq, p *middleware.DataPermission) ([]models.UserConf, int64, int, error) {
	var data models.UserConf
	var list []models.UserConf
	var count int64
	var err error
	if c.Mobile != "" {
		c.Mobile, err = encrypt.AesEncrypt(c.Mobile, []byte(config.AuthConfig.Secret))
		if err != nil {
			c.Mobile = ""
		}
	}
	if c.Email != "" {
		c.Email, err = encrypt.AesEncrypt(c.Email, []byte(config.AuthConfig.Secret))
		if err != nil {
			c.Email = ""
		}
	}

	err = e.Orm.Joins("User").Order("created_at desc").Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}

	for index, item := range list {
		if item.User != nil && item.User.Mobile != "" {
			mobile, err := encrypt.AesDecrypt(item.User.Mobile, []byte(config.AuthConfig.Secret))
			if err == nil {
				if c.ShowInfo {
					list[index].User.Mobile = mobile
				} else {
					list[index].User.Mobile = strutils.HidePartStr(mobile, 3)
				}
			}
		}

		if item.User != nil && item.User.Email != "" {
			email, err := encrypt.AesDecrypt(item.User.Email, []byte(config.AuthConfig.Secret))
			if err == nil {
				if c.ShowInfo {
					list[index].User.Email = email
				} else {
					list[index].User.Email = strutils.HidePartStr(email, 5)
				}
			}
		}
	}
	return list, count, lang.SuccessCode, nil
}

// Get app-获取用户配置管理详情
func (e *UserConf) Get(id int64, p *middleware.DataPermission) (*models.UserConf, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.UserConf{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return data, lang.SuccessCode, nil
}

// QueryOne app-获取用户配置管理一条记录
func (e *UserConf) QueryOne(queryCondition *dto.UserConfQueryReq, p *middleware.DataPermission) (*models.UserConf, int, error) {
	data := &models.UserConf{}
	err := e.Orm.Scopes(
		cDto.MakeCondition(queryCondition.GetNeedSearch()),
		middleware.Permission(data.TableName(), p),
	).First(data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return data, lang.SuccessCode, nil
}

// Count admin-获取用户配置数据总数
func (e *UserConf) Count(queryCondition *dto.UserConfQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.UserConf{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return 0, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return count, lang.SuccessCode, nil
}

// Insert app-新增用户配置管理
func (e *UserConf) Insert(c *dto.UserConfInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	now := time.Now()
	var data models.UserConf
	data.UserId = c.UserId
	data.CanLogin = c.CanLogin
	data.Status = global.SysStatusOk
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.CreatedAt = &now
	data.UpdatedAt = &now
	err := e.Orm.Create(&data).Error
	if err != nil {
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

// Update app-更新用户配置管理
func (e *UserConf) Update(c *dto.UserConfUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	e.Orm = e.Orm.Begin()
	defer func() {
		if err != nil {
			e.Orm.Rollback()
		} else {
			e.Orm.Commit()
		}
	}()

	//最小化变更改动过的数据
	updates := map[string]interface{}{}
	if c.CanLogin != "" && data.CanLogin != c.CanLogin {
		updates["can_login"] = c.CanLogin
		actionType := constant.UserActionTypeForbLogin
		status := global.SysStatusNotOk
		if c.CanLogin == global.SysStatusOk {
			actionType = constant.UserActionTypeAllowLogin
			status = global.SysStatusOk
		}

		//操作行为日志
		userOperLogService := NewUserOperLogService(&e.Service)
		userOperLogInsertReq := dto.UserOperLogInsertReq{}
		userOperLogInsertReq.UserId = c.UserId
		userOperLogInsertReq.CurrUserId = c.CurrUserId
		userOperLogInsertReq.ActionType = actionType
		_, respCode, err = userOperLogService.Insert(&userOperLogInsertReq)
		if err != nil {
			return false, respCode, err
		}

		//todo 清除已登录用户的session,此处需要配合前端用户接口变更逻辑

		//更改用户状态
		userStatusUpdateReq := dto.UserStatusUpdateReq{}
		userStatusUpdateReq.Status = status
		userStatusUpdateReq.CurrUserId = c.CurrUserId
		userStatusUpdateReq.Id = c.UserId
		userService := NewUserService(&e.Service)
		_, respCode, err = userService.UpdateStatus(&userStatusUpdateReq, p)
		if err != nil {
			return false, respCode, err
		}
	}

	if len(updates) > 0 {
		updates["updated_at"] = time.Now()
		updates["update_by"] = c.CurrUserId
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}
