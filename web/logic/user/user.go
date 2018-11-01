package user

import "im/web/models/userModel"



func checkUserName(userName string) bool{
	status := userModel.CheckoutUserNameExist(userName)
	if status == false {
		return
	}
}

