package controllers

type IndexController struct {
	BaseController
}

func (c *IndexController) Get() {
	c.TplName = "tabs.html"
}