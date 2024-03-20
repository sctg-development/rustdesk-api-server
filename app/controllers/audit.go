package controllers

import (
	"rustdesk-api-server/app/dto"
	"rustdesk-api-server/app/services"
	"rustdesk-api-server/utils/beegoHelper"
	"time"
)

var Audit = new(AuditController)

type AuditController struct {
	BaseController
}

// Operational feedback
func (ctl *AuditController) AuditConn() {
	//  {"action":"close","conn_id":129,"id":"1089363550","session_id":9166591467229392641,"uuid":"M0Y4MkI3N0MtMDMwMy01N0EwLTg5MzAtNDcwNUI4NUNFNUZD"}
	//  {"action":"new","conn_id":129,"id":"1089363550","ip":"10.10.102.105","session_id":0,"uuid":"M0Y4MkI3N0MtMDMwMy01N0EwLTg5MzAtNDcwNUI4NUNFNUZD"}
	//  {"conn_id":129,"id":"1089363550","peer":["1089363550","xxxx"],"session_id":9166591467229392641,"type":0,"uuid":"M0Y4MkI3N0MtMDMwMy01N0EwLTg5MzAtNDcwNUI4NUNFNUZD"}
}

// Heartbeat detection POST
func (ctl *AuditController) Audit() {
	req := dto.AuditReq{}

	if err := ctl.BindJSON(&req); err != nil {
		ctl.JSON(beegoHelper.H{
			"error": "The request parameter is abnormal",
		})
		return
	}

	// Set the current user's online information
	tokenInfo := services.Token.FindToken(req.Id, req.Id1, req.Uuid)
	if tokenInfo != nil {
		// Change the active time of the token
		tokenInfo.ActiveTime = time.Now().Unix()
		if !services.Token.Save(tokenInfo) {
			ctl.JSON(beegoHelper.H{
				"error": "Save login heartbeat error",
			})
		}

		ctl.JSON(beegoHelper.H{
			"data": "online",
		})
	} else {
		ctl.JSON(beegoHelper.H{
			"error": "The device is abnormal",
		})
	}

}
