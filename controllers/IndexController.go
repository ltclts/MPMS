package controllers

type IndexController struct {
	Controller
}

func (i *IndexController) Get() {
	i.RenderHtml("测试", "html", "test/html.tpl", "test/css.tpl", "test/scripts.tpl", "")
}
