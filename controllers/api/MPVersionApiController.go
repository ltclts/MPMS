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

func (mpv *MPVersionApiController) Get() {

	req := struct {
		Id int64 `form:"id"`
	}{}

	if err := mpv.ParseForm(&req); err != nil {
		mpv.ApiReturn(structure.Response{Error: 1, Msg: "参数获取失败，请重试！", Info: structure.StringToObjectMap{}})
		return
	}

	where := structure.StringToObjectMap{"mpv.`is_deleted`": models.UnDeleted, "c.`is_deleted`": models.UnDeleted, "mpv.`id`": req.Id}
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

	if len(list) == 0 {
		mpv.ApiReturn(structure.Response{Error: 4, Msg: "没有获取到该版本信息", Info: structure.StringToObjectMap{}})
		return
	}

	info := list[0]
	info.StatusName, _ = models.GetMiniProgramVersionStatusNameByStatus(info.Status)
	resource := models.Resource{}
	fields := []string{"id", "relative_path", "sort", "store_type"}
	whereResource := structure.StringToObjectMap{"is_deleted": models.UnDeleted, "refer_id": info.Id}

	whereResource["refer_type"] = models.ResourceReferTypeMiniProgramVersionSharedImg
	shareImgList, err := resource.Select(fields, whereResource)
	if err != nil {
		mpv.ApiReturn(structure.Response{Error: 5, Msg: err.Error(), Info: structure.StringToObjectMap{}})
		return
	}

	whereResource["refer_type"] = models.ResourceReferTypeMiniProgramVersionBusinessCardCarousel
	carouselImgList, err := resource.Select(fields, whereResource)
	if err != nil {
		mpv.ApiReturn(structure.Response{Error: 6, Msg: err.Error(), Info: structure.StringToObjectMap{}})
		return
	}

	whereResource["refer_type"] = models.ResourceReferTypeMiniProgramVersionBusinessCardElegantDemeanor
	elegantDemeanorImgList, err := resource.Select(fields, whereResource)
	if err != nil {
		mpv.ApiReturn(structure.Response{Error: 7, Msg: err.Error(), Info: structure.StringToObjectMap{}})
		return
	}

	rspInfo := structure.StringToObjectMap{"Version": info}
	for indexList, itemList := range map[string][]models.Resource{"ShareImgList": shareImgList, "CarouselImgList": carouselImgList, "ElegantDemeanorImgList": elegantDemeanorImgList} {
		var itemCopyList []structure.StringToObjectMap
		for _, item := range itemList {
			var itemCopy = structure.StringToObjectMap{}
			itemCopy["Id"] = item.Id
			itemCopy["Path"] = item.GetRealPath()
			itemCopyList = append(itemCopyList, itemCopy)
		}
		rspInfo[indexList] = itemCopyList
	}

	mpv.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: rspInfo})
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
	mpv.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{"list": list}})
}

type ShareInfo struct {
	ShareWords   string `form:"share_info[share_words]"`
	ImgToAddList []int  `form:"share_info[img_to_add][]"` //slice 只支持int 和string 两种类型 参数获取
	ImgToDelList []int  `form:"share_info[img_to_del][]"`
}

type CarouselInfo struct {
	ImgToAddList  []int `form:"param[carousel_info][img_to_add][]"`
	ImgToDelList  []int `form:"param[carousel_info][img_to_del][]"`
	ImgToSortList []int `form:"param[carousel_info][img_to_sort][]"`
}

type ElegantDemeanorInfo struct {
	ImgToAddList  []int `form:"param[elegant_demeanor_info][img_to_add][]"`
	ImgToDelList  []int `form:"param[elegant_demeanor_info][img_to_del][]"`
	ImgToSortList []int `form:"param[elegant_demeanor_info][img_to_sort][]"`
}

type VersionContentInfo struct {
	Name    string `form:"param[content][name]"`
	Flag    string `form:"param[content][flag]"`
	Tel     string `form:"param[content][tel]"`
	Address string `form:"param[content][address]"`
	Lng     string `form:"param[content][lng]"`
	Lat     string `form:"param[content][lat]"`
}

type MpVersionEditReq struct {
	Id          int64 `form:"id"`
	MpId        int64 `form:"mp_id"`
	OperateType int   `form:"operate_type"`
	Type        int   `form:"type"`
	ShareInfo

	//名片展示
	CarouselInfo
	ElegantDemeanorInfo
	VersionContentInfo
}

