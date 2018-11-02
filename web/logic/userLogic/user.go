package userLogic

import (
	"im/web/models/userModel"
	"im/web/define"
)

func CheckUserName(userName string) (code int, msg string){
	status := userModel.CheckoutUserNameExist(userName)
	if status ==  false {
		code = define.ERR_USER_EXIST_CODE
		msg = define.ERR_USER_EXIST_MSG
	}

	return
}

func Login(userName string, password string) (code int , msg string){
	if len(userName) == 0 || len(password) == 0 {
		code = define.ERR_USER_EXIST_CODE
		msg = define.ERR_USER_EXIST_MSG
	}

	return
}

