package jwtauth

import (
	"errors"
	jwt "github.com/appleboy/gin-jwt/v3"
	"github.com/appleboy/gin-jwt/v3/core"
	"github.com/casbin/casbin/v2/util"
	"github.com/gin-gonic/gin"
	jwtIn "github.com/golang-jwt/jwt/v5"
	"go-admin/config/base/constant"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/config"
	"go-admin/core/dto/response"
	"go-admin/core/lang"
	"go-admin/core/middleware/auth/authdto"
	"go-admin/core/middleware/auth/casbin"
	"go-admin/core/runtime"
	"go-admin/core/utils/encrypt"
	"go-admin/core/utils/idgen"
	"go-admin/core/utils/log"
	"go-admin/core/utils/strutils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const JwtPayloadKey = "JWT_PAYLOAD"
const JWTLoginPrefix = "admin:jwt"
const JWTBlacklistPrefix = "admin:jwt:blacklist"
const JWTDevicesPrefix = "admin:jwt:devices"
const JWTActivityPrefix = "admin:jwt:activity"
const JTI = "jti"
const EXP = "exp"

type JwtAuth struct {
	mw                *jwt.GinJWTMiddleware
	enableDeviceCheck bool
	enableBlacklist   bool
	maxDevices        int
}

func NewJwtAuth() (*JwtAuth, error) {
	if config.AuthConfig.Secret == "" {
		return nil, errors.New("jwt secret is required")
	}

	timeout := time.Duration(config.AuthConfig.Timeout) * time.Second
	if timeout <= 0 {
		timeout = time.Hour
	}

	jwtAuth := &JwtAuth{
		enableDeviceCheck: config.AuthConfig.EnableDeviceCheck,
		enableBlacklist:   config.AuthConfig.EnableBlacklist,
		maxDevices:        max(1, config.AuthConfig.MaxDeviceCount),
	}

	mw, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            config.ApplicationConfig.Name,
		SigningAlgorithm: "HS256",
		Key:              []byte(config.AuthConfig.Secret),
		Timeout:          timeout,
		MaxRefresh:       timeout,

		IdentityKey:     authdto.LoginUserId,
		Authenticator:   jwtAuth.Authenticator,
		Authorizer:      jwtAuth.Authorizer,
		PayloadFunc:     jwtAuth.PayloadFunc,
		IdentityHandler: jwtAuth.IdentityHandler,

		LoginResponse:   jwtAuth.LoginResponse,
		LogoutResponse:  jwtAuth.LogoutResponse,
		RefreshResponse: jwtAuth.RefreshResponse,
		Unauthorized:    jwtAuth.Unauthorized,

		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		SendCookie:    false,
		TimeFunc:      time.Now,
	})
	if err != nil {
		return nil, err
	}

	jwtAuth.mw = mw
	return jwtAuth, nil
}

func (j *JwtAuth) Login(c *gin.Context) {
	j.mw.LoginHandler(c)
}

func (j *JwtAuth) Logout(c *gin.Context) {
	j.mw.LogoutHandler(c)
}

func (j *JwtAuth) LoginResponse(c *gin.Context, token *core.Token) {
	lg := lang.GetAcceptLanguage(c)
	userID := c.GetInt64(authdto.LoginUserId)
	userIDStr := strconv.FormatInt(userID, 10)

	if config.ApplicationConfig.IsSingleLogin {
		j.RevokeAllTokens(userIDStr)
	}

	// 添加设备信息
	if j.enableDeviceCheck {
		deviceFP := j.extractDeviceFingerprint(c)
		// 记录设备
		if deviceFP == "" {
			log.GetRequestLogger(c).Warn("device fingerprint empty")
		} else {
			j.recordDevice(userID, deviceFP)
		}
	}

	// set 用于单点登录，更新登录信息
	err := runtime.RuntimeConfig.GetCacheAdapter().Set(
		JWTLoginPrefix,
		userIDStr,
		token.AccessToken,
		config.AuthConfig.Timeout,
	)
	if err != nil {
		j.unauthorized(c, http.StatusUnauthorized, baseLang.AuthErrLogCode, lang.MsgErrf(baseLang.AuthErrLogCode, lg, err).Error())
		return
	}

	// 缓存记录用户最新登录状态
	j.updateLastActivity(userIDStr)

	userName := c.GetString(authdto.UserName)
	response.OK(c, gin.H{
		"token":    token.AccessToken,
		"username": userName,
		"expire":   config.AuthConfig.Timeout,
	}, http.StatusOK, lang.MsgByCode(baseLang.SuccessCode, lg))
}

