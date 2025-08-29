package service

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	adminService "go-admin/app/admin/sys/service"
	"go-admin/app/app/user/models"
	"go-admin/app/app/user/service/dto"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/config"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/utils/encrypt"
	"go-admin/core/utils/strutils"
	"gorm.io/gorm"
)

type UserAccountLog struct {
	service.Service
}

// NewUserAccountLogService app-实例化账变记录
func NewUserAccountLogService(s *service.Service) *UserAccountLog {
	var srv = new(UserAccountLog)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage app-获取账变记录分页列表
func (e *UserAccountLog) GetPage(c *dto.UserAccountLogQueryReq, p *middleware.DataPermission) ([]models.UserAccountLog, int64, int, error) {
	var data models.UserAccountLog
	var list []models.UserAccountLog
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
		list[index].User.Mobile = cacheUsers[item.User.Id].Mobile
		list[index].User.Email = cacheUsers[item.User.Id].Email
	}
	return list, count, baseLang.SuccessCode, nil
}

// Get app-获取账变记录详情
func (e *UserAccountLog) Get(id int64, p *middleware.DataPermission) (*models.UserAccountLog, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.UserAccountLog{}
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

// QueryOne app-获取账变记录一条记录
func (e *UserAccountLog) QueryOne(queryCondition *dto.UserAccountLogQueryReq, p *middleware.DataPermission) (*models.UserAccountLog, int, error) {
	data := &models.UserAccountLog{}
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

// Count admin-获取账变记录数据总数
func (e *UserAccountLog) Count(queryCondition *dto.UserAccountLogQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.UserAccountLog{}).
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

// Export app-导出账变记录
func (e *UserAccountLog) Export(list []models.UserAccountLog) ([]byte, error) {
	sheetName := "UserAccountLog"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "L", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"编号", "状态"})
	dictService := adminService.NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		status := dictService.GetLabel("admin_sys_status", item.Status)

		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, status,
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
