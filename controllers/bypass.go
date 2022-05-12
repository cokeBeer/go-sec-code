package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type BypassController struct {
	beego.Controller
}

func (c *BypassController) Get() {
	c.Ctx.ResponseWriter.Write([]byte("bypass"))
}
