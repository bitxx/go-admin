package service

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	adminService "go-admin/app/admin/service"
	"go-admin/app/app/user/constant"
	appLang "go-admin/app/app/user/lang"
	"go-admin/app/app/user/models"
	"go-admin/app/app/user/service/dto"
	"go-admin/common/core/service"
	cDto "go-admin/common/dto"
	"go-admin/common/global"
	"go-admin/common/middleware"
	"go-admin/common/utils/dateUtils"
	"go-admin/common/utils/encrypt"
	"go-admin/common/utils/strutils"
	"go-admin/config/config"

	"go-admin/config/lang"
	"gorm.io/gorm"
	"time"
)

type UserOperLog struct {
	service.Service
}

// NewUserOperLogService
// @Description: 实例化UserOperLog
// @param s
// @return *UserOperLog
func NewUserOperLogService(s *service.Service) *UserOperLog {
	var srv = new(UserOperLog)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage
// @Description: 获取UserOperLog列表
// @receiver e
// @param c
// @param p
// @return []models.UserOperLog
// @return int64
// @return int
// @return error
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

// Get
// @Description: 获取UserOperLog对象
// @receiver e
// @param id 编号
// @param p
// @return *models.UserOperLog
// @return int
// @return error
func (e *UserOperLog) Get(id int64, p *middleware.DataPermission) (*models.UserOperLog, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.UserOperLog{}
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

// QueryOne
// @Description: 通过自定义条件获取UserOperLog一条记录
// @receiver e
// @param queryCondition 条件
// @return *models.UserOperLog
// @return error
func (e *UserOperLog) QueryOne(queryCondition *dto.UserOperLogQueryReq, p *middleware.DataPermission) (*models.UserOperLog, int, error) {
	data := &models.UserOperLog{}
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

// Count
//
//	@Description: 获取条数
//	@receiver e
//	@param c
//	@return int64
//	@return int
//	@return error
func (e *UserOperLog) Count(queryCondition *dto.UserOperLogQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.UserOperLog{}).
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

// Insert
// @Description: 创建UserOperLog对象
// @receiver e
// @param c
// @return int64 插入数据的主键
// @return int
// @return error
func (e *UserOperLog) Insert(c *dto.UserOperLogInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 || c.UserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.ActionType == "" {
		return 0, appLang.AppUserActionTypeEmptyCode, lang.MsgErr(appLang.AppUserActionTypeEmptyCode, e.Lang)
	}
	if c.UserId <= 0 {
		return 0, appLang.AppUserIdEmptyCode, lang.MsgErr(appLang.AppUserIdEmptyCode, e.Lang)
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
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

// GetExcel
// @Description: GetExcel 导出UserOperLog excel数据
// @receiver e
// @param list
// @return []byte
// @return int
// @return error
func (e *UserOperLog) GetExcel(list []models.UserOperLog) ([]byte, error) {
	sheetName := "UserOperLog"
	xlsx := excelize.NewFile()
	no := xlsx.NewSheet(sheetName)
	xlsx.SetColWidth(sheetName, "A", "G", 25)
	xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
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
		xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, email, mobile, userName, actionType, byType, item.UpdateBy, dateUtils.ConvertToStr(*item.UpdatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
