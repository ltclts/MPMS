package uris

const (
	HtmlUriIndex          = "/"
	HtmlUriUserInfoChange = "/html/user/info_change"
	HtmlUriUserLogin      = "/html/user/login"     //登陆页面
	HtmlUriCompany        = "/html/company"        //公司管理-账户管理
	HtmlUriCompanyCreate  = "/html/company/create" //公司创建
	HtmlUriCompanyEdit    = "/html/company/edit"   //公司编辑

	HtmlUriMiniProgram       = "/html/mini_program"        //小程序-管理
	HtmlUriMiniProgramCreate = "/html/mini_program/create" //小程序-创建页面
	HtmlUriMiniProgramEdit   = "/html/mini_program/edit"   //小程序-编辑页面

	HtmlUriMiniProgramVersion       = "/html/mini_program_version"        //小程序版本-管理
	HtmlUriMiniProgramVersionCreate = "/html/mini_program_version/create" //小程序版本-创建页面
	HtmlUriMiniProgramVersionEdit   = "/html/mini_program_version/edit"   //小程序版本-编辑页面

	ApiUriUserLogin       = "/api/user/login"  //登陆接口
	ApiUriUserLogout      = "/api/user/logout" //登出接口
	ApiUriUserInfoChange  = "/api/user/info_change"
	ApiUriUserGetUserInfo = "/api/user/get_user_info"

	ApiUriUserGetCheckCode = "/api/user/get_check_code"

	ApiUriCompanyList         = "/api/company/list"         //公司管理-账户管理-页面数据
	ApiUriCompanyEdit         = "/api/company/edit"         //公司-创建/编辑
	ApiUriCompanyUpdateStatus = "/api/company/updateStatus" //公司状态变更
	ApiUriCompanyGetEditInfo  = "/api/company/getEditInfo"  //公司-获取 单个

	ApiUriMiniProgramList = "/api/mini_program/list" //小程序列表
	ApiUriMiniProgramEdit = "/api/mini_program/edit" //小程序-创建/编辑

	ApiUriMiniProgramVersionList   = "/api/mini_program_version/list"   //小程序版本列表
	ApiUriMiniProgramVersionEdit   = "/api/mini_program_version/edit"   //小程序版本-创建/编辑
	ApiUriMiniProgramVersionUpload = "/api/mini_program_version/upload" //小程序版本-轮播图上传
	ApiUriMiniProgramVersionGet    = "/api/mini_program_version/get"    //小程序版本-获取版本数据

	ApiUriMpOutInfoGet = "/api/mp/out/info_get"

	ApiUriHelperRefreshDBConPools = "/api/helper/refresh_db_con_pools"
)

//不需要登陆的路由
func GetUnCheckLoginUris() []string {
	return []string{
		HtmlUriUserLogin,
		ApiUriUserLogin,

		ApiUriMpOutInfoGet,
	}
}
