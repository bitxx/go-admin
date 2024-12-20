package models

import (
	"errors"
	"fmt"
	"go-admin/config/base/constant"
	"go-admin/core/global"
	"go-admin/core/runtime"
	"go/ast"
	"go/parser"
	"go/token"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type SysApi struct {
	Id          int64      `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Description string     `json:"description" gorm:"size:256;comment:功能描述"`
	Path        string     `json:"path" gorm:"size:128;comment:地址"`
	Method      string     `json:"method" gorm:"size:16;comment:请求类型"`
	ApiType     string     `json:"apiType" gorm:"size:16;comment:接口类型"`
	Remark      string     `json:"remark" gorm:"size:128;comment:备注"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	CreateBy    int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy    int64      `json:"updateBy" gorm:"index;comment:更新者"`
	SysMenu     []SysMenu  `json:"sysMenu" gorm:"many2many:admin_sys_menu_api_rule;foreignKey:id;joinForeignKey:admin_sys_api_id;references:Id;joinReferences:admin_sys_menu_menu_id;"`
}

func (SysApi) TableName() string {
	return "admin_sys_api"
}

func SaveSysApi(db *gorm.DB, routers []runtime.Router) (err error) {
	tx := db.Begin()
	var dbApiCacheMap = make(map[string]bool)
	var handlerApiCacheMap = make(map[string]*runtime.Router) //缓存路由中实际包含的地址，用于对比数据库，删除库中已经实效的路由地址
	defer func() {
		dbApiCacheMap = nil
		handlerApiCacheMap = nil
		if err != nil {
			tx.Rollback()
			return
		} else {
			tx.Commit()
		}
	}()

	apiInfos, err := getApiDescriptions()

	//读取库中所有接口并加入map缓存
	var dbApilist []SysApi
	err = db.Model(&SysApi{}).Find(&dbApilist).Error
	if err != nil {
		err = errors.New(fmt.Sprintf("get Api dbApilist error: %s \r\n ", err.Error()))
		return
	}
	for _, item := range dbApilist {
		dbApiCacheMap[item.Path+"-"+item.Method] = true
	}
	for _, v := range routers {
		handlerApiCacheMap[v.RelativePath+"-"+v.HttpMethod] = &v
	}

	//根据实际路由对比库中路由，将新路由加入库中
	var newSysApis []SysApi
	for k, v := range handlerApiCacheMap {

		if v.HttpMethod == "HEAD" {
			continue
		}
		//缓存
		if dbApiCacheMap[k] {
			continue
		}
		paths := strings.Split(v.RelativePath, "/")
		apiType := ""
		if len(paths) >= 4 {
			if strings.HasPrefix(paths[3], "admin") {
				apiType = constant.ApiTypeSys
			} else if strings.HasPrefix(paths[3], "plugins") {
				apiType = constant.ApiTypePlugin
			} else if strings.HasPrefix(paths[3], "app") {
				apiType = constant.ApiTypeApp
			}
		}
		newSysApi := SysApi{Path: v.RelativePath, Method: v.HttpMethod}
		if apiType != "" {
			newSysApi.ApiType = apiType
		}
		if apiInfos[v.Handler] != "" {
			newSysApi.Description = apiInfos[v.Handler]
		}
		newSysApis = append(newSysApis, newSysApi)
	}
	if len(newSysApis) > 0 {
		//事务批量插入，提高效率
		if err = tx.Debug().Model(&SysApi{}).Create(&newSysApis).Error; err != nil {
			return err
		}
		if err != nil {
			err = errors.New(fmt.Sprintf("Models SaveSysApi error: %s \r\n ", err.Error()))
			return
		}
		for _, item := range newSysApis {
			dbApiCacheMap[item.Path+"-"+item.Method] = true
		}
	}

	// 删除库中无效接口
	var delIds []int64
	for _, item := range dbApilist {
		if handlerApiCacheMap[item.Path+"-"+item.Method] == nil {
			delIds = append(delIds, item.Id)
		}
	}
	if len(delIds) > 0 {
		if err = tx.Table("admin_sys_menu_api_rule").Where("admin_sys_api_id in (?)", delIds).Delete(nil).Error; err != nil {
			return
		}
		// 删除主表数据
		if err = tx.Model(&SysApi{}).Delete(&SysApi{}, delIds).Error; err != nil {
			return
		}
		if err != nil {
			err = errors.New(fmt.Sprintf("sync delete api error: %s \r\n ", err.Error()))
			return
		}
	}
	return
}

// getApiDescriptions 获取所有api接口的说明
// 使用文件解析获取注释，通过拼接生成handler关联gin获取的handler，进而得到注释
func getApiDescriptions() (map[string]string, error) {
	dirs, err := findAllApiFileDirs("./")
	if err != nil {
		return nil, err
	}
	apiInfos := map[string]string{}
	for _, dir := range dirs {
		apiParseInfos, err := parseApiInfo(dir)
		if err != nil {
			return nil, err
		}

		handlerBase := filepath.Dir(global.ModelName+string(filepath.Separator)+dir) + "."
		for _, apiInfo := range apiParseInfos {
			handler := handlerBase + apiInfo.ClassName + "." + apiInfo.MethodName + "-fm"
			apiInfos[handler] = apiInfo.Description
		}
	}
	return apiInfos, nil
}

type ApiParseInfo struct {
	ClassName   string
	MethodName  string
	Description string
}

// ParseApiInfo 解析每个接口文件中类名、方法和注释
func parseApiInfo(filePath string) (apiParseInfos []ApiParseInfo, err error) {
	var infos []ApiParseInfo
	fset := token.NewFileSet()

	// 打开 Go 文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 解析文件，获取 AST
	node, err := parser.ParseFile(fset, filePath, file, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// 遍历 AST 的声明部分，查找方法
	for _, decl := range node.Decls {
		// 检查是否是函数声明
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue // 如果不是函数声明，跳过
		}

		// 获取函数所属的结构体名称（类名）
		className := ""
		if funcDecl.Recv != nil && len(funcDecl.Recv.List) > 0 {
			switch recvType := funcDecl.Recv.List[0].Type.(type) {
			case *ast.Ident:
				// 值类型
				className = recvType.Name
			case *ast.StarExpr:
				// 指针类型，递归检查
				if ident, ok := recvType.X.(*ast.Ident); ok {
					className = ident.Name
				}
			default:
				className = ""
			}
		}

		// 获取函数注释
		description := ""
		if funcDecl.Doc != nil {
			for _, comment := range funcDecl.Doc.List {
				if strings.HasPrefix(comment.Text, "//") {
					// 方法上面有多行注释，则只提取最上面一行的
					description = strings.TrimSpace(strings.TrimPrefix(comment.Text, "// "+funcDecl.Name.Name))
					break
				}
			}
		}

		// 记录方法信息
		if className != "" && funcDecl.Name.Name != "" && description != "" {
			infos = append(infos, ApiParseInfo{
				ClassName:   className,
				MethodName:  funcDecl.Name.Name,
				Description: description,
			})
		}

	}

	return infos, nil
}

// 获取所有api文件的go文件的路径
func findAllApiFileDirs(rootDir string) ([]string, error) {
	var goFiles []string

	// 使用 Walk 遍历整个目录
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查目录名称是否包含"apis"，并且是目录
		if info.IsDir() && strings.Contains(info.Name(), "apis") {
			// 在当前apis目录下，查找所有的Go文件
			err := filepath.Walk(path, func(subPath string, subInfo os.FileInfo, subErr error) error {
				if subErr != nil {
					return subErr
				}

				// 只处理Go文件
				if !subInfo.IsDir() && strings.HasSuffix(subInfo.Name(), ".go") {
					goFiles = append(goFiles, subPath)
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return goFiles, nil
}
