package models

import (
	"log"

	"github.com/beego/beego/v2/client/orm"
)

type Clients struct {
	DeviceId int32  `json:"device_id" orm:"column(deviceid);auto"`
	Cpu      string `json:"cpu"`
	Uuid     int32  `json:"uuid"`
	Username string `json:"username"`
	Hostname string `json:"hostname"`
	Os       string `json:"os"`
	Version  string `json:"version"`
	Heart    string `json:"int64"`
}

func (u *Clients) TableName() string {
	return "rustdesk_clients"
}

func init() {
	log.Printf("Initialize the model")
	// Initialize the model
	orm.RegisterModel(new(Clients))
}
