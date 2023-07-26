package service

import (
	"errors"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/google/uuid"
	adminService "go-admin/app/admin/service"
	"go-admin/app/plugins/filemgr/constant"
	fLang "go-admin/app/plugins/filemgr/lang"
	"go-admin/app/plugins/filemgr/models"
	"go-admin/app/plugins/filemgr/service/dto"
	"go-admin/common/core/service"
	cDto "go-admin/common/dto"
	"go-admin/common/middleware"
	"go-admin/common/utils/dateUtils"
	"go-admin/common/utils/fileUtils/ossUtils"
	"go-admin/config/config"
	"mime/multipart"
	"path"

	"go-admin/config/lang"
	"gorm.io/gorm"
	"time"
)

type FilemgrApp struct {
	service.Service
}

// NewFilemgrAppService
// @Description: 实例化FilemgrApp
// @param s
// @return *FilemgrApp
func NewFilemgrAppService(s *service.Service) *FilemgrApp {
	var srv = new(FilemgrApp)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage
// @Description: 获取FilemgrApp列表
// @receiver e
// @param c
// @param p
// @return []models.FilemgrApp
// @return int64
// @return int
// @return error
func (e *FilemgrApp) GetPage(c *dto.FilemgrAppQueryReq, p *middleware.DataPermission) ([]models.FilemgrApp, int64, int, error) {
	var data models.FilemgrApp
	var list []models.FilemgrApp
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

	for index, item := range list {
		if item.DownloadType == constant.AppDownloadTypeOss {
			url := ""
			url, err = e.generateAppOssUrl(&item)
			if err == nil {
				item.DownloadUrl = url
			}
		}
		list[index] = item
	}
	return list, count, lang.SuccessCode, nil
}

// Get
// @Description: 获取FilemgrApp对象
// @receiver e
// @param id 编号
// @param p
// @return *models.FilemgrApp
// @return int
// @return error
func (e *FilemgrApp) Get(id int64, p *middleware.DataPermission) (*models.FilemgrApp, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.FilemgrApp{}
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
// @Description: 通过自定义条件获取FilemgrApp一条记录
// @receiver e
// @param queryCondition 条件
// @return *models.FilemgrApp
// @return error
func (e *FilemgrApp) QueryOne(queryCondition *dto.FilemgrAppQueryReq, p *middleware.DataPermission) (*models.FilemgrApp, int, error) {
	data := &models.FilemgrApp{}
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
func (e *FilemgrApp) Count(queryCondition *dto.FilemgrAppQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.FilemgrApp{}).
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
// @Description: 创建FilemgrApp对象
// @receiver e
// @param c
// @return int64 插入数据的主键
// @return int
// @return error
func (e *FilemgrApp) Insert(c *dto.FilemgrAppInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.Platform == "" {
		return 0, fLang.PluginsAppPlatformEmptyCode, lang.MsgErr(fLang.PluginsAppPlatformEmptyCode, e.Lang)
	}
	if c.Version == "" {
		return 0, fLang.PluginsAppVersionEmptyCode, lang.MsgErr(fLang.PluginsAppVersionEmptyCode, e.Lang)
	}
	if c.DownloadType == "" {
		return 0, fLang.PluginsAppDownloadTypeEmptyCode, lang.MsgErr(fLang.PluginsAppDownloadTypeEmptyCode, e.Lang)
	}
	if c.DownloadType == constant.AppDownloadTypeOss || c.DownloadType == constant.AppDownloadTypeLocal {
		if c.AppType == "" {
			return 0, fLang.PluginsAppTypeCode, lang.MsgErr(fLang.PluginsAppTypeCode, e.Lang)
		}
		if c.LocalAddress == "" {
			return 0, fLang.PluginsAppUploadCode, lang.MsgErr(fLang.PluginsAppUploadCode, e.Lang)
		}
	}
	if c.Remark == "" {
		return 0, fLang.PluginsAppRemarkCode, lang.MsgErr(fLang.PluginsAppRemarkCode, e.Lang)
	}

	query := dto.FilemgrAppQueryReq{}
	query.Platform = c.Platform
	query.AppType = c.AppType
	query.Version = c.Version
	count, respCode, err := e.Count(&query)
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, fLang.PluginsAppExistCode, lang.MsgErr(fLang.PluginsAppExistCode, e.Lang)
	}

	//oss上传
	if c.DownloadType == constant.AppDownloadTypeOss {
		err = e.uploadOssFile(c.AppType, c.Version, c.Platform, c.LocalAddress)
		if err != nil {
			return 0, fLang.PluginsAppOssUploadLogCode, lang.MsgLogErrf(e.Log, e.Lang, fLang.PluginsAppUploadCode, fLang.PluginsAppOssUploadLogCode, err)
		}
	}
	if c.DownloadType == constant.AppDownloadTypeLocal {
		if c.LocalRootUrl == "" {
			return 0, fLang.PluginsAppLocalUrlEmptyCode, lang.MsgErr(fLang.PluginsAppLocalUrlEmptyCode, e.Lang)
		}
		//c.LocalAddress = strings.Replace(c.LocalAddress, config.ApplicationConfig.FileRootPath, "", -1)
		c.DownloadUrl = c.LocalRootUrl + c.LocalAddress
	}

	now := time.Now()
	var data models.FilemgrApp
	data.Version = c.Version
	data.Platform = c.Platform
	data.AppType = c.AppType
	data.LocalAddress = c.LocalAddress
	data.DownloadNum = 0
	data.DownloadType = c.DownloadType
	data.DownloadUrl = c.DownloadUrl
	data.Remark = c.Remark
	data.Status = constant.AppPublishWait
	data.CreateBy = c.CurrUserId
	data.CreatedAt = &now
	data.UpdateBy = c.CurrUserId
	data.UpdatedAt = &now
	err = e.Orm.Create(&data).Error
	if err != nil {
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

// Update
// @Description: 修改FilemgrApp对象
// @receiver e
// @param c
// @param p
// @return bool 是否有数据更新
// @return error
func (e *FilemgrApp) Update(c *dto.FilemgrAppUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	//最小化变更改动过的数据
	updates := map[string]interface{}{}
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
// @Description: 删除FilemgrApp
// @receiver e
// @param ids
// @param p
// @return int
// @return error
func (e *FilemgrApp) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	for _, id := range ids {
		result, respCode, err := e.Get(id, p)
		if respCode != lang.DataNotFoundCode && err != nil {
			return respCode, err
		}

		//同一个完全相同的版本，可能因为网路有多条记录，但这些记录都指向一个oss资源，此时只有最后一条记录，才能删除oss资源
		query := dto.FilemgrAppQueryReq{}
		query.Platform = result.Platform
		query.AppType = result.AppType
		query.Version = result.Version
		var count int64
		count, respCode, err = e.Count(&query)
		if respCode != lang.DataNotFoundCode && err != nil {
			return respCode, err
		}
		if count <= 1 {
			//oss删除对应资源,无论删除成功与否
			objectKey, _ := e.generateAppOssObjectKey(result)
			oss, _ := e.getOssClient()
			if oss != nil && objectKey != "" {
				_ = oss.Bucket.DeleteObject(objectKey, nil)
			}
		}
	}
	var data models.FilemgrApp
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// GetSingleUploadFileInfo
//
//	@Description: 获取单个上传文件信息
//	@receiver e
//	@param form
//	@param file
//	@param dst
//	@return int
//	@return error
func (e *FilemgrApp) GetSingleUploadFileInfo(form *multipart.Form, file *multipart.FileHeader, dst *string) (int, error) {
	if len(form.File) != 1 {
		return fLang.PluginsAppSelectOneFileUploadCode, lang.MsgErr(fLang.PluginsAppSelectOneFileUploadCode, e.Lang)
	}
	for _, files := range form.File {
		if len(files) != 1 {
			return fLang.PluginsAppSelectOneFileUploadCode, lang.MsgErr(fLang.PluginsAppSelectOneFileUploadCode, e.Lang)
		}
		for _, item := range files {
			*dst = config.ApplicationConfig.FileRootPath + "app/" + uuid.New().String() + path.Ext(item.Filename)
			*file = *item
			return lang.SuccessCode, nil
		}
	}
	return lang.SuccessCode, nil
}

// GetExcel
// @Description: GetExcel 导出FilemgrApp excel数据
// @receiver e
// @param list
// @return []byte
// @return int
// @return error
func (e *FilemgrApp) GetExcel(list []models.FilemgrApp) ([]byte, error) {
	sheetName := "FilemgrApp"
	xlsx := excelize.NewFile()
	no := xlsx.NewSheet(sheetName)
	xlsx.SetColWidth(sheetName, "A", "J", 25)
	xlsx.SetSheetRow(sheetName, "A1", &[]interface{}{
		"App编号", "版本号", "系统平台", "App类型", "下载类型", "发布状态", "下载地址", "服务器本地地址",
		"更新内容", "创建时间"})
	dictService := adminService.NewSysDictDataService(&e.Service)
	for i, item := range list {
		axis := fmt.Sprintf("A%d", i+2)
		platform := dictService.GetLabel("plugin_filemgr_app_platform", item.Platform)              //平台
		appType := dictService.GetLabel("plugin_filemgr_app_type", item.AppType)                    //app类型
		downloadType := dictService.GetLabel("plugin_filemgr_app_download_type", item.DownloadType) //下载类型
		publishStatus := dictService.GetLabel("plugin_filemgr_publish_status", item.Status)         //下载类型
		//按标签对应输入数据
		xlsx.SetSheetRow(sheetName, axis, &[]interface{}{
			item.Id, item.Version, platform, appType, downloadType, publishStatus, item.DownloadUrl, item.Remark, dateUtils.ConvertToStrByPrt(item.CreatedAt, -1),
		})
	}
	xlsx.SetActiveSheet(no)
	data, _ := xlsx.WriteToBuffer()
	return data.Bytes(), nil
}

// uploadOssFile
// @Description:
// @receiver e
// @param appType
// @param version
// @param platform
// @param localAddress
// @return error
func (e *FilemgrApp) uploadOssFile(appType, version, platform, localAddress string) error {
	app := models.FilemgrApp{}
	app.AppType = appType
	app.Version = version
	app.Platform = platform
	key, err := e.generateAppOssObjectKey(&app)
	if err != nil {
		return err
	}
	client, err := e.getOssClient()
	if err != nil {
		return err
	}
	err = client.UploadWithSpace(key, localAddress)
	if err != nil {
		e.Log.Errorf("上传失败，失败原因:%s \r\n", err)
		return err
	}
	return nil
}

// generateAppOssUrl
// @Description: 获取app下载链接
// @receiver e
// @param App
// @return string
// @return error
func (e *FilemgrApp) generateAppOssUrl(App *models.FilemgrApp) (string, error) {
	appPath, err := e.generateAppOssObjectKey(App)
	if err != nil {
		return "", err
	}
	oss, err := e.getOssClient()
	if err != nil {
		return "", err
	}
	return oss.GeneratePresignedUrl(appPath)
}

// getOssClient
// @Description: 获取oss客户端
// @receiver e
// @param App
// @return *ossUtils.ALiYunOSS
// @return error
func (e *FilemgrApp) getOssClient() (*ossUtils.ALiYunOSS, error) {
	var sysConfService = adminService.NewSysConfigService(&e.Service)
	endPoint, _, _ := sysConfService.GetWithKeyStr("plugin_filemgr_app_oss_endpoint")
	key, _, _ := sysConfService.GetWithKeyStr("plugin_filemgr_app_oss_access_key_id")
	secret, _, _ := sysConfService.GetWithKeyStr("plugin_filemgr_app_oss_access_key_secret")
	bucketName, _, _ := sysConfService.GetWithKeyStr("plugin_filemgr_app_oss_bucket")
	if endPoint == "" || key == "" || secret == "" || bucketName == "" {
		return nil, errors.New("oss param config empty")
	}
	oss := ossUtils.ALiYunOSS{}
	err := oss.InitOssClient(key, secret, endPoint, bucketName)
	if err != nil {
		return nil, err
	}
	return &oss, nil
}

// generateAppOssObjectKey
// @Description: 生成oss key
// @receiver e
// @param App
// @return string
// @return error
func (e *FilemgrApp) generateAppOssObjectKey(App *models.FilemgrApp) (string, error) {
	var sysConfService = adminService.NewSysConfigService(&e.Service)
	var dictDataService = adminService.NewSysDictDataService(&e.Service)

	//app目录
	appPath, _, err := sysConfService.GetWithKeyStr("plugin_filemgr_app_oss_root_path")
	if err != nil {
		return "", errors.New(fmt.Sprintf("oss config err: %s", err))
	}
	appPath += dictDataService.GetLabel("plugin_filemgr_app_type", App.AppType) + "_"
	appPath += App.Version

	switch App.Platform {
	case constant.AppPlatformAndroid:
		appPath += ".apk"
	case constant.AppPlatformIOS:
		appPath += ".ipa"
	default:
		return "", errors.New("app platform error")
	}
	return appPath, nil
}
