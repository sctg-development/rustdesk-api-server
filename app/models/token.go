package models

import (
	"log"

	"github.com/beego/beego/v2/client/orm"
)

type Token struct {
	Id          int32  `json:"id" orm:"auto"`
	Username    string `json:"username"`
	Uid         int32  `json:"uid"`
	ClientId    string `json:"client_id"`
	Uuid        string `json:"uuid"`
	AccessToken string `json:"access_token"`
	ActiveTime  int64  `json:"login_time"`
	LoginTime   int64  `json:"active_time"`
	ExpireTime  int64  `json:"expire_time"`
}

func (u *Token) TableName() string {
	return "rustdesk_token"
}

// Multi-field unique keys
func (u *Token) TableUnique() [][]string {
	return [][]string{
		{"uid", "client_id", "uuid"},
	}
}

func init() {
	log.Printf("Initialize the model")
	// Initialize the model
	orm.RegisterModel(new(Token))
}
