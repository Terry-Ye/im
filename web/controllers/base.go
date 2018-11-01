package controllers

import (
	"github.com/astaxie/beego"
	"terry/meisha_account/logic/sync"
	"github.com/smallnest/rpcx/log"
	"fmt"
)

// Operations about Users
type BaseController struct {
	beego.Controller
}

type RetData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (c *BaseController) RenderData(code int, msg string, data interface{}, ctx *fasthttp.RequestCtx) {
	var retData RetData
	retData.Code = code
	retData.Msg = msg
	retData.Data = data

	// fmt.Println("log: ", retData)

	c.send(ctx, retData)
}


func (c *BaseController) send(ctx *fasthttp.RequestCtx, retData RetData) {
	//sync user info by webhook
	sync.SyncUserInfo(c.UserId, string(ctx.Path()))

	ctx.SetContentType("application/json")
	retJson, _ := util.JsonEncode(retData)
	log.Info("response data: " + retJson + ", api: " + string(ctx.Path()))
	fmt.Fprintf(ctx, retJson)
}