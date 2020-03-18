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
	high := [] *models.HighBad{}
	c.o.QueryTable( new(models.FloorBad).TableName() ).All(&floors)
	c.o.QueryTable( new(models.HighBad).TableName() ).All(&high)
	c.Data["floor"] = &floors
	c.Data["high"] = &high
	c.TplName = "tabs.html"
	spew.Dump(floors)
}