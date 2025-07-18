package service

import (
	"archive/zip"
	"bytes"
	"fmt"
	"go-admin/config/base/constant"

	"go-admin/app/admin/sys/models"
	"go-admin/app/admin/sys/service/dto"
	baseLang "go-admin/config/base/lang"
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

// NewSysGenTableService admin-实例化表管理
func NewSysGenTableService(s *service.Service) *SysGenTable {
	var srv = new(SysGenTable)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage admin-获取表管理分页列表
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
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	return list, count, baseLang.SuccessCode, nil
}

// Get admin-获取表管理详情
func (e *SysGenTable) Get(id int64, p *middleware.DataPermission) (*models.SysGenTable, int, error) {
	if id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
	}
	data := &models.SysGenTable{}
	err := e.Orm.Preload("SysGenColumns").Scopes(
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

// QueryOne admin-获取表管理一条记录
func (e *SysGenTable) QueryOne(queryCondition *dto.SysGenTableQueryReq, p *middleware.DataPermission) (*models.SysGenTable, int, error) {
	data := &models.SysGenTable{}
	err := e.Orm.Model(&models.SysGenTable{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).First(data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return data, baseLang.SuccessCode, nil
}

// Count admin-获取表管理数据总数
func (e *SysGenTable) Count(c *dto.SysGenTableQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysGenTable{}).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Limit(-1).Offset(-1).
		Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return 0, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}
	return count, baseLang.SuccessCode, nil
}

// Insert admin-新增表管理
func (e *SysGenTable) Insert(c *dto.SysGenTableInsertReq) (int, error) {
	if len(c.DbTableNames) <= 0 {
		return baseLang.SysGenTableSelectCode, lang.MsgErr(baseLang.SysGenTableSelectCode, e.Lang)
	}
	req := dto.SysGenTableQueryReq{}
	req.TableNames = c.DbTableNames
	count, respCode, err := e.Count(&req)
	if count > 0 {
		return baseLang.SysGenTableInsertExistCode, lang.MsgErr(baseLang.SysGenTableInsertExistCode, e.Lang)
	}
	sysTables, respCode, err := e.genTables(c.DbTableNames)
	if err != nil {
		return respCode, err
	}

	err = e.Orm.Transaction(func(tx *gorm.DB) error {
		if err = e.Orm.Create(&sysTables).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return baseLang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataInsertCode, baseLang.DataInsertLogCode, err)
	}
	return baseLang.SuccessCode, nil
}

// Update admin-更新表管理
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
			return false, baseLang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataUpdateCode, baseLang.DataUpdateLogCode, err)
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

	return isUpdate, baseLang.SuccessCode, nil
}

// Delete admin-删除表管理
func (e *SysGenTable) Delete(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
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
		return baseLang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataDeleteCode, baseLang.DataDeleteLogCode, err)
	}
	columnsService := NewSysColumnsService(&e.Service)
	columnReq := dto.SysGenColumnDeleteReq{}
	columnReq.TableIds = ids
	respCode, err := columnsService.Delete(columnReq, p)
	if err != nil {
		return respCode, err
	}
	return baseLang.SuccessCode, nil
}

