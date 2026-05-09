package constant

const (
	//登录退出
	UserLoginStatus  = "1"
	UserLogoutStatus = "2"

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
		ReactApiName:       "static/template/react.api.ts.template",
		ReactFormModalName: "static/template/react.formmodal.tsx.template",
		ReactViewName:      "static/template/react.view.tsx.template",
	}
)
