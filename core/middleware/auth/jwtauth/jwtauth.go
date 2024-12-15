package jwtauth

import (
	"fmt"
	"github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
	"go-admin/config/base/constant"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/config"
	"go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware/auth/authdto"
	"go-admin/core/middleware/auth/casbin"
	"go-admin/core/runtime"
	"go-admin/core/utils/log"
	"go-admin/core/utils/strutils"
	"net/http"
	"time"
)

var jwtAuthMiddleware = &GinJWTMiddleware{}

type JwtAuth struct{}

func (j *JwtAuth) Init() {
	timeout := time.Hour
	if config.ApplicationConfig.Mode == "dev" {
		timeout = time.Duration(876010) * time.Hour
	} else {
		if config.AuthConfig.Timeout != 0 {
			timeout = time.Duration(config.AuthConfig.Timeout) * time.Second
		}
	}
	var err error
	jwtAuthMiddleware, err = New(&GinJWTMiddleware{
		Realm:           config.ApplicationConfig.Name,
		Key:             []byte(config.AuthConfig.Secret),
		Timeout:         timeout,
		MaxRefresh:      time.Hour,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
		Authenticator:   Authenticator,
		Authorizator:    Authorizator,
		Unauthorized:    Unauthorized,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   authdto.HeaderTokenName,
		TimeFunc:        time.Now,
	}) //TokenHeadName必须有，不能为空，否则权限识别异常
	if err != nil {
		log.Errorf(fmt.Sprintf("JWT Init Error, %s", err.Error()))
	}
}

func (j *JwtAuth) Login(c *gin.Context) {
	jwtAuthMiddleware.LoginHandler(c)
}

func (j *JwtAuth) Logout(c *gin.Context) {

}

func (j *JwtAuth) Get(c *gin.Context, key string) (interface{}, int, error) {
	var err error
	defer func() {
		if err != nil {
			rLog := log.GetRequestLogger(c)
			rLog.Error(strutils.GetCurrentTimeStr() + " [ERROR] " + c.Request.Method + " " + c.Request.URL.Path + " Get no " + key)
		}
	}()
	data := ExtractClaims(c)
	if data[key] != nil {
		return data[key], baseLang.SuccessCode, nil
	}
	err = lang.MsgErr(baseLang.AuthErr, lang.GetAcceptLanguage(c))
	return nil, baseLang.AuthErr, err
}

func (j *JwtAuth) GetUserId(c *gin.Context) (int64, int, error) {
	result, respCode, err := j.Get(c, authdto.LoginUserId)
	if err != nil {
		return 0, respCode, err
	}
	return int64(result.(float64)), baseLang.SuccessCode, nil
}

func (j *JwtAuth) GetRoleId(c *gin.Context) (int64, int, error) {
	result, respCode, err := j.Get(c, authdto.RoleId)
	if err != nil {
		return 0, respCode, err
	}
	return int64(result.(float64)), baseLang.SuccessCode, nil
}

func (j *JwtAuth) GetRoleKey(c *gin.Context) string {
	result, _, _ := j.Get(c, authdto.RoleKey)
	if result == nil {
		return ""
	}
	return result.(string)
}

func (j *JwtAuth) GetDeptId(c *gin.Context) (int64, int, error) {
	result, respCode, err := j.Get(c, authdto.DeptId)
	if err != nil {
		return 0, respCode, err
	}
	return int64(result.(float64)), baseLang.SuccessCode, nil
}

func (j *JwtAuth) GetUserName(c *gin.Context) string {
	result, _, _ := j.Get(c, authdto.UserName)
	if result == nil {
		return ""
	}
	return result.(string)
}

func (j *JwtAuth) AuthMiddlewareFunc() gin.HandlerFunc {
	return jwtAuthMiddleware.MiddlewareFunc()
}

