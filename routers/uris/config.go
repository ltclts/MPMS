package uris

const (
	HtmlUriIndex = "/" //主页

	HtmlUriLogin             = "/html/user/login"          //登陆页面
	HtmlUriCompany           = "/html/company"             //公司管理-账户管理
	HtmlUriMiniProgram       = "/html/mini_program"        //小程序-管理
	HtmlUriMiniProgramCreate = "/html/mini_program/create" //小程序-创建
	HtmlUriMiniProgramEdit   = "/html/mini_program/edit"   //小程序-创建

	ApiUriLogin  = "/api/user/login"  //登陆接口
	ApiUriLogout = "/api/user/logout" //登出接口

	ApiUriCompanyList     = "/api/company/list"      //公司管理-账户管理-页面数据
	ApiUriMiniProgramEdit = "/api/mini_program/edit" //小程序-创建/编辑

	ApiUriMpOutPageConfigRequest = "/api/mp/out/page_config_request"
)

//不需要登陆的路由
func GetUnCheckLoginUris() []string {
	return []string{
		HtmlUriLogin,
		ApiUriLogin,

		ApiUriMpOutPageConfigRequest,
	}
}
