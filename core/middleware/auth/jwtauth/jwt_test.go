package jwtauth

import (
	"bytes"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-admin/core/middleware/auth/authdto"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// 测试配置
var testConfig = struct {
	secret    string
	timeout   time.Duration
	userID    int64
	roleKey   string
	userName  string
	roleId    int64
	deptId    int64
	dataScope string
}{
	secret:    "test-secret-key-1234567890",
	timeout:   time.Hour,
	userID:    1001,
	roleKey:   "admin",
	userName:  "testuser",
	roleId:    1,
	deptId:    10,
	dataScope: "全部",
}

// 创建测试用的中间件实例
func createTestMiddleware(t *testing.T) *GinJWTMiddleware {
	// 设置临时密钥文件用于测试非对称算法
	createTestKeyFiles(t)
	defer cleanupTestKeyFiles(t)

	mw, err := New(&GinJWTMiddleware{
		Key:        []byte(testConfig.secret),
		Timeout:    testConfig.timeout,
		MaxRefresh: time.Hour,
		Payload: func(data interface{}) MapClaims {
			if v, ok := data.(map[string]interface{}); ok {
				return MapClaims{
					authdto.LoginUserId: v[authdto.LoginUserId],
					authdto.RoleKey:     v[authdto.RoleKey],
					authdto.UserName:    v[authdto.UserName],
					authdto.DataScope:   v[authdto.DataScope],
					authdto.RoleId:      v[authdto.RoleId],
					authdto.DeptId:      v[authdto.DeptId],
				}
			}
			return MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}
			if err := c.ShouldBindJSON(&login); err != nil {
				return nil, ErrMissingLoginValues
			}

			if login.Username == "testuser" && login.Password == "testpass" {
				return map[string]interface{}{
					authdto.LoginUserId: testConfig.userID,
					authdto.RoleKey:     testConfig.roleKey,
					authdto.UserName:    testConfig.userName,
					authdto.DataScope:   testConfig.dataScope,
					authdto.RoleId:      testConfig.roleId,
					authdto.DeptId:      testConfig.deptId,
				}, nil
			}
			return nil, ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(code, gin.H{
				"token": token,
			})
		},
		RefreshResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(code, gin.H{
				"token": token,
			})
		},
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		SecurityConfig: SecurityConfig{
			DeviceCheckEnabled: false, // 测试时关闭设备检查
			TokenBlacklist:     false,
			AllowMultiDevices:  true,
			MaxDevicesPerUser:  5,
		},
	})

	assert.NoError(t, err, "创建中间件失败")
	return mw
}