func (mpv *MPVersionApiController) Edit() {
	req := MpVersionEditReq{}
	if err := mpv.ParseForm(&req); err != nil {
		mpv.ApiReturn(structure.Response{Error: 1, Msg: "参数解析失败！", Info: structure.StringToObjectMap{}})
		return
	}

	if req.Type == models.MiniProgramVersionBusinessCard {
		mpvIns, err := mpv.businessCardEdit(req)
		if err != nil {
			mpv.ApiReturn(structure.Response{Error: 2, Msg: err.Error(), Info: structure.StringToObjectMap{}})
			return
		}
		mpv.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{"id": mpvIns.Id}})
		return
	} else {
		mpv.ApiReturn(structure.Response{Error: 3, Msg: "参数错误，请刷新重试", Info: structure.StringToObjectMap{}})
		return
	}
}

func (mpv *MPVersionApiController) businessCardEdit(req MpVersionEditReq) (mpvIns models.MiniProgramVersion, err error) {
	shareInfo := req.ShareInfo
	carouselInfo := req.CarouselInfo
	elegantDemeanorInfo := req.ElegantDemeanorInfo
	creatorId := mpv.GetSession(session.UUID).(int64)

	_, err = mpvIns.StartTrans()
	if err != nil {
		return mpvIns, err
	}
	var operateType uint8
	operateInfo := structure.StringToObjectMap{}
	if req.OperateType == helper.OperateTypeCreate { //创建
		if req.MpId == 0 {
			_ = mpvIns.Rollback()
			return mpvIns, helper.CreateNewError("参数错误，请刷新重试！")
		}

		toInsert := structure.StringToObjectMap{
			"mini_program_id": req.MpId,
			"type":            req.Type,
			"status":          models.MiniProgramVersionStatusApproved,
			"creator_id":      creatorId,
			"share_words":     shareInfo.ShareWords,
			"content": structure.StringToObjectMap{
				"name":    req.VersionContentInfo.Name,
				"flag":    req.VersionContentInfo.Flag,
				"address": req.VersionContentInfo.Address,
				"tel":     req.VersionContentInfo.Tel,
				"lng":     req.VersionContentInfo.Lng,
				"lat":     req.VersionContentInfo.Lat,
			},
			"is_deleted": models.UnDeleted,
		}
		mpvId, err := mpvIns.Insert(toInsert)
		if err != nil {
			_ = mpvIns.Rollback()
			return mpvIns, err
		}
		req.Id = mpvId
		operateInfo = toInsert
		operateType = models.FlowStatusCreate
	} else if req.OperateType == helper.OperateTypeEdit { //更新
		if req.Id == 0 {
			_ = mpvIns.Rollback()
			return mpvIns, helper.CreateNewError("参数错误，请刷新重试！")
		}
		toUpdate := structure.StringToObjectMap{
			"share_words": shareInfo.ShareWords,
			"content": structure.StringToObjectMap{
				"name":    req.VersionContentInfo.Name,
				"flag":    req.VersionContentInfo.Flag,
				"address": req.VersionContentInfo.Address,
				"tel":     req.VersionContentInfo.Tel,
				"lng":     req.VersionContentInfo.Lng,
				"lat":     req.VersionContentInfo.Lat,
			},
		}
		where := structure.StringToObjectMap{"id": req.Id}
		updateCount, err := mpvIns.Update(toUpdate, where)
		if err != nil {
			_ = mpvIns.Rollback()
			return mpvIns, err
		}
		if updateCount == 0 {
			_ = mpvIns.Rollback()
			return mpvIns, helper.CreateNewError("更新版本失败，请重试！")
		}
		operateInfo["toUpdate"] = toUpdate
		operateInfo["where"] = where
		operateType = models.FlowStatusEdit
	} else {
		_ = mpvIns.Rollback()
		return mpvIns, helper.CreateNewError("未知操作类型，请刷新重试！")
	}

	flow := models.Flow{}
	_, err = flow.Insert(req.Id, models.FlowReferTypeMinProgramVersion, operateType, creatorId, operateInfo)
	if err != nil {
		_ = mpvIns.Rollback()
		return mpvIns, err
	}

	mpvIns, err = mpvIns.SelectOne([]string{}, structure.StringToObjectMap{"id": req.Id})
	if err != nil {
		_ = mpvIns.Rollback()
		return mpvIns, err
	}

	resource := models.Resource{}
	toUpdate := structure.StringToObjectMap{}
	where := structure.StringToObjectMap{}

	//删除
	imgToDelListAll := [][]int{shareInfo.ImgToDelList, carouselInfo.ImgToDelList, elegantDemeanorInfo.ImgToDelList}
	for _, imgToDelList := range imgToDelListAll {
		for _, imgId := range imgToDelList {
			where["id"] = imgId
			toUpdate["is_deleted"] = models.Deleted
			updateCount, err := resource.Update(toUpdate, where)
			if err != nil {
				_ = mpvIns.Rollback()
				return mpvIns, err
			}
			if updateCount == 0 {
				_ = mpvIns.Rollback()
				return mpvIns, helper.CreateNewError("图片删除失败，请重试！")
			}
			_, err = flow.Insert(
				int64(imgId),
				models.FlowReferTypeResource,
				models.FlowStatusDelete,
				creatorId,
				structure.StringToObjectMap{
					"toUpdate": toUpdate, "where": where,
				})
			if err != nil {
				_ = mpvIns.Rollback()
				return mpvIns, err
			}
		}
	}

	//添加
	toUpdate = structure.StringToObjectMap{}
	where = structure.StringToObjectMap{}
	imgToAddListAll := [][]int{shareInfo.ImgToAddList, carouselInfo.ImgToAddList, elegantDemeanorInfo.ImgToAddList}
	for _, imgToAddList := range imgToAddListAll {
		for _, imgId := range imgToAddList {
			where["id"] = imgId
			toUpdate["refer_id"] = mpvIns.Id
			updateCount, err := resource.Update(toUpdate, where)
			if err != nil {
				_ = mpvIns.Rollback()
				return mpvIns, err
			}
			if updateCount == 0 {
				_ = mpvIns.Rollback()
				return mpvIns, helper.CreateNewError("图片保存失败，请重试！")
			}
			_, err = flow.Insert(
				int64(imgId),
				models.FlowReferTypeResource,
				models.FlowStatusEdit,
				creatorId,
				structure.StringToObjectMap{
					"toUpdate": toUpdate, "where": where,
				})
			if err != nil {
				_ = mpvIns.Rollback()
				return mpvIns, err
			}
		}
	}

	//排序
	toUpdate = structure.StringToObjectMap{}
	where = structure.StringToObjectMap{}
	imgToSortListAll := [][]int{carouselInfo.ImgToSortList, elegantDemeanorInfo.ImgToSortList}
	for _, imgToSortList := range imgToSortListAll {
		for index, imgId := range imgToSortList {
			where["id"] = imgId
			toUpdate["sort"] = index
			_, err := resource.Update(toUpdate, where)
			if err != nil {
				_ = mpvIns.Rollback()
				return mpvIns, err
			}
			_, err = flow.Insert(
				int64(imgId),
				models.FlowReferTypeResource,
				models.FlowStatusEdit,
				creatorId,
				structure.StringToObjectMap{
					"toUpdate": toUpdate, "where": where,
				})
			if err != nil {
				_ = mpvIns.Rollback()
				return mpvIns, err
			}
		}
	}

	return mpvIns, mpvIns.Commit()
}

