package constant

const (
	//app下载方式
	SysConfIsFrontend = "1"     //前台显示
	RoleKeyAdmin      = "admin" //角色类型，超级管理员

	//角色状态
	//RoleStatusStart = "1" //启用
	//RoleStatusStop  = "2" //停用

	//登录退出
	UserLoginStatus  = "1"
	UserLogoutStatus = "2"

	//数据权限类型
	DataScope1 = "1" //全部数据
	DataScope2 = "2" //自定数据权限
	DataScope3 = "3" //本部门数据权限
	DataScope4 = "4" //本部门及以下数据权限
	DataScope5 = "5" //仅本人数据权限

	// MenuM 目录
	MenuM string = "1"
	// MenuC 菜单
	MenuC string = "2"
	// MenuF 按钮
	MenuF string = "3"

	ApiTypeApp    = "3"
	ApiTypeSys    = "1"
	ApiTypePlugin = "2"
)

// 模板相关
const (
	RouterName         = "router.go"
	BusinessRouterName = "businessRouter.go"
	ApiName            = "api.go"
	ModelName          = "model.go"
	DtoName            = "dto.go"
	ServiceName        = "service.go"
	VueApiJsName       = "vue.js"              //Vue
	VueIndexName       = "vue.index"           //Vue
	ReactApiName       = "react.api.ts"        //React
	ReactFormModalName = "react.formmodal.tsx" //React
	ReactViewName      = "react.view.tsx"      //React
)

// 模板相关
var (
	TemplatInfo = map[string]string{
		RouterName:         "static/template/router.go.template",
		BusinessRouterName: "static/template/business_router.go.template",
		ApiName:            "static/template/apis.go.template",
		ModelName:          "static/template/model.go.template",
		DtoName:            "static/template/dto.go.template",
		ServiceName:        "static/template/service.go.template",
		VueApiJsName:       "static/template/vue.api.js.template",
		VueIndexName:       "static/template/vue.index.template",
		ReactApiName:       "static/template/react.api.ts.template",
		ReactFormModalName: "static/template/react.formmodal.tsx.template",
		ReactViewName:      "static/template/react.view.tsx.template",
	}
)