func (j *JwtAuth) LogoutResponse(c *gin.Context) {
	_, err := j.RevokeToken(c)
	if err != nil {
		log.GetRequestLogger(c).Warnf("Failed to revoke token during logout: %v", err)
	}
	response.OK(c, nil, baseLang.SysUseLogoutSuccessCode, lang.MsgByCode(baseLang.SysUseLogoutSuccessCode, lang.GetAcceptLanguage(c)))
}

// RevokeToken 撤销Token
func (j *JwtAuth) RevokeToken(c *gin.Context) (int, error) {
	claims, err := j.mw.GetClaimsFromJWT(c)
	if err != nil {
		return baseLang.AuthErr, lang.MsgErr(baseLang.AuthErr, lang.GetAcceptLanguage(c))
	}
	userIDStr, errCode, err := j.GetUserIdStr(c)
	if err != nil {
		return errCode, err
	}

	// 1. 清除单点登录缓存
	if config.ApplicationConfig.IsSingleLogin {
		j.RevokeAllTokens(userIDStr)
	}

	// 2. 将token加入黑名单
	if j.enableBlacklist {
		tokenID, ok := claims[JTI].(string)
		if ok && tokenID != "" {
			// 黑名单有效期比token短
			blacklistTTL := j.mw.Timeout
			exp, ok := claims[EXP].(float64)
			if ok {
				expireTime := time.Unix(int64(exp), 0)
				remaining := time.Until(expireTime)
				if remaining > 0 {
					// 使用剩余时间 + 缓冲（如5分钟）
					blacklistTTL = remaining + time.Minute*5
				}
			}

			runtime.RuntimeConfig.GetCacheAdapter().Set(
				JWTBlacklistPrefix,
				tokenID,
				"1",
				int(blacklistTTL.Seconds()),
			)
		}
	}

	// 3. 清除设备记录中的当前设备
	if j.enableDeviceCheck {
		currentDeviceFP := j.extractDeviceFingerprint(c)
		savedDeviceFP, ok := claims[authdto.DeviceFingerprint].(string)

		if ok && currentDeviceFP != "" && savedDeviceFP != "" && currentDeviceFP == savedDeviceFP {
			j.removeDevice(userIDStr, currentDeviceFP)
		}
	}

	return http.StatusOK, nil
}

