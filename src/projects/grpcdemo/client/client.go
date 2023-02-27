package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"gopractice/lib/proto/protocol_demo"
	"runtime/debug"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithChainUnaryInterceptor(LoggerInterceptor, RecoveryInterceptor))

	if err != nil {
		fmt.Printf("new grpc client failed: %s \n", err)
		return
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
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

func LoggerInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	t := time.Now()
	fmt.Printf("gRpc begin method: method: %s | req: %v | time: %s", method, req, t.Format("2006-01-02 15:04:05.000000"))
	fmt.Println()
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("gRpc finish method: %s | rsp: %v | time: %s | durations: %s", method, reply, t.Format("2006-01-02 15:04:05.000000"), time.Since(t))
	fmt.Println()
	return err
}

func RecoveryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v | %s", e, string(debug.Stack()))
			fmt.Println()
		}
	}()
	fmt.Printf("RecoveryInterceptor in\n")
	err = invoker(ctx, method, req, reply, cc, opts...)
	fmt.Printf("RecoveryInterceptor out\n")
	return
}
