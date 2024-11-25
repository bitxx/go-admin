package authdto

const (
	UserInfo    = "UserInfo"
	LoginUserId = "LoginUserId"
	RoleId      = "RoleId"
	RoleKey     = "RoleKey"
	DeptId      = "DeptId"
	UserName    = "UserName"
	DataScope   = "DataScope"

	HeaderAuthorization = "Authorization"
	HeaderTokenName     = "Bearer"
)

type Resp struct {
	RequestId string      `json:"requestId"`
	Msg       string      `json:"msg"`
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
}

type Data struct {
	Token    string `json:"token"`
	UserName string `json:"username"`
	//Expire   string      `json:"expire"`
	//UserInfo interface{} `json:"userInfo"`
}
