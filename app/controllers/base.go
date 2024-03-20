package controllers

import (
	"log"
	"rustdesk-api-server/utils/beegoHelper"

	"github.com/beego/beego/v2/client/orm"

	//beego "github.com/beego/beego/v2/adapter"
	"rustdesk-api-server/app/models"
	"rustdesk-api-server/utils"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
	controllerName string
	actionName     string
	loginUserInfo  *models.User
}

func (ctl *BaseController) Prepare() {
	controllerName, actionName := ctl.GetControllerAndAction()
	ctl.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	ctl.actionName = strings.ToLower(actionName)
	log.Println("Request interface", ctl.Ctx.Input.URL(), ctl.Ctx.Input.Method(), string(ctl.Ctx.Input.RequestBody))
	// Get a token
	token := ctl.Ctx.Input.Header("Authorization")
	if ctl.controllerName != "login" && ctl.controllerName != "index" && !(ctl.controllerName == "user" && (ctl.actionName == "reg" || ctl.actionName == "setpwd")) {
		if token == "" {
			ctl.JSON(beegoHelper.H{
				"error": "User authorization verification failed",
			})
		} else {
			// Verify the user login
			if !ctl.CheckLogin() {
				ctl.JSON(beegoHelper.H{
					"error": "The user authorization information is incorrect",
				})
			}
		}

	}
}

type JsonResult struct {
	Code  int         `json:"code"`  // Response Code: 0 Success 401 Please Login 403 No Permission 500 Error
	Msg   string      `json:"msg"`   // Message prompts
	Data  interface{} `json:"data"`  // Data objects
	Count int64       `json:"count"` // Total number of records
}

func (this *BaseController) JSON(obj interface{}) {
	this.Data["json"] = obj
	//Serialize the output of JSON
	this.ServeJSON()
	this.StopRun()
}

// Verify the login information
func (ctl *BaseController) CheckLogin() bool {
	token := ctl.Ctx.Input.Header("Authorization")
	token = strings.ReplaceAll(token, "Bearer ", "")
	// Decrypt the token
	parseToken, err := utils.ParseToken(token)
	if err != nil {
		return false
	}

	// Find the user's login information
	var loginTokenInfo models.Token
	err = orm.NewOrm().QueryTable(new(models.Token)).
		Filter("uid", parseToken.UserId).
		Filter("client_id", parseToken.ClientId).
		Filter("access_token", parseToken.AccessToken).
		One(&loginTokenInfo)

	if err != nil {
		return false
	}

	var loginInfo models.User
	err = orm.NewOrm().QueryTable(new(models.User)).
		Filter("id", loginTokenInfo.Uid).
		One(&loginInfo)

	if err != nil {
		return false
	}
	// Determine whether a user is disabled
	ctl.loginUserInfo = &loginInfo
	if ctl.loginUserInfo.Status != 1 {
		ctl.JSON(beegoHelper.H{
			"error": "The user has been disabled",
		})
		return false
	}

	return true
}
