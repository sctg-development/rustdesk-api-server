package controllers

import (
	"rustdesk-api-server/utils/beegoHelper"
	"time"
)

type HeartController struct {
	BaseController
}

// Heartbeat detection POST
func (ctl *HeartController) Heart() {

	ctl.JSON(beegoHelper.H{
		"modified_at": time.Now().Unix(),
	})
}
