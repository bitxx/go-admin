package jwtauth

import (
	"crypto/rsa"
	"errors"
	baseLang "go-admin/config/base/lang"
	"go-admin/core/config"
	"go-admin/core/lang"
	"go-admin/core/middleware/auth/authdto"
	"go-admin/core/runtime"
	"go-admin/core/utils/encrypt"
	"go-admin/core/utils/idgen"
	"go-admin/core/utils/log"
	"go-admin/core/utils/strutils"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const JwtPayloadKey = "JWT_PAYLOAD"
const JWTLoginPrefix = "admin:jwt"
const JWTBlacklistPrefix = "admin:jwt:blacklist"
const JWTDevicesPrefix = "admin:jwt:devices"
const DeviceFingerprint = "dev_fp"
const LoginIP = "login_ip"
const ClientInfo = "client_info"
const TokenID = "jti"

type MapClaims map[string]interface{}
type SecurityConfig struct {
	DeviceCheckEnabled bool // 设备检查
	TokenBlacklist     bool // Token黑名单
	AllowMultiDevices  bool // 允许多设备登录
	MaxDevicesPerUser  int  // 每个用户最大设备数
}

// GinJWTMiddleware provides a Json-Web-Token authentication implementation. On failure, a 401 HTTP response
// is returned. On success, the wrapped middleware is called, and the userID is made available as
// c.GetArticle("userID").(string).
// Users can get a token by posting a json request to LoginHandler. The token then needs to be passed in
// the Authentication header. Example: Authorization XXX_TOKEN_XXX
type GinJWTMiddleware struct {

	// signing algorithm - possible values are HS256, HS384, HS512
	// Optional, default is HS256.
	SigningAlgorithm string

	// Secret key used for signing. Required.
	Key []byte

	// Duration that a jwt token is valid. Optional, defaults to one hour.
	Timeout time.Duration

	// This field allows clients to refresh their token until MaxRefresh has passed.
	// Note that clients can refresh their token in the last moment of MaxRefresh.
	// This means that the maximum validity timespan for a token is TokenTime + MaxRefresh.
	// Optional, defaults to 0 meaning not refreshable.
	MaxRefresh time.Duration

	// Callback function that should perform the authentication of the user based on login info.
	// Must return user data as user identifier, it will be stored in Claim Array. Required.
	// Check error (e) to determine the appropriate error message.
	Authenticator func(c *gin.Context) (interface{}, error)

	// Callback function that should perform the authorization of the authenticated user. Called
	// only after an authentication success. Must return true on success, false on failure.
	// Optional, default to success.
	Authorizator func(data interface{}, c *gin.Context) bool

	// Callback function that will be called during login.
	// Using this function it is possible to add additional payload data to the webtoken.
	// The data is then made available during requests via c.GetArticle("JWT_PAYLOAD").
	// Note that the payload is not encrypted.
	// The attributes mentioned on jwt.io can't be used as keys for the map.
	// Optional, by default no additional data will be set.
	PayloadFunc func(data interface{}) MapClaims

	// User can define own Unauthorized func.
	Unauthorized func(*gin.Context, int, string)

	// User can define own LoginResponse func.
	LoginResponse func(*gin.Context, int, string, time.Time)

	// User can define own RefreshResponse func.
	RefreshResponse func(*gin.Context, int, string, time.Time)

	// Set the identity handler function
	IdentityHandler func(*gin.Context) interface{}

	// Set the identity key
	IdentityKey string

	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract token from the request.
	// Optional. Default value "header:Authorization".
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	TokenLookup string

	// TokenHeadName is a string in the header. Default value is
	TokenHeadName string

	// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
	TimeFunc func() time.Time

	// HTTP Status messages for when something in the JWT middleware fails.
	// Check error (e) to determine the appropriate error message.
	HTTPStatusMessageFunc func(e error, c *gin.Context) string

	// Private key file for asymmetric algorithms
	PrivKeyFile string

	// Public key file for asymmetric algorithms
	PubKeyFile string

	// Private key
	privKey *rsa.PrivateKey

	// Public key
	pubKey *rsa.PublicKey

	SecurityConfig SecurityConfig
}

var (
	// ErrMissingSecretKey indicates Secret key is required
	ErrMissingSecretKey = errors.New("secret key is required")

	// ErrForbidden when HTTP status 403 is given
	ErrForbidden = errors.New("you don't have permission to access this resource")

	// ErrMissingAuthenticatorFunc indicates Authenticator is required
	ErrMissingAuthenticatorFunc = errors.New("ginJWTMiddleware.Authenticator func is undefined")

	// ErrMissingLoginValues indicates a user tried to authenticate without username or password
	ErrMissingLoginValues = errors.New("missing Username or Password or Code")

	// ErrFailedAuthentication indicates authentication failed, could be faulty username or password
	ErrFailedAuthentication = errors.New("incorrect Username or Password")

	// ErrFailedTokenCreation indicates JWT Token failed to create, reason unknown
	ErrFailedTokenCreation = errors.New("failed to create JWT Token")

	// ErrExpiredToken indicates JWT token has expired. Can't refresh.
	ErrExpiredToken = errors.New("token is expired")

	// ErrEmptyAuthHeader can be thrown if authing with a HTTP header, the Auth header needs to be set
	ErrEmptyAuthHeader = errors.New("auth header is empty")

	// ErrMissingExpField missing exp field in token
	ErrMissingExpField = errors.New("missing exp field")

	// ErrWrongFormatOfExp field must be float64 format
	ErrWrongFormatOfExp = errors.New("exp must be float64 format")

	// ErrInvalidAuthHeader indicates auth header is invalid
	ErrInvalidAuthHeader = errors.New("auth header is invalid")

	// ErrEmptyQueryToken can be thrown if authing with URL Query, the query token variable is empty
	ErrEmptyQueryToken = errors.New("query token is empty")

	// ErrEmptyParamToken can be thrown if authing with parameter in path, the parameter in path is empty
	ErrEmptyParamToken = errors.New("parameter token is empty")

	// ErrInvalidSigningAlgorithm indicates signing algorithm is invalid, needs to be HS256, HS384, HS512, RS256, RS384 or RS512
	ErrInvalidSigningAlgorithm = errors.New("invalid signing algorithm")

	ErrInvalidVerificationode = errors.New("验证码错误")

	// ErrNoPrivKeyFile indicates that the given private key is unreadable
	ErrNoPrivKeyFile = errors.New("private key file unreadable")

	// ErrNoPubKeyFile indicates that the given public key is unreadable
	ErrNoPubKeyFile = errors.New("public key file unreadable")

	// ErrInvalidPrivKey indicates that the given private key is invalid
	ErrInvalidPrivKey = errors.New("private key invalid")

	// ErrInvalidPubKey indicates the the given public key is invalid
	ErrInvalidPubKey = errors.New("public key invalid")

	ErrDeviceMismatch     = errors.New("device mismatch")
	ErrTokenRevoked       = errors.New("token revoked")
	ErrTooManyDevices     = errors.New("too many devices")
	ErrSuspiciousActivity = errors.New("suspicious activity detected")
)

// IdentityKey default identity key
var IdentityKey = authdto.LoginUserId

// New for check error with GinJWTMiddleware
func New(m *GinJWTMiddleware) (*GinJWTMiddleware, error) {
	if err := m.MiddlewareInit(); err != nil {
		return nil, err
	}

	return m, nil
}

// MiddlewareInit initialize jwt configs.
func (mw *GinJWTMiddleware) MiddlewareInit() error {

	if mw.TokenLookup == "" {
		mw.TokenLookup = "header:Authorization"
	}

	if mw.SigningAlgorithm == "" {
		mw.SigningAlgorithm = "HS256"
	}

	if mw.TimeFunc == nil {
		mw.TimeFunc = time.Now
	}

	mw.TokenHeadName = strings.TrimSpace(mw.TokenHeadName)
	if len(mw.TokenHeadName) == 0 {
		mw.TokenHeadName = authdto.HeaderTokenName
	}

	if mw.Authorizator == nil {
		mw.Authorizator = func(data interface{}, c *gin.Context) bool {
			return true
		}
	}

	if mw.Unauthorized == nil {
		mw.Unauthorized = func(c *gin.Context, code int, message string) {
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": message,
			})
		}
	}

	if mw.LoginResponse == nil {
		mw.LoginResponse = func(c *gin.Context, code int, token string, expire time.Time) {
			userName, _ := c.Get(authdto.UserName)
			c.JSON(http.StatusOK, gin.H{
				"requestId": strutils.GenerateMsgIDFromContext(c),
				"msg":       "",
				"code":      http.StatusOK,
				"data": gin.H{
					"token":    token,
					"username": userName.(string),
					//"expire":   expire.Format(time.RFC3339),
					//"userInfo": userInfo,
				},
			})
		}
	}

	if mw.RefreshResponse == nil {
		mw.RefreshResponse = func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, gin.H{
				"requestId": strutils.GenerateMsgIDFromContext(c),
				"msg":       "",
				"code":      http.StatusOK,
				"data": gin.H{
					"token": token,
					//"expire": expire.Format(time.RFC3339),
				},
			})
		}
	}

	if mw.IdentityKey == "" {
		mw.IdentityKey = IdentityKey
	}

	if mw.IdentityHandler == nil {
		mw.IdentityHandler = func(c *gin.Context) interface{} {
			claims := ExtractClaims(c)
			return claims
		}
	}

	if mw.HTTPStatusMessageFunc == nil {
		mw.HTTPStatusMessageFunc = func(e error, c *gin.Context) string {
			return e.Error()
		}
	}

	// 初始化安全配置，一个用户最多可登录设备数
	if mw.SecurityConfig.MaxDevicesPerUser <= 0 {
		mw.SecurityConfig.MaxDevicesPerUser = 5
	}

	if mw.usingPublicKeyAlgo() {
		return mw.readKeys()
	}

	if mw.Key == nil {
		return ErrMissingSecretKey
	}

	// 验证Timeout设置
	if mw.Timeout <= 0 {
		mw.Timeout = time.Hour
	}

	// 验证MaxRefresh，确保不大于Timeout
	if mw.MaxRefresh > mw.Timeout {
		mw.MaxRefresh = mw.Timeout / 2
	}
	return nil
}

