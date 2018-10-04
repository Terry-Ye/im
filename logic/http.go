package main

import (
	"net/http"
	"im/libs/proto"
	inet "im/libs/net"
	log "github.com/sirupsen/logrus"
	"net"
	"io/ioutil"
)

var (
	router *proto.Router
)

func InitHTTP() (err error) {
	// ServrMux 本质上是一个 HTTP 请求路由器
	var network, addr string

	for i := 0; i < len(Conf.Base.HttpAddrs); i++ {

		httpServeMux := http.NewServeMux()
		httpServeMux.HandleFunc("/api/v1/push", Push)

		if network, addr, err = inet.ParseNetwork(Conf.Base.HttpAddrs[i]); err!=nil {
			log.Errorf("inet.ParseNetwork() error(%v)", err)
			return
		}

		log.Infof("start http listen:\"%s\"", Conf.Base.HttpAddrs[i])

		go httpListen(httpServeMux, network, addr)
		select {

		}

	}
	return
}
func Push(w http.ResponseWriter, r *http.Request) {
	// log.Info("yes")
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
	}

	var (
		auth      = r.URL.Query().Get("auth")
		err       error
		bodyBytes []byte
		body	string
	)

	if router, err = getRouter(auth); err != nil {

		log.Errorf("get router error : %s", err)
		return
	}

	log.Infof("router info %v", router)

	if bodyBytes, err = ioutil.ReadAll(r.Body); err != nil {
		log.Errorf("get router error : %s", err)
	}
	r.Body.Close()
	body = string(bodyBytes)

	log.Infof("get bodyBytes : %s", body)

	if err := RedisPublishCh(router.ServerId, router.UserId, bodyBytes); err != nil {
		log.Errorf("redis Publish err: %s", err)
	}



}

func httpListen(mux *http.ServeMux, network, addr string) {

	httpServer := &http.Server{Handler: mux, ReadTimeout: Conf.Base.HTTPReadTimeout, WriteTimeout: Conf.Base.HTTPWriteTimeout}
	httpServer.SetKeepAlivesEnabled(true)

	l, err := net.Listen(network, addr)
	if err != nil {
		log.Errorf("net.Listen(\"%s\", \"%s\") error(%v)", network, addr, err)
		panic(err)
	}
	if err := httpServer.Serve(l); err != nil {
		log.Errorf("server.Serve() error(%v)", err)
		panic(err)
	}
}
