package service

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	adminService "go-admin/internal/app/admin/sys/service"
	"go-admin/internal/app/biz/user/models"
	"go-admin/internal/app/biz/user/service/dto"
	"go-admin/internal/common/constant"
	clang "go-admin/internal/common/lang"
	"go-admin/pkg/config"
	cDto "go-admin/pkg/dto"
	"go-admin/pkg/dto/service"
	"go-admin/pkg/global"
	"go-admin/pkg/lang"
	"go-admin/pkg/middleware"
	"go-admin/pkg/utils/dateutils"
	"go-admin/pkg/utils/encrypt"
	"go-admin/pkg/utils/strutils"
	"gorm.io/gorm"
	"time"
)

type UserOperLog struct {
	service.Service
}

// NewUserOperLogService biz-实例化用户操作日志
func NewUserOperLogService(s *service.Service) *UserOperLog {
	var srv = new(UserOperLog)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage biz-获取用户操作日志分页列表
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

	err = e.Orm.Preload("User").Order("created_at desc").Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
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
	return list, count, clang.SuccessCode, nil
}

// Get biz-获取用户操作日志详情
func (e *UserOperLog) Get(id int64, p *middleware.DataPermission) (*models.UserOperLog, int, error) {
	if id <= 0 {
		return nil, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	data := &models.UserOperLog{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataNotFoundCode, lang.MsgErr(clang.DataNotFoundCode, e.Lang)
	}
	return data, clang.SuccessCode, nil
}

// QueryOne biz-获取用户操作记录一条记录
func (e *UserOperLog) QueryOne(queryCondition *dto.UserOperLogQueryReq, p *middleware.DataPermission) (*models.UserOperLog, int, error) {
	data := &models.UserOperLog{}
	err := e.Orm.Scopes(
		cDto.MakeCondition(queryCondition.GetNeedSearch()),
		middleware.Permission(data.TableName(), p),
	).First(data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, clang.DataNotFoundCode, lang.MsgErr(clang.DataNotFoundCode, e.Lang)
	}
	return data, clang.SuccessCode, nil
}

// Count admin-获取用户操作记录数据总数
func (e *UserOperLog) Count(queryCondition *dto.UserOperLogQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.UserOperLog{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
		).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, clang.DataNotFoundCode, lang.MsgErr(clang.DataNotFoundCode, e.Lang)
	}
	return count, clang.SuccessCode, nil
}

// Insert admin-新增用户操作记录
func (e *UserOperLog) Insert(c *dto.UserOperLogInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 || c.UserId <= 0 {
		return 0, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.ActionType == "" {
		return 0, clang.UserActionTypeEmptyCode, lang.MsgErr(clang.UserActionTypeEmptyCode, e.Lang)
	}
	if c.UserId <= 0 {
		return 0, clang.UserIdEmptyCode, lang.MsgErr(clang.UserIdEmptyCode, e.Lang)
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
		return 0, clang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataInsertCode, clang.DataInsertLogCode, err)
	}
	return data.Id, clang.SuccessCode, nil
}

// Export biz-导出用户操作日志
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
		actionType := dictDataService.GetLabel("biz_user_action_type", item.ActionType) //行为类型
		byType := dictDataService.GetLabel("biz_user_by_type", item.ByType)             //更新用户类型
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
