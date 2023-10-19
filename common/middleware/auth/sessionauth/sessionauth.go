package sessionauth

import (
	"encoding/json"
	"github.com/casbin/casbin/v2/util"
	uuid "github.com/satori/go.uuid"
	"go-admin/app/admin/constant"
	"go-admin/common/core"
	"go-admin/common/core/api"
	"go-admin/common/core/pkg"
	"go-admin/common/core/pkg/response"
	"go-admin/common/middleware/auth/authdto"
	"go-admin/common/middleware/auth/casbin"
	"go-admin/common/utils/i18n"
	"go-admin/common/utils/strutils"
	"go-admin/config/config"
	"go-admin/config/lang"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SessionLoginPrefixTmp = "admin:login:session:tmp" //登录中转
	SessionLoginPrefix    = "admin:login:session:user"
)

type SessionAuth struct{}

func (s *SessionAuth) Init() {}

func (s *SessionAuth) Login(c *gin.Context) {
	errResp := authdto.Resp{
		RequestId: strutils.GenerateMsgIDFromContext(c),
		Msg:       lang.MsgByCode(lang.RequestErr, i18n.GetAcceptLanguage(c)),
		Code:      lang.RequestErr,
		Data:      nil,
	}

	userId := c.GetInt64(authdto.LoginUserId)
	if userId <= 0 {
		c.JSON(lang.RequestErr, errResp)
		return
	}

	//获取sid，并用sid保存userId
	sid := uuid.NewV4().String()
	err := core.Runtime.GetCacheAdapter().Set(SessionLoginPrefixTmp, sid, userId, config.AuthConfig.Timeout)
	log := api.GetRequestLogger(c)
	if err != nil {
		log.Error(err)
		c.JSON(lang.RequestErr, errResp)
		return
	}
	if config.ApplicationConfig.IsSingleLogin {
		_ = core.Runtime.GetCacheAdapter().Del(SessionLoginPrefix, strconv.FormatInt(userId, 10))
	}

	//session信息
	roleId, _ := c.Get(authdto.RoleId)
	roleKey, _ := c.Get(authdto.RoleKey)
	deptId, _ := c.Get(authdto.DeptId)
	userName, _ := c.Get(authdto.UserName)
	dataScope, _ := c.Get(authdto.DataScope)
	sessionInfo, err := json.Marshal(map[string]interface{}{
		authdto.LoginUserId: userId,
		authdto.RoleKey:     roleKey,
		authdto.UserName:    userName,
		authdto.DataScope:   dataScope,
		authdto.RoleId:      roleId,
		authdto.DeptId:      deptId,
	})
	if err != nil {
		log.Error(err)
		c.JSON(lang.RequestErr, errResp)
		return
	}
	values := map[string]interface{}{
		sid: string(sessionInfo),
	}
	//用userId保存sid，记录登录状态（此操作可用于多点登录）
	err = core.Runtime.GetCacheAdapter().HashSet(config.AuthConfig.Timeout, SessionLoginPrefix, strconv.FormatInt(userId, 10), values)
	log = api.GetRequestLogger(c)
	if err != nil {
		log.Error(err)
		c.JSON(lang.RequestErr, errResp)
		return
	}
	userInfo, _ := c.Get(authdto.UserInfo)
	resp := authdto.Resp{
		RequestId: strutils.GenerateMsgIDFromContext(c),
		Msg:       "",
		Code:      http.StatusOK,
		Data: authdto.Data{
			Token:    sid,
			Expire:   time.Now().Add(time.Duration(config.AuthConfig.Timeout) * time.Second).Format(time.RFC3339),
			UserInfo: userInfo,
		},
	}
	c.JSON(http.StatusOK, resp)
}

func (s *SessionAuth) Logout(c *gin.Context) {
	userId := c.GetInt64(authdto.LoginUserId)
	if userId <= 0 {
		return
	}
	_ = core.Runtime.GetCacheAdapter().Del(SessionLoginPrefix, strconv.FormatInt(userId, 10))
	c.JSON(http.StatusOK, authdto.Resp{
		RequestId: strutils.GenerateMsgIDFromContext(c),
		Msg:       "",
		Code:      http.StatusOK,
		Data:      nil,
	})
}

