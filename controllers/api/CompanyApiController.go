package api

import (
	"MPMS/models"
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
	if listReq.Id != 0 {
		companyWhereMap["id"] = listReq.Id
	}
	companies, err := company.Select([]string{"id", "short_name", "creator_id", "expire_at"}, companyWhereMap)
	if err != nil {
		c.ApiReturn(structure.Response{Error: 2, Msg: fmt.Sprintf("获取数据失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
		return
	}

	for _, item := range companies {
		statusName, _ := item.GetStatusName()
		listItem := structure.StringToObjectMap{"id": item.Id, "name": item.ShortName, "status": statusName, "expire_at": item.ExpireAt}
		user, err = user.GetContactUserByCompanyId(item.Id)
		if err != nil {
			c.ApiReturn(structure.Response{Error: 3, Msg: fmt.Sprintf("获取联系人数据失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
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
			c.ApiReturn(structure.Response{Error: 4, Msg: fmt.Sprintf("获取创建人数据失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
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
