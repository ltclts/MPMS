package api

import (
	"MPMS/helper"
	"MPMS/models"
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
	companies, err := company.Select([]string{"id", "short_name", "creator_id", "expire_at"}, companyWhereMap)
	if err != nil {
		c.ApiReturn(structure.Response{Error: 3, Msg: fmt.Sprintf("获取数据失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
		return
	}

	for _, item := range companies {
		statusName, _ := item.GetStatusName()
		listItem := structure.StringToObjectMap{"id": item.Id, "name": item.ShortName, "status": statusName, "expire_at": item.ExpireAt}
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
	if req.OperateType == helper.OperateTypeCreate {
		//todo 校验码检查
		if models.UserTypeAdmin != c.GetSession(session.UserType).(uint8) {
			c.ApiReturn(structure.Response{Error: 2, Msg: "您没有创建公司的权限！", Info: structure.StringToObjectMap{}})
			return
		}

		company, err := c.create(req.CompanyInfo, req.UserInfo)
		if err != nil {
			c.ApiReturn(structure.Response{Error: 3, Msg: err.Error(), Info: structure.StringToObjectMap{}})
			return
		}
		c.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{"id": company.Id}})
		return
	}
	c.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{}})

}

func (c *CompanyApiController) create(companyInfo CompanyInfo, userInfo UserInfo) (company models.Company, err error) {
	creatorId := c.GetSession(session.UUID).(int64)
	_, err = company.StartTrans()
	if err != nil {
		return company, err
	}
	companyId, err := company.Insert(structure.StringToObjectMap{
		"name":       companyInfo.Name,
		"short_name": companyInfo.ShortName,
		"remark":     companyInfo.Remark,
		"expire_at":  companyInfo.ExpireAt,
		"creator_id": creatorId,
	})
	if err != nil {
		_ = company.Rollback()
		return company, err
	}
	//写入流水
	flow := models.Flow{}
	_, err = flow.Insert(companyId, models.FlowReferTypeCompany, models.FlowStatusCreate, creatorId, structure.StringToObjectMap{})
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
		return company, helper.CreateNewError("该邮箱已存在！")
	}

	user.Phone = userInfo.Phone
	user.Email = userInfo.Email
	user.Name = userInfo.Name
	_, err = user.CreateContactUser(companyId, creatorId)
	if err != nil {
		_ = company.Rollback()
		return company, err
	}

	err = company.Commit()
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
