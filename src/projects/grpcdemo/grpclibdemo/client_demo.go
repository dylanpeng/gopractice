package grpclibdemo

import (
	"context"
	"errors"
	"fmt"
	"gopractice/lib/proto/common"
	"strings"
	"time"
)

const HelloWordClientName = "common"

var serverMap map[string]*Config
var HelloWorldClient = &helloWorldClient{}
var ConnPool *Pool

type helloWorldClient struct {
}

func (c *helloWorldClient) GetClient() (client common.CommonServiceClient, err error) {
	serverConf := serverMap[HelloWordClientName]
	conn, err := ConnPool.GetConnection(context.Background(), serverConf.GetAddress())

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

func CallGrpc(method string, req interface{}, rsp interface{}, timeout time.Duration) error {
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	groupName, err := GetGroupName(method)

	if err != nil {
		return err
	}

	serverConf, exist := serverMap[groupName]

	if !exist {
		return errors.New("not init grpc client")
	}

	conn, err := ConnPool.GetConnection(ctx, serverConf.GetAddress())

	if err != nil {
		fmt.Printf("GetClient GetConnection fail. | err: %s", err)
		return errors.New("GetClient GetConnection fail")
	}

	err = conn.Invoke(ctx, common.CommonService_CommonTest_FullMethodName, req, rsp)
	if err != nil {
		fmt.Printf("invoke fail.")
		return errors.New("invoke fail")
	}

	return nil
}

func GetGroupName(method string) (string, error) {
	item := strings.Split(method, "/")

	if len(item) < 3 {
		return "", errors.New("undefined grpc service")
	}

	return item[1], nil
}
