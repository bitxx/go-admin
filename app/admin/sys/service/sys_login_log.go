package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"go-admin/app/admin/sys/models"
	"go-admin/app/admin/sys/service/dto"
	baseLang "go-admin/config/base/lang"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/utils/dateutils"
	"gorm.io/gorm"
)

type SysLoginLog struct {
	service.Service
}

// NewSysLoginLogService admin-实例化登录日志
func NewSysLoginLogService(s *service.Service) *SysLoginLog {
	var srv = new(SysLoginLog)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage admin-获取登录日志分页列表
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
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

// Get admin-获取登录日志详情
func (e *SysLoginLog) Get(id int64, p *middleware.DataPermission) (*models.SysLoginLog, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.SysLoginLog{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// Delete admin-删除登录日志
func (e *SysLoginLog) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	var err error
	var data models.SysLoginLog
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// Export admin-导出登录日志
func (e *SysLoginLog) Export(list []models.SysLoginLog) ([]byte, error) {
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
		loginLogStatus := dictService.GetLabel("admin_sys_loginlog_status", item.Status) //平台
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
