package api

import (
	"MPMS/helper"
	"MPMS/models"
	"MPMS/services/email"
	"MPMS/session"
	"MPMS/structure"
	"fmt"
)

type CompanyApiController struct {
	Controller
}

func (c *CompanyApiController) List() {
	listReq := struct {
		Name       string  `form:"name"`        //公司名称
		StatusList []uint8 `form:"status_list"` //状态
		PageSize   uint64  `form:"page_size"`   //每页数量
		PageIndex  uint64  `form:"page_index"`  //页码
		Id         int64   `form:"id"`
	}{}
	if err := c.ParseForm(&listReq); err != nil {
		c.ApiReturn(structure.Response{Error: 1, Msg: "参数获取失败，请重试！", Info: structure.StringToObjectMap{}})
		return
	}

	var list []structure.StringToObjectMap
	company := models.Company{}
	user := models.User{}
	program := models.MiniProgram{}
	companyWhereMap := structure.StringToObjectMap{"is_deleted": models.UnDeleted}

	if c.GetSession(session.UserType).(uint8) == models.UserTypeCustomer && listReq.Id == 0 { //如果是用户登陆并且没有id 则获取
		company, err := c.getSessionCompanyInfo()
		if err != nil {
			c.ApiReturn(structure.Response{Error: 2, Msg: "参数获取失败，请重试！", Info: structure.StringToObjectMap{}})
			return
		}
		listReq.Id = company.Id
	}

	if listReq.Id != 0 {
		companyWhereMap["id"] = listReq.Id
	}
	companies, err := company.Select([]string{"id", "short_name", "creator_id", "status", "expire_at"}, companyWhereMap)
	if err != nil {
		c.ApiReturn(structure.Response{Error: 3, Msg: fmt.Sprintf("获取数据失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
		return
	}

	for _, item := range companies {
		statusName, _ := item.GetStatusName()
		listItem := structure.StringToObjectMap{"id": item.Id, "name": item.ShortName, "status": item.Status, "status_name": statusName, "expire_at": item.ExpireAt}
		user, err = user.GetContactUserByCompanyId(item.Id)
		if err != nil {
			c.ApiReturn(structure.Response{Error: 4, Msg: fmt.Sprintf("获取联系人数据失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
			return
		}

		listItem["company_contact_user"] = user.Name
		listItem["company_contact_user_phone"] = user.Phone

		//创建人
		user, err = user.SelectOne([]string{"name", "phone"}, structure.StringToObjectMap{
			"is_deleted": models.UnDeleted,
			"id":         item.CreatorId,
		})
		if err != nil {
			c.ApiReturn(structure.Response{Error: 5, Msg: fmt.Sprintf("获取创建人数据失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
			return
		}

		listItem["creator"] = user.Name

		//小程序个数获取
		count, _ := program.Count(structure.StringToObjectMap{
			"is_deleted": models.UnDeleted,
			"company_id": item.Id,
		})
		listItem["mp_count"] = count
		list = append(list, listItem)
	}
	c.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{"list": list}})
}

type CompanyInfo struct {
	Id        int64  `form:"company_info[id]"`
	Name      string `form:"company_info[name]"`
	ShortName string `form:"company_info[short_name]"`
	ExpireAt  string `form:"company_info[expire_at]"`
	Remark    string `form:"company_info[remark]"`
}
type UserInfo struct {
	Name      string `form:"user_info[name]"`
	Email     string `form:"user_info[email]"`
	CheckCode string `form:"user_info[check_code]"`
	Phone     string `form:"user_info[phone]"`
}

func (c *CompanyApiController) Edit() {
	req := struct {
		OperateType uint8 `form:"operate_type"`
		CompanyInfo
		UserInfo
	}{}
	if err := c.ParseForm(&req); err != nil {
		c.ApiReturn(structure.Response{Error: 1, Msg: "参数获取失败，请重试！", Info: structure.StringToObjectMap{}})
		return
	}
	fmt.Println(req)
	if helper.OperateTypeCreate == req.OperateType { //创建
		if err := c.checkRegisterCode(req.UserInfo.CheckCode); err != nil {
			c.ApiReturn(structure.Response{Error: 2, Msg: err.Error(), Info: structure.StringToObjectMap{}})
			return
		}

		if models.UserTypeAdmin != c.GetSession(session.UserType).(uint8) {
			c.ApiReturn(structure.Response{Error: 3, Msg: "您没有创建公司的权限！", Info: structure.StringToObjectMap{}})
			return
		}

		company, err := c.create(req.CompanyInfo, req.UserInfo)
		if err != nil {
			c.ApiReturn(structure.Response{Error: 4, Msg: err.Error(), Info: structure.StringToObjectMap{}})
			return
		}
		c.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{"id": company.Id}})
		return
	} else if helper.OperateTypeEdit == req.OperateType { //编辑
		if req.CompanyInfo.Id == 0 {
			c.ApiReturn(structure.Response{Error: 5, Msg: "参数错误，请刷新重试！", Info: structure.StringToObjectMap{}})
			return
		}

		if req.UserInfo.CheckCode != "" {
			if err := c.checkRegisterCode(req.UserInfo.CheckCode); err != nil {
				c.ApiReturn(structure.Response{Error: 6, Msg: err.Error(), Info: structure.StringToObjectMap{}})
				return
			}
		}
		_, err := c.update(req.CompanyInfo, req.UserInfo)
		if err != nil {
			c.ApiReturn(structure.Response{Error: 7, Msg: err.Error(), Info: structure.StringToObjectMap{}})
			return
		}
		c.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{}})
		return
	} else {
		c.ApiReturn(structure.Response{Error: 1, Msg: "参数错误，请刷新重试", Info: structure.StringToObjectMap{}})
		return
	}

}

func (c *CompanyApiController) checkRegisterCode(checkCode string) error {

	if checkCodeStored := c.GetSession(session.UserRegisterCheckCode); checkCodeStored != nil {
		checkCodeStored := checkCodeStored.(string)
		if checkCode != checkCodeStored {
			return helper.CreateNewError("验证码不正确！")
		}

		c.DelSession(session.UserRegisterCheckCode)
		return nil
	}

	return helper.CreateNewError("请获取验证码")
}

func (c *CompanyApiController) update(companyInfo CompanyInfo, userInfo UserInfo) (company models.Company, err error) {
	operatorId := c.GetSession(session.UUID).(int64)
	_, err = company.StartTrans()
	if err != nil {
		return company, err
	}
	infoToUpdate := structure.StringToObjectMap{
		"name":       companyInfo.Name,
		"short_name": companyInfo.ShortName,
		"remark":     companyInfo.Remark,
		"expire_at":  companyInfo.ExpireAt,
	}
	updateCount, err := company.Update(infoToUpdate, structure.StringToObjectMap{"id": companyInfo.Id})
	if err != nil {
		_ = company.Rollback()
		return company, err
	}
	if 1 != updateCount {
		_ = company.Rollback()
		return company, helper.CreateNewError("更新公司信息失败！")
	}
	//写入流水
	flow := models.Flow{}
	_, err = flow.Insert(companyInfo.Id, models.FlowReferTypeCompany, models.FlowStatusEdit, operatorId, infoToUpdate)
	if err != nil {
		_ = company.Rollback()
		return company, err
	}

	user := models.User{}
	user, err = user.GetContactUserByCompanyId(companyInfo.Id)
	if err != nil {
		_ = company.Rollback()
		return company, err
	}

	userExisted := models.User{}
	if user.Email != userInfo.Email {
		_, err := userExisted.SelectOne([]string{"id"}, structure.StringToObjectMap{"email": userInfo.Email, "is_deleted": models.UnDeleted})
		if err != nil {
			_ = company.Rollback()
			return company, err
		}

		if userExisted.Id != user.Id {
			_ = company.Rollback()
			return company, helper.CreateNewError("该邮箱已注册！")
		}
	}

	userInfoToUpdate := structure.StringToObjectMap{"name": userInfo.Name, "email": userInfo.Email, "phone": userInfo.Phone}
	updateCount, err = user.Update(userInfoToUpdate, structure.StringToObjectMap{"id": user.Id})
	if err != nil {
		_ = company.Rollback()
		return company, err
	}
	if 1 != updateCount {
		_ = company.Rollback()
		return company, helper.CreateNewError("更新联系人信息失败！")
	}

	_, err = flow.Insert(user.Id, models.FlowReferTypeContactUser, models.FlowStatusEdit, operatorId, userInfoToUpdate)
	if err != nil {
		_ = company.Rollback()
		return company, err
	}

	err = company.Commit()
	return company, err
}

func (c *CompanyApiController) create(companyInfo CompanyInfo, userInfo UserInfo) (company models.Company, err error) {
	operatorId := c.GetSession(session.UUID).(int64)
	_, err = company.StartTrans()
	if err != nil {
		return company, err
	}

	infoToSave := structure.StringToObjectMap{
		"name":       companyInfo.Name,
		"short_name": companyInfo.ShortName,
		"remark":     companyInfo.Remark,
		"expire_at":  companyInfo.ExpireAt,
		"creator_id": operatorId,
	}
	companyId, err := company.Insert(infoToSave)
	if err != nil {
		_ = company.Rollback()
		return company, err
	}
	//写入流水
	flow := models.Flow{}
	_, err = flow.Insert(companyId, models.FlowReferTypeCompany, models.FlowStatusCreate, operatorId, infoToSave)
	if err != nil {
		_ = company.Rollback()
		return company, err
	}
	company.Id = companyId
	user := models.User{}
	user, err = user.SelectOne([]string{"id"}, structure.StringToObjectMap{"email": userInfo.Email, "is_deleted": models.UnDeleted})
	if err != nil {
		_ = company.Rollback()
		return company, err
	}
	if user.Id != 0 {
		_ = company.Rollback()
		return company, helper.CreateNewError("该邮箱已注册！")
	}

	user.Phone = userInfo.Phone
	user.Email = userInfo.Email
	user.Name = userInfo.Name
	_, err = user.CreateContactUser(companyId, operatorId)
	if err != nil {
		_ = company.Rollback()
		return company, err
	}

	err = company.Commit()
	if err == nil {
		//注册邮件
		email.SetMsg(email.NoticePasswordEmail{Tos: []email.To{{Name: user.Name, Addr: user.Email}}, Password: user.Password})
	}
	return company, err
}

func (c *CompanyApiController) GetEditInfo() {
	req := struct {
		Id int64 `form:"id"`
	}{}
	if err := c.ParseForm(&req); err != nil {
		c.ApiReturn(structure.Response{Error: 1, Msg: "参数获取失败，请重试！", Info: structure.StringToObjectMap{}})
		return
	}

	if c.GetSession(session.UserType).(uint8) == models.UserTypeCustomer && req.Id == 0 { //如果是用户登陆并且没有id 则获取
		company, err := c.getSessionCompanyInfo()
		if err != nil {
			c.ApiReturn(structure.Response{Error: 2, Msg: "参数获取失败，请重试！", Info: structure.StringToObjectMap{}})
			return
		}
		req.Id = company.Id
	}

	companyWhereMap := structure.StringToObjectMap{"is_deleted": models.UnDeleted, "id": req.Id}
	company := models.Company{}
	_, err := company.SelectOne([]string{}, companyWhereMap)
	if err != nil {
		c.ApiReturn(structure.Response{Error: 3, Msg: fmt.Sprintf("获取数据失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
		return
	}

	if company.Id == 0 {
		c.ApiReturn(structure.Response{Error: 3, Msg: "没有获取到公司信息！", Info: structure.StringToObjectMap{}})
		return
	}

	user := models.User{}
	user, err = user.GetContactUserByCompanyId(company.Id)
	if err != nil {
		c.ApiReturn(structure.Response{Error: 4, Msg: fmt.Sprintf("获取联系人数据失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
		return
	}
	c.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{
		"company_info": company,
		"user_info":    user,
	}})
}

func (c *CompanyApiController) UpdateStatus() {

	if models.UserTypeAdmin != c.GetSession(session.UserType).(uint8) {
		c.ApiReturn(structure.Response{Error: 3, Msg: "您没有修改权限！", Info: structure.StringToObjectMap{}})
		return
	}
	operatorId := c.GetSession(session.UUID).(int64)
	req := struct {
		Id         int64 `form:"id"`
		ToStatus   uint8 `form:"to_status"`
		FromStatus uint8 `form:"from_status"`
	}{}
	if err := c.ParseForm(&req); err != nil {
		c.ApiReturn(structure.Response{Error: 1, Msg: "参数获取失败，请重试！", Info: structure.StringToObjectMap{}})
		return
	}

	if 0 == req.Id || 0 == req.ToStatus {
		c.ApiReturn(structure.Response{Error: 2, Msg: "参数错误，请重试！", Info: structure.StringToObjectMap{}})
		return
	}
	if _, err := models.GetCompanyStatusNameByStatus(req.ToStatus); err != nil {
		c.ApiReturn(structure.Response{Error: 3, Msg: "非法操作", Info: structure.StringToObjectMap{}})
		return
	}

	company := models.Company{}
	if _, err := company.StartTrans(); err != nil {
		c.ApiReturn(structure.Response{Error: 4, Msg: "状态变更失败，请联系技术人员！", Info: structure.StringToObjectMap{}})
		return
	}
	toUpdate := structure.StringToObjectMap{"status": req.ToStatus}
	where := structure.StringToObjectMap{"id": req.Id, "status": req.FromStatus}
	updateCount, err := company.Update(toUpdate, where)
	if err != nil {
		_ = company.Rollback()
		c.ApiReturn(structure.Response{Error: 5, Msg: "状态变更失败，请联系技术人员！", Info: structure.StringToObjectMap{}})
		return
	}
	if 1 != updateCount {
		_ = company.Rollback()
		c.ApiReturn(structure.Response{Error: 6, Msg: "状态变更失败，请联系技术人员！", Info: structure.StringToObjectMap{}})
		return
	}
	//写入流水
	flow := models.Flow{}
	_, err = flow.Insert(req.Id, models.FlowReferTypeCompany, models.FlowStatusEdit, operatorId, structure.StringToObjectMap{"update": toUpdate, "where": where})
	if err != nil {
		_ = company.Rollback()
		c.ApiReturn(structure.Response{Error: 7, Msg: "状态变更失败，请联系技术人员！", Info: structure.StringToObjectMap{}})
		return
	}

	err = company.Commit()
	if err != nil {
		c.ApiReturn(structure.Response{Error: 8, Msg: "状态变更失败，请联系技术人员！", Info: structure.StringToObjectMap{}})
		return
	}
	c.ApiReturn(structure.Response{Error: 0, Msg: "ok！", Info: structure.StringToObjectMap{}})
}
