package user

import (
	"im/web/models/userModel"
	"im/web/module/util"
	"im/web/controllers"
	"encoding/json"
	"im/web/logic/userLogic"
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	controllers.BaseController
}

// type User struct {
// 	Id     string
// 	Username   string `valid:"Required;MinSize(3);MaxSize(32)"`
// 	Password    int    `valid:"Required;MinSize(6);MaxSize(20)"`
// 	CreateTime int64
// }


// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	var (
		user userModel.User
	)
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	// userInfo := userModel.GetUserInfoByUserName(user.UserName)





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
	var (
		user userModel.User
	)
	valid := validation.Validation{}
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)



	b, err := valid.Valid(&user)
	if err != nil {
		beego.Debug("valid err %v", err)
		return
	}
	if !b {
		// validation does not pass
		// blabla...
		for _, err := range valid.Errors {
			beego.Debug("err %v", err.Key)
			// log.Println(err.Key, err.Message)
		}
	}

	code, msg  := userLogic.CheckUserName(user.UserName)
	if code != 0 {
		u.Data["json"] = u.RenderDataSimple(code, msg)
		u.ServeJSON()
		return
	}
	user.Id = util.GenUuid()
	code, msg  = userModel.AddOne(user)
	u.Data["json"] = u.RenderDataSimple(code, msg)
	u.ServeJSON()
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

