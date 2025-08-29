package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	adminService "go-admin/app/admin/sys/service"
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
	"go-admin/core/utils/dateutils"
	"go-admin/core/utils/encrypt"
	"go-admin/core/utils/strutils"
	"gorm.io/gorm"
	"time"
)

type UserOperLog struct {
	service.Service
}

// NewUserOperLogService app-实例化用户操作日志
func NewUserOperLogService(s *service.Service) *UserOperLog {
	var srv = new(UserOperLog)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage app-获取用户操作日志分页列表
func (e *UserOperLog) GetPage(c *dto.UserOperLogQueryReq, p *middleware.DataPermission) ([]models.UserOperLog, int64, int, error) {
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

	var data models.UserOperLog
	var list []models.UserOperLog
	var count int64

	err = e.Orm.Joins("User").Order("created_at desc").Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
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
	return list, count, baseLang.SuccessCode, nil
}

// Get app-获取用户操作日志详情
func (e *UserOperLog) Get(id int64, p *middleware.DataPermission) (*models.UserOperLog, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.UserOperLog{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// QueryOne app-获取用户操作记录一条记录
func (e *UserOperLog) QueryOne(queryCondition *dto.UserOperLogQueryReq, p *middleware.DataPermission) (*models.UserOperLog, int, error) {
	data := &models.UserOperLog{}
	err := e.Orm.Scopes(
		cDto.MakeCondition(queryCondition.GetNeedSearch()),
		middleware.Permission(data.TableName(), p),
	).First(data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// Count admin-获取用户操作记录数据总数
func (e *UserOperLog) Count(queryCondition *dto.UserOperLogQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.UserOperLog{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return 0, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return count, baseLang.SuccessCode, nil
}

// Insert admin-新增用户操作记录
func (e *UserOperLog) Insert(c *dto.UserOperLogInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 || c.UserId <= 0 {
		return 0, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	if c.ActionType == "" {
		return 0, baseLang.UserActionTypeEmptyCode, lang.MsgErr(baseLang.UserActionTypeEmptyCode, e.Lang)
	}
	if c.UserId <= 0 {
		return 0, baseLang.UserIdEmptyCode, lang.MsgErr(baseLang.UserIdEmptyCode, e.Lang)
	}
	now := time.Now()
	var data models.UserOperLog
	data.UserId = c.UserId
	data.ActionType = c.ActionType
	data.ByType = constant.UserByTypeBack
	data.CreateBy = c.CurrUserId
	data.CreatedAt = &now
	data.Status = global.SysStatusOk
	data.UpdateBy = c.CurrUserId
	data.UpdatedAt = &now
	err := e.Orm.Create(&data).Error
	if err != nil {
		return 0, baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return data.Id, baseLang.SuccessCode, nil
}

// Export app-导出用户操作日志
func (e *UserOperLog) Export(list []models.UserOperLog) ([]byte, error) {
	sheetName := "UserOperLog"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "G", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"编号", "用户邮箱", "用户手机号", "昵称", "用户行为类型", "更新用户类型", "更新者编号", "更新时间"})
	var dictDataService = adminService.NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		email := ""
		mobile := ""
		userName := ""                                                                  //真实姓名
		actionType := dictDataService.GetLabel("app_user_action_type", item.ActionType) //行为类型
		byType := dictDataService.GetLabel("app_user_by_type", item.ByType)             //更新用户类型
		if item.User != nil {
			email = item.User.Email
			mobile = item.User.Mobile
		}
		if item.User != nil {
			userName = item.User.UserName
		}

		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, email, mobile, userName, actionType, byType, item.UpdateBy, dateutils.ConvertToStr(*item.UpdatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