// MiddlewareFunc makes GinJWTMiddleware implement the Middleware interface.
func (mw *GinJWTMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		mw.middlewareImpl(c)
	}
}

// GetClaimsFromJWT get claims from JWT token
func (mw *GinJWTMiddleware) GetClaimsFromJWT(c *gin.Context) (MapClaims, error) {
	token, tokenStr, err := mw.parseToken(c)
	if err != nil {
		return nil, err
	}

	claims := MapClaims{}
	for key, value := range token.Claims.(jwt.MapClaims) {
		claims[key] = value
	}

	// get user id
	userID, ok := claims[authdto.LoginUserId].(float64)
	if !ok {
		return nil, lang.MsgErr(baseLang.AuthErr, mw.getAcceptLanguage(c))
	}
	userIDStr := strconv.FormatInt(int64(userID), 10)

	if config.ApplicationConfig.IsSingleLogin {
		saveTokenStr := mw.getCacheString(JWTLoginPrefix, userIDStr)
		if saveTokenStr != tokenStr {
			mw.logSecurityEvent(c, "single_login_violation", userID)
			return nil, lang.MsgErr(baseLang.AuthErr, mw.getAcceptLanguage(c))
		}
	}

	// 3. check blocklist
	if mw.SecurityConfig.TokenBlacklist {
		tokenID, ok := claims[TokenID].(string)
		if ok && tokenID != "" {
			isBlacklisted := mw.getCacheString(JWTBlacklistPrefix, tokenID)
			if isBlacklisted != "" {
				mw.logSecurityEvent(c, "blacklisted_token_used", userID)
				return nil, lang.MsgErr(baseLang.AuthErr, mw.getAcceptLanguage(c))
			}
		}
	}

	// 4. device check
	if mw.SecurityConfig.DeviceCheckEnabled {
		currentDeviceFP := mw.extractDeviceFingerprint(c)

		// 必须要有设备指纹
		if currentDeviceFP == "" {
			mw.logSecurityEvent(c, "missing_current_device_fingerprint", userID)
			return nil, lang.MsgErr(baseLang.AuthErr, mw.getAcceptLanguage(c))
		}

		// 获取设备列表
		devices := mw.getCacheString(JWTDevicesPrefix, userIDStr)
		if devices == "" {
			// 没有设备记录，拒绝访问
			mw.logSecurityEvent(c, "no_device_record", userID)
			return nil, lang.MsgErr(baseLang.AuthErr, mw.getAcceptLanguage(c))
		}

		// 1. 检查当前设备是否在允许的设备列表中
		deviceExists := false
		deviceList := strings.Split(devices, ",")
		for _, d := range deviceList {
			if d == currentDeviceFP {
				deviceExists = true
				break
			}
		}

		// 2. 如果不允许多设备，当前设备必须在列表中
		if !mw.SecurityConfig.AllowMultiDevices {
			if !deviceExists {
				// 设备不在列表中，拒绝
				mw.logSecurityEvent(c, "device_not_in_list_single_device", userID)
				return nil, lang.MsgErr(baseLang.AuthErr, mw.getAcceptLanguage(c))
			}
			// 单设备模式，设备在列表中，允许访问
		} else {
			// 3. 允许多设备，检查设备数限制
			if !deviceExists {
				// 新设备
				if devices == "" {
					// 第一个设备，直接添加
					deviceList = []string{currentDeviceFP}
					runtime.RuntimeConfig.GetCacheAdapter().Set(
						JWTDevicesPrefix,
						userIDStr,
						currentDeviceFP,
						config.AuthConfig.Timeout,
					)
				} else {
					// 已有设备，检查设备数
					if len(deviceList) >= mw.SecurityConfig.MaxDevicesPerUser {
						// 设备数已达上限
						mw.logSecurityEvent(c, "too_many_devices", userID)
						return nil, lang.MsgErr(baseLang.AuthErr, mw.getAcceptLanguage(c))
					}

					// 添加新设备到列表
					deviceList = append(deviceList, currentDeviceFP)
					runtime.RuntimeConfig.GetCacheAdapter().Set(
						JWTDevicesPrefix,
						userIDStr,
						strings.Join(deviceList, ","),
						config.AuthConfig.Timeout,
					)
				}
			}
			// 设备已在列表中，允许访问
		}

		// 4. 可选：记录token中的设备指纹（用于审计）
		savedDeviceFP, ok := claims[DeviceFingerprint].(string)
		if ok && savedDeviceFP != "" && currentDeviceFP != savedDeviceFP {
			// 设备指纹不匹配，记录安全事件（但不拒绝）
			mw.logSecurityEvent(c, "token_device_mismatch_audit", userID)
		}
	}

	// 5. 更新最后活动时间
	mw.updateLastActivity(userIDStr)

	return claims, nil
}

