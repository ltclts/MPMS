package api

import (
	"MPMS/helper"
	"MPMS/models"
	"MPMS/session"
	"MPMS/structure"
)

/**
小程序接口
*/
type MPApiController struct {
	Controller
}

type MPInfoReq struct {
	Name      string `form:"mp_info[name]"` //!!!这里很奇怪 不能直接用上层接收
	Appid     string `form:"mp_info[appid]"`
	Remark    string `form:"mp_info[remark]"`
	CompanyId int64  `form:"mp_info[company_id]"`
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
		return
	}

	mp.ApiReturn(structure.Response{Error: 4, Msg: "编辑失败", Info: structure.StringToObjectMap{}})
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

	//创建
	mpId, err := mpIns.Insert(structure.StringToObjectMap{
		"name":       req.Name,
		"remark":     req.Remark,
		"appid":      req.Appid,
		"creator_id": creatorId,
		"company_id": req.CompanyId,
	})
	if err != nil {
		_ = mpIns.Rollback()
		return mpIns, err
	}
	mpIns.Id = mpId
	//写入流水
	flow := models.Flow{}
	_, err = flow.Insert(mpIns.Id, models.FlowReferTypeMinProgram, models.FlowStatusCreate, creatorId, structure.StringToObjectMap{})
	if err != nil {
		_ = mpIns.Rollback()
		return mpIns, err
	}
	err = mpIns.Commit()
	return mpIns, err
}
