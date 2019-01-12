package api

import (
	"MPMS/models"
	"MPMS/structure"
	"fmt"
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

	var list []structure.StringToObjectMap
	company := models.Company{}
	user := models.User{}
	var users []models.User
	relation := models.Relation{}
	var relations []models.Relation
	program := models.MiniProgram{}

	companies, err := company.Select([]string{"id", "short_name", "creator_id", "expire_at"}, structure.StringToObjectMap{"is_deleted": models.UnDeleted})
	if err != nil {
		u.ApiReturn(structure.Response{Error: 2, Msg: fmt.Sprintf("获取数据失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
		return
	}

	for _, item := range companies {
		statusName, _ := item.GetStatusName()
		listItem := structure.StringToObjectMap{"id": item.Id, "name": item.ShortName, "status": statusName, "expire_at": item.ExpireAt}
		relations, err = relation.Select([]string{"refer_id_others"}, structure.StringToObjectMap{
			"is_deleted": models.UnDeleted,
			"refer_type": models.RelationReferTypeCompanyContactUser,
			"refer_id":   item.Id,
		})
		if err != nil {
			u.ApiReturn(structure.Response{Error: 3, Msg: fmt.Sprintf("获取联系人数据失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
			return
		}
		if len(relations) >= 1 {
			//联系人
			users, err = user.Select([]string{"name", "phone"}, structure.StringToObjectMap{
				"is_deleted": models.UnDeleted,
				"id":         relations[0].ReferIdOthers,
			})
			if err != nil {
				u.ApiReturn(structure.Response{Error: 4, Msg: fmt.Sprintf("获取联系人数据失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
				return
			}
			if len(users) >= 1 {
				listItem["company_contact_user"] = users[0].Name
				listItem["company_contact_user_phone"] = users[0].Phone
			}
		}

		//创建人
		users, err = user.Select([]string{"name", "phone"}, structure.StringToObjectMap{
			"is_deleted": models.UnDeleted,
			"id":         item.CreatorId,
		})
		if err != nil {
			u.ApiReturn(structure.Response{Error: 4, Msg: fmt.Sprintf("获取创建人数据失败：%s", err.Error()), Info: structure.StringToObjectMap{}})
			return
		}
		if len(users) >= 1 {
			listItem["creator"] = users[0].Name
		}

		//小程序个数获取
		count, _ := program.Count(structure.StringToObjectMap{
			"is_deleted": models.UnDeleted,
			"company_id": item.Id,
		})
		listItem["mp_count"] = count
		list = append(list, listItem)
	}
	u.ApiReturn(structure.Response{Error: 0, Msg: "ok", Info: structure.StringToObjectMap{"list": list}})
}
