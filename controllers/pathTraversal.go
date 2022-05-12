package controllers

import (
	"go-sec-code/utils"
	"io/ioutil"

	beego "github.com/beego/beego/v2/server/web"
)

type PathTraversalVuln1Controller struct {
	beego.Controller
}

type PathTraversalSafe1Controller struct {
	beego.Controller
}

func (c *PathTraversalVuln1Controller) Get() {
	file := c.GetString("file")
	output, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	c.Ctx.ResponseWriter.Write(output)
}

func (c *PathTraversalSafe1Controller) Get() {
	file := c.GetString("file")
	pathTraversalFilter := utils.PathTraversalFilter{}
	if pathTraversalFilter.DoFilter(file) {
		c.Ctx.ResponseWriter.Write([]byte("evil input"))
	} else {
		output, err := ioutil.ReadFile("static/" + file)
		if err != nil {
			panic(err)
		}
		c.Ctx.ResponseWriter.Write(output)
	}
}
