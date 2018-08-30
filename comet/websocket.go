package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func InitWebsocket(bind string) (err error) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(DefaultServer, w, r)
	})
	err = http.ListenAndServe(bind, nil)
	return err

}

// serveWs handles websocket requests from the peer.
func serveWs(server *Server, w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {

		return
	}

}
