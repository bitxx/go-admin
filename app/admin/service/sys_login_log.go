package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/core/service"
	cDto "go-admin/common/dto"
	"go-admin/common/middleware"
	"go-admin/common/utils/dateutils"
	"go-admin/config/lang"
	"gorm.io/gorm"
)

type SysLoginLog struct {
	service.Service
}

func NewSysLoginLogService(s *service.Service) *SysLoginLog {
	var srv = new(SysLoginLog)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetSysLoginLogPage 获取SysLoginLog列表
func (e *SysLoginLog) GetPage(c *dto.SysLoginLogQueryReq, p *middleware.DataPermission) ([]models.SysLoginLog, int64, int, error) {
	var list []models.SysLoginLog
	var data models.SysLoginLog
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

// Get 获取SysLoginLog对象
func (e *SysLoginLog) Get(id int64, p *middleware.DataPermission) (*models.SysLoginLog, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.SysLoginLog{}
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

// Remove 删除SysLoginLog
func (e *SysLoginLog) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var err error
	var data models.SysLoginLog
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetExcel 导出Category
func (e *SysLoginLog) GetExcel(list []models.SysLoginLog) ([]byte, error) {
	//sheet名称
	sheetName := "LoginLog"
	xlsx := excelize.NewFile()
	no, _ := xlsx.NewSheet(sheetName)
	//各列间隔
	_ = xlsx.SetColWidth(sheetName, "A", "P", 25)
	//头部描述
	_ = xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"编号", "用户编号", "日志状态", "ip地址", "归属地", "代理", "浏览器", "系统", "固件", "登录时间", "备注"})
	dictService := NewSysDictDataService(&e.Service)

	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		loginLogStatus := dictService.GetLabel("sys_loginlog_status", item.Status) //平台
		//按标签对应输入数据
		_ = xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.UserId, loginLogStatus, item.Ipaddr, item.LoginLocation, item.Agent,
			item.Browser, item.Os, item.Platform, dateutils.ConvertToStrByPrt(item.LoginTime, -1), item.Remark,
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}
