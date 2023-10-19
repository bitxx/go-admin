package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	adminService "go-admin/app/admin/service"
	"go-admin/app/app/user/models"
	"go-admin/app/app/user/service/dto"
	cDto "go-admin/common/dto"
	"go-admin/common/dto/service"
	"go-admin/common/middleware"
	"go-admin/common/utils/encrypt"
	"go-admin/common/utils/strutils"
	"go-admin/config/config"
	"go-admin/config/lang"
	"gorm.io/gorm"
)

type UserAccountLog struct {
	service.Service
}

// NewUserAccountLogService
// @Description: 实例化UserAccountLog
// @param s
// @return *UserAccountLog
func NewUserAccountLogService(s *service.Service) *UserAccountLog {
	var srv = new(UserAccountLog)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage
// @Description: 获取UserAccountLog列表
// @receiver e
// @param c
// @param p
// @return []models.UserAccountLog
// @return int64
// @return int
// @return error
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
// @Description: 获取UserAccountLog对象
// @receiver e
// @param id 编号
// @param p
// @return *models.UserAccountLog
// @return int
// @return error
func (e *UserAccountLog) Get(id int64, p *middleware.DataPermission) (*models.UserAccountLog, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.UserAccountLog{}
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
// @Description: 通过自定义条件获取UserAccountLog一条记录
// @receiver e
// @param queryCondition 条件
// @return *models.UserAccountLog
// @return error
func (e *UserAccountLog) QueryOne(queryCondition *dto.UserAccountLogQueryReq, p *middleware.DataPermission) (*models.UserAccountLog, int, error) {
	data := &models.UserAccountLog{}
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
func (e *UserAccountLog) Count(queryCondition *dto.UserAccountLogQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.UserAccountLog{}).
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

// GetExcel
// @Description: GetExcel 导出UserAccountLog excel数据
// @receiver e
// @param list
// @return []byte
// @return int
// @return error
func (e *UserAccountLog) GetExcel(list []models.UserAccountLog) ([]byte, error) {
	sheetName := "UserAccountLog"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	_ = xlsx.SetColWidth(sheetName, "A", "L", 25)
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"编号", "状态"})
	dictService := adminService.NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		status := dictService.GetLabel("sys_status", item.Status)

		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, status,
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
