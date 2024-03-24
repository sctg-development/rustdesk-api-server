package routers

import (
	"rustdesk-api-server/app/controllers"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

// Initialize the routing service
func init() {
	// Cross-domain solutions
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// Allow access to all sources
		AllowAllOrigins: true,
		// Optional parameters "GET", "POST", "PUT", "DELETE", "OPTIONS" (*For all)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// Refers to the types of headers that are allowed
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// A list of exposed HTTP headers
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// If set, it allows the sharing of authentication credentials, such as cookies
		AllowCredentials: true,
	}))

	// Set the routing information
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/api/heartbeat", &controllers.HeartController{}, "post:Heart")
	beego.Router("/api/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/api/ab", &controllers.AddressBookController{}, "post:Update")
	// for v1.2.2 client compatiblity
	beego.Router("/api/ab", &controllers.AddressBookController{}, "get:List")
	beego.Router("/api/ab/get", &controllers.AddressBookController{}, "post:List")
	beego.Router("/api/audit", &controllers.AuditController{}, "post:Audit")
	beego.Router("/api/audit/conn", &controllers.AuditController{}, "post:AuditConn")
	beego.Router("/api/logout", &controllers.LogoutController{}, "post:Logout")
	beego.Router("/api/currentUser", &controllers.UserController{}, "post:CurrentUser")
	beego.Router("/api/reg", &controllers.UserController{}, "get:Reg")
	beego.Router("/api/set-pwd", &controllers.UserController{}, "get:SetPwd")
	beego.Router("/api/users", &controllers.UserController{}, "get:Users")
	beego.Router("/api/peers", &controllers.UserController{}, "get:Peers")
	beego.Router("/api/software/info", &controllers.SoftwareController{}, "get:GetSoftwareInfo")
	beego.Router("/api/software/client-download-link/w64", &controllers.SoftwareController{}, "get:GetClientDownloadLinkW64")
	beego.Router("/api/software/client-download-link/w32", &controllers.SoftwareController{}, "get:GetClientDownloadLinkW32")
	beego.Router("/api/software/client-download-link/osx", &controllers.SoftwareController{}, "get:GetClientDownloadLinkOSX")
	beego.Router("/api/software/client-download-link/osxarm64", &controllers.SoftwareController{}, "get:GetClientDownloadLinkOSXArm64")
	// Set up an error route
	beego.ErrorController(&controllers.ErrorController{})
}
