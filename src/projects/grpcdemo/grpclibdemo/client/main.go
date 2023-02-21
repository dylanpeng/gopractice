package main

import (
	"context"
	"fmt"
	goPool "github.com/processout/grpc-go-pool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gopractice/lib/proto/common"
	"gopractice/projects/grpcdemo/grpclibdemo"
	"sync"
	"time"
)

func main() {
	CallClient()
}

func CallClient() {
	clientConf := &grpclibdemo.ClientConfig{
		AddrMap:  make(map[string][]*grpclibdemo.Config),
		Idle:     120,
		Ttl:      120,
		Capacity: 100,
		Timeout:  3,
	}

	groupName, _ := grpclibdemo.GetGroupName(common.CommonService_CommonTest_FullMethodName)
	clientConf.AddrMap[groupName] = []*grpclibdemo.Config{{
		Host: "localhost",
		Port: "23000",
	}, {
		Host: "localhost",
		Port: "23001",
	}}

	client := grpclibdemo.NewClient(clientConf)

	a := sync.WaitGroup{}
	a.Add(10000)

	for i := 0; i < 10000; i++ {
		go func() {
			req := &common.CommonReq{Message: "hello"}
			rsp := &common.Response{}
			client.CallGrpc(common.CommonService_CommonTest_FullMethodName, req, rsp, 3*time.Second)
			a.Done()
		}()
	}

	a.Wait()
}

func CallClientDemo() {
	client, err := grpclibdemo.HelloWorldClient.GetClient()

	if err != nil {
		fmt.Printf("GetClient fail. | err: %s\n", err)
		return
	}

	req := &common.CommonReq{Message: "hello"}
	rsp := &common.Response{}

	rsp, err = client.CommonTest(context.Background(), req)

	if err != nil {
		fmt.Printf("CommonTest fail. | err: %s", err)
		return
	}

	fmt.Printf("client success. | rsp: %+v\n", *rsp)
}

func CallGoPoll() {
	pool, err := goPool.New(GoPoolFactory, 10, 10, 10*time.Minute)

	if err != nil {
		fmt.Printf("new pool fail. | err: %s\n", err)
		return
	}

	for i := 0; i < 10000; i++ {
		conn, err := pool.Get(context.Background())

		if err != nil {
			fmt.Printf("Get conn fail. | err: %s\n", err)
			return
		}

		client := common.NewCommonServiceClient(conn)

		req := &common.CommonReq{Message: "hello"}
		rsp := &common.Response{}
		rsp, err = client.CommonTest(context.Background(), req)

		if err != nil {
			fmt.Printf("CommonTest fail. | err: %s\n", err)
			return
		}

		//conn.Close()
		fmt.Printf("call success | rsp: %+v\n", *rsp)
	}
}

func GoPoolFactory() (*grpc.ClientConn, error) {
	return grpc.Dial("localhost:23000", grpc.WithTransportCredentials(insecure.NewCredentials()))
}
