package controllers

import (
	"rustdesk-api-server/utils/beegoHelper"
	"time"
)

type SysinfoController struct {
	BaseController
}

// Heartbeat detection POST
func (ctl *SysinfoController) Sysinfo() {

	ctl.JSON(beegoHelper.H{
		"modified_at": time.Now().Unix(),
	})
}
