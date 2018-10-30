package controllers

import (
	"im/web/models/UserModel"
	"im/web/module/util"
	"encoding/json"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

type User struct {
	Username   string `valid:"Required;MinSize(3);MaxSize(32)"`
	Password    int    `valid:"Required;MinSize(6);MaxSize(20)"`
}



// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	if UserModel.Login(username, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

// @Title Register
// @Description Register user
// @Param	username		formData 	string	true		"The username for register"
// @Param	password		formData 	string	true		"The password for register"
// @Success 200 {string} register success
// @Failure 403 user not exist
// @router /register [post]
func (u *UserController) Register() {

	var user UserModel.User

	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	user.Id = util.GenUuid()

	UserModel.AddOne(user)
	return
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {

}


// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (u *UserController) GetAll() {

}

