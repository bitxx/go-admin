package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	adminService "go-admin/app/admin/sys/service"
	aLang "go-admin/app/app/user/lang"
	"go-admin/app/app/user/models"
	"go-admin/app/app/user/service/dto"
	cLang "go-admin/app/plugins/content/lang"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"gorm.io/gorm"
	"time"
)

type UserLevel struct {
	service.Service
}

// NewUserLevelService app-实例化用户等级管理
func NewUserLevelService(s *service.Service) *UserLevel {
	var srv = new(UserLevel)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage app-获取用户等级管理分页列表
func (e *UserLevel) GetPage(c *dto.UserLevelQueryReq, p *middleware.DataPermission) ([]models.UserLevel, int64, int, error) {
	var data models.UserLevel
	var list []models.UserLevel
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

// Get app-获取用户等级管理详情
func (e *UserLevel) Get(id int64, p *middleware.DataPermission) (*models.UserLevel, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.UserLevel{}
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

// QueryOne app-获取用户等级管理一条记录
func (e *UserLevel) QueryOne(queryCondition *dto.UserLevelQueryReq, p *middleware.DataPermission) (*models.UserLevel, int, error) {
	data := &models.UserLevel{}
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

// Count admin-获取用户等级管理数据总数
func (e *UserLevel) Count(queryCondition *dto.UserLevelQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.UserLevel{}).
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

// Insert app-新增用户等级管理详情
func (e *UserLevel) Insert(c *dto.UserLevelInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.LevelType == "" {
		return 0, aLang.UserLevelTypeEmptyCode, lang.MsgErr(aLang.UserLevelTypeEmptyCode, e.Lang)
	}
	if c.Name == "" {
		return 0, aLang.UserLevelNameEmptyCode, lang.MsgErr(aLang.UserLevelNameEmptyCode, e.Lang)
	}
	if c.Level <= 0 {
		return 0, aLang.UserLevelEmptyCode, lang.MsgErr(aLang.UserLevelEmptyCode, e.Lang)
	}

	//若存在等级名称和类型对应的信息，则不可继续添加
	queryReq := dto.UserLevelQueryReq{}
	queryReq.Name = c.Name
	queryReq.LevelType = c.LevelType
	count, respCode, err := e.Count(&queryReq)
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, aLang.UserLevelNameAndTypeExistCode, lang.MsgErr(aLang.UserLevelNameAndTypeExistCode, e.Lang)
	}

	now := time.Now()
	var data models.UserLevel
	data.Name = c.Name
	data.LevelType = c.LevelType
	data.Level = c.Level
	data.Status = global.SysStatusOk
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

// Update app-更新用户等级管理详情
func (e *UserLevel) Update(c *dto.UserLevelUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}
	req := dto.UserLevelQueryReq{}
	req.Name = c.Name
	req.LevelType = c.LevelType
	resp, respCode, err := e.QueryOne(&req, nil)
	if err != nil && respCode != lang.DataNotFoundCode {
		return false, respCode, err
	}
	if respCode == lang.SuccessCode && resp.Id != data.Id {
		return false, aLang.UserLevelNameAndTypeExistCode, lang.MsgErr(aLang.UserLevelNameAndTypeExistCode, e.Lang)
	}

	//最小化变更改动过的数据
	updates := map[string]interface{}{}
	if c.Name != "" && data.Name != c.Name {
		updates["name"] = c.Name
	}
	if c.LevelType != "" && data.LevelType != c.LevelType {
		updates["level_type"] = c.LevelType
	}
	if c.Level > 0 && data.Level != c.Level {
		updates["level"] = c.Level
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

// Delete app-删除用户等级管理详情
func (e *UserLevel) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}

	//用户是否使用该等级
	userService := NewUserService(&e.Service)
	userReq := dto.UserQueryReq{}
	userReq.LevelIds = ids
	count, respCode, err := userService.Count(&userReq)
	if err != nil && respCode != lang.DataNotFoundCode {
		return respCode, err
	}
	if count > 0 {
		return aLang.UserLevelNameAndTypeExistCode, lang.MsgErr(cLang.PluginsCategoryNameHasUsedCode, e.Lang)
	}

	//
	var data models.UserLevel
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// Export app-导出用户等级管理详情
func (e *UserLevel) Export(list []models.UserLevel) ([]byte, error) {
	sheetName := "UserLevel"
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
