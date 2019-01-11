package controllers

import (
	"MPMS/models"
	"MPMS/routers/uris"
	"MPMS/session"
	"MPMS/structure"
	"fmt"
	"github.com/astaxie/beego"
)

/**
定义基类  共通方法可以放在这边写
*/
type Controller struct {
	beego.Controller
}

func (b *Controller) RenderHtml(title string, pageName string, tplName string, htmlCssName string, scriptsName string, sidebarName string) {
	b.TplName = tplName
	b.Data["Title"] = title
	b.Data["CurrentPageName"] = pageName
	b.Data["xsrfdata"] = b.XSRFToken()
	b.Data["ApiUriLogout"] = uris.ApiUriLogout
	b.Data["CompanyName"] = "两分钱"
	b.Data["LoginUserName"] = b.GetSession(session.UserName)
	b.getMenuList()
	b.Layout = "layout.tpl"
	b.LayoutSections = map[string]string{"HtmlCss": htmlCssName, "Scripts": scriptsName, "Sidebar": sidebarName}
	fmt.Println(b.Data)
}

type dashBoard struct {
	//dataType string  //列表类型：单个；下拉菜单
	Name   string  `json:"name"`
	Icon   string  `json:"icon"`
	Routes []route `json:"routes"`
	Active bool    `json:"active"`
}

type route struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	PageName string `json:"page_name"`
	Active   bool   `json:"active"`
}

/**
菜单生成
*/
func (b *Controller) getMenuList() {
	var menu []dashBoard
	var routes []route
	var currentPageName = b.Data["CurrentPageName"]

	m := models.Menu{}
	active := false
	if menuList, e := m.Select([]string{}, structure.StringToObjectMap{"is_deleted": models.UnDeleted, "type": models.MenuTypeFirst}); e == nil {
		for _, menuItem := range menuList {
			if menuItemList, e := m.Select([]string{"name", "name_en", "uri"}, structure.StringToObjectMap{"is_deleted": models.UnDeleted, "type": models.MenuTypeSecond, "parent_id": menuItem.Id}); e == nil {
				routes = []route{}
				active = false
				for _, routeItem := range menuItemList {
					if !active {
						if active = routeItem.NameEn == currentPageName; active {
							b.Data["MenuFirstName"] = menuItem.Name
							b.Data["MenuSecondName"] = routeItem.Name
						}
					}
					routes = append(routes, route{routeItem.Name, routeItem.Uri, routeItem.NameEn, routeItem.NameEn == currentPageName})
				}
				menu = append(menu, dashBoard{menuItem.Name, menuItem.Uri, routes, active})
			}
		}
	}

	b.Data["MenuList"] = menu
}
