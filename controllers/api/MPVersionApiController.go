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
type MPVersionApiController struct {
	Controller
}

func (mpv *MPVersionApiController) List() {
	req := struct {
		Id            int64 `form:"id"`
		CompanyId     int64 `form:"company_id"`
		MiniProgramId int64 `form:"mini_program_id"`
	}{}
	if err := mpv.ParseForm(&req); err != nil {
		mpv.ApiReturn(structure.Response{Error: 1, Msg: "参数获取失败，请重试！", Info: structure.StringToObjectMap{}})
		return
	}

	where := structure.StringToObjectMap{"mpv.`is_deleted`": models.UnDeleted, "c.`is_deleted`": models.UnDeleted}
	if req.Id != 0 {
		where["mpv.`id`"] = req.Id
	}
	if req.CompanyId != 0 {
		where["c.`id`"] = req.CompanyId
	}

	if req.MiniProgramId != 0 {
		where["mp.`id`"] = req.MiniProgramId
	}

	//用户只能看到本公司的
	if mpv.GetSession(session.UserType).(uint8) == models.UserTypeCustomer {
		company, err := mpv.getSessionCompanyInfo()
		if err != nil {
			mpv.ApiReturn(structure.Response{Error: 2, Msg: "参数获取失败，请重试！", Info: structure.StringToObjectMap{}})
			return
		}
		where["c.id"] = company.Id
	}

	version := models.MiniProgramVersion{}
	list, err := version.GetList(where)
	if err != nil {
		mpv.ApiReturn(structure.Response{Error: 3, Msg: err.Error(), Info: structure.StringToObjectMap{}})
		return
	}

	for index, item := range list {
		statusName, _ := models.GetMiniProgramVersionStatusNameByStatus(item.Status)
		list[index].StatusName = statusName
	}
	fmt.Println(list)
	mpv.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{"list": list}})
}

func (mpv *MPVersionApiController) Edit() {
	req := struct {
		OperateType int8 `form:"operate_type"`
		MPInfoReq
	}{}
	if err := mpv.ParseForm(&req); err != nil {
		mpv.ApiReturn(structure.Response{Error: 1, Msg: "参数解析失败！", Info: structure.StringToObjectMap{}})
		return
	}

	if helper.OperateTypeCreate == req.OperateType { //创建
		mpIns, err := mpv.create(req.MPInfoReq)
		if err != nil {
			mpv.ApiReturn(structure.Response{Error: 2, Msg: err.Error(), Info: structure.StringToObjectMap{}})
			return
		}
		mpv.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{"id": mpIns.Id}})
	} else if helper.OperateTypeEdit == req.OperateType {
		_, err := mpv.edit(req.MPInfoReq)
		if err != nil {
			mpv.ApiReturn(structure.Response{Error: 3, Msg: err.Error(), Info: structure.StringToObjectMap{}})
			return
		}
		mpv.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{}})
	} else {
		mpv.ApiReturn(structure.Response{Error: 4, Msg: "参数错误，请刷新重试", Info: structure.StringToObjectMap{}})
	}

}

func (mpv *MPVersionApiController) create(req MPInfoReq) (mpIns models.MiniProgram, err error) {
	creatorId := mpv.GetSession(session.UUID).(int64)
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

func (mpv *MPVersionApiController) edit(req MPInfoReq) (mpIns models.MiniProgram, err error) {
	operatorId := mpv.GetSession(session.UUID).(int64)
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

/**
轮播图上传接口
*/
func (mpv *MPVersionApiController) CarouselUpload() {
	mpv.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{}})
}
