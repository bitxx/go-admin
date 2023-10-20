package api

import (
	"errors"
	"fmt"
	"github.com/bitxx/logger/logbase"
	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go-admin/common/dto/response"
	"go-admin/common/dto/service"
	"go-admin/common/utils/ginutils"
	"go-admin/common/utils/langutils"
	"gorm.io/gorm"
	"net/http"
)

var DefaultLanguage = "zh-CN"

type Api struct {
	Context *gin.Context
	Logger  *logbase.Helper
	Orm     *gorm.DB
	Lang    string //语言 en 英文 zh-cn中文
	Errors  error
}

func (e *Api) AddError(err error) {
	if e.Errors == nil {
		e.Errors = err
	} else if err != nil {
		e.Logger.Error(err)
		e.Errors = fmt.Errorf("%v; %w", e.Errors, err)
	}
}

// MakeContext 设置http上下文
func (e *Api) MakeContext(c *gin.Context) *Api {
	e.Context = c
	e.Logger = GetRequestLogger(c)
	return e
}

// GetLogger 获取上下文提供的日志
func (e Api) GetLogger() *logbase.Helper {
	return GetRequestLogger(e.Context)
}

// Bind 参数校验
func (e *Api) Bind(d interface{}, bindings ...binding.Binding) *Api {
	var err error
	if len(bindings) == 0 {
		bindings = constructor.GetBindingForGin(d)
	}
	for i := range bindings {
		if bindings[i] == nil {
			err = e.Context.ShouldBindUri(d)
		} else {
			err = e.Context.ShouldBindWith(d, bindings[i])
		}
		if err != nil && err.Error() == "EOF" {
			e.Logger.Warn("request body is not present anymore. ")
			err = nil
			continue
		}
		if err != nil {
			e.AddError(err)
			break
		}
	}
	//vd.SetErrorFactory(func(failPath, msg string) error {
	//	return fmt.Errorf(`"validation failed: %s %s"`, failPath, msg)
	//})
	if err1 := vd.Validate(d); err1 != nil {
		e.AddError(err1)
	}
	return e
}

// GetOrm 获取Orm DB
func (e Api) GetOrm() (*gorm.DB, error) {
	db, err := ginutils.GetOrm(e.Context)
	if err != nil {
		e.Logger.Error(http.StatusInternalServerError, err, "数据库连接获取失败")
		return nil, err
	}
	return db, nil
}

// MakeOrm 设置Orm DB
func (e *Api) MakeOrm() *Api {
	var err error
	if e.Logger == nil {
		err = errors.New("at MakeOrm logger is nil")
		e.AddError(err)
		return e
	}
	db, err := ginutils.GetOrm(e.Context)
	if err != nil {
		e.Logger.Error(http.StatusInternalServerError, err, "数据库连接获取失败")
		e.AddError(err)
	}
	e.Orm = db
	return e
}

func (e *Api) MakeService(c *service.Service) *Api {
	c.Log = e.Logger
	c.Orm = e.Orm
	return e
}

// Error 通常错误数据处理
func (e Api) Error(code int, msg string) {
	response.Error(e.Context, code, msg)
}

func (e Api) DownloadZip(fileName string, data []byte) {
	response.Download(e.Context, data, fileName, "application/zip")
}

func (e Api) OKByCode(data interface{}, code int, msg string) {
	response.OKByCode(e.Context, data, code, msg)
}

// OK 通常成功数据处理
func (e Api) OK(data interface{}, msg string) {
	response.OK(e.Context, data, msg)
}

// PageOK 分页数据处理
func (e Api) PageOK(result, extend interface{}, count int64, pageIndex int, pageSize int, msg string) {
	response.PageOK(e.Context, result, extend, count, pageIndex, pageSize, msg)
}

func (e Api) DownloadExcel(fileName string, data []byte) {
	response.Download(e.Context, data, fileName, "application/vnd.ms-excel")
}

// Custom 兼容函数
func (e Api) Custom(data gin.H) {
	response.Custum(e.Context, data)
}

// getAcceptLanguage 获取当前语言
func (e *Api) getAcceptLanguage() string {
	languages := langutils.ParseAcceptLanguage(e.Context.GetHeader("Accept-Language"), nil)
	if len(languages) == 0 {
		return DefaultLanguage
	}
	return languages[0]
}
