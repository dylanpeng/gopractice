package grpclibdemo

import (
	"fmt"
	"gopractice/lib/proto/protocol_demo"
)

const HelloWordClientName = "common"

var serverMap map[string]*Config
var HelloWorldClient *helloWorldClient
var ConnPool *Pool

type helloWorldClient struct {
}

func (c *helloWorldClient) GetClient() (client protocol_demo.HelloWorldClient, err error) {
	serverConf := serverMap[HelloWordClientName]
	conn, err := ConnPool.GetConnection(serverConf.GetAddress())

	if err != nil {
		fmt.Printf("GetClient GetConnection fail. | err: %s", err)
		return
	}

	client = protocol_demo.NewHelloWorldClient(conn.ClientConn)

	return
}

func init() {
	serverMap = make(map[string]*Config, 8)
	serverMap[HelloWordClientName] = &Config{
		Host: "localhost",
		Port: "69000",
	}

	ConnPool = NewPool(20, 20, 120)
	return
}
