package main

import (
	"net/http"
	"im/libs/proto"
	inet "im/libs/net"
	log "github.com/sirupsen/logrus"
	"net"
	"io/ioutil"
	"strconv"
	"time"
	"encoding/json"
	"im/libs/define"
)

type retData struct {
	Code int
	Msg string
}



func InitHTTP() (err error) {
	// ServrMux 本质上是一个 HTTP 请求路由器
	var network, addr string

	for i := 0; i < len(Conf.Base.HttpAddrs); i++ {

		httpServeMux := http.NewServeMux()
		httpServeMux.HandleFunc("/api/v1/push", Push)
		httpServeMux.HandleFunc("/api/v1/pushRoom", PushRoom)

		if network, addr, err = inet.ParseNetwork(Conf.Base.HttpAddrs[i]); err != nil {
			log.Errorf("inet.ParseNetwork() error(%v)", err)
			return
		}

		log.Infof("start http listen:\"%s\"", Conf.Base.HttpAddrs[i])

		go httpListen(httpServeMux, network, addr)
		select {}

	}
	return
}

func PushRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
	}
	var (
		auth = r.URL.Query().Get("auth")
		err       error
		bodyBytes []byte
		body      string
		sendData *proto.Send
		formUserInfo *proto.Router
	)

	// get roomId
	rid, err := strconv.ParseInt(r.URL.Query().Get("rid"), 10, 32)
	if err != nil {
		log.Errorf("rid invalid : %s", rid)
	}
	if formUserInfo, err = getRouter(auth); err != nil {
		log.Errorf("get router error : %s", err)
		return
	}

	if bodyBytes, err = ioutil.ReadAll(r.Body); err != nil {
		log.Errorf("get router error : %s", err)
	}

	defer r.Body.Close()
	body = string(bodyBytes)
	log.Infof("PushRoom get bodyBytes : %s", body)

	json.Unmarshal(bodyBytes, &sendData)
	sendData.FormUserName = formUserInfo.UserName
	sendData.FormUserId = formUserInfo.UserId
	if bodyBytes, err = json.Marshal(sendData); err != nil {
		log.Errorf("redis Publish room err: %s", err)
	}



	if err := RedisPublishRoom(int32(rid), bodyBytes); err != nil {
		log.Errorf("redis Publish room err: %s", err)
	}
}
func Push(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method Not Allowed", 405)
	}

	var (
		auth      = r.URL.Query().Get("auth")
		acceptUserId     = r.URL.Query().Get("userId")
		err       error
		bodyBytes []byte
		body      string
		formUserInfo *proto.Router
		res       = map[string]interface{}{"code": define.SUCCESS_REPLY, "msg":define.SUCCESS_REPLY_MSG}
		sendData *proto.Send

	)


	if formUserInfo, err = getRouter(auth); err != nil {
		log.Errorf("get router error : %s", err)
		return
	}

	log.Infof("push round userId %s", formUserInfo.UserId)


	if bodyBytes, err = ioutil.ReadAll(r.Body); err != nil {
		res["code"] = define.NETWORK_ERR
		res["msg"] = define.NETWORK_ERR_MSG
		log.Errorf("get router error : %s", err)
	}

	serverId := RedisCli.Get(getKey(acceptUserId)).Val()
	sid, err := strconv.ParseInt(serverId, 10, 8)
	if err != nil {
		res["code"] = define.SEND_ERR
		res["msg"] = define.SEND_ERR_MSG
		log.Infof("router err %v", err)
		return
	}

	defer retPWrite(w, r, res, &body, time.Now())

	json.Unmarshal(bodyBytes, &sendData)
	sendData.FormUserName = formUserInfo.UserName
	sendData.FormUserId = formUserInfo.UserId
	if bodyBytes, err = json.Marshal(sendData); err != nil {
		log.Errorf("redis Publish err: %s", err)
	}
	body = string(bodyBytes)
	// log.Infof("get bodyBytes : %s", body)
	// log.Infof("get sendUserId : %s", acceptUserId)
	// log.Infof("get userId : %s", formUserInfo.UserId)
	// log.Infof("get sid : %d", sid)


	if err := RedisPublishCh(int8(sid), acceptUserId, bodyBytes); err != nil {
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

func retPWrite(w http.ResponseWriter, r *http.Request, res map[string]interface{}, body *string, start time.Time) {
	data, err := json.Marshal(res)
	if err != nil {
		log.Errorf("json.Marshal(\"%v\") error(%v)", res, err)
		return
	}
	dataStr := string(data)
	if _, err := w.Write([]byte(dataStr)); err != nil {
		log.Errorf("w.Write(\"%s\") error(%v)", dataStr, err)
	}

	log.Infof("req: \"%s\", post: \"%s\", res:\"%s\", ip:\"%s\", time:\"%fs\"", r.URL.String(), *body, dataStr, r.RemoteAddr, time.Now().Sub(start).Seconds())
}