/**
上传接口
*/
func (mpv *MPVersionApiController) Upload() {
	req := struct {
		Id           int64 `form:"id"`
		ReferType    uint8 `form:"refer_type"`
		CurrentCount int64 `form:"current_count"`
	}{}
	if err := mpv.ParseForm(&req); err != nil {
		mpv.ApiReturn(structure.Response{Error: 1, Msg: "参数解析失败！", Info: structure.StringToObjectMap{}})
		return
	}

	typeToAllowedCountMap := structure.Uint8ToInt64{ //todo 写入配置
		models.ResourceReferTypeMiniProgramVersionSharedImg:                   1,
		models.ResourceReferTypeMiniProgramVersionBusinessCardCarousel:        4,
		models.ResourceReferTypeMiniProgramVersionBusinessCardElegantDemeanor: 4,
	}

	if allowedCount := typeToAllowedCountMap[req.ReferType]; allowedCount == 0 || allowedCount <= req.CurrentCount {
		mpv.ApiReturn(structure.Response{Error: 2, Msg: fmt.Sprintf("当前已上传%d张，无法继续上传！", req.CurrentCount), Info: structure.StringToObjectMap{}})
		return
	}

	f, h, err := mpv.GetFile("file")
	if err != nil {
		mpv.ApiReturn(structure.Response{Error: 2, Msg: err.Error(), Info: structure.StringToObjectMap{}})
		return
	}

	resourceId, url, err := Upload(
		f, h, req.ReferType, 0, mpv.GetSession(session.UUID).(int64),
	)
	if err != nil {
		mpv.ApiReturn(structure.Response{Error: 3, Msg: err.Error(), Info: structure.StringToObjectMap{}})
		return
	}

	mpv.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{"url": url, "resource_id": resourceId}})
}
