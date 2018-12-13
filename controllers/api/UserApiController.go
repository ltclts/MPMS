package api

import (
	"MPMS/structure"
)

type UserApiController struct {
	Controller
}

func (u *UserApiController) Login() {
	u.ApiReturn(structure.Response{Error: 1, Msg: "用户名不正确！", Info: structure.Map{}})
}
