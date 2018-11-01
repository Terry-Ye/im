package models

import (
	"github.com/astaxie/beego/orm"
	"im/web/module/config"

	"github.com/astaxie/beego"
	"strings"
	"im/web/models/UserModel"
	_ "github.com/go-sql-driver/mysql"
)

var confDbStr string

func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	config, err := config.Reader("database.conf")
	if err != nil  {
		beego.Error("config reader err: %v", err)
	}

	host := config.String("mysql::host")

	port := config.String("mysql::port")
	userStr := config.String("mysql::username")
	password := config.String("mysql::password")

	confDbStr = userStr+":"+ password +"@(" + host + ":" + port + ")/{{DB_NAME}}?charset=utf8"


	orm.RegisterDataBase("default", "mysql", getDbStr(UserModel.USER_DB), 30)
	orm.RegisterModelWithPrefix("tb_", new(UserModel.User))
	// orm.RegisterModel(new(UserModel.User))

}



func getDbStr(dbname string) string {
	return strings.Replace(confDbStr, "{{DB_NAME}}", dbname, -1)
}