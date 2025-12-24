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
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const HeaderTokenName = "Bearer"
const HeaderAuthorization = "Authorization"

const JwtPayloadKey = "JWT_PAYLOAD"
const JWTLoginPrefix = "admin:jwt"
const JWTBlacklistPrefix = "admin:jwt:blacklist"
const JWTDevicesPrefix = "admin:jwt:devices"
const DeviceFingerprint = "dev_fp"
const LoginIP = "login_ip"
const ClientInfo = "client_info"
const TokenID = "jti"

type SecurityConfig struct {
	DeviceCheckEnabled bool // 设备检查
	TokenBlacklist     bool // Token黑名单
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
	Payload func(data interface{}) jwt.MapClaims

	// User can define own Unauthorized func.
	Unauthorized func(*gin.Context, int, int, string)

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

	// ErrMissingAuthorizatorFunc indicates Authenticator is required
	ErrMissingAuthorizatorFunc = errors.New("ginJWTMiddleware.Authorizator func is undefined")

	// ErrMissingAuthenticatorFunc indicates Authenticator is required
	ErrMissingAuthenticatorFunc = errors.New("ginJWTMiddleware.Authenticator func is undefined")

	// ErrMissingUnauthorizedFunc indicates Unauthorized is required
	ErrMissingUnauthorizedFunc = errors.New("ginJWTMiddleware.Unauthorized func is undefined")

	// ErrMissingLoginResponseFunc indicates LoginResponse is required
	ErrMissingLoginResponseFunc = errors.New("ginJWTMiddleware.LoginResponse func is undefined")

	// ErrMissingPayloadFunc indicates Payload is required
	ErrMissingPayloadFunc = errors.New("ginJWTMiddleware.Payload func is undefined")

	// ErrMissingRefreshResponseFunc indicates RefreshResponse is required
	ErrMissingRefreshResponseFunc = errors.New("ginJWTMiddleware.RefreshResponse func is undefined")

	// ErrMissingIdentityHandlerFunc indicates IdentityHandler is required
	ErrMissingIdentityHandlerFunc = errors.New("ginJWTMiddleware.IdentityHandler func is undefined")

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
		mw.TokenLookup = "header:" + HeaderAuthorization
	}

	if mw.SigningAlgorithm == "" {
		mw.SigningAlgorithm = "HS256"
	}

	if mw.TimeFunc == nil {
		mw.TimeFunc = time.Now
	}

	mw.TokenHeadName = strings.TrimSpace(mw.TokenHeadName)
	if len(mw.TokenHeadName) == 0 {
		mw.TokenHeadName = HeaderTokenName
	}

	if mw.Authorizator == nil {
		return ErrMissingAuthorizatorFunc
	}

	if mw.Authenticator == nil {
		return ErrMissingAuthenticatorFunc
	}

	if mw.Unauthorized == nil {
		return ErrMissingUnauthorizedFunc
	}

	if mw.LoginResponse == nil {
		return ErrMissingLoginResponseFunc
	}

	if mw.RefreshResponse == nil {
		return ErrMissingRefreshResponseFunc
	}

	if mw.IdentityHandler == nil {
		return ErrMissingIdentityHandlerFunc
	}

	if mw.Payload == nil {
		return ErrMissingPayloadFunc
	}

	if mw.Key == nil {
		return ErrMissingSecretKey
	}

	if mw.IdentityKey == "" {
		mw.IdentityKey = IdentityKey
	}

	// 初始化安全配置，一个用户最多可登录设备数
	if mw.SecurityConfig.MaxDevicesPerUser <= 0 {
		mw.SecurityConfig.MaxDevicesPerUser = 5
	}

	if mw.usingPublicKeyAlgo() {
		return mw.readKeys()
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
func (mw *GinJWTMiddleware) GetClaimsFromJWT(c *gin.Context) (jwt.MapClaims, error) {
	token, tokenStr, err := mw.parseToken(c)
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(jwt.MapClaims)

	// get user id
	userID, ok := claims[authdto.LoginUserId].(float64)
	if !ok {
		return nil, lang.MsgErr(baseLang.AuthErr, mw.getAcceptLanguage(c))
	}
	userIDStr := strconv.FormatInt(int64(userID), 10)

	if config.ApplicationConfig.IsSingleLogin {
		// 从缓存获取该用户最新的token
		savedToken := mw.getCacheString(JWTLoginPrefix, userIDStr)
		if savedToken != tokenStr {
			// 当前token不是最新的，说明用户在其他地方登录了
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
		savedDeviceFP, ok := claims[DeviceFingerprint].(string)
		if currentDeviceFP == "" || savedDeviceFP == "" || !ok || currentDeviceFP != savedDeviceFP {
			mw.logSecurityEvent(c, "device_fingerprint_err", userID)
			return nil, lang.MsgErr(baseLang.AuthErr, mw.getAcceptLanguage(c))
		}

		// 获取设备列表
		devices := mw.getCacheString(JWTDevicesPrefix, userIDStr)

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
				if len(deviceList) >= mw.SecurityConfig.MaxDevicesPerUser {
					// 设备数已达上限
					mw.logSecurityEvent(c, "too_many_devices", userID)
					return nil, lang.MsgErr(baseLang.AuthErr, mw.getAcceptLanguage(c))
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
	mw.updateLastActivity(userIDStr)

	return claims, nil
}

// LoginHandler can be used by clients to get a jwt token.
// Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}.
// Reply will be of the form {"token": "TOKEN"}.
func (mw *GinJWTMiddleware) LoginHandler(c *gin.Context) {
	lg := mw.getAcceptLanguage(c)
	data, err := mw.Authenticator(c)
	if err != nil {
		mw.unauthorized(c, http.StatusUnauthorized, baseLang.AuthErrLogCode, lang.MsgErrf(baseLang.AuthErrLogCode, lg, err).Error())
		return
	}

	userID, ok := data.(map[string]interface{})[authdto.LoginUserId].(int64)
	if !ok {
		mw.unauthorized(c, http.StatusUnauthorized, baseLang.AuthErrLogCode, lang.MsgErrf(baseLang.AuthErrLogCode, lg, err).Error())
		return
	}
	userIDStr := strconv.FormatInt(userID, 10)

	if config.ApplicationConfig.IsSingleLogin {
		runtime.RuntimeConfig.GetCacheAdapter().Del(JWTLoginPrefix, userIDStr)
	}

	// Create the token
	token := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	claims := token.Claims.(jwt.MapClaims)

	if mw.Payload != nil {
		for key, value := range mw.Payload(data) {
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
		mw.recordDevice(userID, deviceFP)
	}

	tokenString, err := mw.signedString(token)
	if err != nil {
		mw.unauthorized(c, http.StatusUnauthorized, baseLang.AuthErrLogCode, lang.MsgErrf(baseLang.AuthErrLogCode, lg, err).Error())
		return
	}

	// set
	err = runtime.RuntimeConfig.GetCacheAdapter().Set(
		JWTLoginPrefix,
		userIDStr,
		tokenString,
		config.AuthConfig.Timeout,
	)
	if err != nil {
		mw.unauthorized(c, http.StatusUnauthorized, baseLang.AuthErrLogCode, lang.MsgErrf(baseLang.AuthErrLogCode, lg, err).Error())
		return
	}

	mw.LoginResponse(c, http.StatusOK, tokenString, expire)
}

func (mw *GinJWTMiddleware) LogoutHandler(c *gin.Context, httpCode, code int, message string) {
	mw.Unauthorized(c, httpCode, code, message)
}

// RefreshHandler can be used to refresh a token. The token still needs to be valid on refresh.
// Shall be put under an endpoint that is using the GinJWTMiddleware.
// Reply will be of the form {"token": "TOKEN"}.
func (mw *GinJWTMiddleware) RefreshHandler(c *gin.Context) {
	tokenString, expire, err := mw.RefreshToken(c)
	if err != nil {
		lg := mw.getAcceptLanguage(c)
		mw.unauthorized(c, http.StatusUnauthorized, baseLang.AuthErrLogCode, lang.MsgErrf(baseLang.AuthErrLogCode, lg, err).Error())
		return
	}

	mw.RefreshResponse(c, http.StatusOK, tokenString, expire)
}

// RefreshToken refresh token and check if token is expired
func (mw *GinJWTMiddleware) RefreshToken(c *gin.Context) (string, time.Time, error) {
	// 获取当前请求的token字符串
	_, tokenStr, parseErr := mw.parseToken(c)
	if parseErr != nil {
		return "", time.Now(), parseErr
	}
	claims, err := mw.CheckIfTokenExpire(c)
	if err != nil {
		return "", time.Now(), err
	}

	// 在刷新token时也需要检查单点登录
	if config.ApplicationConfig.IsSingleLogin {
		userID, ok := claims[authdto.LoginUserId].(float64)
		if ok {
			userIDStr := strconv.FormatInt(int64(userID), 10)
			savedToken := mw.getCacheString(JWTLoginPrefix, userIDStr)
			if savedToken != tokenStr {
				// token已失效（被新登录踢掉）
				return "", time.Now(), lang.MsgErr(baseLang.AuthErr, mw.getAcceptLanguage(c))
			}
		}
	}

	// 刷新时检查设备
	if mw.SecurityConfig.DeviceCheckEnabled {
		currentDeviceFP := mw.extractDeviceFingerprint(c)
		savedDeviceFP, ok := claims[DeviceFingerprint].(string)
		if currentDeviceFP == "" || savedDeviceFP == "" || !ok || currentDeviceFP != savedDeviceFP {
			mw.logSecurityEvent(c, "device_fingerprint_err", "")
			return "", time.Now(), lang.MsgErr(baseLang.AuthErr, mw.getAcceptLanguage(c))
		}
	}

	// Create the token
	newToken := jwt.New(jwt.GetSigningMethod(mw.SigningAlgorithm))
	newClaims := jwt.MapClaims{}

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

	if mw.Payload != nil {
		for key, value := range mw.Payload(data) {
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
	lg := mw.getAcceptLanguage(c)

	claims, err := mw.GetClaimsFromJWT(c)
	if err != nil {
		mw.unauthorized(c, http.StatusUnauthorized, baseLang.AuthErrLogCode, lang.MsgErrf(baseLang.AuthErrLogCode, lg, err).Error())
		return
	}

	if claims["exp"] == nil {
		mw.unauthorized(c, http.StatusUnauthorized, baseLang.AuthErrLogCode, lang.MsgErrf(baseLang.AuthErrLogCode, lg, err).Error())
		return
	}

	if _, ok := claims["exp"].(float64); !ok {
		mw.unauthorized(c, http.StatusUnauthorized, baseLang.AuthErrLogCode, lang.MsgErrf(baseLang.AuthErrLogCode, lg, err).Error())
		return
	}
	if int64(claims["exp"].(float64)) < mw.TimeFunc().Unix() {
		mw.unauthorized(c, http.StatusUnauthorized, baseLang.AuthErrLogCode, lang.MsgErrf(baseLang.AuthErrLogCode, lg, err).Error())
		return
	}

	c.Set(JwtPayloadKey, claims)
	identity := mw.IdentityHandler(c)

	if identity != nil {
		c.Set(mw.IdentityKey, identity)
	}

	if !mw.Authorizator(identity, c) {
		mw.unauthorized(c, http.StatusUnauthorized, baseLang.AuthErrLogCode, lang.MsgErrf(baseLang.AuthErrLogCode, lg, err).Error())
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
func ExtractClaims(c *gin.Context) jwt.MapClaims {
	claims, exists := c.Get(JwtPayloadKey)
	if !exists {
		return make(jwt.MapClaims)
	}

	return claims.(jwt.MapClaims)
}

// ExtractClaimsFromToken help to extract the JWT claims from token
func ExtractClaimsFromToken(token *jwt.Token) jwt.MapClaims {
	if token == nil {
		return make(jwt.MapClaims)
	}

	return token.Claims.(jwt.MapClaims)
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

func (mw *GinJWTMiddleware) unauthorized(c *gin.Context, httpCode, code int, message string) {
	c.Abort()
	mw.LogoutHandler(c, httpCode, code, message)
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