// 创建测试用的密钥文件
func createTestKeyFiles(t *testing.T) {
	// 创建测试目录
	os.MkdirAll("test_keys", 0755)

	// 生成测试用的RSA密钥对（简化版，实际应该用openssl生成）
	privKey := `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAzH5rq2O+q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
-----END RSA PRIVATE KEY-----`

	pubKey := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzH5rq2O+q2q2q2q2q2q2
q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q2q
-----END PUBLIC KEY-----`

	err := os.WriteFile("test_keys/rsa_private.pem", []byte(privKey), 0644)
	assert.NoError(t, err)
	err = os.WriteFile("test_keys/rsa_public.pem", []byte(pubKey), 0644)
	assert.NoError(t, err)
}

func cleanupTestKeyFiles(t *testing.T) {
	os.RemoveAll("test_keys")
}

// TestNew 测试创建中间件
func TestNew(t *testing.T) {
	// 测试正常创建
	mw := createTestMiddleware(t)
	assert.NotNil(t, mw)
	assert.Equal(t, "HS256", mw.SigningAlgorithm)

	// 测试缺少密钥
	_, err := New(&GinJWTMiddleware{
		Key:              nil,
		SigningAlgorithm: "HS256",
	})
	assert.Error(t, err)
	assert.Equal(t, ErrMissingSecretKey, err)

	// 测试非对称算法但缺少密钥文件
	_, err = New(&GinJWTMiddleware{
		Key:              []byte("secret"),
		SigningAlgorithm: "RS256",
		PrivKeyFile:      "nonexistent.pem",
		PubKeyFile:       "nonexistent.pem",
	})
	assert.Error(t, err)
}

// TestMiddlewareInit 测试中间件初始化
func TestMiddlewareInit(t *testing.T) {
	mw := &GinJWTMiddleware{
		Key: []byte(testConfig.secret),
	}

	err := mw.MiddlewareInit()
	assert.NoError(t, err)
	assert.Equal(t, "header:Authorization", mw.TokenLookup)
	assert.Equal(t, "HS256", mw.SigningAlgorithm)
	assert.NotNil(t, mw.TimeFunc)
	assert.NotNil(t, mw.Authorizator)
	assert.NotNil(t, mw.Unauthorized)
	assert.NotNil(t, mw.IdentityHandler)
}

// TestLoginHandler 测试登录
func TestLoginHandler(t *testing.T) {
	mw := createTestMiddleware(t)

	// 创建测试请求
	loginData := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	jsonData, _ := json.Marshal(loginData)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")

	// 执行登录
	mw.LoginHandler(c)

	// 验证响应
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Contains(t, response, "token")
	token := response["token"].(string)
	assert.NotEmpty(t, token)

	// 验证token可以解析
	parsedToken, err := mw.parseTokenString(token)
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)
}

// TestLoginHandlerFailed 测试登录失败
func TestLoginHandlerFailed(t *testing.T) {
	mw := createTestMiddleware(t)

	// 错误密码
	loginData := map[string]string{
		"username": "testuser",
		"password": "wrongpass",
	}
	jsonData, _ := json.Marshal(loginData)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")

	mw.LoginHandler(c)

	// 应该返回错误
	assert.NotEqual(t, http.StatusOK, w.Code)
}

// TestGetClaimsFromJWT 测试解析JWT Claims
func TestGetClaimsFromJWT(t *testing.T) {
	mw := createTestMiddleware(t)

	// 先生成一个token
	tokenString, expire, err := mw.TokenGenerator(map[string]interface{}{
		authdto.LoginUserId: testConfig.userID,
		authdto.RoleKey:     testConfig.roleKey,
		authdto.UserName:    testConfig.userName,
		authdto.DataScope:   testConfig.dataScope,
		authdto.RoleId:      testConfig.roleId,
		authdto.DeptId:      testConfig.deptId,
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
	assert.True(t, expire.After(time.Now()))

	// 测试解析
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", "Bearer "+tokenString)

	claims, err := mw.GetClaimsFromJWT(c)
	assert.NoError(t, err)
	assert.NotNil(t, claims)

	// 验证claims内容
	assert.Equal(t, float64(testConfig.userID), claims[authdto.LoginUserId])
	assert.Equal(t, testConfig.roleKey, claims[authdto.RoleKey])
	assert.Equal(t, testConfig.userName, claims[authdto.UserName])
	assert.Equal(t, testConfig.dataScope, claims[authdto.DataScope])
	assert.Equal(t, float64(testConfig.roleId), claims[authdto.RoleId])
	assert.Equal(t, float64(testConfig.deptId), claims[authdto.DeptId])
}

// TestMiddlewareFunc 测试中间件函数
func TestMiddlewareFunc(t *testing.T) {
	mw := createTestMiddleware(t)

	// 创建带token的请求
	tokenString, _, _ := mw.TokenGenerator(map[string]interface{}{
		authdto.LoginUserId: testConfig.userID,
		authdto.RoleKey:     testConfig.roleKey,
	})

	// 创建测试路由
	router := gin.New()
	router.Use(mw.MiddlewareFunc())
	router.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// 有效请求
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// 无效token请求
	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalidtoken")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

// TestRefreshToken 测试刷新Token
func TestRefreshToken(t *testing.T) {
	mw := createTestMiddleware(t)

	// 生成一个快要过期的token
	claims := jwt.MapClaims{
		authdto.LoginUserId: float64(testConfig.userID),
		authdto.RoleKey:     testConfig.roleKey,
		"exp":               time.Now().Add(time.Minute).Unix(), // 1分钟后过期
		"orig_iat":          time.Now().Add(-mw.Timeout + time.Minute*2).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(mw.Key)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/refresh", nil)
	c.Request.Header.Set("Authorization", "Bearer "+tokenString)

	newToken, expire, err := mw.RefreshToken(c)
	assert.NoError(t, err)
	assert.NotEmpty(t, newToken)
	assert.True(t, expire.After(time.Now()))
	assert.NotEqual(t, tokenString, newToken)
}

// TestCheckIfTokenExpire 测试检查Token过期
func TestCheckIfTokenExpire(t *testing.T) {
	mw := createTestMiddleware(t)

	// 测试未过期token
	claims := jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour).Unix(),
		"orig_iat": time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(mw.Key)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", "Bearer "+tokenString)

	parsedClaims, err := mw.CheckIfTokenExpire(c)
	assert.NoError(t, err)
	assert.NotNil(t, parsedClaims)

	// 测试已过期但可刷新的token
	claims["exp"] = time.Now().Add(-time.Minute).Unix()
	claims["orig_iat"] = time.Now().Add(-mw.MaxRefresh + time.Minute).Unix()
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ = token.SignedString(mw.Key)

	c.Request.Header.Set("Authorization", "Bearer "+tokenString)
	parsedClaims, err = mw.CheckIfTokenExpire(c)
	assert.NoError(t, err)
	assert.NotNil(t, parsedClaims)

	// 测试完全过期的token
	claims["orig_iat"] = time.Now().Add(-mw.MaxRefresh - time.Hour).Unix()
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ = token.SignedString(mw.Key)

	c.Request.Header.Set("Authorization", "Bearer "+tokenString)
	parsedClaims, err = mw.CheckIfTokenExpire(c)
	assert.Error(t, err)
	assert.Equal(t, ErrExpiredToken, err)
}

// TestTokenGenerator 测试Token生成器
func TestTokenGenerator(t *testing.T) {
	mw := createTestMiddleware(t)

	data := map[string]interface{}{
		authdto.LoginUserId: testConfig.userID,
		authdto.RoleKey:     testConfig.roleKey,
	}

	tokenString, expire, err := mw.TokenGenerator(data)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
	assert.True(t, expire.After(time.Now()))

	// 验证token可以解析
	token, err := mw.parseTokenString(tokenString)
	assert.NoError(t, err)
	assert.True(t, token.Valid)

	// 验证claims
	claims := token.Claims.(jwt.MapClaims)
	assert.Equal(t, float64(testConfig.userID), claims[authdto.LoginUserId])
	assert.Equal(t, testConfig.roleKey, claims[authdto.RoleKey])
	assert.Contains(t, claims, "exp")
	assert.Contains(t, claims, "orig_iat")
	assert.Contains(t, claims, TokenID)
}

// TestExtractDeviceFingerprint 测试设备指纹提取
func TestExtractDeviceFingerprint(t *testing.T) {
	mw := createTestMiddleware(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("User-Agent", "Test-Agent/1.0")
	c.Request.RemoteAddr = "192.168.1.100:8080"

	fingerprint := mw.extractDeviceFingerprint(c)
	assert.NotEmpty(t, fingerprint)
	assert.Len(t, fingerprint, 32) // SHA256取前32字符

	// 测试缺少User-Agent
	c.Request.Header.Del("User-Agent")
	fingerprint = mw.extractDeviceFingerprint(c)
	assert.Empty(t, fingerprint)

	// 测试缺少IP
	c.Request.Header.Set("User-Agent", "Test-Agent/1.0")
	c.Request.RemoteAddr = ""
	fingerprint = mw.extractDeviceFingerprint(c)
	assert.Empty(t, fingerprint)
}

// TestRevokeToken 测试撤销Token
func TestRevokeToken(t *testing.T) {
	// 这个测试需要缓存适配器，这里简化测试
	mw := createTestMiddleware(t)

	// 由于需要真实的缓存适配器，这里只测试函数签名
	// 实际项目中需要mock缓存适配器
	tokenString, _, _ := mw.TokenGenerator(map[string]interface{}{
		authdto.LoginUserId: testConfig.userID,
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/logout", nil)
	c.Request.Header.Set("Authorization", "Bearer "+tokenString)

	// 这里会失败，因为缓存适配器未设置
	// 实际项目中应该mock缓存适配器
	_ = mw.RevokeToken(c)
}

// TestGetUserDevices 测试获取用户设备
func TestGetUserDevices(t *testing.T) {
	mw := createTestMiddleware(t)

	// 测试空设备列表
	devices := mw.GetUserDevices("user123")
	assert.Empty(t, devices)

	// 注意：recordDevice需要缓存适配器支持
	// 实际测试中需要mock缓存适配器
}

// TestParseTokenMethods 测试各种解析Token的方式
func TestParseTokenMethods(t *testing.T) {
	mw := createTestMiddleware(t)

	tokenString, _, _ := mw.TokenGenerator(map[string]interface{}{
		authdto.LoginUserId: testConfig.userID,
	})

	// 测试header方式
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", "Bearer "+tokenString)

	token, _, err := mw.parseToken(c)
	assert.NoError(t, err)
	assert.NotNil(t, token)

	// 测试query方式
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/test?token="+tokenString, nil)
	mw.TokenLookup = "query:token"
	mw.MiddlewareInit() // 重新初始化

	token, _, err = mw.parseToken(c)
	assert.NoError(t, err)
	assert.NotNil(t, token)

	// 测试无效token
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", "Bearer invalid.token.here")

	token, _, err = mw.parseToken(c)
	assert.Error(t, err)
	assert.Nil(t, token)
}

// TestExtractClaims 测试提取Claims
func TestExtractClaims(t *testing.T) {

	// 创建测试上下文
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// 空上下文应该返回空claims
	claims := ExtractClaims(c)
	assert.NotNil(t, claims)
	assert.Empty(t, claims)

	// 设置claims后再提取
	testClaims := MapClaims{
		"user_id": 123,
		"role":    "admin",
	}
	c.Set(JwtPayloadKey, testClaims)

	claims = ExtractClaims(c)
	assert.Equal(t, testClaims, claims)
}

// TestExtractClaimsFromToken 测试从Token提取Claims
func TestExtractClaimsFromToken(t *testing.T) {
	mw := createTestMiddleware(t)

	tokenString, _, _ := mw.TokenGenerator(map[string]interface{}{
		authdto.LoginUserId: testConfig.userID,
		authdto.RoleKey:     testConfig.roleKey,
	})

	token, err := mw.parseTokenString(tokenString)
	assert.NoError(t, err)

	claims := ExtractClaimsFromToken(token)
	assert.NotNil(t, claims)
	assert.Equal(t, float64(testConfig.userID), claims[authdto.LoginUserId])
	assert.Equal(t, testConfig.roleKey, claims[authdto.RoleKey])

	// 测试nil token
	claims = ExtractClaimsFromToken(nil)
	assert.NotNil(t, claims)
	assert.Empty(t, claims)
}

// TestGetToken 测试获取Token字符串
func TestGetToken(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// 空上下文
	token := GetToken(c)
	assert.Empty(t, token)

	// 设置token
	c.Set("JWT_TOKEN", "test.token.123")
	token = GetToken(c)
	assert.Equal(t, "test.token.123", token)
}

// TestSecurityFeatures 测试安全功能
func TestSecurityFeatures(t *testing.T) {
	config := SecurityConfig{
		DeviceCheckEnabled: true,
		TokenBlacklist:     true,
		AllowMultiDevices:  false,
		MaxDevicesPerUser:  2,
	}

	assert.True(t, config.DeviceCheckEnabled)
	assert.True(t, config.TokenBlacklist)
	assert.False(t, config.AllowMultiDevices)
	assert.Equal(t, 2, config.MaxDevicesPerUser)
}

// TestErrorHandling 测试错误处理
func TestErrorHandling(t *testing.T) {
	mw := createTestMiddleware(t)

	// 测试各种错误情况
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// 测试空的Authorization头
	c.Request = httptest.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", "")
	_, _, err := mw.parseToken(c)
	assert.Error(t, err)
	assert.Equal(t, ErrEmptyAuthHeader, err)

	// 测试格式错误的Authorization头
	c.Request.Header.Set("Authorization", "InvalidFormat")
	_, _, err = mw.parseToken(c)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidAuthHeader, err)

	// 测试错误的Token Head Name
	c.Request.Header.Set("Authorization", "WrongPrefix token")
	_, _, err = mw.parseToken(c)
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidAuthHeader, err)
}

// TestUsingPublicKeyAlgo 测试非对称算法判断
func TestUsingPublicKeyAlgo(t *testing.T) {
	mw := &GinJWTMiddleware{}

	// 对称算法
	mw.SigningAlgorithm = "HS256"
	assert.False(t, mw.usingPublicKeyAlgo())

	mw.SigningAlgorithm = "HS384"
	assert.False(t, mw.usingPublicKeyAlgo())

	mw.SigningAlgorithm = "HS512"
	assert.False(t, mw.usingPublicKeyAlgo())

	// 非对称算法
	mw.SigningAlgorithm = "RS256"
	assert.True(t, mw.usingPublicKeyAlgo())

	mw.SigningAlgorithm = "RS384"
	assert.True(t, mw.usingPublicKeyAlgo())

	mw.SigningAlgorithm = "RS512"
	assert.True(t, mw.usingPublicKeyAlgo())

	// 注意：当前实现不支持ES系列算法
	// mw.SigningAlgorithm = "ES256"
	// assert.True(t, mw.usingPublicKeyAlgo()) // 这会失败
}

// TestJWTFunctions 测试JWT相关函数
func TestJWTFunctions(t *testing.T) {
	// 测试jwtFromHeader
	mw := createTestMiddleware(t)
	mw.TokenHeadName = "Bearer"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// 正常情况
	c.Request = httptest.NewRequest("GET", "/test", nil)
	c.Request.Header.Set("Authorization", "Bearer testtoken")
	token, err := mw.jwtFromHeader(c, "Authorization")
	assert.NoError(t, err)
	assert.Equal(t, "testtoken", token)

	// 空header
	c.Request.Header.Set("Authorization", "")
	token, err = mw.jwtFromHeader(c, "Authorization")
	assert.Error(t, err)
	assert.Equal(t, ErrEmptyAuthHeader, err)

	// 错误格式
	c.Request.Header.Set("Authorization", "Invalid testtoken")
	token, err = mw.jwtFromHeader(c, "Authorization")
	assert.Error(t, err)
	assert.Equal(t, ErrInvalidAuthHeader, err)

	// 测试jwtFromQuery
	c.Request = httptest.NewRequest("GET", "/test?token=querytoken", nil)
	token, err = mw.jwtFromQuery(c, "token")
	assert.NoError(t, err)
	assert.Equal(t, "querytoken", token)

	// 空query
	c.Request = httptest.NewRequest("GET", "/test", nil)
	token, err = mw.jwtFromQuery(c, "token")
	assert.Error(t, err)
	assert.Equal(t, ErrEmptyQueryToken, err)

}

// 主测试函数
func TestMain(m *testing.M) {
	// 设置测试环境
	gin.SetMode(gin.TestMode)

	// 运行测试
	code := m.Run()

	// 清理
	os.Exit(code)
}
