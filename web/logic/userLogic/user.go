package userLogic

import (
	"im/web/models/userModel"
	"im/web/libs/define"
	"im/web/module/util"
	"im/web/module/redis"
	"github.com/astaxie/beego"
)

type ReturnInfo struct {
	Auth string
	UserId string
	UserName string
}
func CheckUserName(userName string) (code int, msg string){
	status := userModel.CheckoutUserNameExist(userName)
	if status ==  false {
		code = define.ERR_USER_EXIST_CODE
		msg = define.ERR_USER_EXIST_MSG
	}

	return
}

func AddOne(user userModel.User)  (code int, msg string){
	user.Id = util.GenUuid()
	user.Password = util.Md5(user.Password)
	code, msg  = userModel.AddOne(user)
	return

}

func Login(user userModel.User) (code int , msg string, RetData ReturnInfo){
	userInfo := userModel.GetUserInfoByUserName(user.UserName)
	// password err
	if util.Md5(user.Password) != userInfo.Password {
		code = define.ERR_USER_PASSWORD_CODE
		msg = define.ERR_USER_PASSWORD_MSG
		return
	}
	err := redis.InitRedis()
	if err != nil  {
		beego.Error("redis init  err: %s", err)
	}
	auth := util.GenUuid()
	err = redis.Set(auth, user, 86400)
	RetData = ReturnInfo{auth, userInfo.Id,userInfo.UserName}
	if err != nil {
		beego.Error("redis set auth err: %s", err)
	}
	return
}





