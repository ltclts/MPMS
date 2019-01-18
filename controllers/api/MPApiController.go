package api

import (
	"MPMS/helper"
	"MPMS/models"
	"MPMS/session"
	"MPMS/structure"
	"fmt"
)

/**
小程序接口
*/
type MPApiController struct {
	Controller
}

type MPInfoReq struct {
	Id        int64  `form:"mp_info[id]"`
	Name      string `form:"mp_info[name]"` //!!!这里很奇怪 不能直接用上层接收
	Appid     string `form:"mp_info[appid]"`
	Remark    string `form:"mp_info[remark]"`
	CompanyId int64  `form:"mp_info[company_id]"`
}

func (mp *MPApiController) List() {
	listReq := struct {
		Id int64 `form:"id"`
	}{}
	if err := mp.ParseForm(&listReq); err != nil {
		mp.ApiReturn(structure.Response{Error: 1, Msg: "参数获取失败，请重试！", Info: structure.StringToObjectMap{}})
		return
	}

	var list []structure.StringToObjectMap
	company := models.Company{}
	user := models.User{}
	program := models.MiniProgram{}

	mpWhere := structure.StringToObjectMap{"is_deleted": models.UnDeleted}
	if mp.GetSession(session.UserType).(uint8) == models.UserTypeCustomer { //如果是用户登陆并且没有id 则获取
		company, err := mp.getSessionCompanyInfo()
		if err != nil {
			mp.ApiReturn(structure.Response{Error: 2, Msg: "参数获取失败，请重试！", Info: structure.StringToObjectMap{}})
			return
		}
		mpWhere["company_id"] = company.Id
	}

	if listReq.Id != 0 {
		mpWhere["id"] = listReq.Id
	}

	mpList, err := program.Select([]string{"id", "name", "remark", "appid", "creator_id", "company_id", "created_at"}, mpWhere)
	if err != nil {
		mp.ApiReturn(structure.Response{Error: 3, Msg: fmt.Sprintf("获取数据失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
		return
	}

	for _, item := range mpList {

		listItem := structure.StringToObjectMap{
			"id":         item.Id,
			"name":       item.Name,
			"remark":     item.Remark,
			"appid":      item.Appid,
			"created_at": item.CreatedAt,
		}
		user, err = user.SelectOne([]string{"name"}, structure.StringToObjectMap{"id": item.CreatorId})
		if err != nil {
			mp.ApiReturn(structure.Response{Error: 4, Msg: fmt.Sprintf("获取创建人信息失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
			return
		}

		listItem["creator"] = user.Name

		company, err = company.SelectOne([]string{"short_name", "expire_at"}, structure.StringToObjectMap{"id": item.CompanyId})
		if err != nil {
			mp.ApiReturn(structure.Response{Error: 4, Msg: fmt.Sprintf("获取所属公司信息失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
			return
		}
		listItem["company_name"] = company.ShortName
		listItem["expire_at"] = company.ExpireAt
		list = append(list, listItem)
	}
	mp.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{"list": list}})
}

func (mp *MPApiController) Edit() {
	req := struct {
		OperateType int8 `form:"operate_type"`
		MPInfoReq
	}{}
	if err := mp.ParseForm(&req); err != nil {
		mp.ApiReturn(structure.Response{Error: 1, Msg: "参数解析失败！", Info: structure.StringToObjectMap{}})
		return
	}

	if helper.OperateTypeCreate == req.OperateType { //创建
		mpIns, err := mp.create(req.MPInfoReq)
		if err != nil {
			mp.ApiReturn(structure.Response{Error: 2, Msg: err.Error(), Info: structure.StringToObjectMap{}})
			return
		}
		mp.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{"id": mpIns.Id}})
	} else if helper.OperateTypeEdit == req.OperateType {
		_, err := mp.edit(req.MPInfoReq)
		if err != nil {
			mp.ApiReturn(structure.Response{Error: 3, Msg: err.Error(), Info: structure.StringToObjectMap{}})
			return
		}
		mp.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{}})
	} else {
		mp.ApiReturn(structure.Response{Error: 4, Msg: "参数错误，请刷新重试", Info: structure.StringToObjectMap{}})
	}

}

func (mp *MPApiController) create(req MPInfoReq) (mpIns models.MiniProgram, err error) {
	creatorId := mp.GetSession(session.UUID).(int64)
	_, err = mpIns.StartTrans()
	if err != nil {
		return mpIns, err
	}

	_, err = mpIns.SelectOne([]string{"id"}, structure.StringToObjectMap{"appid": req.Appid, "is_deleted": models.UnDeleted})
	if err != nil {
		_ = mpIns.Rollback()
		return mpIns, err
	}
	if mpIns.Id != 0 {
		_ = mpIns.Rollback()
		return mpIns, helper.CreateNewError("该appid已存在！")
	}

	toInsert := structure.StringToObjectMap{
		"name":       req.Name,
		"remark":     req.Remark,
		"appid":      req.Appid,
		"creator_id": creatorId,
		"company_id": req.CompanyId,
	}
	//创建
	mpId, err := mpIns.Insert(toInsert)
	if err != nil {
		_ = mpIns.Rollback()
		return mpIns, err
	}
	mpIns.Id = mpId
	//写入流水
	flow := models.Flow{}
	_, err = flow.Insert(mpIns.Id, models.FlowReferTypeMinProgram, models.FlowStatusCreate, creatorId, toInsert)
	if err != nil {
		_ = mpIns.Rollback()
		return mpIns, err
	}
	err = mpIns.Commit()
	return mpIns, err
}

func (mp *MPApiController) edit(req MPInfoReq) (mpIns models.MiniProgram, err error) {
	operatorId := mp.GetSession(session.UUID).(int64)
	_, err = mpIns.StartTrans()
	if err != nil {
		return mpIns, err
	}

	toUpdate := structure.StringToObjectMap{
		"name":   req.Name,
		"remark": req.Remark,
	}
	//编辑
	mpId, err := mpIns.Update(toUpdate, structure.StringToObjectMap{"id": req.Id})
	if err != nil {
		_ = mpIns.Rollback()
		return mpIns, err
	}
	mpIns.Id = mpId
	//写入流水
	flow := models.Flow{}
	_, err = flow.Insert(mpIns.Id, models.FlowReferTypeMinProgram, models.FlowStatusEdit, operatorId, toUpdate)
	if err != nil {
		_ = mpIns.Rollback()
		return mpIns, err
	}
	err = mpIns.Commit()
	return mpIns, err
}
