package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"gopractice/lib/proto/protocol_demo"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		fmt.Printf("new grpc client failed: %s \n", err)
		return
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()

	status := defaultReadyCheck(ctx, conn)
	if status == connectivity.Ready {
		fmt.Printf("client beready \n")
	} else {
		fmt.Printf("client not beready \n")
		return
	}

	c := protocol_demo.NewHelloWorldClient(conn)

	r, err := c.GetHelloWorld(context.Background(), &protocol_demo.HelloWorldReq{Id: 100})

	if err != nil {
		fmt.Printf("request GetHelloWorld faild: %s \n", err)
		return
	}

	fmt.Printf("success, message is : %s \n", r.Message)

}

func defaultReadyCheck(ctx context.Context, conn *grpc.ClientConn) connectivity.State {
	for {
		s := conn.GetState()

		if s == connectivity.Ready || s == connectivity.Shutdown {
			return s
		}

		if !conn.WaitForStateChange(ctx, s) {
			return connectivity.Idle
		}
	}
}
