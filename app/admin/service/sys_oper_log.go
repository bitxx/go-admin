package service

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/core/service"
	cDto "go-admin/common/dto"
	"go-admin/common/middleware"
	"go-admin/config/lang"
	"gorm.io/gorm"
)

type SysOperLog struct {
	service.Service
}

func NewSysOperLogService(s *service.Service) *SysOperLog {
	var srv = new(SysOperLog)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage 获取SysOperLog列表
func (e *SysOperLog) GetPage(c *dto.SysOperLogQueryReq, p *middleware.DataPermission) ([]models.SysOperLog, int64, int, error) {
	var list []models.SysOperLog
	var data models.SysOperLog
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

// Get 获取SysOperLog对象
func (e *SysOperLog) Get(id int64, p *middleware.DataPermission) (*models.SysOperLog, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.SysOperLog{}
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

// Remove 删除SysOperLog
func (e *SysOperLog) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var err error
	var data models.SysOperLog
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetExcel 导出OperLog
func (e *SysOperLog) GetExcel(list []models.SysOperLog) ([]byte, error) {
	//sheet名称
	sheetName := "OperLog"
	xlsx := excelize.NewFile()
	no := xlsx.NewSheet(sheetName)
	//各列间隔
	xlsx.SetColWidth(sheetName, "A", "J", 25)
	//头部描述
	xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"日志编号", "用户编号", "请求方法", "请求地址", "请求IP", "访问位置", "返回码",
		"耗时", "用户代理", "操作时间"})

	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		//按标签对应输入数据
		xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.UserId, item.RequestMethod, item.OperUrl, item.OperIp,
			item.OperLocation, item.Status, item.LatencyTime, item.UserAgent, item.OperTime,
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
