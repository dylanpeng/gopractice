package main

import (
	"context"
	"fmt"
	"gopractice/lib/proto/common"
	"gopractice/projects/grpcdemo/grpclibdemo"
)

func main() {
	client, err := grpclibdemo.HelloWorldClient.GetClient()

	if err != nil {
		fmt.Printf("GetClient fail. | err: %s\n", err)
		return
	}

	req := &common.Empty{}
	rsp := &common.Response{}

	rsp, err = client.CommonTest(context.Background(), req)

	if err != nil {
		fmt.Printf("CommonTest fail. | err: %s", err)
		return
	}

	fmt.Printf("client success. | rsp: %+v\n", *rsp)
}
