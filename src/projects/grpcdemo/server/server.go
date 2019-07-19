package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gopractice/lib/proto/protocol_demo"
	"net"
)

func main(){
	lis, err := net.Listen("tcp", ":50052")
	if err != nil{
		fmt.Printf("failed to listen: %s \n", err)
		return
	}

	s := grpc.NewServer()
	protocol_demo.RegisterHelloWorldServer(s, &server{})
	fmt.Println("success")
	err = s.Serve(lis)

	if err != nil{
		fmt.Printf("failed to start grpc server: %s \n", err)
		return
	}
}

type server struct{
}

func (s *server) GetHelloWorld(ctx context.Context, req *protocol_demo.HelloWorldReq) (rsp *protocol_demo.HelloWorldRsp, err error){
	rsp = &protocol_demo.HelloWorldRsp{}
	rsp.Message = fmt.Sprintf("hello user: %d", req.Id)
	return rsp, nil
}