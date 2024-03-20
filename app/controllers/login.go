package controllers

import (
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/services"
	"rustdesk-api-server/utils/beegoHelper"
	"rustdesk-api-server/utils/common"
	"strings"
)

var Login = new(LoginController)

type LoginController struct {
	BaseController
}

// login
func (ctl *LoginController) Login() {
	if ctl.Ctx.Input.IsPost() {
		// Get the request parameters
		var req dto.LoginReq
		if err := ctl.BindJSON(&req); err != nil {
			ctl.JSON(common.JsonResult{
				Error: err.Error(),
			})
		}
		req.Username = strings.TrimSpace(req.Username)
		if len(req.Username) == 0 {
			ctl.JSON(common.JsonResult{
				Code:  -1,
				Error: "The username cannot be empty",
			})
		}
		req.Password = strings.TrimSpace(req.Password)
		if len(req.Password) == 0 {
			ctl.JSON(common.JsonResult{
				Code:  -1,
				Error: "The password cannot be empty",
			})
		}
		req.Id = strings.TrimSpace(req.Id)
		if len(req.Id) == 0 {
			ctl.JSON(common.JsonResult{
				Code:  -1,
				Error: "The client ID cannot be empty",
			})
		}

		// Check whether the account and password in the database are valid
		token, err := services.Login.UserLogin(req.Username, req.Password, req.Id, req.Uuid, ctl.Ctx)
		if err != nil {
			//     return json({"type": "access_token","access_token":token,"user":{"name":username,"email":res['email'],"note":res['note'],"status":res['status'],"grp":res['group'],"is_admin":True if res['is_admin']==1 else False }})
			ctl.JSON(common.JsonResult{
				//Code:  -1,
				Error: err.Error(),
			})
		}

		ctl.JSON(beegoHelper.H{
			"type":         "access_token",
			"access_token": token,
			"user": beegoHelper.H{
				"name": req.Username,
			},
		})

	}

}
