package api

import (
	"MPMS/structure"
	"fmt"
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

func (mp *MPOutApiController) PageConfigRequest() {
	req := struct {
		Appid string `form:"appid"`
	}{}

	if err := mp.ParseJsonData(&req); err != nil {
		fmt.Println("请求：", req)
		mp.ApiReturn(structure.Response{Error: 1, Msg: "没有获取到请求参数！", Info: structure.StringToObjectMap{}})
		return
	}

	if req.Appid == "" {
		fmt.Println("请求：", req)
		mp.ApiReturn(structure.Response{Error: 1, Msg: "没有获取到appid！", Info: structure.StringToObjectMap{}})
		return
	}

	mp.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{"page_type": req.Appid}})
}
