package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gopractice/lib/proto/protocol_demo"
)

func main(){
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil{
		fmt.Printf("new grpc client failed: %s \n", err)
		return
	}

	defer conn.Close()

	c := protocol_demo.NewHelloWorldClient(conn)

	r, err := c.GetHelloWorld(context.Background(), &protocol_demo.HelloWorldReq{Id:100})

	if err != nil{
		fmt.Printf("request GetHelloWorld faild: %s \n", err)
		return
	}

	fmt.Printf("success, message is : %s \n", r.Message)

}
