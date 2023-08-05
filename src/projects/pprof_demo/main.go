package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var conn *amqp.Connection

func main() {
	go getConn()

	log.Println(http.ListenAndServe(":6060", nil))
}

func getConn() *amqp.Connection {
	// 创建链接
	//url := fmt.Sprintf("%s://%s:%s@%s:%s%s", PROTOCOL, LOGIN, PASSWORD, HOST, PORT, VIRTUALHOST)
	var err error
	url := "amqp://guest:guest@www.jsifeise.ksdjfoiej/toutiao"
	for {
		conn, err = amqp.DialConfig(url, amqp.Config{Heartbeat: 60 * time.Second})
		if err != nil {
			fmt.Printf("Failed to open Connection: %s | time: %+v \n", err.Error(), time.Now())
			// 5秒重连
			time.Sleep(1 * time.Second)
			continue
		}
		return conn
	}
}
