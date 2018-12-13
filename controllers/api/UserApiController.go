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
	users, err := user.Select([]string{}, structure.Map{"email": param.Email, "is_deleted": models.UNDELETED})
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

	u.SetSession(session.UUID, user.Id)
	u.SetSession(session.UseName, user.Name)

	var redirect string
	requestUri := u.GetSession(session.RequestUri)
	if requestUri != nil {
		redirect = requestUri.(string)
	} else {
		redirect = uris.HtmlUriIndex
	}

	u.ApiReturn(structure.Response{Error: 0, Msg: "ok！", Info: structure.Map{"uri": redirect}})
}
