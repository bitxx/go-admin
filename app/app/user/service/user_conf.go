package service

import (
	"errors"
	"go-admin/app/app/user/models"
	"go-admin/app/app/user/service/dto"
	"go-admin/config/base/constant"
	baseLang "go-admin/config/base/lang"
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
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}

	cacheUsers := map[int64]*models.User{}
	for _, u := range list {
		if u.User == nil || cacheUsers[u.User.Id] != nil {
			continue
		}
		cacheUsers[u.User.Id] = u.User
		if u.User.Mobile != "" {
			mobile, err := encrypt.AesDecrypt(u.User.Mobile, []byte(config.AuthConfig.Secret))
			if err == nil {
				if c.ShowInfo {
					cacheUsers[u.User.Id].Mobile = mobile
				} else {
					cacheUsers[u.User.Id].Mobile = strutils.HidePartStr(mobile, 3)
				}
			}
		}

		if u.User.Email != "" {
			email, err := encrypt.AesDecrypt(u.User.Email, []byte(config.AuthConfig.Secret))
			if err == nil {
				if c.ShowInfo {
					cacheUsers[u.User.Id].Email = email
				} else {
					cacheUsers[u.User.Id].Email = strutils.HidePartStr(email, 5)
				}
			}
		}
	}

	for index, item := range list {
		if item.User == nil || cacheUsers[item.User.Id] == nil {
			continue
		}
		list[index].User.Mobile = cacheUsers[item.User.Id].Mobile
		list[index].User.Email = cacheUsers[item.User.Id].Email
	}
	return list, count, baseLang.SuccessCode, nil
}

// Get app-获取用户配置管理详情
func (e *UserConf) Get(id int64, p *middleware.DataPermission) (*models.UserConf, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.UserConf{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// QueryOne app-获取用户配置管理一条记录
func (e *UserConf) QueryOne(queryCondition *dto.UserConfQueryReq, p *middleware.DataPermission) (*models.UserConf, int, error) {
	data := &models.UserConf{}
	err := e.Orm.Scopes(
		cDto.MakeCondition(queryCondition.GetNeedSearch()),
		middleware.Permission(data.TableName(), p),
	).First(data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// Count admin-获取用户配置数据总数
func (e *UserConf) Count(queryCondition *dto.UserConfQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.UserConf{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return count, baseLang.SuccessCode, nil
}

// Insert app-新增用户配置管理
func (e *UserConf) Insert(c *dto.UserConfInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
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
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Update app-更新用户配置管理
func (e *UserConf) Update(c *dto.UserConfUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
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
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
		}
		return true, baseLang.SuccessCode, nil
	}
	return false, baseLang.SuccessCode, nil
}