// GetDBTablePage admin-获取表管理的DB表分页列表
func (e *SysGenTable) GetDBTablePage(c dto.DBTableQueryReq) ([]dto.DBTableResp, int64, int, error) {
	var list []models.DBTable
	var count int64
	pageSize := c.GetPageSize()
	pageIndex := c.GetPageIndex()

	// 公共部分
	excludeTables := []interface{}{
		"admin_sys_role_menu", "admin_sys_role_dept", "admin_sys_menu_api_rule",
		"admin_sys_gen_column", "admin_sys_casbin_rule",
	}
	limitOffset := []interface{}{pageSize, (pageIndex - 1) * pageSize}

	var querySql, countSql string
	var args, countArgs []interface{}

	if config.DatabaseConfig.Driver == global.DBDriverPostgres {
		querySql = `
SELECT 
	tablename AS "table_name",
	obj_description(('"' || tablename || '"')::regclass, 'pg_class') AS "table_comment",
	NULL::timestamp AS "create_time"
FROM pg_tables 
WHERE schemaname = 'public'
  AND tablename NOT IN (?, ?, ?, ?, ?)
  AND tablename NOT IN (SELECT table_name FROM admin_sys_gen_table)
LIMIT ? OFFSET ?
`
		countSql = `
SELECT COUNT(*) FROM pg_tables 
WHERE schemaname = 'public'
  AND tablename NOT IN (?, ?, ?, ?, ?)
  AND tablename NOT IN (SELECT table_name FROM admin_sys_gen_table)
`
		args = append(excludeTables, limitOffset...)
		countArgs = excludeTables
	} else {
		db := e.Orm.Migrator().CurrentDatabase()
		querySql = `
SELECT 
	table_name as table_name,
	table_comment as table_comment,
	create_time as create_time
FROM information_schema.tables 
WHERE table_schema = ?
  AND table_name NOT IN (?, ?, ?, ?, ?)
  AND table_name NOT IN (SELECT table_name FROM admin_sys_gen_table)
LIMIT ? OFFSET ?
`
		countSql = `
SELECT COUNT(*) FROM information_schema.tables 
WHERE table_schema = ?
  AND table_name NOT IN (?, ?, ?, ?, ?)
  AND table_name NOT IN (SELECT table_name FROM admin_sys_gen_table)
`
		args = append([]interface{}{db}, excludeTables...)
		args = append(args, limitOffset...)
		countArgs = append([]interface{}{db}, excludeTables...)
	}

	// 查询数据
	if err := e.Orm.Raw(querySql, args...).Scan(&list).Error; err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err := e.Orm.Raw(countSql, countArgs...).Scan(&count).Error; err != nil {
		return nil, 0, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}

	// 构造响应
	respList := make([]dto.DBTableResp, 0, len(list))
	for _, item := range list {
		respList = append(respList, dto.DBTableResp{
			TableName:    item.TBName,
			TableComment: item.TableComment,
			CreatedAt:    dateutils.ConvertToStrByPrt(item.CreateTime, -1),
		})
	}

	return respList, count, baseLang.SuccessCode, nil
}

// genTables admin-根据表名称生成表结构集合
func (e *SysGenTable) genTables(dbTableNames []string) ([]models.SysGenTable, int, error) {
	if len(dbTableNames) <= 0 {
		return nil, baseLang.SysGenTableSelectCode, lang.MsgErr(baseLang.SysGenTableSelectCode, e.Lang)
	}
	dbTables, resp, err := e.getDBTableList(dbTableNames)
	if err != nil {
		return nil, resp, err
	}

	var sysTables []models.SysGenTable
	now := time.Now()
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
			//must cmp pk at first
			if config.DatabaseConfig.Driver == global.DBDriverPostgres {
				var isPK bool
				sql := `
						SELECT EXISTS (
							SELECT 1
							FROM pg_index i
							JOIN pg_attribute a ON a.attrelid = i.indrelid AND a.attnum = ANY(i.indkey)
							WHERE i.indrelid = ?::regclass
							  AND i.indisprimary
							  AND a.attname = ?
						) AS is_primary_key;
    					`
				_ = e.Orm.Raw(sql, table.TBName, column.ColumnName).Scan(&isPK).Error
				if isPK {
					sysColumn.IsPk = global.SysStatusOk
				}
			} else if strings.Contains(column.ColumnKey, "PR") {
				sysColumn.IsPk = global.SysStatusOk
			}
			sysColumn.IsRequired = global.SysStatusNotOk
			if strings.Contains(column.IsNullable, "NO") {
				sysColumn.IsRequired = global.SysStatusOk
			}
			if strings.Contains(column.ColumnType, "int") || strings.Contains(column.ColumnType, "BIGINT") {
				sysColumn.GoType = "int64"
				sysColumn.HtmlType = "numInput"
			} else if strings.Contains(column.ColumnType, "decimal") {
				sysColumn.GoType = "decimal.Decimal"
				sysColumn.HtmlType = "input"
			} else if strings.Contains(column.ColumnType, "timestamp") || strings.Contains(column.ColumnType, "datetime") {
				sysColumn.GoType = "*time.Time"
				sysColumn.HtmlType = "datetime"
			} else {
				sysColumn.GoType = "string"
				sysColumn.HtmlType = "input"
			}
			sysColumn.CreatedAt = &now
			sysColumn.UpdatedAt = &now
			sysTable.SysGenColumns = append(sysTable.SysGenColumns, sysColumn)
		}

		sysTable.CreatedAt = &now
		sysTable.UpdatedAt = &now
		sysTables = append(sysTables, sysTable)
	}
	return sysTables, baseLang.SuccessCode, nil
}

