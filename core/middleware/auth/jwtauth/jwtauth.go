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
const DeviceFingerprint = "dev_fp"
const TokenID = "jti"

type SecurityConfig struct {
	DeviceCheckEnabled bool // 设备检查
	TokenBlacklist     bool // Token黑名单
	MaxDevicesPerUser  int  // 每个用户最大设备数
}
type JwtAuth struct {
	ginJwtMiddleware *jwt.GinJWTMiddleware
	securityConfig   SecurityConfig
}

var jwtAuth JwtAuth

func (j *JwtAuth) Init() error {
	timeout := time.Hour
	if config.AuthConfig.Timeout != 0 {
		timeout = time.Duration(config.AuthConfig.Timeout) * time.Second
	}

	ginJwtMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            config.ApplicationConfig.Name,
		SigningAlgorithm: "HS256",
		Key:              []byte(config.AuthConfig.Secret),
		Timeout:          timeout,
		MaxRefresh:       time.Hour,
		Authenticator:    Authenticator,
		Authorizer:       Authorizer,
		Unauthorized:     Unauthorized,
		LoginResponse:    LoginResponse,
		LogoutResponse:   LogoutResponse,
		RefreshResponse:  RefreshResponse,
		PayloadFunc:      PayloadFunc,
		IdentityHandler:  IdentityHandler,
		IdentityKey:      authdto.LoginUserId,
		TokenLookup:      "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:    "Bearer",
		SendCookie:       true,
		TimeFunc:         time.Now,
	})
	if err != nil {
		return err
	}

	securityConfig := SecurityConfig{
		DeviceCheckEnabled: config.AuthConfig.DeviceCheck,
		TokenBlacklist:     config.AuthConfig.TokenBlacklist,
		MaxDevicesPerUser:  config.AuthConfig.MaxDevicesPerUser,
	}

	if securityConfig.MaxDevicesPerUser <= 0 {
		securityConfig.MaxDevicesPerUser = 5
	}
	jwtAuth = JwtAuth{
		ginJwtMiddleware: ginJwtMiddleware,
		securityConfig:   securityConfig,
	}
	return err
}

func (j *JwtAuth) Login(c *gin.Context) {
	j.ginJwtMiddleware.LoginHandler(c)
}

func (j *JwtAuth) Logout(c *gin.Context) {
	j.ginJwtMiddleware.LogoutHandler(c)
	Unauthorized(c, http.StatusOK, strconv.Itoa(baseLang.SuccessCode)+"_"+lang.MsgByCode(baseLang.SuccessCode, lang.GetAcceptLanguage(c)))
}

