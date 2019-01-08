package controllers

import (
	"MPMS/routers/uris"
	"fmt"
	"github.com/astaxie/beego"
)

/**
定义基类  共通方法可以放在这边写
*/
type Controller struct {
	beego.Controller
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

func (b *Controller) RenderHtml(title string, pageName string, tplName string, htmlCssName string, scriptsName string, sidebarName string) {
	b.TplName = tplName
	b.Data["Title"] = title
	b.Data["CurrentPageName"] = pageName
	b.Data["xsrfdata"] = b.XSRFToken()
	b.Data["ApiUriLogout"] = uris.ApiUriLogout
	b.getMenuList()
	b.Layout = "layout.tpl"
	b.LayoutSections = map[string]string{"HtmlCss": htmlCssName, "Scripts": scriptsName, "Sidebar": sidebarName}
	fmt.Println(b.Data)
}

/**
菜单生成
*/
func (b *Controller) getMenuList() {
	var menu []dashBoard
	var routes []route
	var currentPageName = b.Data["CurrentPageName"]

	/**
	财务管理菜单
	*/
	routes = []route{}
	routes = append(routes, route{"财务管理", "/finance", "financeFlow", false})
	routes = append(routes, route{"类型管理", "/finance/type", "financeType", false})
	menu = append(menu, dashBoard{"财务管理", "user", routes, false})

	/**
	系统设置
	*/
	routes = []route{}
	routes = append(routes, route{"角色管理", "", "role", false})
	routes = append(routes, route{"权限管理", "", "privilege", false})
	menu = append(menu, dashBoard{"系统设置", "cogs", routes, false})

	/**
	系统设置
	*/
	routes = []route{}
	routes = append(routes, route{"测试", "/html", "html", false})
	menu = append(menu, dashBoard{"测试", "cogs", routes, false})

	/**
	筛选出指定active的链接
	*/
	for m, v := range menu {
		for n, route := range v.Routes {
			if route.PageName == currentPageName {
				menu[m].Routes[n].Active = true
				menu[m].Active = true
				goto end
			}
		}
	}

end:
	b.Data["MenuList"] = menu
}