// getDBTableList admin-从数据库中获取表指定表的完整结构
func (e *SysGenTable) getDBTableList(tableNames []string) ([]models.DBTable, int, error) {
	if len(tableNames) == 0 {
		return nil, baseLang.SysGenTableSelectCode, lang.MsgErr(baseLang.SysGenTableSelectCode, e.Lang)
	}

	var list []models.DBTable
	var sql string
	var args []interface{}

	if config.DatabaseConfig.Driver == global.DBDriverPostgres {
		placeholders := make([]string, len(tableNames))
		for i := range tableNames {
			placeholders[i] = fmt.Sprintf("$%d", i+1)
			args = append(args, tableNames[i])
		}
		sql = fmt.Sprintf(`
SELECT 
	tablename AS "table_name",
	obj_description(('"' || tablename || '"')::regclass, 'pg_class') AS "table_comment",
	NULL::text AS "create_time"
FROM pg_tables
WHERE schemaname = 'public'
  AND tablename IN (%s)`, strings.Join(placeholders, ", "))
	} else {
		sql = `
SELECT 
	table_name as table_name,
	table_comment as table_comment,
	create_time as create_time
FROM information_schema.tables
WHERE table_schema = ?
  AND table_name IN (?)`
		args = append([]interface{}{e.Orm.Migrator().CurrentDatabase()}, tableNames)
	}

	err := e.Orm.Raw(sql, args...).Scan(&list).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, baseLang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.DataQueryCode, baseLang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, baseLang.DataNotFoundCode, lang.MsgErr(baseLang.DataNotFoundCode, e.Lang)
	}

	return list, baseLang.SuccessCode, nil
}

