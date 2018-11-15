package userLogic

import (
	"im/web/models/userModel"
	"im/web/libs/define"
	"im/web/module/util"
	"im/web/module/redis"
	"github.com/astaxie/beego"
	"encoding/json"
	"time"
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
	code = define.SUCCESS_CODE
	msg = define.SUCCESS_MSG
	return
}

func CheckAuth(auth string) (code int, msg string) {
	var user userModel.User

	err := redis.InitRedis()
	if err != nil  {
		beego.Error("redis init  err: %s", err)
	}
	jsonStr := redis.Get(auth)
	// auth info %s get im_api_7d69404ea4ab8ef6: {"Id":"06324815a645d98e","UserName":"t4444422","Password":"","CreateTime":0}
	beego.Debug("auth info %v", jsonStr)

	if err := json.Unmarshal([]byte(jsonStr), &user); err != nil {
		beego.Debug("json err %s", err)
		code = define.ERR_USER_NO_EXIST_CODE
		msg = define.ERR_USER_NO_EXIST_MSG
		return
	}
	beego.Debug("auth info %v", user)
	if user.Id == "" {
		code = define.ERR_USER_NO_EXIST_CODE
		msg = define.ERR_USER_NO_EXIST_MSG
		return
	}
	code = define.SUCCESS_CODE
	msg = define.SUCCESS_MSG
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
	user.Password = ""
	user.Id = userInfo.Id
	user.UserName = userInfo.UserName

	jsonUser, err := json.Marshal(user)
	if err != nil {
		beego.Error("jsonUse  err: %s", err)
	}
	err = redis.Set(auth, jsonUser, time.Hour)
	RetData = ReturnInfo{auth, userInfo.Id,userInfo.UserName}
	if err != nil {
		beego.Error("redis set auth err: %s", err)
	}



	code = define.SUCCESS_CODE
	msg = define.SUCCESS_MSG
	return
}





