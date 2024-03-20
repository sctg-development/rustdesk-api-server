package services

import (
	"errors"
	"rustdesk-api-server/app/models"
	"rustdesk-api-server/utils"
	"rustdesk-api-server/utils/gmd5"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
)

var Login = new(LoginService)

type LoginService struct {
}

func (s *LoginService) UserLogin(username, password, clientId, uuid string, ctx *context.Context) (token string, err error) {

	// Query whether a user exists
	var user models.User
	err = orm.NewOrm().QueryTable(new(models.User)).
		Filter("username", username).One(&user)
	if err != nil {
		return "", errors.New("The username or password is incorrect")
	}

	// Generate passwords
	pwd := User.GenPwd(password)
	// Check if the password is correct
	if user.Password != pwd {
		return "", errors.New("The username or password is incorrect")
	}
	// Determine whether a user is disabled
	if user.Status != 1 {
		return "", errors.New("The current user is disabled")
	}

	m := orm.NewOrm()
	entity := models.User{Id: user.Id}
	entity.LastLoginTime = time.Now().Unix()
	entity.LastLoginIp = ctx.Input.IP()
	entity.UpdateTime = time.Now().Unix()
	m.Update(&entity, "LastLoginTime", "LastLoginIp", "UpdateTime")

	// Generate a login token
	token2 := gmd5.EncryptNE(user.Password + clientId + uuid)

	// Return to JWT
	token, _ = utils.GenerateJwtToken(int(user.Id), user.Username, token2, clientId, uuid)

	// Save your current computer login information
	Token.Login(&user, clientId, uuid, token2)
	return token, nil
}
