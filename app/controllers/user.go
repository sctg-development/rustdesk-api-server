package controllers

import (
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/services"
	"rustdesk-api-server/global"
	"rustdesk-api-server/utils/beegoHelper"
)

type UserController struct {
	BaseController
}

// Current user information
func (ctl *UserController) CurrentUser() {
	ctl.JSON(beegoHelper.H{
		"name": ctl.loginUserInfo.Username,
	})
}

// Registered Users
func (ctl *UserController) Reg() {
	req := dto.UserRegReq{}
	req.Username = ctl.GetString("username")
	req.Password = ctl.GetString("password")
	req.AuthKey = ctl.GetString("auth_key")
	if len(req.Username) < 4 || len(req.Username) > 20 {
		ctl.JSON(beegoHelper.H{
			"error": "The username is between 4 and 20 digits",
		})
	}

	if len(req.AuthKey) == 0 {
		ctl.JSON(beegoHelper.H{
			"error": "Please enter the authorization code",
		})
	}

	// Determine whether the registration key is legitimate
	if req.AuthKey != global.ConfigVar.App.AuthKey {
		ctl.JSON(beegoHelper.H{
			"error": "The authorization code is incorrect",
		})
	}

	// Go to register an account
	if services.User.Reg(req.Username, req.Password) {
		ctl.JSON(beegoHelper.H{
			"msg": "Registration is successful",
		})
	} else {
		ctl.JSON(beegoHelper.H{
			"error": "Registration failed",
		})
	}
}

// Change the user password
func (ctl *UserController) SetPwd() {
	req := dto.UserSetPwdReq{}
	req.Username = ctl.GetString("username")
	req.Password = ctl.GetString("password")
	req.AuthKey = ctl.GetString("auth_key")
	if len(req.Username) < 4 || len(req.Username) > 20 {
		ctl.JSON(beegoHelper.H{
			"error": "The username is between 4 and 20 digits",
		})
	}

	if len(req.AuthKey) == 0 {
		ctl.JSON(beegoHelper.H{
			"error": "Please enter the authorization code",
		})
	}

	// Determine whether the registration key is legitimate
	if req.AuthKey != global.ConfigVar.App.AuthKey {
		ctl.JSON(beegoHelper.H{
			"error": "The authorization code is incorrect",
		})
	}

	// Go to register an account
	if services.User.ResetPassword(req.Username, req.Password) {
		ctl.JSON(beegoHelper.H{
			"msg": "The modification was successful",
		})
	} else {
		ctl.JSON(beegoHelper.H{
			"error": "The modification failed",
		})
	}
}

// grouping
func (ctl *UserController) Users() {
	ctl.JSON(beegoHelper.H{
		"msg":   "success",
		"total": 1,
		"data": []beegoHelper.H{
			{
				"name":     "Default user",
				"email":    "ff",
				"note":     "note",
				"status":   1,
				"is_admin": true,
			},
		},
	})
}

func (ctl *UserController) Peers() {
	ctl.JSON(beegoHelper.H{
		"msg":   "success",
		"total": 1,
		"data": []beegoHelper.H{
			{
				"id": "test",
				"info": beegoHelper.H{
					"username": "",
					"os":       "", // windows
					//linux
					//macos
					//android
					"device_name": "",
				},
				"user":      "ff",
				"user_name": "Occupancy",
				"node":      "tt",
				"is_admin":  true,
			},
		},
	})
}
