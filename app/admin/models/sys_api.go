package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-admin/app/admin/constant"
	"go-admin/core/runtime"
	"go-admin/core/utils/log"
	"go-admin/core/utils/storage"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var IsSync = false

type SysApi struct {
	Id          int        `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	Description string     `json:"description" gorm:"size:256;comment:功能描述"`
	Path        string     `json:"path" gorm:"size:128;comment:地址"`
	Method      string     `json:"method" gorm:"size:16;comment:请求类型"`
	ApiType     string     `json:"apiType" gorm:"size:16;comment:接口类型"`
	CreatedAt   *time.Time `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"comment:最后更新时间"`
	CreateBy    int64      `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy    int64      `json:"updateBy" gorm:"index;comment:更新者"`
}

func (SysApi) TableName() string {
	return "sys_api"
}

func SaveSysApi(message storage.Messager) (err error) {
	var apiCacheMap = make(map[string]bool)
	IsSync = true
	defer func() {
		apiCacheMap = nil
		IsSync = false
	}()

	var rb []byte
	rb, err = json.Marshal(message.GetValues())
	if err != nil {
		return err
	}

	var l runtime.Routers
	err = json.Unmarshal(rb, &l)
	if err != nil {
		return err
	}

	dbList := runtime.RuntimeConfig.GetDb()
	for _, d := range dbList {
		var list []SysApi
		err = d.Model(&SysApi{}).Find(&list).Error
		if err != nil {
			return errors.New(fmt.Sprintf("get Api list error: %s \r\n ", err.Error()))
		}
		for _, item := range list {
			apiCacheMap[item.Path+"-"+item.Method] = true
		}
		for _, v := range l.List {
			if v.HttpMethod == "HEAD" {
				continue
			}
			//缓存
			if apiCacheMap[v.RelativePath+"-"+v.HttpMethod] {
				continue
			}
			paths := strings.Split(v.RelativePath, "/")
			apiType := ""
			if len(paths) >= 4 {
				if strings.HasPrefix(paths[3], "sys") {
					apiType = constant.ApiTypeSys
				} else if strings.HasPrefix(paths[3], "plugin") {
					apiType = constant.ApiTypePlugin
				} else if strings.HasPrefix(paths[3], "app") {
					apiType = constant.ApiTypeApp
				}
			}
			newSysApi := SysApi{Path: v.RelativePath, Method: v.HttpMethod}
			if apiType != "" {
				newSysApi.ApiType = apiType
			}
			err = d.Debug().Model(&SysApi{}).Create(&newSysApi).Error
			if err != nil {
				log.Errorf("Models SaveSysApi error: %s \r\n ", err.Error())
				continue
			}
			apiCacheMap[v.RelativePath+"-"+v.HttpMethod] = true
		}
	}
	return nil
}

type ApiParseInfo struct {
	ClassName   string
	MethodName  string
	Description string
}

// ParseApiInfo 解析每个接口文件中类名、方法和注释
func ParseApiInfo(filePath string) (apiParseInfos []ApiParseInfo, err error) {
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
					// 提取普通方法描述注释
					description = strings.TrimSpace(strings.TrimPrefix(comment.Text, "// "+funcDecl.Name.Name))
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

func FindAllApiFileDirs(rootDir string) ([]string, error) {
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
