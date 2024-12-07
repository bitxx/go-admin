package service

import (
	"archive/zip"
	"bytes"
	"go-admin/app/admin/constant"
	sysLang "go-admin/app/admin/lang"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/core/config"
	cDto "go-admin/core/dto"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"go-admin/core/utils/dateutils"
	"go-admin/core/utils/fileutils"
	"gorm.io/gorm"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type SysGenTable struct {
	service.Service
}

func NewSysGenTableService(s *service.Service) *SysGenTable {
	var srv = new(SysGenTable)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

func (e *SysGenTable) GetPage(c *dto.SysGenTableQueryReq, p *middleware.DataPermission) ([]models.SysGenTable, int64, int, error) {
	var list []models.SysGenTable
	var data models.SysGenTable
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
func (e *SysGenTable) Get(id int64, p *middleware.DataPermission) (*models.SysGenTable, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.SysGenTable{}
	err := e.Orm.Preload("SysGenColumns").Scopes(
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

// QueryOne 通过自定义条件获取一条记录
func (e *SysGenTable) QueryOne(queryCondition *dto.SysGenTableQueryReq, p *middleware.DataPermission) (*models.SysGenTable, int, error) {
	data := &models.SysGenTable{}
	err := e.Orm.Model(&models.SysGenTable{}).
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

// Count 获取条数
func (e *SysGenTable) Count(c *dto.SysGenTableQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysGenTable{}).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Limit(-1).Offset(-1).
		Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return 0, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return count, lang.SuccessCode, nil
}

func (e *SysGenTable) Insert(c *dto.SysGenTableInsertReq) (int, error) {
	if len(c.DbTableNames) <= 0 {
		return sysLang.SysGenTableSelectCode, lang.MsgErr(sysLang.SysGenTableSelectCode, e.Lang)
	}
	req := dto.SysGenTableQueryReq{}
	req.TableNames = c.DbTableNames
	count, respCode, err := e.Count(&req)
	if count > 0 {
		return sysLang.SysGenTableInsertExistCode, lang.MsgErr(sysLang.SysGenTableInsertExistCode, e.Lang)
	}
	sysTables, respCode, err := e.genTables(c.DbTableNames)
	if err != nil {
		return respCode, err
	}

	e.Orm = e.Orm.Begin()
	defer func() {
		if err != nil {
			e.Orm.Rollback()
		} else {
			e.Orm.Commit()
		}
	}()
	err = e.Orm.Create(&sysTables).Error
	if err != nil {
		return lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return lang.SuccessCode, nil
}

func (e *SysGenTable) Update(c *dto.SysGenTableUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	e.Orm = e.Orm.Begin()
	defer func() {
		if err != nil {
			e.Orm.Rollback()
		} else {
			e.Orm.Commit()
		}
	}()

	//最小化变更改动过的数据
	updates := map[string]interface{}{}
	if c.FunctionAuthor != "" && data.FunctionAuthor != c.FunctionAuthor {
		updates["function_author"] = c.FunctionAuthor
	}
	if c.TableComment != "" && data.TableComment != c.TableComment {
		updates["table_comment"] = c.TableComment
	}
	if c.ClassName != "" && data.ClassName != c.ClassName {
		updates["class_name"] = c.ClassName
	}
	if c.BusinessName != "" && data.BusinessName != c.BusinessName {
		updates["business_name"] = c.BusinessName
	}
	if c.PackageName != "" && data.PackageName != c.PackageName {
		updates["package_name"] = c.PackageName
	}
	if c.ModuleName != "" && data.ModuleName != c.ModuleName {
		updates["module_name"] = c.ModuleName
	}
	if c.FunctionName != "" && data.FunctionName != c.FunctionName {
		updates["function_name"] = c.FunctionName
	}
	if c.Remark != "" && data.Remark != c.Remark {
		updates["remark"] = c.Remark
	}

	isUpdate := false
	if len(updates) > 0 {
		updates["updated_at"] = time.Now()
		updates["update_by"] = c.CurrUserId
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		isUpdate = true
	}
	columnsService := NewSysColumnsService(&e.Service)
	for _, column := range c.Columns {
		column.CurrUserId = c.CurrUserId
		var b bool
		b, respCode, err = columnsService.Update(&column, p)
		if err != nil {
			return false, respCode, err
		}
		if b {
			isUpdate = true
		}
	}

	return isUpdate, lang.SuccessCode, nil
}

func (e *SysGenTable) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	var err error
	e.Orm = e.Orm.Begin()
	defer func() {
		if err != nil {
			e.Orm.Rollback()
		} else {
			e.Orm.Commit()
		}
	}()
	var data models.SysGenTable
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	columnsService := NewSysColumnsService(&e.Service)
	columnReq := dto.SysGenColumnDeleteReq{}
	columnReq.TableIds = ids
	respCode, err := columnsService.Remove(columnReq, p)
	if err != nil {
		return respCode, err
	}
	return lang.SuccessCode, nil
}

func (e *SysGenTable) GetDBTablePage(c dto.DBTableQueryReq) ([]dto.DBTableResp, int64, int, error) {
	var list []models.DBTable
	var data models.DBTable
	var count int64

	err := e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Where("table_name not like 'sys_%'").
		Where("table_name not in (select table_name from sys_gen_table)").
		Where("table_schema= ? ", e.Orm.Migrator().CurrentDatabase()).
		Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	var respList []dto.DBTableResp
	for _, item := range list {
		dbTableResp := dto.DBTableResp{}
		dbTableResp.CreatedAt = dateutils.ConvertToStrByPrt(item.CreateTime, -1)
		dbTableResp.TableName = item.TBName
		dbTableResp.TableComment = item.TableComment
		respList = append(respList, dbTableResp)
	}
	return respList, count, lang.SuccessCode, nil
}

// genTables 根据表名称生成表结构集合
func (e *SysGenTable) genTables(dbTableNames []string) ([]models.SysGenTable, int, error) {
	if len(dbTableNames) <= 0 {
		return nil, sysLang.SysGenTableSelectCode, lang.MsgErr(sysLang.SysGenTableSelectCode, e.Lang)
	}
	dbTables, resp, err := e.getDBTableList(dbTableNames)
	if err != nil {
		return nil, resp, err
	}

	var sysTables []models.SysGenTable
	for _, table := range dbTables {
		sysTable := models.SysGenTable{}

		// 默认去除表前缀后再去初始化modelName、packageName、businessName
		tableNoPrefix := table.TBName
		packageName := ""
		businessName := ""
		tempBusinessName := "" //备选packageName
		splits := strings.Split(table.TBName, "_")
		if len(splits) > 0 {
			tableNoPrefix = strings.Replace(table.TBName, splits[0]+"_", "", 1)
			packageName = splits[0]
			if len(splits) > 2 {
				businessName = splits[1]
			}
		}

		tbNameSplits := strings.Split(tableNoPrefix, "_")
		for index, _ := range tbNameSplits {
			strStart := string([]byte(tbNameSplits[index])[:1])
			strend := string([]byte(tbNameSplits[index])[1:])
			// 大驼峰表名 结构体使用
			sysTable.ClassName += strings.ToUpper(strStart) + strend
			// 小驼峰表名 js函数名和权限标识使用
			if index == 0 {
				tempBusinessName += strings.ToLower(strStart) + strend
			} else {
				tempBusinessName += strings.ToUpper(strStart) + strend
			}
		}
		if businessName == "" {
			businessName = tempBusinessName
		}
		if packageName == "" {
			packageName = tempBusinessName
		}
		sysTable.PackageName = packageName
		sysTable.BusinessName = businessName
		sysTable.ModuleName = strings.Replace(tableNoPrefix, "_", "-", -1)

		sysTable.TBName = table.TBName
		sysTable.TableComment = table.TableComment
		if sysTable.TableComment == "" {
			sysTable.TableComment = sysTable.ClassName
		}
		sysTable.FunctionName = sysTable.TableComment
		sysTable.FunctionAuthor = config.ApplicationConfig.Author

		columnsService := NewSysColumnsService(&e.Service)
		columns, respCode, err := columnsService.GetDBColumnList(table.TBName)
		if err != nil {
			return nil, respCode, err
		}
		for index, column := range columns {
			sysColumn := models.SysGenColumn{}
			sysColumn.ColumnComment = column.ColumnComment
			sysColumn.ColumnName = column.ColumnName
			sysColumn.ColumnType = column.ColumnType
			sysColumn.Sort = index + 1
			sysColumn.QueryType = "EQ"
			sysColumn.IsPk = global.SysStatusNotOk
			sysColumn.IsQuery = global.SysStatusNotOk
			sysColumn.IsList = global.SysStatusNotOk

			namelist := strings.Split(sysColumn.ColumnName, "_")
			for i := 0; i < len(namelist); i++ {
				strStart := string([]byte(namelist[i])[:1])
				strend := string([]byte(namelist[i])[1:])
				sysColumn.GoField += strings.ToUpper(strStart) + strend
				if i == 0 {
					sysColumn.JsonField = strings.ToLower(strStart) + strend
				} else {
					sysColumn.JsonField += strings.ToUpper(strStart) + strend
				}
			}
			if strings.Contains(column.ColumnKey, "PR") {
				sysColumn.IsPk = global.SysStatusOk
			}
			sysColumn.IsRequired = global.SysStatusNotOk
			if strings.Contains(column.IsNullable, "NO") {
				sysColumn.IsRequired = global.SysStatusOk
			}
			if strings.Contains(column.ColumnType, "int") {
				sysColumn.GoType = "int64"
				sysColumn.HtmlType = "input"
			} else if strings.Contains(column.ColumnType, "decimal") {
				sysColumn.GoType = "decimal.Decimal"
				sysColumn.HtmlType = "input"
			} else if strings.Contains(column.ColumnType, "timestamp") {
				sysColumn.GoType = "*time.Time"
				sysColumn.HtmlType = "datetime"
			} else if strings.Contains(column.ColumnType, "datetime") {
				sysColumn.GoType = "*time.Time"
				sysColumn.HtmlType = "datetime"
			} else {
				sysColumn.GoType = "string"
				sysColumn.HtmlType = "input"
			}
			sysTable.SysGenColumns = append(sysTable.SysGenColumns, sysColumn)
		}
		sysTables = append(sysTables, sysTable)
	}
	return sysTables, lang.SuccessCode, nil
}

// getDBTableList 从数据库中获取表指定表的完整结构
func (e *SysGenTable) getDBTableList(tableNames []string) ([]models.DBTable, int, error) {
	if len(tableNames) <= 0 {
		return nil, sysLang.SysGenTableSelectCode, lang.MsgErr(sysLang.SysGenTableSelectCode, e.Lang)
	}
	var data []models.DBTable
	err := e.Orm.Where("TABLE_NAME in (?)", tableNames).Find(&data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return data, lang.SuccessCode, nil
}

func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// Preview
//
//	@Description: 代码预览
//	@receiver e
//	@param c
//	@param p
//	@return []dto.TemplateResp
//	@return int
//	@return error
func (e *SysGenTable) Preview(c dto.SysGenTableGenCodeReq, p *middleware.DataPermission) ([]dto.TemplateResp, int, error) {
	if c.Id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if config.GenConfig.Type != global.GenTypeVue && config.GenConfig.Type != global.GenTypeReact {
		return nil, sysLang.SysGenFrontTypeErrCode, lang.MsgErr(sysLang.SysGenFrontTypeErrCode, e.Lang)
	}
	table, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return nil, respCode, err
	}
	var resp []dto.TemplateResp
	for k, v := range constant.TemplatInfo {
		tpl, err := template.New(filepath.Base(v)).Funcs(template.FuncMap{
			"contains": strings.Contains,
		}).ParseFiles(v)
		if err != nil {
			return nil, sysLang.SysGenTemplateModelReadLogErrCode, lang.MsgLogErrf(e.Log, e.Lang, sysLang.SysGenTemplateModelReadErrCode, sysLang.SysGenTemplateModelReadLogErrCode, err)
		}

		var content bytes.Buffer
		err = tpl.Execute(&content, table)
		if err != nil {
			return nil, sysLang.SysGenTemplateModelDecodeLogErrCode, lang.MsgLogErrf(e.Log, e.Lang, sysLang.SysGenTemplateModelDecodeErrCode, sysLang.SysGenTemplateModelDecodeLogErrCode, err)
		}

		//生成文件的路径
		defaultPath := "./app/"
		path := defaultPath
		tableName := strings.Replace(table.ModuleName, "-", "_", -1) //golang 文件名使用下划线
		if k == constant.ModelName {
			path = path + table.PackageName + "/" + table.BusinessName + "/models/" + tableName + ".go"
		}
		if k == constant.ApiName {
			path = path + table.PackageName + "/" + table.BusinessName + "/apis/" + tableName + ".go"
		}
		if k == constant.BusinessRouterName {
			path = path + table.PackageName + "/" + table.BusinessName + "/router/" + tableName + ".go"
		}
		if k == constant.DtoName {
			path = path + table.PackageName + "/" + table.BusinessName + "/service/dto/" + tableName + ".go"
		}
		if k == constant.ServiceName {
			path = path + table.PackageName + "/" + table.BusinessName + "/service/" + tableName + ".go"
		}
		if k == constant.RouterName {
			path = path + table.PackageName + "/" + table.BusinessName + "/router/router.go.bk"
		}
		if k == constant.LangName {
			path = path + table.PackageName + "/" + table.BusinessName + "/lang/lang.go.bk"
		}
		if k == constant.ConstantName {
			path = path + table.PackageName + "/" + table.BusinessName + "/constant/constant.go.bk"
		}
		if config.GenConfig.Type == global.GenTypeVue {
			if k == constant.VueApiJsName {
				path = config.GenConfig.FrontPath + "/api/" + table.PackageName + "/" + table.BusinessName + "/" + table.ModuleName + ".js"
			}
			if k == constant.VueIndexName {
				path = config.GenConfig.FrontPath + "/views/" + table.PackageName + "/" + table.BusinessName + "/" + table.ModuleName + "/index.vue"
			}
		}

		if config.GenConfig.Type == global.GenTypeReact {
			if k == constant.ReactApiName {
				path = config.GenConfig.FrontPath + "/api/" + table.PackageName + "/" + table.BusinessName + "/" + table.ModuleName + "/index.ts"
			}
			if k == constant.ReactFormModalName {
				path = config.GenConfig.FrontPath + "/views/" + table.PackageName + "/" + table.BusinessName + "/" + table.ModuleName + "/components/FormModal.tsx"
			}
			if k == constant.ReactViewName {
				path = config.GenConfig.FrontPath + "/views/" + table.PackageName + "/" + table.BusinessName + "/" + table.ModuleName + "/index.tsx"
			}
		}

		if path != defaultPath {
			tplResp := dto.TemplateResp{
				Name:    k,
				Path:    path,
				Content: content.String(),
			}
			resp = append(resp, tplResp)
		}

	}
	return resp, lang.SuccessCode, nil
}

// GenCode
//
//	@Description: 生成代码文件
//	@receiver e
//	@param c
func (e *SysGenTable) GenCode(c dto.SysGenTableGenCodeReq, p *middleware.DataPermission) (*bytes.Buffer, int, error) {
	if c.Id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}

	templateResp, respCode, err := e.Preview(c, p)
	if err != nil {
		return nil, respCode, err
	}

	//如果是下载zpi压缩包代码
	if c.IsDownload == global.SysStatusOk {
		buf := new(bytes.Buffer)
		writer := zip.NewWriter(buf)
		defer writer.Close()
		for _, tpl := range templateResp {
			err = fileutils.ZipFilCreate(writer, *bytes.NewBufferString(tpl.Content), tpl.Path)
		}
		return buf, lang.SuccessCode, nil
	}
	//如果是直接生成代码
	for _, tpl := range templateResp {
		err = fileutils.CreateDirFromFilePath(tpl.Path)
		if err != nil {
			e.Log.Warn(err)
		}
		err = fileutils.FileCreate(*bytes.NewBufferString(tpl.Content), tpl.Path)
		if err != nil {
			e.Log.Warn(err)
		}
	}
	return nil, lang.SuccessCode, nil
}

// GenDB 插入菜单到数据库
func (e *SysGenTable) GenDB(c dto.SysGenTableGetReq, p *middleware.DataPermission) (int, error) {
	if c.Id <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}

	var err error
	e.Orm = e.Orm.Begin()
	defer func() {
		if err != nil {
			e.Orm.Rollback()
		} else {
			e.Orm.Commit()
		}
	}()
	table, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return respCode, err
	}
	menuService := NewSysMenuService(&e.Service)

	//插入菜单
	premission := table.PackageName + ":" + table.BusinessName + ":" + table.ModuleName
	path := "/" + table.PackageName + "/" + table.BusinessName + "/" + table.ModuleName
	cMenuInsertReq := dto.SysMenuInsertReq{
		CurrUserId:  c.CurrUserId,
		Title:       table.TableComment,
		Icon:        "pass",
		Path:        path,
		MenuType:    constant.MenuC,
		ParentId:    0,
		IsKeepAlive: global.SysStatusNotOk,
		Element:     path + "/index.vue",
		IsHidden:    global.SysStatusNotOk,
		IsFrame:     global.SysStatusOk,
		IsAffix:     global.SysStatusNotOk,
	}
	cMenuId, respCode, err := menuService.Insert(&cMenuInsertReq)
	if err != nil {
		return respCode, err
	}

	//查询按钮
	mMenuQueryInsertReq := dto.SysMenuInsertReq{
		CurrUserId:  c.CurrUserId,
		Title:       "分页获取" + table.TableComment,
		MenuType:    constant.MenuF,
		Permission:  premission + ":query",
		ParentId:    cMenuId,
		IsKeepAlive: global.SysStatusNotOk,
		IsHidden:    global.SysStatusNotOk,
		IsFrame:     global.SysStatusOk,
		IsAffix:     global.SysStatusNotOk,
	}
	_, respCode, err = menuService.Insert(&mMenuQueryInsertReq)
	if err != nil {
		return respCode, err
	}

	//新增按钮
	mMenuAddInsertReq := dto.SysMenuInsertReq{
		CurrUserId:  c.CurrUserId,
		Title:       "创建" + table.TableComment,
		MenuType:    constant.MenuF,
		Permission:  premission + ":add",
		ParentId:    cMenuId,
		IsKeepAlive: global.SysStatusNotOk,
		IsHidden:    global.SysStatusNotOk,
		IsFrame:     global.SysStatusOk,
		IsAffix:     global.SysStatusNotOk,
	}
	_, respCode, err = menuService.Insert(&mMenuAddInsertReq)
	if err != nil {
		return respCode, err
	}

	//修改按钮
	mMenuUpdateInsertReq := dto.SysMenuInsertReq{
		CurrUserId:  c.CurrUserId,
		Title:       "修改" + table.TableComment,
		MenuType:    constant.MenuF,
		Permission:  premission + ":edit",
		ParentId:    cMenuId,
		IsKeepAlive: global.SysStatusNotOk,
		IsHidden:    global.SysStatusNotOk,
		IsFrame:     global.SysStatusOk,
		IsAffix:     global.SysStatusNotOk,
	}
	_, respCode, err = menuService.Insert(&mMenuUpdateInsertReq)
	if err != nil {
		return respCode, err
	}

	//删除按钮
	mMenuDelInsertReq := dto.SysMenuInsertReq{
		CurrUserId:  c.CurrUserId,
		Title:       "删除" + table.TableComment,
		MenuType:    constant.MenuF,
		Permission:  premission + ":del",
		ParentId:    cMenuId,
		IsKeepAlive: global.SysStatusNotOk,
		IsHidden:    global.SysStatusNotOk,
		IsFrame:     global.SysStatusOk,
		IsAffix:     global.SysStatusNotOk,
	}
	_, respCode, err = menuService.Insert(&mMenuDelInsertReq)
	if err != nil {
		return respCode, err
	}

	//导出按钮
	mMenuExportInsertReq := dto.SysMenuInsertReq{
		CurrUserId:  c.CurrUserId,
		Title:       "导出" + table.TableComment,
		MenuType:    constant.MenuF,
		Permission:  premission + ":export",
		ParentId:    cMenuId,
		IsKeepAlive: global.SysStatusNotOk,
		IsHidden:    global.SysStatusNotOk,
		IsFrame:     global.SysStatusOk,
		IsAffix:     global.SysStatusNotOk,
	}
	_, respCode, err = menuService.Insert(&mMenuExportInsertReq)
	if err != nil {
		return respCode, err
	}
	return lang.SuccessCode, nil
}
