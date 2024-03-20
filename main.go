package main

import (
	"log"
	_ "rustdesk-api-server/boot/config"
	_ "rustdesk-api-server/boot/mysql"
	_ "rustdesk-api-server/routers"
	"rustdesk-api-server/utils/flogs"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/server/web"
)

func main() {

	log.Println("RustDesk Api Server")

	flogs.InitLogger()
	//Set the request content to a replica
	web.BConfig.CopyRequestBody = true
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	err := config.InitGlobalInstance("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	web.Run()
}
