package userModel

import (
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego"
)


const USER_DB  = "test"

func init() {

}

type User struct {
	Id       string `orm:"pk"`
	UserName string
	Password string
	CreateTime int64

}

func Login(username, password string) bool {

	return false
}

func GetOne(userId string) (User *User, err error)  {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	return
}


func CheckoutUserNameExist(userName string) bool {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	user := User{UserName: userName}
	err := o.Read(&user, "UserName")
	if err == orm.ErrNoRows {
		return true
	}
	return false
}


func AddOne(user User) (aa int64){
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	user.CreateTime  = time.Now().Unix()
	beego.Debug("user %v", user)
	aa, err := o.Insert(&user)
	if err != nil  {
		beego.Error("insert err :%v", err)
	}
	beego.Debug("insert %v", aa)
	return aa


}


