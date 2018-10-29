package main

// Start Commond eg: ./client 1 5000 localhost:8080
// first parameter：beginning userId
// second parameter: amount of clients
// third parameter: comet server ip

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	//mrand "math/rand"
	"net"
	"os"
	"runtime"
	"strconv"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/gorilla/websocket"
	"net/http"
	"encoding/json"
	"os/signal"
	"net/url"
)

const (
	OP_HANDSHARE        = int32(0)
	OP_HANDSHARE_REPLY  = int32(1)
	OP_HEARTBEAT        = int32(2)
	OP_HEARTBEAT_REPLY  = int32(3)
	OP_SEND_SMS         = int32(4)
	OP_SEND_SMS_REPLY   = int32(5)
	OP_DISCONNECT_REPLY = int32(6)
	OP_AUTH             = int32(7)
	OP_AUTH_REPLY       = int32(8)
	OP_TEST             = int32(254)
	OP_TEST_REPLY       = int32(255)
)

const (
	rawHeaderLen = uint16(16)
	heart        = 240 * time.Second //s
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Proto struct {
	Ver       int16           `json:"ver"`  // protocol version
	Operation string           `json:"op"`   // operation for request
	// SeqId     int32           `json:"seq"`  // sequence number chosen by client
	Body      json.RawMessage `json:"body"` // binary body bytes(json.RawMessage is []byte)
}

var (
	countDown int64
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	begin, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	num, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	go result()

	for i := begin; i < begin+num; i++ {
		go client(fmt.Sprintf("%d", i))
	}

	var exit chan bool
	<-exit
}

func result() {
	var (
		lastTimes int64
		diff      int64
		nowCount  int64
		timer     = int64(30)
	)

	for {
		nowCount = atomic.LoadInt64(&countDown)
		diff = nowCount - lastTimes
		lastTimes = nowCount
		fmt.Println(fmt.Sprintf("%s down:%d down/s:%d", time.Now().Format("2006-01-02 15:04:05"), nowCount, diff/timer))
		time.Sleep(time.Duration(timer) * time.Second)
	}
}

func client(key string) {
	for {
		startClient(key)
		time.Sleep(3 * time.Second)
	}
}




func startClient(key string) {
	//time.Sleep(time.Duration(mrand.Intn(30)) * time.Second)
	quit := make(chan bool, 1)
	defer close(quit)
	var (
		w http.ResponseWriter
		r *http.Request
	)



	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
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

	go writePump(ch)
	go readPump(ch)





}


func readPumpvb(ch *Channel) {
	defer func() {

		s.Bucket(ch.uid).delCh(ch)
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
				return
			}
		}
		if message == nil {
			return
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
			return
		}
		if uid == "" {
			log.Error("Invalid Auth ,uid empty")
			return
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
		// ch.broadcast <- message

	}
}


func writePump(ch *Channel) {
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
			log.Printf("message write :%v", message)
			w.Write(message.Body)

			// Add queued chat messages to the current websocket message.
			// n := len(ch.broadcast)
			// for i := 0; i < n; i++ {
			// 	w.Write(newline)
			// 	w.Write(<-ch.broadcast)
			// }

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