func (s *SessionAuth) Get(c *gin.Context, key string) (interface{}, int, error) {
	var err error
	defer func() {
		if err != nil {
			log := api.GetRequestLogger(c)
			log.Error(pkg.GetCurrentTimeStr() + " [ERROR] " + c.Request.Method + " " + c.Request.URL.Path + " Get no " + key)
		}
	}()
	cache := core.Runtime.GetCacheAdapter()
	sid := strings.Replace(c.Request.Header.Get(authdto.HeaderAuthorization), authdto.HeaderTokenName+" ", "", -1)
	uid, err := cache.Get(SessionLoginPrefixTmp, sid)
	if sid == "" || uid == "" || err != nil {
		err = lang.MsgErr(lang.AuthErr, i18n.GetAcceptLanguage(c))
		return "", lang.AuthErr, err
	}
	userInfoStr, err := cache.HashGet(SessionLoginPrefix, uid, sid)
	userInfo := map[string]interface{}{}
	err = json.Unmarshal([]byte(userInfoStr), &userInfo)
	if err != nil || userInfo[key] == nil {
		return "", lang.AuthErr, lang.MsgErr(lang.AuthErr, i18n.GetAcceptLanguage(c))
	}
	return userInfo[key], lang.SuccessCode, nil
}

func (s *SessionAuth) GetUserId(c *gin.Context) (int64, int, error) {
	result, respCode, err := s.Get(c, authdto.LoginUserId)
	if err != nil {
		return 0, respCode, err
	}
	return int64(result.(float64)), respCode, err
}

func (s *SessionAuth) GetRoleId(c *gin.Context) (int64, int, error) {
	result, respCode, err := s.Get(c, authdto.RoleId)
	if err != nil {
		return 0, respCode, err
	}
	return int64(result.(float64)), respCode, err
}

func (s *SessionAuth) GetDeptId(c *gin.Context) (int64, int, error) {
	result, respCode, err := s.Get(c, authdto.DeptId)
	if err != nil {
		return 0, respCode, err
	}
	return int64(result.(float64)), respCode, err
}
func (s *SessionAuth) GetUserName(c *gin.Context) string {
	result, _, _ := s.Get(c, authdto.UserName)
	if result == nil {
		return ""
	}
	return result.(string)
}

func (s *SessionAuth) GetRoleKey(c *gin.Context) string {
	result, _, _ := s.Get(c, authdto.RoleKey)
	if result == nil {
		return ""
	}
	return result.(string)
}
func (s *SessionAuth) AuthMiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		cache := core.Runtime.GetCacheAdapter()
		sid := strings.Replace(c.Request.Header.Get(authdto.HeaderAuthorization), authdto.HeaderTokenName+" ", "", -1)
		isExist := cache.Exist(SessionLoginPrefixTmp, sid)
		errResp := authdto.Resp{
			RequestId: strutils.GenerateMsgIDFromContext(c),
			Msg:       lang.MsgByCode(lang.AuthErr, i18n.GetAcceptLanguage(c)),
			Code:      lang.AuthErr,
			Data:      nil,
		}
		if !isExist {
			c.JSON(lang.AuthErr, errResp)
			c.Abort()
			return
		}

		// 从session中获取用户id,第一次用于缓存拿到uid，第二次用uid检测sid是否有效，可用于多端登录
		uid, err := cache.Get(SessionLoginPrefixTmp, sid)
		if err != nil {
			c.JSON(lang.AuthErr, errResp)
			c.Abort()
			return
		}
		_, err = cache.HashGet(SessionLoginPrefix, uid, sid)
		if err != nil {
			c.JSON(lang.AuthErr, errResp)
			c.Abort()
			return
		}
		c.Set(authdto.LoginUserId, uid)
		_ = cache.Expire(SessionLoginPrefixTmp, sid, config.AuthConfig.Timeout)
		_ = cache.Expire(SessionLoginPrefix, uid, config.AuthConfig.Timeout)
		c.Next()
	}
}

func (s *SessionAuth) AuthCheckRoleMiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		roleKey := s.GetRoleKey(c)
		log := api.GetRequestLogger(c)
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
			log.Infof("Casbin exclusion, no validation method:%s path:%s", c.Request.Method, c.Request.URL.Path)
			c.Next()
			return
		}
		e := core.Runtime.GetCasbinKey(c.Request.Host)
		res, err = e.Enforce(roleKey, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			log.Errorf("AuthCheckRole error:%s method:%s path:%s", err, c.Request.Method, c.Request.URL.Path)
			response.Error(c, lang.ServerErr, lang.MsgByCode(lang.ServerErr, i18n.GetAcceptLanguage(c)))
			return
		}

		if res {
			log.Infof("isTrue: %v role: %s method: %s path: %s", res, roleKey, c.Request.Method, c.Request.URL.Path)
			c.Next()
		} else {
			log.Warnf("isTrue: %v role: %s method: %s path: %s message: %s", res, roleKey, c.Request.Method, c.Request.URL.Path, "The current request has no permission. Please confirm it!")
			response.Error(c, lang.ForbitErr, lang.MsgByCode(lang.ForbitErr, i18n.GetAcceptLanguage(c)))
			c.Abort()
			return
		}
	}
}
