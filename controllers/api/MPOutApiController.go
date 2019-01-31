package api

import (
	"MPMS/models"
	"MPMS/structure"
	"encoding/json"
)

/**
小程序外部调用接口
*/
type MPOutApiController struct {
	Controller
}

//控制器级别的去除校验
func (mp *MPOutApiController) Prepare() {
	mp.EnableXSRF = false
}

func (mp *MPOutApiController) RequestInfo() {
	req := struct {
		Appid string `form:"appid"`
	}{}

	if err := mp.ParseJsonData(&req); err != nil {
		mp.ApiReturn(structure.Response{Error: 1, Msg: "没有获取到请求参数！", Info: structure.StringToObjectMap{}})
		return
	}

	if req.Appid == "" {
		mp.ApiReturn(structure.Response{Error: 2, Msg: "没有获取到appid！", Info: structure.StringToObjectMap{}})
		return
	}

	program := models.MiniProgram{}
	program, err := program.SelectOne([]string{}, structure.StringToObjectMap{"is_deleted": models.UnDeleted, "appid": req.Appid})
	if err != nil {
		mp.ApiReturn(structure.Response{Error: 3, Msg: err.Error(), Info: structure.StringToObjectMap{}})
		return
	}

	if program.Id == 0 {
		mp.ApiReturn(structure.Response{Error: 4, Msg: "没有获取到小程序信息", Info: structure.StringToObjectMap{}})
		return
	}

	version := models.MiniProgramVersion{}
	version, err = version.SelectOne([]string{}, structure.StringToObjectMap{"is_deleted": models.UnDeleted, "mini_program_id": program.Id, "status": models.MiniProgramVersionStatusOnline})
	if err != nil {
		mp.ApiReturn(structure.Response{Error: 5, Msg: err.Error(), Info: structure.StringToObjectMap{}})
		return
	}

	if version.Id == 0 {
		mp.ApiReturn(structure.Response{Error: 6, Msg: "没有获取到小程序版本信息", Info: structure.StringToObjectMap{}})
		return
	}

	if version.Type == models.MiniProgramVersionBusinessCard {
		resource := models.Resource{}
		fields := []string{"id", "relative_path", "sort", "store_type"}
		whereResource := structure.StringToObjectMap{"is_deleted": models.UnDeleted, "refer_id": version.Id}

		whereResource["refer_type"] = models.ResourceReferTypeMiniProgramVersionSharedImg
		shareImgList, err := resource.Select(fields, whereResource)
		if err != nil {
			mp.ApiReturn(structure.Response{Error: 5, Msg: err.Error(), Info: structure.StringToObjectMap{}})
			return
		}

		whereResource["refer_type"] = models.ResourceReferTypeMiniProgramVersionBusinessCardCarousel
		carouselImgList, err := resource.Select(fields, whereResource)
		if err != nil {
			mp.ApiReturn(structure.Response{Error: 6, Msg: err.Error(), Info: structure.StringToObjectMap{}})
			return
		}

		whereResource["refer_type"] = models.ResourceReferTypeMiniProgramVersionBusinessCardElegantDemeanor
		elegantDemeanorImgList, err := resource.Select(fields, whereResource)
		if err != nil {
			mp.ApiReturn(structure.Response{Error: 7, Msg: err.Error(), Info: structure.StringToObjectMap{}})
			return
		}

		var content = structure.StringToObjectMap{}
		_ = json.Unmarshal([]byte(version.Content), &content)
		content["share_words"] = version.ShareWords
		rspInfo := structure.StringToObjectMap{"version": content, "type": models.MiniProgramVersionBusinessCard}
		for indexList, itemList := range map[string][]models.Resource{
			"share_img_list":            shareImgList,
			"carousel_img_list":         carouselImgList,
			"elegant_demeanor_img_list": elegantDemeanorImgList,
		} {
			var itemCopyList []string
			for _, item := range itemList {
				itemCopyList = append(itemCopyList, item.GetRealPath())
			}
			rspInfo[indexList] = itemCopyList
		}

		mp.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: rspInfo})
		return
	}

	mp.ApiReturn(structure.Response{Error: 7, Msg: "未找到任何信息", Info: structure.StringToObjectMap{}})
}
