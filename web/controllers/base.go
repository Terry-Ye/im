package controllers

import (
	"github.com/astaxie/beego"
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

