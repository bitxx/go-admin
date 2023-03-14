package service

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	adminService "go-admin/app/admin/service"
	uLang "go-admin/app/app/user/lang"
	"go-admin/app/app/user/models"
	"go-admin/app/app/user/service/dto"
	"go-admin/common/core/service"
	cDto "go-admin/common/dto"
	"go-admin/common/middleware"

	"go-admin/config/lang"
	"gorm.io/gorm"
	"time"
)

type UserCountryCode struct {
	service.Service
}

// NewUserCountryCodeService
// @Description: 实例化UserCountryCode
// @param s
// @return *UserCountryCode
func NewUserCountryCodeService(s *service.Service) *UserCountryCode {
	var srv = new(UserCountryCode)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage
// @Description: 获取UserCountryCode列表
// @receiver e
// @param c
// @param p
// @return []models.UserCountryCode
// @return int64
// @return int
// @return error
func (e *UserCountryCode) GetPage(c *dto.UserCountryCodeQueryReq, p *middleware.DataPermission) ([]models.UserCountryCode, int64, int, error) {
	var data models.UserCountryCode
	var list []models.UserCountryCode
	var count int64

	err := e.Orm.Order("created_at desc").Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, count, lang.SuccessCode, nil
}

// Get
// @Description: 获取UserCountryCode对象
// @receiver e
// @param id 编号
// @param p
// @return *models.UserCountryCode
// @return int
// @return error
func (e *UserCountryCode) Get(id int64, p *middleware.DataPermission) (*models.UserCountryCode, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.UserCountryCode{}
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
// @Description: 通过自定义条件获取UserCountryCode一条记录
// @receiver e
// @param queryCondition 条件
// @return *models.UserCountryCode
// @return error
func (e *UserCountryCode) QueryOne(queryCondition *dto.UserCountryCodeQueryReq, p *middleware.DataPermission) (*models.UserCountryCode, int, error) {
	data := &models.UserCountryCode{}
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
func (e *UserCountryCode) Count(queryCondition *dto.UserCountryCodeQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.UserCountryCode{}).
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
// @Description: 创建UserCountryCode对象
// @receiver e
// @param c
// @return int64 插入数据的主键
// @return int
// @return error
func (e *UserCountryCode) Insert(c *dto.UserCountryCodeInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Country == "" {
		return 0, uLang.AppUserCountryEmptyCode, lang.MsgErr(uLang.AppUserCountryEmptyCode, e.Lang)
	}
	if c.Code == "" {
		return 0, uLang.AppUserCountryCodeEmptyCode, lang.MsgErr(uLang.AppUserCountryCodeEmptyCode, e.Lang)
	}
	if c.Status == "" {
		return 0, uLang.AppUserCountryStatusEmptyCode, lang.MsgErr(uLang.AppUserCountryStatusEmptyCode, e.Lang)
	}

	//检测国家名称是否存在
	reqName := dto.UserCountryCodeQueryReq{}
	reqName.CountryInner = c.Country
	count, respCode, err := e.Count(&reqName)
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, uLang.AppUserCountryHasExistCode, lang.MsgErr(uLang.AppUserCountryHasExistCode, e.Lang)
	}

	//检测国家区号是否存在
	reqCode := dto.UserCountryCodeQueryReq{}
	reqCode.Code = c.Code
	count, respCode, err = e.Count(&reqCode)
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, uLang.AppUserCountryCodeHasExistCode, lang.MsgErr(uLang.AppUserCountryCodeHasExistCode, e.Lang)
	}

	now := time.Now()
	var data models.UserCountryCode
	data.Country = c.Country
	data.Code = c.Code
	data.Status = c.Status
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.CreatedAt = &now
	data.UpdatedAt = &now
	err = e.Orm.Create(&data).Error
	if err != nil {
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

// Update
// @Description: 修改UserCountryCode对象
// @receiver e
// @param c
// @param p
// @return bool 是否有数据更新
// @return error
func (e *UserCountryCode) Update(c *dto.UserCountryCodeUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Country == "" {
		return false, uLang.AppUserCountryEmptyCode, lang.MsgErr(uLang.AppUserCountryEmptyCode, e.Lang)
	}
	if c.Code == "" {
		return false, uLang.AppUserCountryCodeEmptyCode, lang.MsgErr(uLang.AppUserCountryCodeEmptyCode, e.Lang)
	}
	if c.Status == "" {
		return false, uLang.AppUserCountryStatusEmptyCode, lang.MsgErr(uLang.AppUserCountryStatusEmptyCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	//最小化变更改动过的数据
	updates := map[string]interface{}{}
	if c.Country != "" && data.Country != c.Country {
		req := dto.UserCountryCodeQueryReq{}
		req.CountryInner = c.Country
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, uLang.AppUserCountryHasExistCode, lang.MsgErr(uLang.AppUserCountryHasExistCode, e.Lang)
		}
		updates["country"] = c.Country
	}
	if c.Code != "" && data.Code != c.Code {
		req := dto.UserCountryCodeQueryReq{}
		req.Code = c.Code
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, uLang.AppUserCountryCodeHasExistCode, lang.MsgErr(uLang.AppUserCountryCodeHasExistCode, e.Lang)
		}
		updates["code"] = c.Code
	}
	if c.Status != "" && data.Status != c.Status {
		updates["status"] = c.Status
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

// Remove
// @Description: 删除UserCountryCode
// @receiver e
// @param ids
// @param p
// @return int
// @return error
func (e *UserCountryCode) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var data models.UserCountryCode
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetExcel
// @Description: GetExcel 导出UserCountryCode excel数据
// @receiver e
// @param list
// @return []byte
// @return int
// @return error
func (e *UserCountryCode) GetExcel(list []models.UserCountryCode) ([]byte, error) {
	sheetName := "UserCountryCode"
	xlsx := excelize.NewFile()
	no := xlsx.NewSheet(sheetName)
	xlsx.SetColWidth(sheetName, "A", "L", 25)
	xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"编号", "状态"})
	dictService := adminService.NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		status := dictService.GetLabel("sys_status", item.Status)

		//按标签对应输入数据
		xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, status,
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
