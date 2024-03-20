package models

import (
	"log"

	"github.com/beego/beego/v2/client/orm"
)

type Tags struct {
	Id    int32  `json:"id"`
	Uid   int32  `json:"uid"`
	Tag   string `json:"tag"`
	Color string `json:"color,omitempty"`
}

func (u *Tags) TableName() string {
	return "rustdesk_tags"
}

func init() {
	log.Printf("Initialize the model")
	// Initialize the model
	orm.RegisterModel(new(Tags))
}
