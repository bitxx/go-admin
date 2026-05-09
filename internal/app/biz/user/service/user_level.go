package service

import (
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	adminService "go-admin/internal/app/admin/sys/service"
	"go-admin/internal/app/biz/user/models"
	"go-admin/internal/app/biz/user/service/dto"
	clang "go-admin/internal/common/lang"
	cDto "go-admin/pkg/dto"
	"go-admin/pkg/dto/service"
	"go-admin/pkg/global"
	"go-admin/pkg/lang"
	"go-admin/pkg/middleware"
	"gorm.io/gorm"
	"time"
)

type UserLevel struct {
	service.Service
}

// NewUserLevelService biz-实例化用户等级管理
func NewUserLevelService(s *service.Service) *UserLevel {
	var srv = new(UserLevel)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage biz-获取用户等级管理分页列表
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
		return nil, 0, clang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataQueryCode, clang.DataQueryLogCode, err)
	}
	return list, count, clang.SuccessCode, nil
}

// Get biz-获取用户等级管理详情
func (e *UserLevel) Get(id int64, p *middleware.DataPermission) (*models.UserLevel, int, error) {
	if id <= 0 {
		return nil, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	data := &models.UserLevel{}
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

// QueryOne biz-获取用户等级管理一条记录
func (e *UserLevel) QueryOne(queryCondition *dto.UserLevelQueryReq, p *middleware.DataPermission) (*models.UserLevel, int, error) {
	data := &models.UserLevel{}
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

// Count admin-获取用户等级管理数据总数
func (e *UserLevel) Count(queryCondition *dto.UserLevelQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.UserLevel{}).
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

// Insert biz-新增用户等级管理详情
func (e *UserLevel) Insert(c *dto.UserLevelInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	if c.LevelType == "" {
		return 0, clang.UserLevelTypeEmptyCode, lang.MsgErr(clang.UserLevelTypeEmptyCode, e.Lang)
	}
	if c.Name == "" {
		return 0, clang.UserLevelNameEmptyCode, lang.MsgErr(clang.UserLevelNameEmptyCode, e.Lang)
	}
	if c.Level <= 0 {
		return 0, clang.UserLevelEmptyCode, lang.MsgErr(clang.UserLevelEmptyCode, e.Lang)
	}

	//若存在等级名称和类型对应的信息，则不可继续添加
	queryReq := dto.UserLevelQueryReq{}
	queryReq.Name = c.Name
	queryReq.LevelType = c.LevelType
	count, respCode, err := e.Count(&queryReq)
	if err != nil && respCode != clang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, clang.UserLevelNameAndTypeExistCode, lang.MsgErr(clang.UserLevelNameAndTypeExistCode, e.Lang)
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
		return 0, clang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataInsertCode, clang.DataInsertLogCode, err)
	}
	return data.Id, clang.SuccessCode, nil
}

// Update biz-更新用户等级管理详情
func (e *UserLevel) Update(c *dto.UserLevelUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}
	req := dto.UserLevelQueryReq{}
	req.Name = c.Name
	req.LevelType = c.LevelType
	resp, respCode, err := e.QueryOne(&req, nil)
	if err != nil && respCode != clang.DataNotFoundCode {
		return false, respCode, err
	}
	if respCode == clang.SuccessCode && resp.Id != data.Id {
		return false, clang.UserLevelNameAndTypeExistCode, lang.MsgErr(clang.UserLevelNameAndTypeExistCode, e.Lang)
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
			return false, clang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataUpdateCode, clang.DataUpdateLogCode, err)
		}
		return true, clang.SuccessCode, nil
	}
	return false, clang.SuccessCode, nil
}

// Delete biz-删除用户等级管理详情
func (e *UserLevel) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return clang.ParamErrCode, lang.MsgErr(clang.ParamErrCode, e.Lang)
	}

	//用户是否使用该等级
	userService := NewUserService(&e.Service)
	userReq := dto.UserQueryReq{}
	userReq.LevelIds = ids
	count, respCode, err := userService.Count(&userReq)
	if err != nil && respCode != clang.DataNotFoundCode {
		return respCode, err
	}
	if count > 0 {
		return clang.UserLevelNameAndTypeExistCode, lang.MsgErr(clang.PluginsCategoryNameHasUsedCode, e.Lang)
	}

	//
	var data models.UserLevel
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return clang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, clang.DataDeleteCode, clang.DataDeleteLogCode, err)
	}
	return clang.SuccessCode, nil
}

// Export biz-导出用户等级管理详情
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