// RevokeToken 撤销Token
func (j *JwtAuth) RevokeToken(c *gin.Context) (int, error) {
	claims, err := j.ginJwtMiddleware.GetClaimsFromJWT(c)
	if err != nil {
		return baseLang.AuthErr, lang.MsgErr(baseLang.AuthErr, lang.GetAcceptLanguage(c))
	}
	userIDStr, errCode, err := j.GetUserIdStr(c)
	if err != nil {
		return errCode, err
	}

	// 1. 清除单点登录缓存
	runtime.RuntimeConfig.GetCacheAdapter().Del(JWTLoginPrefix, userIDStr)

	// 2. 将token加入黑名单
	if j.securityConfig.TokenBlacklist {
		tokenID, ok := claims[TokenID].(string)
		if ok && tokenID != "" {
			// 黑名单有效期比token短
			blacklistTTL := j.ginJwtMiddleware.Timeout
			exp, ok := claims["exp"].(float64)
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
	if j.securityConfig.DeviceCheckEnabled {
		currentDeviceFP := j.extractDeviceFingerprint(c)
		savedDeviceFP, ok := claims[DeviceFingerprint].(string)

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
	return strconv.Itoa(int(result.(float64))), baseLang.SuccessCode, lang.MsgErr(baseLang.SuccessCode, lang.GetAcceptLanguage(c))
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
	return func(c *gin.Context) {
		err := j.authCheck(c)
		if err != nil {
			lg := lang.GetAcceptLanguage(c)
			jwtAuth.unauthorized(c, http.StatusUnauthorized, http.StatusUnauthorized, lang.MsgErrf(http.StatusUnauthorized, lg, err).Error())
			c.Abort()
		}
		j.ginJwtMiddleware.MiddlewareFunc()
	}
}

func (j *JwtAuth) RefreshHandler(c *gin.Context) {
	j.ginJwtMiddleware.RefreshHandler(c)
}

func (j *JwtAuth) authCheck(c *gin.Context) error {
	claims, err := jwtAuth.ginJwtMiddleware.CheckIfTokenExpire(c)
	if err != nil {
		return err
	}
	userIDStr, _, err := j.GetUserIdStr(c)
	if err != nil {
		return err
	}

	tokenOld, err := j.ginJwtMiddleware.ParseToken(c)
	if err != nil {
		return err
	}

	if config.ApplicationConfig.IsSingleLogin {
		// 从缓存获取该用户最新的token
		savedToken := j.getCacheString(JWTLoginPrefix, userIDStr)

		if savedToken != tokenOld.Raw {
			return errors.New("only support single login")
		}
	}

	// 3. check blocklist
	if j.securityConfig.TokenBlacklist {
		tokenID, ok := claims[TokenID].(string)
		if ok && tokenID != "" {
			isBlacklisted := j.getCacheString(JWTBlacklistPrefix, tokenID)
			if isBlacklisted != "" {
				return errors.New("token is blacklisted")
			}
		}
	}

	// 4. device check
	if j.securityConfig.DeviceCheckEnabled {
		currentDeviceFP := j.extractDeviceFingerprint(c)
		savedDeviceFP, ok := claims[DeviceFingerprint].(string)
		if currentDeviceFP == "" || savedDeviceFP == "" || !ok || currentDeviceFP != savedDeviceFP {
			return errors.New("device check error")
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
				if len(deviceList) >= j.securityConfig.MaxDevicesPerUser {
					return errors.New("too many devices login")
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
	return nil
}

// updateLastActivity 更新最后活动时间
func (j *JwtAuth) updateLastActivity(userID string) {
	// 可以记录用户最后活动时间，用于会话管理
	runtime.RuntimeConfig.GetCacheAdapter().Set(
		"admin:jwt:activity",
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

// RevokeUserAllTokens 撤销用户的所有Token
func (j *JwtAuth) RevokeUserAllTokens(userID string) {
	// 1. 清除单点登录缓存
	_ = runtime.RuntimeConfig.GetCacheAdapter().Del(JWTLoginPrefix, userID)

	// 2. 清除设备记录
	_ = runtime.RuntimeConfig.GetCacheAdapter().Del(JWTDevicesPrefix, userID)

	// 3. 清除活动记录
	_ = runtime.RuntimeConfig.GetCacheAdapter().Del("admin:jwt:activity", userID)
}

func (j *JwtAuth) unauthorized(c *gin.Context, httpCode, errCode int, message string) {
	c.Header("WWW-Authenticate", "JWT realm=\""+j.ginJwtMiddleware.Realm+"\"")
	if !j.ginJwtMiddleware.DisabledAbort {
		c.Abort()
	}

	j.ginJwtMiddleware.Unauthorized(c, httpCode, strconv.Itoa(errCode)+"_"+message)
}

// recordDevice 记录设备
func (j *JwtAuth) recordDevice(userID int64, deviceFP string) {
	if !j.securityConfig.DeviceCheckEnabled {
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
		if len(deviceList) > j.securityConfig.MaxDevicesPerUser {
			deviceList = deviceList[len(deviceList)-j.securityConfig.MaxDevicesPerUser:]
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

func PayloadFunc(data interface{}) jwtIn.MapClaims {
	if v, ok := data.(map[string]interface{}); ok {
		userId, _ := v[authdto.LoginUserId]
		roleKey, _ := v[authdto.RoleKey]
		userName, _ := v[authdto.UserName]
		dataScope, _ := v[authdto.DataScope]
		roleId, _ := v[authdto.RoleId]
		deptId, _ := v[authdto.DeptId]

		return jwtIn.MapClaims{
			authdto.LoginUserId: userId,
			authdto.RoleKey:     roleKey,
			authdto.UserName:    userName,
			authdto.DataScope:   dataScope,
			authdto.RoleId:      roleId,
			authdto.DeptId:      deptId,
		}
	}
	return jwtIn.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
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

func Authenticator(c *gin.Context) (interface{}, error) {
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
		authdto.LoginUserId: userId,
		authdto.RoleKey:     roleKey,
		authdto.UserName:    userName,
		authdto.DataScope:   dataScope,
		authdto.RoleId:      roleId,
		authdto.DeptId:      deptId,
	}
	return resp, nil
}

func Authorizer(c *gin.Context, data interface{}) bool {
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
		return true
	}
	return false
}

func LoginResponse(c *gin.Context, token *core.Token) {
	lg := lang.GetAcceptLanguage(c)

	userID, errCode, err := jwtAuth.GetUserId(c)
	if err != nil {
		jwtAuth.unauthorized(c, http.StatusUnauthorized, errCode, lang.MsgErrf(errCode, lg, err).Error())
		return
	}
	userIDStr := strconv.FormatInt(userID, 10)

	if config.ApplicationConfig.IsSingleLogin {
		_, _ = jwtAuth.RevokeToken(c)
	}

	// 添加设备信息
	if jwtAuth.securityConfig.DeviceCheckEnabled {
		deviceFP := jwtAuth.extractDeviceFingerprint(c)

		// 记录设备
		jwtAuth.recordDevice(userID, deviceFP)
	}

	// set
	err = runtime.RuntimeConfig.GetCacheAdapter().Set(
		JWTLoginPrefix,
		userIDStr,
		token.AccessToken,
		config.AuthConfig.Timeout,
	)
	if err != nil {
		jwtAuth.unauthorized(c, http.StatusUnauthorized, baseLang.AuthErrLogCode, lang.MsgErrf(baseLang.AuthErrLogCode, lg, err).Error())
		return
	}

	userName := jwtAuth.GetUserName(c)
	response.OK(c, gin.H{
		"data": gin.H{
			"token":    token.AccessToken,
			"username": userName,
			"expire":   token.ExpiresIn(),
			//"userInfo": userInfo,
		},
	}, http.StatusOK, lang.MsgByCode(baseLang.SuccessCode, lg))
}

func LogoutResponse(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func RefreshResponse(c *gin.Context, token *core.Token) {
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

func Unauthorized(c *gin.Context, httpCode int, message string) {
	_, _ = jwtAuth.RevokeToken(c)
	temp := strings.SplitN(message, "_", 1)
	errCode := httpCode
	if len(temp) == 2 {
		code, err := strconv.ParseInt(temp[0], 10, 64)
		if err == nil {
			errCode = int(code)
			message = temp[1]
		}

	}
	response.ErrorByHttpCode(c, httpCode, errCode, message)
}
