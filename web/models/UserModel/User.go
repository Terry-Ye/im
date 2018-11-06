package userModel

import (
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/astaxie/beego"
	"im/web/libs/define"
)


const USER_DB  = "test"

func init() {

}

type User struct {
	Id       string `orm:"pk"`
	UserName string `valid:"Required;MinSize(3);MaxSize(20)"`
	Password string `valid:"Required;MinSize(6);MaxSize(20)"`
	CreateTime int64 `data:"CreateTime"`
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


func GetUserInfoByUserName(userName string) (user User) {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	user = User{UserName: userName}
	err := o.Read(&user, "UserName")
	if err == orm.ErrNoRows {
		return User{}
	}
	return user
}

func AddOne(user User) (code int, msg string){
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	user.CreateTime  = time.Now().Unix()
	_, err := o.Insert(&user)
	if err != nil  {
		code = define.ERR_MYSQL_EXCEPTION_CODE
		msg = define.ERR_MYSQL_EXCEPTION_MSG
		beego.Error("mysql insert err :%v", err)
		return
	}
	code = define.SUCCESS_CODE
	msg = define.SUCCESS_MSG
	return
}


