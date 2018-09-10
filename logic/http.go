package main

import (
	"net/http"
	"im/libs/proto"
)
var (
	router *proto.Router
)
func InitHttp() (err error){
	// ServrMux 本质上是一个 HTTP 请求路由器
	httpServerMux := http.NewServeMux()
	httpServerMux.HandleFunc("/api/v1/push", Push)
	return err
}


func Push(w http.ResponseWriter, r http.Request) {
	if r.Method != "POST" {
		http.Error(w,"Method Not Allowed", 405)
	}

	var (
		auth = r.URL.Query().Get("auth")
	)

	router = getRouter(auth)






}
