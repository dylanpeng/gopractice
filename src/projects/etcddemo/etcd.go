package main

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var client *clientv3.Client
var dir string = "/test/node1"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var err error
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 10 * time.Second,
	})

	if err != nil {
		fmt.Printf("new etcd client fail. | err: %s\n", err)
		return
	}

	user := &User{
		Name: "user_name",
		Age:  20,
	}

	data, _ := json.Marshal(user)

	_ = AddNode(dir, string(data))

	_ = GetNode(dir)
}

func AddNode(key, value string) error {
	rsp, err := client.Put(context.Background(), key, value)

	if err != nil {
		fmt.Printf("put etcd node fail. | err: %s\n", err)
		return err
	}

	fmt.Printf("rsp: %+v", *rsp)
	return nil
}

func GetNode(key string) error {
	rsp, err := client.Get(context.Background(), key)

	if err != nil {
		fmt.Printf("put etcd node fail. | err: %s\n", err)
		return err
	}

	user := &User{}

	_ = json.Unmarshal(rsp.Kvs[0].Value, user)

	fmt.Printf("user: %+v", *user)
	return nil
}