// Preview admin-预览表管理的代码页面
func (e *SysGenTable) Preview(c dto.SysGenTableGenCodeReq, p *middleware.DataPermission) ([]dto.TemplateResp, int, error) {
	if c.Id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
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
			return nil, baseLang.SysGenTemplateModelReadLogErrCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.SysGenTemplateModelReadErrCode, baseLang.SysGenTemplateModelReadLogErrCode, err)
		}

		var content bytes.Buffer
		err = tpl.Execute(&content, table)
		if err != nil {
			return nil, baseLang.SysGenTemplateModelDecodeLogErrCode, lang.MsgLogErrf(e.Log, e.Lang, baseLang.SysGenTemplateModelDecodeErrCode, baseLang.SysGenTemplateModelDecodeLogErrCode, err)
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

		if k == constant.ReactApiName {
			path = config.GenConfig.FrontPath + "/api/" + table.PackageName + "/" + table.BusinessName + "/" + table.ModuleName + "/index.ts"
		}
		if k == constant.ReactFormModalName {
			path = config.GenConfig.FrontPath + "/views/" + table.PackageName + "/" + table.BusinessName + "/" + table.ModuleName + "/components/FormModal.tsx"
		}
		if k == constant.ReactViewName {
			path = config.GenConfig.FrontPath + "/views/" + table.PackageName + "/" + table.BusinessName + "/" + table.ModuleName + "/index.tsx"
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
	return resp, baseLang.SuccessCode, nil
}

// GenCode admin-生成表管理的代码
func (e *SysGenTable) GenCode(c dto.SysGenTableGenCodeReq, p *middleware.DataPermission) (*bytes.Buffer, int, error) {
	if c.Id <= 0 {
		return nil, baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
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
		return buf, baseLang.SuccessCode, nil
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
	return nil, baseLang.SuccessCode, nil
}

// GenDB admin-表管理中生成菜单数据
func (e *SysGenTable) GenDB(c dto.SysGenTableGetReq, p *middleware.DataPermission) (int, error) {
	if c.Id <= 0 {
		return baseLang.ParamErrCode, lang.MsgErr(baseLang.ParamErrCode, e.Lang)
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
	basePremission := table.PackageName + ":" + table.ModuleName
	basePath := "/" + table.PackageName + "/" + table.BusinessName + "/" + table.ModuleName

	menuService := NewSysMenuService(&e.Service)

	//插入主菜单
	cMenuInsertReq := dto.SysMenuInsertReq{
		CurrUserId:  c.CurrUserId,
		Title:       table.TableComment,
		Icon:        "LayoutFilled",
		Path:        basePath,
		Element:     basePath + "/index",
		MenuType:    constant.MenuC,
		ParentId:    0,
		IsKeepAlive: global.SysStatusOk,
		IsFrame:     global.SysStatusOk,
		IsHidden:    global.SysStatusNotOk,
		IsAffix:     global.SysStatusNotOk,
	}
	cMenuId, respCode, err := menuService.Insert(&cMenuInsertReq)
	if err != nil {
		return respCode, err
	}

	//查询按钮
	mMenuQueryInsertReq := dto.SysMenuInsertReq{
		CurrUserId:  c.CurrUserId,
		Title:       "获取" + table.TableComment + "分页列表",
		Icon:        "AppstoreOutlined",
		MenuType:    constant.MenuF,
		Permission:  basePremission + ":query",
		ParentId:    cMenuId,
		IsKeepAlive: global.SysStatusNotOk,
		IsHidden:    global.SysStatusNotOk,
		IsAffix:     global.SysStatusNotOk,
		IsFrame:     global.SysStatusOk,
	}
	_, respCode, err = menuService.Insert(&mMenuQueryInsertReq)
	if err != nil {
		return respCode, err
	}

	//新增按钮
	mMenuAddInsertReq := dto.SysMenuInsertReq{
		CurrUserId:  c.CurrUserId,
		Title:       "新增" + table.TableComment,
		Icon:        "AppstoreOutlined",
		MenuType:    constant.MenuF,
		Permission:  basePremission + ":add",
		ParentId:    cMenuId,
		IsKeepAlive: global.SysStatusNotOk,
		IsHidden:    global.SysStatusNotOk,
		IsAffix:     global.SysStatusNotOk,
		IsFrame:     global.SysStatusOk,
	}
	_, respCode, err = menuService.Insert(&mMenuAddInsertReq)
	if err != nil {
		return respCode, err
	}

	//更新按钮
	mMenuUpdateInsertReq := dto.SysMenuInsertReq{
		CurrUserId:  c.CurrUserId,
		Title:       "更新" + table.TableComment,
		Icon:        "AppstoreOutlined",
		MenuType:    constant.MenuF,
		Permission:  basePremission + ":edit",
		ParentId:    cMenuId,
		IsKeepAlive: global.SysStatusNotOk,
		IsHidden:    global.SysStatusNotOk,
		IsAffix:     global.SysStatusNotOk,
		IsFrame:     global.SysStatusOk,
	}
	_, respCode, err = menuService.Insert(&mMenuUpdateInsertReq)
	if err != nil {
		return respCode, err
	}

	//删除按钮
	mMenuDelInsertReq := dto.SysMenuInsertReq{
		CurrUserId:  c.CurrUserId,
		Title:       "删除" + table.TableComment,
		Icon:        "AppstoreOutlined",
		MenuType:    constant.MenuF,
		Permission:  basePremission + ":del",
		ParentId:    cMenuId,
		IsKeepAlive: global.SysStatusNotOk,
		IsHidden:    global.SysStatusNotOk,
		IsAffix:     global.SysStatusNotOk,
		IsFrame:     global.SysStatusOk,
	}
	_, respCode, err = menuService.Insert(&mMenuDelInsertReq)
	if err != nil {
		return respCode, err
	}

	//导出按钮
	mMenuExportInsertReq := dto.SysMenuInsertReq{
		CurrUserId:  c.CurrUserId,
		Title:       "导出" + table.TableComment,
		Icon:        "AppstoreOutlined",
		MenuType:    constant.MenuF,
		Permission:  basePremission + ":export",
		ParentId:    cMenuId,
		IsKeepAlive: global.SysStatusNotOk,
		IsHidden:    global.SysStatusNotOk,
		IsAffix:     global.SysStatusNotOk,
		IsFrame:     global.SysStatusOk,
	}
	_, respCode, err = menuService.Insert(&mMenuExportInsertReq)
	if err != nil {
		return respCode, err
	}
	return baseLang.SuccessCode, nil
}
