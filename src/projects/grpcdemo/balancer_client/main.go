package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"gopractice/lib/proto/protocol_demo"
)

// Following is an example name resolver implementation. Read the name
// resolution example to learn more about it.

const (
	exampleScheme      = "example"
	exampleServiceName = "lb.example.grpc.lixueduan.com"
)

var addrs = []string{"localhost:50052", "localhost:50053"}

func main() {
	resolver.Register(&exampleResolverBuilder{})

	conn, err := grpc.Dial(fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName),
		// 配置 loadBalancing 策略
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("new grpc client failed: %s \n", err)
		return
	}

	defer conn.Close()

	c := protocol_demo.NewHelloWorldClient(conn)

	for i := 0; i < 100; i++ {
		_, err := c.GetHelloWorld(context.Background(), &protocol_demo.HelloWorldReq{Id: 100})

		if err != nil {
			fmt.Printf("request GetHelloWorld faild: %s \n", err)
			return
		}
	}
}

type exampleResolverBuilder struct{}

func (*exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &exampleResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			exampleServiceName: addrs,
		},
	}
	r.start()
	return r, nil
}
func (*exampleResolverBuilder) Scheme() string { return exampleScheme }

type exampleResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *exampleResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint()]
	ads := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		ads[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: ads})
}
func (*exampleResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (*exampleResolver) Close()                                  {}
