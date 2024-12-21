package main

import (
	"fmt"
	"go-admin/core/global"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	//outputFile := "./handlers.go"

	handlers, err := getApiDescriptions()
	if err != nil {
		panic(err)
	}
	// 将 handlers 写入到 handlers_gen.go 文件中
	err = writeToFile("app/admin/sys/models/sys_api_gen_desc.go", handlers)
	if err != nil {
		panic(err)
	}

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

// 将路由信息写入文件
func writeToFile(filePath string, handlers map[string]string) error {
	fileContent := "// Package models 下方代码由go generate自动生成，勿改动\n"
	fileContent += "package models\n\n"
	fileContent += "var ApiDescMap = map[string]string{\n"
	for handler, description := range handlers {
		fileContent += fmt.Sprintf("\t\"%s\": \"%s\",\n", handler, description)
	}
	fileContent += "}\n"
	return os.WriteFile(filePath, []byte(fileContent), 0644)
}
