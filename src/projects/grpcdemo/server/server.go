package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopractice/lib/proto/protocol_demo"
	"net"
	"runtime/debug"
	"sync"
	"time"
)

var addrs = []string{":50052", ":50053"}

func main() {
	var wg sync.WaitGroup
	for _, addr := range addrs {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			startServer(addr)
		}(addr)
	}
	wg.Wait()
}

func startServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("failed to listen: %s \n", err)
		return
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(RecoveryInterceptor, LoggerInterceptor),
	}

	s := grpc.NewServer(opts...)
	protocol_demo.RegisterHelloWorldServer(s, &server{addr: addr})
	fmt.Println("success")
	err = s.Serve(lis)

	if err != nil {
		fmt.Printf("failed to start grpc server: %s \n", err)
		return
	}
}

type server struct {
	addr string
}

func (s *server) GetHelloWorld(ctx context.Context, req *protocol_demo.HelloWorldReq) (rsp *protocol_demo.HelloWorldRsp, err error) {
	rsp = &protocol_demo.HelloWorldRsp{}
	rsp.Message = fmt.Sprintf("hello user: %d", req.Id)
	fmt.Printf("%s\n", s.addr)
	return rsp, nil
}

// 日志拦截器
func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	t := time.Now()
	fmt.Printf("gRpc begin method: method: %s | req: %v | time: %s", info.FullMethod, req, t.Format("2006-01-02 15:04:05.000000"))
	fmt.Println()
	resp, err = handler(ctx, req)
	fmt.Printf("gRpc finish method: %s | rsp: %v | time: %s | durations: %s", info.FullMethod, resp, t.Format("2006-01-02 15:04:05.000000"), time.Since(t))
	fmt.Println()
	return
}

// recover 拦截器
func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v | %s", e, string(debug.Stack()))
			fmt.Println()
		}
	}()
	fmt.Printf("RecoveryInterceptor in\n")
	resp, err = handler(ctx, req)
	fmt.Printf("RecoveryInterceptor out\n")
	return
}
