package api

import (
	"MPMS/helper"
	"MPMS/models"
	"MPMS/routers/uris"
	"MPMS/services/email"
	"MPMS/services/log"
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

	m := models.User{}
	users, err := m.Select([]string{}, structure.StringToObjectMap{"email": loginReq.Email, "is_deleted": models.UnDeleted})
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

	if !m.CheckPwd(loginReq.Password) {
		u.ApiReturn(structure.Response{Error: 5, Msg: "用户名或密码不正确！", Info: structure.StringToObjectMap{}})
		return
	}

	if m.Status != models.UserStatusInUse {
		currentStatusName, err := m.GetStatusName()
		if err != nil {
			u.ApiReturn(structure.Response{Error: 6, Msg: "您的账户处于异常状态，请联系管理员！", Info: structure.StringToObjectMap{}})
			return
		}
		u.ApiReturn(structure.Response{Error: 6, Msg: fmt.Sprintf("您的账户%s，不能进行登陆！", currentStatusName), Info: structure.StringToObjectMap{}})
		return
	}

	if m.Type == models.UserTypeCustomer {
		company := models.Company{}
		company, err = company.GetCompanyByContactUserId(m.Id)
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
	u.SetSession(session.UUID, m.Id)
	u.SetSession(session.UserName, m.Name)
	u.SetSession(session.UserType, m.Type)

	var redirect string
	requestUri := u.GetSession(session.RequestUri)
	if requestUri != nil {
		redirect = requestUri.(string)
	} else {
		if m.Type == models.UserTypeCustomer {
			redirect = uris.HtmlUriMiniProgram
		} else {
			redirect = uris.HtmlUriCompany
		}
	}
	u.ApiReturn(structure.Response{Error: 0, Msg: "ok！", Info: structure.StringToObjectMap{"uri": redirect}})
}

func (u *UserApiController) Logout() {
	u.DestroySession()
	u.ApiReturn(structure.Response{Error: 0, Msg: "ok！", Info: structure.StringToObjectMap{"uri": uris.HtmlUriUserLogin}})
}

func (u *UserApiController) GetCheckCode() {
	req := struct {
		Email string `form:"email"`
	}{}
	if err := u.ParseForm(&req); err != nil {
		u.ApiReturn(structure.Response{Error: 1, Msg: "获取参数失败！", Info: structure.StringToObjectMap{}})
		return
	}
	log.Info("获取验证码", req, u.GetSession(session.UUID))

	if req.Email == "" {
		u.ApiReturn(structure.Response{Error: 2, Msg: "请输入邮箱！", Info: structure.StringToObjectMap{}})
		return
	}

	m := models.User{}
	m, err := m.SelectOne([]string{}, structure.StringToObjectMap{"email": req.Email, "is_deleted": models.UnDeleted})
	if err != nil {
		u.ApiReturn(structure.Response{Error: 3, Msg: err.Error(), Info: structure.StringToObjectMap{}})
		return
	}
	if m.Id != 0 {
		u.ApiReturn(structure.Response{Error: 4, Msg: "该邮箱已存在！", Info: structure.StringToObjectMap{}})
		return
	}

	checkCode := helper.GetRandomStrBy(4)
	email.SetMsg(email.RegisterEmail{Tos: []email.To{{Addr: req.Email}}, Code: checkCode})
	u.SetSession(session.UserRegisterCheckCode, checkCode)
	u.ApiReturn(structure.Response{Error: 0, Msg: "ok！", Info: structure.StringToObjectMap{}})
}

func (u *UserApiController) GetUserInfo() {
	req := struct {
		Id int64 `form:"id"`
	}{}
	if err := u.ParseForm(&req); err != nil {
		u.ApiReturn(structure.Response{Error: 1, Msg: "获取参数失败！", Info: structure.StringToObjectMap{}})
		return
	}

	if req.Id == 0 {
		req.Id = u.GetSession(session.UUID).(int64)
	}

	user := models.User{}
	user, err := user.SelectOne([]string{"id", "email", "phone"}, structure.StringToObjectMap{"is_deleted": models.UnDeleted, "id": req.Id})
	if err != nil || user.Id == 0 {
		u.ApiReturn(structure.Response{Error: 2, Msg: "获取用户数据失败！", Info: structure.StringToObjectMap{}})
		return
	}
	u.ApiReturn(structure.Response{Error: 0, Msg: "ok！", Info: structure.StringToObjectMap{"user": user}})

}

func (u *UserApiController) InfoChange() {
	req := struct {
		Phone       string `form:"phone"`
		NewPassword string `form:"new_password"`
	}{}
	if err := u.ParseForm(&req); err != nil {
		u.ApiReturn(structure.Response{Error: 1, Msg: "获取参数失败！", Info: structure.StringToObjectMap{}})
		return
	}

	user := models.User{}
	userId := u.GetSession(session.UUID).(int64)
	if _, err := user.StartTrans(); err != nil {
		u.ApiReturn(structure.Response{Error: 2, Msg: "信息修改失败，请联系技术人员！", Info: structure.StringToObjectMap{}})
		return
	}
	toUpdate := structure.StringToObjectMap{"phone": req.Phone, "password": helper.Md5(req.NewPassword)}
	where := structure.StringToObjectMap{"id": userId}
	updateCount, err := user.Update(toUpdate, where)
	if err != nil {
		_ = user.Rollback()
		u.ApiReturn(structure.Response{Error: 3, Msg: "信息修改失败，请联系技术人员！", Info: structure.StringToObjectMap{}})
		return
	}
	if 1 != updateCount {
		_ = user.Rollback()
		u.ApiReturn(structure.Response{Error: 4, Msg: "信息修改失败，请联系技术人员！", Info: structure.StringToObjectMap{}})
		return
	}
	//写入流水
	flow := models.Flow{}
	_, err = flow.Insert(userId, models.FlowReferTypeContactUser, models.FlowStatusEdit, userId, structure.StringToObjectMap{"update": toUpdate, "where": where})
	if err != nil {
		_ = user.Rollback()
		u.ApiReturn(structure.Response{Error: 5, Msg: "信息修改失败，请联系技术人员！", Info: structure.StringToObjectMap{}})
		return
	}

	err = user.Commit()
	if err != nil {
		u.ApiReturn(structure.Response{Error: 6, Msg: "信息修改失败，请联系技术人员！", Info: structure.StringToObjectMap{}})
		return
	}
	u.ApiReturn(structure.Response{Error: 0, Msg: "ok！", Info: structure.StringToObjectMap{}})
}
