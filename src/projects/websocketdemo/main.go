package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
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

var (
	newline    = []byte{'\n'}
	space      = []byte{' '}
	broadcast  chan []byte
	register   chan *websocket.Conn
	conns      map[*websocket.Conn]bool
	unregister chan *websocket.Conn
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	// 广播内容管道
	broadcast = make(chan []byte)
	// 注册连接管道
	register = make(chan *websocket.Conn)
	// 取消注册连接管道
	unregister = make(chan *websocket.Conn)
	// 连接池
	conns = make(map[*websocket.Conn]bool)
	router := gin.Default()
	// home页面，有ws连接javascript代码
	router.GET("/", serveHome)
	// 服务端websocket连接
	router.GET("/ws", serveWs)

	// websocket广播
	go run()

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	return
}

// 主页
func serveHome(ctx *gin.Context) {
	w, r := ctx.Writer, ctx.Request

	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

// websocket连接api
func serveWs(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Printf("create conn fail. | err: %s", err)
		return
	}
	register <- conn

	go readPump(conn)
	go writePump(conn)
}

// 监听读方法
func readPump(conn *websocket.Conn) {
	defer func() {
		unregister <- conn
		conn.Close()
	}()
	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	conn.SetPongHandler(func(string) error { conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		// 广播
		broadcast <- message
	}
}

// ping保持连接
func writePump(conn *websocket.Conn) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		conn.Close()
	}()
	for {
		select {
		case <-ticker.C:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// 广播方法
func run() {
	for {
		select {
		case conn := <-register:
			conns[conn] = true
		case conn := <-unregister:
			if _, ok := conns[conn]; ok {
				delete(conns, conn)
			}
		case message, ok := <-broadcast:
			for conn := range conns {
				conn.SetWriteDeadline(time.Now().Add(writeWait))
				if !ok {
					// The hub closed the channel.
					conn.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}

				w, err := conn.NextWriter(websocket.TextMessage)
				if err != nil {
					return
				}
				w.Write(message)

				if err := w.Close(); err != nil {
					return
				}
			}
		}
	}
}
