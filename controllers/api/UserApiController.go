package api

import (
	"MPMS/models"
	"MPMS/routers/uris"
	"MPMS/session"
	"MPMS/structure"
	"encoding/json"
	"fmt"
)

type UserApiController struct {
	Controller
}

func (u *UserApiController) Login() {
	//表单方式提交
	loginReq := struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}{}
	if err := u.ParseForm(&loginReq); err != nil {
		u.ApiReturn(structure.Response{Error: 1, Msg: "用户名或密码不正确！", Info: structure.StringToObjectMap{}})
		return
	}

	user := models.User{}
	users, err := user.Select([]string{}, structure.StringToObjectMap{"email": loginReq.Email, "is_deleted": models.UnDeleted})
	if err != nil {
		u.ApiReturn(structure.Response{Error: 2, Msg: err.Error(), Info: structure.StringToObjectMap{}})
		return
	}

	if len(users) == 0 {
		u.ApiReturn(structure.Response{Error: 3, Msg: "用户名或密码不正确！", Info: structure.StringToObjectMap{}})
		return
	}

	if len(users) > 1 {
		u.ApiReturn(structure.Response{Error: 4, Msg: "用户名不唯一，请核查！", Info: structure.StringToObjectMap{}})
		return
	}

	if !user.CheckPwd(loginReq.Password) {
		u.ApiReturn(structure.Response{Error: 5, Msg: "用户名或密码不正确！", Info: structure.StringToObjectMap{}})
		return
	}

	if user.Status != models.UserStatusInUse {
		currentStatusName, err := user.GetStatusName()
		if err != nil {
			u.ApiReturn(structure.Response{Error: 6, Msg: "您的账户处于异常状态，请联系管理员！", Info: structure.StringToObjectMap{}})
			return
		}
		u.ApiReturn(structure.Response{Error: 6, Msg: fmt.Sprintf("您的账户%s，不能进行登陆！", currentStatusName), Info: structure.StringToObjectMap{}})
		return
	}

	if user.Type == models.UserTypeCustomer {
		company := models.Company{}
		company, err = company.GetCompanyByContactUserId(user.Id)
		if err != nil {
			u.ApiReturn(structure.Response{Error: 7, Msg: fmt.Sprintf("获取公司信息失败：%s！", err.Error()), Info: structure.StringToObjectMap{}})
			return
		}
		companyStr, err := json.Marshal(company)
		if err != nil {
			u.ApiReturn(structure.Response{Error: 8, Msg: fmt.Sprintf("获取公司信息失败：%s！", err.Error()), Info: structure.StringToObjectMap{}})
			return
		}
		u.SetSession(session.CompanyInfo, companyStr)
	}
	u.SetSession(session.UUID, user.Id)
	u.SetSession(session.UserName, user.Name)
	u.SetSession(session.UserType, user.Type)

	var redirect string
	requestUri := u.GetSession(session.RequestUri)
	if requestUri != nil {
		redirect = requestUri.(string)
	} else {
		redirect = uris.HtmlUriIndex
	}

	u.ApiReturn(structure.Response{Error: 0, Msg: "ok！", Info: structure.StringToObjectMap{"uri": redirect}})
}

func (u *UserApiController) Logout() {
	u.DestroySession()
	u.ApiReturn(structure.Response{Error: 0, Msg: "ok！", Info: structure.StringToObjectMap{"uri": uris.HtmlUriLogin}})
}