// LoginHandler can be used by clients to get a jwt token.
// Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// Reply will be of the form {"token": "TOKEN"}.
func (mw *GinJWTMiddleware) LoginHandler(c *gin.Context) {
	if mw.Authenticator == nil {
		mw.unauthorized(c, http.StatusInternalServerError, mw.HTTPStatusMessageFunc(ErrMissingAuthenticatorFunc, c))
		return
	}

	data, err := mw.Authenticator(c)

	if err != nil {
		mw.unauthorized(c, http.StatusBadRequest, mw.HTTPStatusMessageFunc(err, c))
		return
	}

	// Create the token
	token := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	if mw.PayloadFunc != nil {
		for key, value := range mw.PayloadFunc(data) {
			claims[key] = value
		}
	}

	expire := mw.TimeFunc().Add(mw.Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = mw.TimeFunc().Unix()
	// 生成Token唯一ID
	claims[TokenID] = idgen.UUID()

	// 添加设备信息
	if mw.SecurityConfig.DeviceCheckEnabled {
		deviceFP := mw.extractDeviceFingerprint(c)
		claims[DeviceFingerprint] = deviceFP
		claims[LoginIP] = c.ClientIP()
		claims[ClientInfo] = c.Request.UserAgent()

		// 记录设备
		userID, ok := data.(map[string]interface{})[authdto.LoginUserId].(int64)
		if ok {
			mw.recordDevice(userID, deviceFP)
		}
	}

	tokenString, err := mw.signedString(token)

	if err != nil {
		mw.unauthorized(c, http.StatusInternalServerError, mw.HTTPStatusMessageFunc(ErrFailedTokenCreation, c))
		return
	}

	// set
	userID, ok := data.(map[string]interface{})[authdto.LoginUserId].(int64)
	if ok {
		err = runtime.RuntimeConfig.GetCacheAdapter().Set(
			JWTLoginPrefix,
			strconv.FormatInt(userID, 10),
			tokenString,
			config.AuthConfig.Timeout,
		)
		if err != nil {
			mw.unauthorized(c, http.StatusInternalServerError, mw.HTTPStatusMessageFunc(err, c))
			return
		}
	}
	mw.LoginResponse(c, http.StatusOK, tokenString, expire)
}

// RefreshHandler can be used to refresh a token. The token still needs to be valid on refresh.
// Shall be put under an endpoint that is using the GinJWTMiddleware.
// Reply will be of the form {"token": "TOKEN"}.
func (mw *GinJWTMiddleware) RefreshHandler(c *gin.Context) {
	tokenString, expire, err := mw.RefreshToken(c)
	if err != nil {
		mw.unauthorized(c, http.StatusUnauthorized, mw.HTTPStatusMessageFunc(err, c))
		return
	}

	mw.RefreshResponse(c, http.StatusOK, tokenString, expire)
}

// RefreshToken refresh token and check if token is expired
func (mw *GinJWTMiddleware) RefreshToken(c *gin.Context) (string, time.Time, error) {
	claims, err := mw.CheckIfTokenExpire(c)
	if err != nil {
		return "", time.Now(), err
	}

	// 刷新时检查设备
	if mw.SecurityConfig.DeviceCheckEnabled {
		currentDeviceFP := mw.extractDeviceFingerprint(c)
		savedDeviceFP, ok := claims[DeviceFingerprint].(string)

		if ok && currentDeviceFP != "" && currentDeviceFP != savedDeviceFP {
			if savedDeviceFP == "" {
				// 老token可能没有设备信息，更新claims
				claims[DeviceFingerprint] = currentDeviceFP
			} else if currentDeviceFP != "" && currentDeviceFP != savedDeviceFP {
				// 设备不匹配
				mw.logSecurityEvent(c, "refresh_device_mismatch", claims[authdto.LoginUserId])
				return "", time.Now(), lang.MsgErr(baseLang.AuthErr, mw.getAcceptLanguage(c))
			}
		}
	}

	// Create the token
	newToken := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	newClaims := newToken.Claims.(jwt.MapClaims)

	// 复制原有claims但排除exp和orig_iat
	for key, value := range claims {
		if key != "exp" && key != "orig_iat" {
			newClaims[key] = value
		}
	}

	expire := mw.TimeFunc().Add(mw.Timeout)
	newClaims["exp"] = expire.Unix()
	newClaims["orig_iat"] = mw.TimeFunc().Unix()

	tokenString, err := mw.signedString(newToken)
	if err != nil {
		return "", time.Now(), err
	}

	// 更新缓存中的token
	userID, ok := claims[authdto.LoginUserId].(float64)
	if ok {
		err = runtime.RuntimeConfig.GetCacheAdapter().Set(
			JWTLoginPrefix,
			strconv.FormatInt(int64(userID), 10),
			tokenString,
			config.AuthConfig.Timeout,
		)
		if err != nil {
			return "", time.Now(), err
		}
	}

	return tokenString, expire, nil
}

// CheckIfTokenExpire check if token expire
func (mw *GinJWTMiddleware) CheckIfTokenExpire(c *gin.Context) (jwt.MapClaims, error) {
	token, _, err := mw.parseToken(c)

	if err != nil {
		// If we receive an error, and the error is anything other than a single
		// ValidationErrorExpired, we want to return the error.
		// If the error is just ValidationErrorExpired, we want to continue, as we can still
		// refresh the token if it's within the MaxRefresh time.
		// (see https://github.com/appleboy/gin-jwt/issues/176)
		var validationErr *jwt.ValidationError
		ok := errors.As(err, &validationErr)
		if !ok || validationErr.Errors != jwt.ValidationErrorExpired {
			return nil, err
		}
	}

	claims := token.Claims.(jwt.MapClaims)

	origIat := int64(claims["orig_iat"].(float64))

	if origIat < mw.TimeFunc().Add(-mw.MaxRefresh).Unix() {
		return nil, ErrExpiredToken
	}

	return claims, nil
}

// TokenGenerator method that clients can use to get a jwt token.
func (mw *GinJWTMiddleware) TokenGenerator(data interface{}) (string, time.Time, error) {
	token := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	if mw.PayloadFunc != nil {
		for key, value := range mw.PayloadFunc(data) {
			claims[key] = value
		}
	}

	expire := mw.TimeFunc().UTC().Add(mw.Timeout)
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = mw.TimeFunc().Unix()
	claims[TokenID] = idgen.UUID()
	tokenString, err := mw.signedString(token)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expire, nil
}

func (mw *GinJWTMiddleware) readKeys() error {
	err := mw.privateKey()
	if err != nil {
		return err
	}
	err = mw.publicKey()
	if err != nil {
		return err
	}
	return nil
}

func (mw *GinJWTMiddleware) privateKey() error {
	keyData, err := os.ReadFile(mw.PrivKeyFile)
	if err != nil {
		return ErrNoPrivKeyFile
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return ErrInvalidPrivKey
	}
	mw.privKey = key
	return nil
}

func (mw *GinJWTMiddleware) publicKey() error {
	keyData, err := os.ReadFile(mw.PubKeyFile)
	if err != nil {
		return ErrNoPubKeyFile
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return ErrInvalidPubKey
	}
	mw.pubKey = key
	return nil
}

func (mw *GinJWTMiddleware) usingPublicKeyAlgo() bool {
	switch mw.SigningAlgorithm {
	case "RS256", "RS512", "RS384":
		return true
	}
	return false
}

func (mw *GinJWTMiddleware) middlewareImpl(c *gin.Context) {
	claims, err := mw.GetClaimsFromJWT(c)
	if err != nil {
		mw.unauthorized(c, http.StatusUnauthorized, mw.HTTPStatusMessageFunc(err, c))
		return
	}

	if claims["exp"] == nil {
		mw.unauthorized(c, http.StatusBadRequest, mw.HTTPStatusMessageFunc(ErrMissingExpField, c))
		return
	}

	if _, ok := claims["exp"].(float64); !ok {
		mw.unauthorized(c, http.StatusBadRequest, mw.HTTPStatusMessageFunc(ErrWrongFormatOfExp, c))
		return
	}
	if int64(claims["exp"].(float64)) < mw.TimeFunc().Unix() {
		mw.unauthorized(c, http.StatusUnauthorized, mw.HTTPStatusMessageFunc(ErrExpiredToken, c))
		return
	}

	c.Set(JwtPayloadKey, claims)
	identity := mw.IdentityHandler(c)

	if identity != nil {
		c.Set(mw.IdentityKey, identity)
	}

	if !mw.Authorizator(identity, c) {
		mw.unauthorized(c, http.StatusForbidden, mw.HTTPStatusMessageFunc(ErrForbidden, c))
		return
	}

	c.Next()
}

func (mw *GinJWTMiddleware) signedString(token *jwt.Token) (string, error) {
	if mw.usingPublicKeyAlgo() {
		return token.SignedString(mw.privKey)
	}
	return token.SignedString(mw.Key)
}

// extractDeviceFingerprint 提取设备指纹
func (mw *GinJWTMiddleware) extractDeviceFingerprint(c *gin.Context) string {
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

// logSecurityEvent 记录安全事件
func (mw *GinJWTMiddleware) logSecurityEvent(c *gin.Context, eventType string, userID interface{}) {
	rLog := log.GetRequestLogger(c)
	rLog.Warnf("Security event: %s, UserID: %v, IP: %s, UA: %s",
		eventType, userID, c.ClientIP(), c.Request.UserAgent())

	// 可以在这里添加安全事件到数据库
}

// recordDevice 记录设备
func (mw *GinJWTMiddleware) recordDevice(userID int64, deviceFP string) {
	if !mw.SecurityConfig.DeviceCheckEnabled {
		return
	}

	userIDStr := strconv.FormatInt(userID, 10)

	// 获取现有设备列表
	devices := mw.getCacheString(JWTDevicesPrefix, userIDStr)
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
		if len(deviceList) > mw.SecurityConfig.MaxDevicesPerUser {
			deviceList = deviceList[len(deviceList)-mw.SecurityConfig.MaxDevicesPerUser:]
		}

		runtime.RuntimeConfig.GetCacheAdapter().Set(
			JWTDevicesPrefix,
			userIDStr,
			strings.Join(deviceList, ","),
			config.AuthConfig.Timeout,
		)
	}
}

// RevokeToken 撤销Token
func (mw *GinJWTMiddleware) RevokeToken(c *gin.Context) error {
	claims, err := mw.GetClaimsFromJWT(c)
	if err != nil {
		return err
	}

	// 获取用户ID
	userID, ok := claims[authdto.LoginUserId].(float64)
	if !ok {
		return errors.New("invalid user id in token")
	}
	userIDStr := strconv.FormatInt(int64(userID), 10)

	// 1. 清除单点登录缓存
	runtime.RuntimeConfig.GetCacheAdapter().Del(JWTLoginPrefix, userIDStr)

	// 2. 将token加入黑名单
	if mw.SecurityConfig.TokenBlacklist {
		tokenID, ok := claims[TokenID].(string)
		if ok && tokenID != "" {
			// 黑名单有效期比token短
			blacklistTTL := mw.Timeout / 2
			exp, ok := claims["exp"].(float64)
			if ok {
				expireTime := time.Unix(int64(exp), 0)
				remaining := time.Until(expireTime)
				if remaining > 0 {
					blacklistTTL = remaining / 2
				}
			}
			if blacklistTTL < time.Minute*5 {
				blacklistTTL = time.Minute * 5
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
	if mw.SecurityConfig.DeviceCheckEnabled {
		currentDeviceFP := mw.extractDeviceFingerprint(c)
		savedDeviceFP, ok := claims[DeviceFingerprint].(string)

		if ok && currentDeviceFP != "" && savedDeviceFP != "" && currentDeviceFP == savedDeviceFP {
			mw.removeDevice(userIDStr, currentDeviceFP)
		}
	}

	return nil
}

// ExtractClaims help to extract the JWT claims
func ExtractClaims(c *gin.Context) MapClaims {
	claims, exists := c.Get(JwtPayloadKey)
	if !exists {
		return make(MapClaims)
	}

	return claims.(MapClaims)
}

// ExtractClaimsFromToken help to extract the JWT claims from token
func ExtractClaimsFromToken(token *jwt.Token) MapClaims {
	if token == nil {
		return make(MapClaims)
	}

	claims := MapClaims{}
	for key, value := range token.Claims.(jwt.MapClaims) {
		claims[key] = value
	}

	return claims
}

// GetToken help to get the JWT token string
func GetToken(c *gin.Context) string {
	token, exists := c.Get("JWT_TOKEN")
	if !exists {
		return ""
	}

	return token.(string)
}

// updateLastActivity 更新最后活动时间
func (mw *GinJWTMiddleware) updateLastActivity(userID string) {
	// 可以记录用户最后活动时间，用于会话管理
	runtime.RuntimeConfig.GetCacheAdapter().Set(
		"admin:jwt:activity",
		userID,
		strconv.FormatInt(time.Now().Unix(), 10),
		config.AuthConfig.Timeout,
	)
}

// removeDevice 移除设备
func (mw *GinJWTMiddleware) removeDevice(userID string, deviceFP string) {
	devices := mw.getCacheString(JWTDevicesPrefix, userID)
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

// GetUserDevices 获取用户的所有设备
func (mw *GinJWTMiddleware) GetUserDevices(userID string) []string {
	devices := mw.getCacheString(JWTDevicesPrefix, userID)
	if devices == "" {
		return []string{}
	}

	return strings.Split(devices, ",")
}

// RevokeUserAllTokens 撤销用户的所有Token
func (mw *GinJWTMiddleware) RevokeUserAllTokens(userID string) error {
	// 1. 清除单点登录缓存
	runtime.RuntimeConfig.GetCacheAdapter().Del(JWTLoginPrefix, userID)

	// 2. 清除设备记录
	runtime.RuntimeConfig.GetCacheAdapter().Del(JWTDevicesPrefix, userID)

	// 3. 清除活动记录
	runtime.RuntimeConfig.GetCacheAdapter().Del("admin:jwt:activity", userID)

	return nil
}

func (mw *GinJWTMiddleware) jwtFromHeader(c *gin.Context, key string) (string, error) {
	authHeader := c.Request.Header.Get(key)

	if authHeader == "" {
		return "", ErrEmptyAuthHeader
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == mw.TokenHeadName) {
		return "", ErrInvalidAuthHeader
	}

	return parts[1], nil
}

func (mw *GinJWTMiddleware) jwtFromQuery(c *gin.Context, key string) (string, error) {
	token := c.Query(key)

	if token == "" {
		return "", ErrEmptyQueryToken
	}

	return token, nil
}

func (mw *GinJWTMiddleware) jwtFromParam(c *gin.Context, key string) (string, error) {
	token := c.Param(key)

	if token == "" {
		return "", ErrEmptyParamToken
	}

	return token, nil
}

// parseToken parse jwt token from gin context
func (mw *GinJWTMiddleware) parseToken(c *gin.Context) (*jwt.Token, string, error) {
	var token string
	var err error

	methods := strings.Split(mw.TokenLookup, ",")
	for _, method := range methods {
		if len(token) > 0 {
			break
		}
		parts := strings.Split(strings.TrimSpace(method), ":")
		if len(parts) != 2 {
			continue // 或者返回错误
		}
		k := strings.TrimSpace(parts[0])
		v := strings.TrimSpace(parts[1])
		switch k {
		case "header":
			token, err = mw.jwtFromHeader(c, v)
		case "query":
			token, err = mw.jwtFromQuery(c, v)
		case "param":
			token, err = mw.jwtFromParam(c, v)
		}
	}

	if err != nil {
		return nil, "", err
	}

	tk, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(mw.SigningAlgorithm) != t.Method {
			return nil, ErrInvalidSigningAlgorithm
		}
		if mw.usingPublicKeyAlgo() {
			return mw.pubKey, nil
		}
		c.Set("JWT_TOKEN", token)

		return mw.Key, nil
	})
	if err != nil {
		return nil, "", err
	}
	return tk, token, nil
}

// parseTokenString parse jwt token string
func (mw *GinJWTMiddleware) parseTokenString(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(mw.SigningAlgorithm) != t.Method {
			return nil, ErrInvalidSigningAlgorithm
		}
		if mw.usingPublicKeyAlgo() {
			return mw.pubKey, nil
		}

		return mw.Key, nil
	})
}

func (mw *GinJWTMiddleware) unauthorized(c *gin.Context, code int, message string) {
	c.Abort()
	msg := lang.MsgByCode(code, mw.getAcceptLanguage(c))
	if msg != "" {
		message = msg
	}

	mw.Unauthorized(c, code, message)
}

// getAcceptLanguage 获取当前语言
func (mw *GinJWTMiddleware) getAcceptLanguage(c *gin.Context) string {
	languages := lang.ParseAcceptLanguage(c.GetHeader("Accept-Language"), nil)
	if len(languages) == 0 {
		return "zh-CN"
	}
	return languages[0]
}

func (mw *GinJWTMiddleware) getCacheString(prefix, key string) string {
	val, err := runtime.RuntimeConfig.GetCacheAdapter().Get(prefix, key)
	if err != nil || val == "" {
		return ""
	}
	return val
}
