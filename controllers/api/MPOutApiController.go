package api

import (
	"MPMS/models"
	"MPMS/structure"
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

		rspInfo := structure.StringToObjectMap{"Version": version, "Type": models.MiniProgramVersionBusinessCard}
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

		mp.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: rspInfo})
		return
	}

	mp.ApiReturn(structure.Response{Error: 7, Msg: "未找到任何信息", Info: structure.StringToObjectMap{}})
}
