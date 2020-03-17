package controllers

import (
	"Calculator/models"
	"github.com/davecgh/go-spew/spew"
)

type IndexController struct {
	BaseController
}

//查询地台床
func (c *IndexController) Get() {
	floors := [] *models.FloorBad{}
	c.o.QueryTable( new(models.FloorBad).TableName() ).RelatedSel().All(&floors)
	c.Data["floor"] = &floors
	c.TplName = "tabs.html"

	spew.Dump(floors)
}