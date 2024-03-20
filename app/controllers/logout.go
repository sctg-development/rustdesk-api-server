package controllers

import (
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/services"
	"rustdesk-api-server/utils/beegoHelper"
)

type LogoutController struct {
	BaseController
}

// Sign out
func (ctl *LogoutController) Logout() {
	req := dto.LogoutReq{}
	if err := ctl.BindJSON(&req); err != nil {
		ctl.JSON(beegoHelper.H{
			"error": err.Error(),
		})
	}

	if services.User.Logout(ctl.loginUserInfo, req.Id) {
		ctl.JSON(beegoHelper.H{"data": "The exit is complete"})
	} else {
		ctl.JSON(beegoHelper.H{"error": "Exit failed"})
	}
}
