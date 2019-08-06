package controller

import (
	"com.phh/generator/service"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type IndexController struct {
	Ctx              iris.Context
	GeneratorService service.GeneratorService
}

//localhost:8080/
func (c *IndexController) Get() mvc.Result {
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "首页",
			"navIndex": "layui-this",
		},
	}
}

//localhost:8080/about
func (c *IndexController) GetAbout() mvc.Result {
	return mvc.View{
		Name: "about.html",
		Data: iris.Map{
			"Title":    "关于",
			"navAbout": "layui-this",
		},
	}
}

//获取表列表
func (c *IndexController) GetTables() {
	name := c.Ctx.URLParam("name")
	list := c.GeneratorService.QueryTableList(name)
	var result = make(map[string]interface{})
	result["code"] = "0"
	result["msg"] = ""
	result["count"] = len(list)
	result["data"] = list
	c.Ctx.JSON(result)
}
