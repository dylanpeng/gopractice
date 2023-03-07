package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"gopractice/lib/proto/protocol_demo"
	"log"
	"net"
	"testing"
)

func dialer() func(ctx context.Context, str string) (net.Conn, error) {
	listner := bufconn.Listen(1024 * 1024)

	serv := grpc.NewServer()

	protocol_demo.RegisterHelloWorldServer(serv, &server{})

	go func() {
		if err := serv.Serve(listner); err != nil {
			log.Fatal(err)
		}
	}()

	return func(ctx context.Context, str string) (net.Conn, error) {
		return listner.Dial()
	}
}

func TestServer_GetHelloWorld(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithContextDialer(dialer()))

	if err != nil {
		t.Fatalf("grpc dial fail. | err: %s", err)
	}

	client := protocol_demo.NewHelloWorldClient(conn)

	rsp, err := client.GetHelloWorld(ctx, &protocol_demo.HelloWorldReq{
		Id: 100,
	})

	if err != nil {
		t.Fatalf("grpc call GetHelloWorld fail. | err: %s", err)
	}

	t.Logf("call success rsp.Message: %s", rsp.Message)
}
