package uris

const (
	HtmlUriIndex = "/" //主页

	HtmlUriLogin = "/html/user/login" //登陆页面

	ApiUriLogin  = "/api/user/login"  //登陆接口
	ApiUriLogout = "/api/user/logout" //登出接口

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
