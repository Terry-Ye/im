package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"im/web/libs/define"
)

type DataNil struct {
}

// Operations about Users
type BaseController struct {
	beego.Controller
}

type RetData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}



func (c *BaseController) RenderData(code int, msg string, data interface{}) (retData RetData) {
	retData.Code = code
	retData.Msg = msg
	retData.Data = data
	return retData
}

func (c *BaseController) RenderDataSimple(code int, msg string) (retData RetData){
	beego.Debug("code %v", code)
	retData.Code = code
	retData.Msg = msg
	retData.Data = DataNil{}
	return

}

func (c *BaseController) CheckParams(dataModel interface{}) (code int, msg string){
	// beego.Error("dataModel Valid err : %v", dataModel)
	valid := validation.Validation{}
	b, err := valid.Valid(dataModel)
	if err != nil {
		code = define.ERR_SYSTEM_EXCEPTION_CODE
		msg = define.ERR_SYSTEM_EXCEPTION_MSG
		beego.Error("Register Valid err : %v", err)
		return
	}
	if !b {
		for _, err := range valid.Errors {
			code = define.ERR_PARAM_VAILD_CODE
			msg = err.Key+ ":" + err.Message
			return
		}
	}
	return
}

