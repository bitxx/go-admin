package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"go-admin/common/dto/service"
	"go-admin/common/middleware"
	"go-admin/common/utils/dateutils"
	"go-admin/config/lang"
	"gorm.io/gorm"
	"time"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	cDto "go-admin/common/dto"
)

type SysApi struct {
	service.Service
}

func NewSysApiService(s *service.Service) *SysApi {
	var srv = new(SysApi)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage 获取SysApi列表
func (e *SysApi) GetPage(c *dto.SysApiQueryReq, p *middleware.DataPermission) ([]models.SysApi, int64, int, error) {
	var list []models.SysApi
	var data models.SysApi
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

// Get 获取SysApi对象with id
func (e *SysApi) Get(id int64, p *middleware.DataPermission) (*models.SysApi, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.SysApi{}
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

func (e *SysApi) QueryOne(queryCondition *dto.SysApiQueryReq, p *middleware.DataPermission) (*models.SysApi, int, error) {
	data := &models.SysApi{}
	err := e.Orm.Model(&models.SysApi{}).
		Scopes(
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

// Update 修改SysApi对象
func (e *SysApi) Update(c *dto.SysApiUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	updates := map[string]interface{}{}
	if c.Title != "" && data.Title != c.Title {
		updates["title"] = c.Title
	}
	if c.ApiType != "" && data.ApiType != c.ApiType {
		updates["api_type"] = c.ApiType
	}
	/*	if data.Action != c.Action {
		updates["action"] = c.Action
	}*/
	if len(updates) > 0 {
		updates["update_by"] = c.CurrUserId
		updates["updated_at"] = time.Now()
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// Remove 删除SysApi
func (e *SysApi) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var err error
	var data models.SysApi
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetExcel 导出SysApi
func (e *SysApi) GetExcel(list []models.SysApi) ([]byte, error) {
	//sheet名称
	sheetName := "Api"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	//各列间隔
	_ = xlsx.SetColWidth(sheetName, "A", "F", 25)
	//头部描述
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"编号", "标题", "请求地址", "请求方法", "请求类型", "创建时间"})

	dictService := NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		action := dictService.GetLabel("sys_api_action", item.Action)
		apiType := dictService.GetLabel("sys_api_type", item.ApiType)
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.Title, item.Path, action, apiType, dateutils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
