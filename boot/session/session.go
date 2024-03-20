package session

import beego "github.com/beego/beego/v2/server/web"

// Initialize the beego session settings
func init() {
	//Session expiration time, default is 3600 seconds
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 7200

	//Session defaults to the time the client's cookie exists, which is 3600 seconds
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 7200
}
