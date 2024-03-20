package controllers

import "rustdesk-api-server/utils/beegoHelper"

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.JSON(beegoHelper.H{
		"error": "Page not found",
	})
	//c.Ctx.WriteString("customize page not found")
	//c.Data= "page not found"
	//c.TplName = "404.tpl"
}

func (c *ErrorController) Error501() {
	c.JSON(beegoHelper.H{
		"error": "An error is reported on the server",
	})
	//c.Ctx.WriteString("customize server error")
	//c.Data["content"] = "server error"
	//c.TplName = "501.tpl"
}
