package api

import (
	"MPMS/structure"
)

type CompanyApiController struct {
	Controller
}

func (u *CompanyApiController) List() {
	param := struct {
		Name       string  `form:"name"`        //公司名称
		StatusList []uint8 `form:"status_list"` //状态
		PageSize   uint64  `form:"page_size"`   //每页数量
		PageIndex  uint64  `form:"page_index"`  //页码
	}{}

	if err := u.ParseForm(&param); err != nil {
		u.ApiReturn(structure.Response{Error: 1, Msg: "获取数据失败，请重试！", Info: structure.StringToObjectMap{}})
		return
	}
	u.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{}})
}
