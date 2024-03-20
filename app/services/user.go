package services

import (
	"rustdesk-api-server/app/models"
	"rustdesk-api-server/global"
	"rustdesk-api-server/utils/gmd5"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

var User = new(UserService)

type UserService struct {
}

func (u *UserService) Reg(username, password string) bool {
	// Generate passwords
	hashPwd := u.GenPwd(password)

	// Insert or modify
	m := &models.User{
		Username:   username,
		Password:   hashPwd,
		Status:     1,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
	_, err := orm.NewOrm().Insert(m)
	if err != nil {
		return false
	}

	return true
}

// Generate a saved password
func (u *UserService) GenPwd(password string) string {
	// Verify that the password is correct
	pwd, err := gmd5.Encrypt(password + global.ConfigVar.App.CryptKey)
	if err != nil {
		panic("md5 encrypt Err" + err.Error())
	}

	return pwd
}

// Reset your password
func (u *UserService) ResetPassword(username string, password string) bool {
	// Generate passwords
	hashPwd := u.GenPwd(password)

	m := User.FindByUserName(username)
	if m == nil {
		return false
	}
	m.Password = hashPwd
	_, err := orm.NewOrm().Update(m, "password")
	if err != nil {
		return false
	}

	return true
}

// Query user information based on user name
func (u *UserService) FindByUserName(username string) *models.User {
	ret := models.User{}
	err := orm.NewOrm().QueryTable(new(models.User)).Filter("username", username).One(&ret)
	if err != nil {
		return nil
	}
	return &ret
}

func (u *UserService) Logout(info *models.User, clientId string) bool {
	// Delete the login token
	token := &models.Token{}

	_, err := orm.NewOrm().Raw("delete from "+token.TableName()+" where uid = ? and client_id = ?", info.Id, clientId).Exec()
	if err != nil {
		return false
	}
	return true
}
