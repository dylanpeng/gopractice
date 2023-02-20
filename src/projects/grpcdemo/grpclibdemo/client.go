package grpclibdemo

import (
	"fmt"
	"gopractice/lib/proto/common"
)

const HelloWordClientName = "common"

var serverMap map[string]*Config
var HelloWorldClient = &helloWorldClient{}
var ConnPool *Pool

type helloWorldClient struct {
}

func (c *helloWorldClient) GetClient() (client common.CommonServiceClient, err error) {
	serverConf := serverMap[HelloWordClientName]
	conn, err := ConnPool.GetConnection(serverConf.GetAddress())

	if err != nil {
		fmt.Printf("GetClient GetConnection fail. | err: %s", err)
		return
	}

	client = common.NewCommonServiceClient(conn.ClientConn)

	return
}

func init() {
	serverMap = make(map[string]*Config, 8)
	serverMap[HelloWordClientName] = &Config{
		Host: "localhost",
		Port: "23000",
	}

	ConnPool = NewPool(20, 20, 120)
	return
}