func (j *JwtAuth) Get(c *gin.Context, key string) (interface{}, int, error) {
	var err error
	defer func() {
		if err != nil {
			rLog := log.GetRequestLogger(c)
			rLog.Error(strutils.GetCurrentTimeStr() + " [ERROR] " + c.Request.Method + " " + c.Request.URL.Path + " Get no " + key)
		}
	}()
	data := jwt.ExtractClaims(c)
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

func (j *JwtAuth) GetUserIdStr(c *gin.Context) (string, int, error) {
	result, respCode, err := j.Get(c, authdto.LoginUserId)
	if err != nil {
		return "", respCode, err
	}
	return strconv.Itoa(int(result.(float64))), baseLang.SuccessCode, nil
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
	return j.mw.MiddlewareFunc()
}

func (j *JwtAuth) RefreshHandler(c *gin.Context) {
	j.mw.RefreshHandler(c)
}

func (j *JwtAuth) authCheck(c *gin.Context) bool {
	rLog := log.GetRequestLogger(c)
	claims, err := j.mw.CheckIfTokenExpire(c)
	if err != nil {
		rLog.Error(err)
		return false
	}
	userIDStr, _, err := j.GetUserIdStr(c)
	if err != nil {
		rLog.Error(err)
		return false
	}

	tokenOld, err := j.mw.ParseToken(c)
	if err != nil {
		rLog.Error(err)
		return false
	}

	if config.ApplicationConfig.IsSingleLogin {
		// 从缓存获取该用户最新的token
		cacheToken, _ := runtime.RuntimeConfig.GetCacheAdapter().Get(JWTLoginPrefix, userIDStr)
		if cacheToken != tokenOld.Raw {
			rLog.Error(errors.New("only support single login"))
			j.RevokeAllTokens(userIDStr)
			return false
		}
	}

	// 3. check blocklist
	if j.enableBlacklist {
		jti, ok := claims[JTI].(string)
		if ok && jti != "" {
			isBlacklisted := j.getCacheString(JWTBlacklistPrefix, jti)
			if isBlacklisted != "" {
				rLog.Error(errors.New("token is blacklisted"))
				return false
			}
		}
	}

	// 4. device check
	if j.enableDeviceCheck {
		currentDeviceFP := j.extractDeviceFingerprint(c)
		savedDeviceFP, ok := claims[authdto.DeviceFingerprint].(string)
		if currentDeviceFP == "" || savedDeviceFP == "" || !ok || currentDeviceFP != savedDeviceFP {
			rLog.Error(errors.New("device check error"))
			return false
		}

		// 获取设备列表
		devices := j.getCacheString(JWTDevicesPrefix, userIDStr)

		// 1). 检查当前设备是否在允许的设备列表中
		deviceExists := false
		deviceList := strings.Split(devices, ",")
		for _, d := range deviceList {
			if d == currentDeviceFP {
				deviceExists = true
				break
			}
		}

		// 2). 允许多设备登录，但设备列表中不存在当前设备，则判断是否增加设备
		if !deviceExists {
			// 检查设备数量
			if devices == "" {
				// 无设备，直接添加
				deviceList = []string{currentDeviceFP}
			} else {
				// 有设备，则判断数量
				if len(deviceList) >= j.maxDevices {
					rLog.Error(errors.New("too many devices login"))
					return false
				}
				// 添加新设备到列表
				deviceList = append(deviceList, currentDeviceFP)
			}
			runtime.RuntimeConfig.GetCacheAdapter().Set(
				JWTDevicesPrefix,
				userIDStr,
				strings.Join(deviceList, ","),
				config.AuthConfig.Timeout,
			)

		}
		// 设备已在列表中，允许访问
	}

	// 4. 更新最后活动时间
	j.updateLastActivity(userIDStr)
	return true
}

// updateLastActivity 更新最后活动时间
func (j *JwtAuth) updateLastActivity(userID string) {
	// 可以记录用户最后活动时间，用于会话管理
	runtime.RuntimeConfig.GetCacheAdapter().Set(
		JWTActivityPrefix,
		userID,
		strconv.FormatInt(time.Now().Unix(), 10),
		config.AuthConfig.Timeout,
	)
}

// AuthCheckRoleMiddlewareFunc 权限检查中间件
func (j *JwtAuth) AuthCheckRoleMiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Get(JwtPayloadKey)
		v := data.(jwtIn.MapClaims)
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

// GetUserDevices 获取用户的所有设备
func (j *JwtAuth) GetUserDevices(userID string) []string {
	devices := j.getCacheString(JWTDevicesPrefix, userID)
	if devices == "" {
		return []string{}
	}

	return strings.Split(devices, ",")
}

// RevokeAllTokens 撤销用户的所有Token
func (j *JwtAuth) RevokeAllTokens(userID string) {
	// 1. 清除单点登录缓存
	_ = runtime.RuntimeConfig.GetCacheAdapter().Del(JWTLoginPrefix, userID)

	// 2. 清除设备记录
	_ = runtime.RuntimeConfig.GetCacheAdapter().Del(JWTDevicesPrefix, userID)

	// 3. 清除活动记录
	_ = runtime.RuntimeConfig.GetCacheAdapter().Del(JWTActivityPrefix, userID)
}

func (j *JwtAuth) unauthorized(c *gin.Context, httpCode, errCode int, message string) {
	c.Header("WWW-Authenticate", "JWT realm=\""+j.mw.Realm+"\"")
	if !j.mw.DisabledAbort {
		c.Abort()
	}

	j.mw.Unauthorized(c, httpCode, strconv.Itoa(errCode)+"_"+message)
}

// recordDevice 记录设备
func (j *JwtAuth) recordDevice(userID int64, deviceFP string) {
	if !j.enableDeviceCheck {
		return
	}

	userIDStr := strconv.FormatInt(userID, 10)

	// 获取现有设备列表
	devices := j.getCacheString(JWTDevicesPrefix, userIDStr)
	var deviceList []string

	if devices != "" {
		deviceList = strings.Split(devices, ",")
	}

	// 添加新设备
	found := false
	for _, d := range deviceList {
		if d == deviceFP {
			found = true
			break
		}
	}

	if !found {
		deviceList = append(deviceList, deviceFP)

		// 限制设备数量
		if len(deviceList) > j.maxDevices {
			deviceList = deviceList[len(deviceList)-j.maxDevices:]
		}

		runtime.RuntimeConfig.GetCacheAdapter().Set(
			JWTDevicesPrefix,
			userIDStr,
			strings.Join(deviceList, ","),
			config.AuthConfig.Timeout,
		)
	}
}

// extractDeviceFingerprint 提取设备指纹
func (j *JwtAuth) extractDeviceFingerprint(c *gin.Context) string {
	userAgent := c.Request.UserAgent()
	clientIP := c.ClientIP()

	// 如果这些为空，生成默认值
	if userAgent == "" {
		userAgent = "unknown"
	}
	if clientIP == "" {
		clientIP = "0.0.0.0"
	}

	// 确保有足够的数据生成指纹
	data := userAgent + "|" + clientIP

	acceptLanguage := c.GetHeader("Accept-Language")
	if acceptLanguage != "" {
		data += "|" + acceptLanguage
	}

	acceptEncoding := c.GetHeader("Accept-Encoding")
	if acceptEncoding != "" {
		data += "|" + acceptEncoding
	}

	return encrypt.SHA256VString(data)
}

// removeDevice 移除设备
func (j *JwtAuth) removeDevice(userID string, deviceFP string) {
	devices := j.getCacheString(JWTDevicesPrefix, userID)
	if devices == "" {
		return
	}

	deviceList := strings.Split(devices, ",")
	var newList []string

	for _, d := range deviceList {
		if d != deviceFP {
			newList = append(newList, d)
		}
	}

	if len(newList) > 0 {
		runtime.RuntimeConfig.GetCacheAdapter().Set(
			JWTDevicesPrefix,
			userID,
			strings.Join(newList, ","),
			config.AuthConfig.Timeout,
		)
	} else {
		runtime.RuntimeConfig.GetCacheAdapter().Del(JWTDevicesPrefix, userID)
	}
}

func (j *JwtAuth) getCacheString(prefix, key string) string {
	val, err := runtime.RuntimeConfig.GetCacheAdapter().Get(prefix, key)
	if err != nil || val == "" {
		return ""
	}
	return val
}

func (j *JwtAuth) IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return map[string]interface{}{
		authdto.LoginUserId: claims[authdto.LoginUserId],
		authdto.RoleKey:     claims[authdto.RoleKey],
		authdto.UserName:    claims[authdto.UserName],
		authdto.DataScope:   claims[authdto.DataScope],
		authdto.RoleId:      claims[authdto.RoleId],
		authdto.DeptId:      claims[authdto.DeptId],
	}
}