// AuthCheckRoleMiddlewareFunc 权限检查中间件
func (j *JwtAuth) AuthCheckRoleMiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get(JwtPayloadKey)
		v := data.(MapClaims)
		roleKey := v[authdto.RoleKey]

		rLog := log.GetRequestLogger(c)
		var res, casbinExclude bool
		var err error
		//检查权限
		if roleKey == constant.RoleKeyAdmin {
			res = true
			c.Next()
			return
		}
		for _, i := range casbin.CasbinExclude {
			if util.KeyMatch2(c.Request.URL.Path, i.Url) && c.Request.Method == i.Method {
				casbinExclude = true
				break
			}
		}
		if casbinExclude {
			rLog.Infof("Casbin exclusion, no validation method:%s path:%s", c.Request.Method, c.Request.URL.Path)
			c.Next()
			return
		}
		e := runtime.RuntimeConfig.GetCasbinKey(c.Request.Host)
		res, err = e.Enforce(roleKey, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			rLog.Errorf("AuthCheckRole error:%s method:%s path:%s", err, c.Request.Method, c.Request.URL.Path)
			response.Error(c, baseLang.ServerErr, lang.MsgByCode(baseLang.ServerErr, lang.GetAcceptLanguage(c)))
			return
		}

		if res {
			rLog.Infof("isTrue: %v role: %s method: %s path: %s", res, roleKey, c.Request.Method, c.Request.URL.Path)
			c.Next()
		} else {
			rLog.Warnf("isTrue: %v role: %s method: %s path: %s message: %s", res, roleKey, c.Request.Method, c.Request.URL.Path, "The current request has no permission. Please confirm it!")
			response.Error(c, baseLang.ForbitErr, lang.MsgByCode(baseLang.ForbitErr, lang.GetAcceptLanguage(c)))
			c.Abort()
			return
		}
	}
}

func PayloadFunc(data interface{}) MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		userId, _ := v[authdto.LoginUserId]
		roleKey, _ := v[authdto.RoleKey]
		userName, _ := v[authdto.UserName]
		dataScope, _ := v[authdto.DataScope]
		roleId, _ := v[authdto.RoleId]
		deptId, _ := v[authdto.DeptId]

		return MapClaims{
			authdto.LoginUserId: userId,
			authdto.RoleKey:     roleKey,
			authdto.UserName:    userName,
			authdto.DataScope:   dataScope,
			authdto.RoleId:      roleId,
			authdto.DeptId:      deptId,
		}
	}
	return MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := ExtractClaims(c)
	return map[string]interface{}{
		authdto.LoginUserId: claims[authdto.LoginUserId],
		authdto.RoleKey:     claims[authdto.RoleKey],
		authdto.UserName:    claims[authdto.UserName],
		authdto.DataScope:   claims[authdto.DataScope],
		authdto.RoleId:      claims[authdto.RoleId],
		authdto.DeptId:      claims[authdto.DeptId],
	}
}

func Authenticator(c *gin.Context) (interface{}, error) {
	userId, b := c.Get(authdto.LoginUserId)
	if !b || userId == nil {
		return nil, ErrFailedAuthentication
	}

	roleId, _ := c.Get(authdto.RoleId)
	roleKey, _ := c.Get(authdto.RoleKey)
	deptId, _ := c.Get(authdto.DeptId)
	userName, _ := c.Get(authdto.UserName)
	dataScope, _ := c.Get(authdto.DataScope)

	resp := map[string]interface{}{
		authdto.LoginUserId: userId,
		authdto.RoleKey:     roleKey,
		authdto.UserName:    userName,
		authdto.DataScope:   dataScope,
		authdto.RoleId:      roleId,
		authdto.DeptId:      deptId,
	}
	return resp, nil
}

func Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(map[string]interface{}); ok {
		userId, _ := v[authdto.LoginUserId]
		if userId != nil {
			c.Set(authdto.LoginUserId, int64(userId.(float64))) //这里一定要用string保存userId，以防取出Interface转换复杂
		}
		roleKey, _ := v[authdto.RoleKey]
		if roleKey != nil {
			c.Set(authdto.RoleKey, roleKey)
		}
		roleId, _ := v[authdto.RoleId]
		if roleId != nil {
			c.Set(authdto.RoleId, int64(roleId.(float64))) //这里一定要用string保存userId，以防取出Interface转换复杂
		}
		deptId, _ := v[authdto.DeptId]
		if roleId != nil {
			c.Set(authdto.DeptId, int64(deptId.(float64))) //这里一定要用string保存userId，以防取出Interface转换复杂
		}
		userName, _ := v[authdto.UserName]
		if userName != nil {
			c.Set(authdto.UserName, userName)
		}
		dataScope, _ := v[authdto.DataScope]
		if dataScope != nil {
			c.Set(authdto.DataScope, dataScope)
		}
		return true
	}
	return false
}

func Unauthorized(c *gin.Context, code int, message string) {
	resp := &authdto.Resp{
		Msg:  message,
		Code: code,
	}
	c.JSON(http.StatusOK, resp)
}
