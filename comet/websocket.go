package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"time"
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
	upgrader := websocket.Upgrader{
		ReadBufferSize:  DefaultServer.Options.ReadBufferSize,
		WriteBufferSize: DefaultServer.Options.WriteBufferSize,
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Error(err)
		return
	}


	go server.writePump(conn)
	go server.readPump(conn)
}


func (s *Server) writePump(conn *websocket.Conn) {
	defer func(){
		conn.Close()
	}()
}

func (s *Server) readPump(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()

	conn.SetReadLimit(s.Options.MaxMessageSize)
	conn.SetReadDeadline(time.Now().Add(s.Options.PongWait))
	conn.SetPongHandler(func(string) error {
		conn.SetReadDeadline(time.Now().Add(s.Options.PongWait));
		return nil
	})

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway,websocket.CloseAbnormalClosure) {
				log.Errorf("readPump ReadMessage err:%v", err)
			}
		}
	}
}


// func (server *Server) run() {
// 	for{
// 		select {
// 		// case server.Buckets
// 		}
// 	}
// }