func (j *JwtAuth) PayloadFunc(data interface{}) jwtIn.MapClaims {
	claims := jwtIn.MapClaims{}

	if v, ok := data.(map[string]interface{}); ok {
		claims[authdto.LoginUserId] = v[authdto.LoginUserId]
		claims[authdto.RoleKey] = v[authdto.RoleKey]
		claims[authdto.UserName] = v[authdto.UserName]
		claims[authdto.DataScope] = v[authdto.DataScope]
		claims[authdto.RoleId] = v[authdto.RoleId]
		claims[authdto.DeptId] = v[authdto.DeptId]
		claims[authdto.JTI] = v[authdto.JTI]
		claims[authdto.DeviceFingerprint] = v[authdto.DeviceFingerprint]
		claims[authdto.LoginIP] = v[authdto.LoginIP]
		claims[authdto.UserAgent] = v[authdto.UserAgent]
	}

	return claims
}

func (j *JwtAuth) Authenticator(c *gin.Context) (interface{}, error) {
	userId, b := c.Get(authdto.LoginUserId)
	if !b || userId == nil {
		return nil, errors.New("incorrect Username or Password")
	}

	roleId, _ := c.Get(authdto.RoleId)
	roleKey, _ := c.Get(authdto.RoleKey)
	deptId, _ := c.Get(authdto.DeptId)
	userName, _ := c.Get(authdto.UserName)
	dataScope, _ := c.Get(authdto.DataScope)

	resp := map[string]interface{}{
		authdto.LoginUserId:       userId,
		authdto.RoleKey:           roleKey,
		authdto.UserName:          userName,
		authdto.DataScope:         dataScope,
		authdto.RoleId:            roleId,
		authdto.DeptId:            deptId,
		authdto.JTI:               idgen.UUID(),
		authdto.DeviceFingerprint: j.extractDeviceFingerprint(c),
		authdto.LoginIP:           c.ClientIP(),
		authdto.UserAgent:         c.Request.UserAgent(),
	}
	return resp, nil
}

func (j *JwtAuth) Authorizer(c *gin.Context, data interface{}) bool {
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
		if deptId != nil {
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
		return j.authCheck(c)
	}
	return false
}

func (j *JwtAuth) RefreshResponse(c *gin.Context, token *core.Token) {
	c.JSON(http.StatusOK, gin.H{
		"requestId": strutils.GenerateMsgIDFromContext(c),
		"msg":       "",
		"code":      http.StatusOK,
		"data": gin.H{
			"token":  token,
			"expire": token.ExpiresAt,
		},
	})
}

func (j *JwtAuth) Unauthorized(c *gin.Context, httpCode int, message string) {
	_, _ = j.RevokeToken(c)
	temp := strings.SplitN(message, "_", 1)
	errCode := httpCode
	if len(temp) == 2 {
		code, err := strconv.ParseInt(temp[0], 10, 64)
		if err == nil {
			errCode = int(code)
			message = temp[1]
		}
	}
	log.GetRequestLogger(c).Errorf("Cache set JWTLoginPrefix failed: %d-%s", errCode, message)
	response.ErrorByHttpCode(c, httpCode, errCode, message)
}
