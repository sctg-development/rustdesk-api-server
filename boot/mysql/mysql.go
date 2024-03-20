package mysql

import (
	"fmt"
	"rustdesk-api-server/global"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// Register the MySQL driver
func init() {
	orm.Debug = true

	if global.ConfigVar.DBType == "mysql" {
		logs.Info("Database registration type MySQL")

		err := orm.RegisterDriver("mysql", orm.DRMySQL)
		if err != nil {
			logs.Error("MySQL Driver Registration Failed:", err)
		}

		// Format the connector
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
			global.ConfigVar.Mysql.Username,
			global.ConfigVar.Mysql.Password,
			global.ConfigVar.Mysql.Host,
			global.ConfigVar.Mysql.Port,
			global.ConfigVar.Mysql.Database,
		)

		// Sign up for a linked database
		err = orm.RegisterDataBase("default", "mysql", connStr)
		if err != nil {
			logs.Error("MySQL database registration failed", err)
		}

	} else {
		logs.Info("Database registration type sqlite3")

		err := orm.RegisterDriver("sqlite", orm.DRSqlite)
		if err != nil {
			logs.Error("sqlite3 registration driver failed:", err)
		}

		err = orm.RegisterDataBase("default", "sqlite3", "sqlite3.db")
		if err != nil {
			logs.Error("sqlite3 database registration failed", err)
		}
	}

}
