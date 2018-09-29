package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"time"
	"im/libs/proto"
	"encoding/json"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func InitWebsocket(bind string) (err error) {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(DefaultServer, w, r)
	})

	err = http.ListenAndServe(bind, nil)
	return err
}

// serveWs handles websocket requests from the peer.
func serveWs(server *Server, w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  DefaultServer.Options.ReadBufferSize,
		WriteBufferSize: DefaultServer.Options.WriteBufferSize,
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Error(err)
		return
	}
	var ch *Channel
	// 写入配置
	ch = NewChannel(server.Options.BroadcastSize)
	ch.conn = conn

	go server.writePump(ch)
	go server.readPump(ch)
}

func (s *Server) readPump(ch *Channel) {
	defer func() {
		ch.conn.Close()
	}()

	ch.conn.SetReadLimit(s.Options.MaxMessageSize)
	ch.conn.SetReadDeadline(time.Now().Add(s.Options.PongWait))
	ch.conn.SetPongHandler(func(string) error {
		ch.conn.SetReadDeadline(time.Now().Add(s.Options.PongWait));
		return nil
	})

	for {
		_, message, err := ch.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Errorf("readPump ReadMessage err:%v", err)
			}
		}
		var(
			connArg *proto.ConnArg

		)
		log.Infof("message :%s", message)
		if err := json.Unmarshal([]byte(message), &connArg); err != nil {
			log.Errorf("message struct %b", connArg)
		}
		uid, err := s.operator.Connect(connArg)
		log.Infof("websocket uid:%s", uid)
		if err != nil {
			log.Errorf("s.operator.Connect error %s", err)
		}


		b := s.Bucket(uid)

		// rpc 操作获取uid 存入ch 存入Server 未写

		// b.broadcast <- message
		log.Infof("connArg roomId : %d", connArg.RoomId)
		err = b.Put(uid, connArg.RoomId, ch)
		if err != nil {
			log.Errorf("conn close err: %s", err)
			ch.conn.Close()
		}
		log.Infof("message  333 :%s", message)
		ch.broadcast <- message

	}
}

func (s *Server) writePump(ch *Channel) {
	ticker := time.NewTicker(s.Options.PingPeriod)
	log.Printf("ticker :%v", ticker)

	defer func() {
		ticker.Stop()
		ch.conn.Close()
	}()
	for {
		select {
		case message, ok := <-ch.broadcast:
			ch.conn.SetWriteDeadline(time.Now().Add(s.Options.WriteWait))
			if !ok {
				// The hub closed the channel.
				ch.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			log.Printf("TextMessage :%v", websocket.TextMessage)
			w, err := ch.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			log.Printf("message :%v", message)
			w.Write(message)
			// Add queued chat messages to the current websocket message.
			n := len(ch.broadcast)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-ch.broadcast)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			ch.conn.SetWriteDeadline(time.Now().Add(s.Options.WriteWait))
			log.Printf("websocket.PingMessage :%v", websocket.PingMessage)
			if err := ch.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
