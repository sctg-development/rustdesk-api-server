package controllers

import "rustdesk-api-server/utils/common"

type IndexController struct {
	BaseController
}

func (ctl *IndexController) Index() {
	ctl.JSON(common.JsonResult{
		Code:  1,
		Msg:   "Welcome, author github: https://github.com/sctg-development",
		Data:  nil,
		Count: 0,
	})
}
