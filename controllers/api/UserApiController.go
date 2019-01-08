package api

import (
	"MPMS/models"
	"MPMS/routers/uris"
	"MPMS/session"
	"MPMS/structure"
)

type UserApiController struct {
	Controller
}

type loginParas struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (u *UserApiController) Login() {
	param := loginParas{}
	if err := u.ParseForm(&param); err != nil {
		u.ApiReturn(structure.Response{Error: 1, Msg: "用户名或密码不正确！", Info: structure.Map{}})
	}

	user := models.User{}
	users, err := user.Select([]string{}, structure.Map{"email": param.Email, "is_deleted": models.UnDeleted})
	if err != nil {
		u.ApiReturn(structure.Response{Error: 2, Msg: err.Error(), Info: structure.Map{}})
	}

	if len(users) == 0 {
		u.ApiReturn(structure.Response{Error: 3, Msg: "用户名或密码不正确！", Info: structure.Map{}})
	}

	if len(users) > 1 {
		u.ApiReturn(structure.Response{Error: 4, Msg: "用户名不唯一，请核查！", Info: structure.Map{}})
	}

	if !user.CheckPwd(param.Password) {
		u.ApiReturn(structure.Response{Error: 5, Msg: "用户名或密码不正确！", Info: structure.Map{}})
	}

	if user.Status != models.UserStatusInUse {
		currentStatusName := (user.GetUserStatusToNameMap())[user.Status]
		if currentStatusName == "" {
			u.ApiReturn(structure.Response{Error: 6, Msg: "您的账户处于异常状态，请联系管理员！", Info: structure.Map{}})
		}
		u.ApiReturn(structure.Response{Error: 6, Msg: "您的账户" + currentStatusName + "，不能进行登陆！", Info: structure.Map{}})
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

	u.ApiReturn(structure.Response{Error: 0, Msg: "ok！", Info: structure.Map{"uri": redirect}})
}

func (u *UserApiController) Logout() {
	u.SetSession(session.UUID, nil)
	u.SetSession(session.UserName, nil)
	u.SetSession(session.UserType, nil)
	u.ApiReturn(structure.Response{Error: 0, Msg: "ok！", Info: structure.Map{"uri": uris.HtmlUriLogin}})
}
