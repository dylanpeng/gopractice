package grpclibdemo

import (
	"context"
	"errors"
	"fmt"
	"gopractice/lib/proto/common"
	"math/rand"
	"strings"
	"time"
)

type ClientConfig struct {
	AddrMap  map[string][]*Config `json:"addr_map" toml:"addr_map"`
	Idle     int                  `json:"idle" toml:"idle"`
	Ttl      int                  `json:"ttl" toml:"ttl"`
	Capacity int                  `json:"capacity" toml:"capacity"`
	Timeout  int                  `json:"timeout" toml:"timeout"`
}

type Client struct {
	conf *ClientConfig
	pool *Pool
}

var failCount = 0

func (c *Client) CallGrpc(method string, req interface{}, rsp interface{}, timeout time.Duration) error {
	//if timeout == 0 {
	//	timeout = 3 * time.Second
	//}
	//
	//ctx, _ := context.WithTimeout(context.Background(), timeout)

	groupName, err := c.GetGroupName(method)

	if err != nil {
		return err
	}

	addrs, exist := c.conf.AddrMap[groupName]

	if !exist || len(addrs) == 0 {
		return errors.New("not init grpc client\n")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(len(addrs))

	conn, err := c.pool.GetConnection(context.Background(), addrs[i].GetAddress())

	if err != nil {
		fmt.Printf("GetClient GetConnection fail. | err: %s\n", err)
		return errors.New("GetClient GetConnection fail")
	}

	fmt.Printf("before conn id: %d | state: %s\n", conn.id, conn.GetState())

	err = conn.Invoke(context.Background(), common.CommonService_CommonTest_FullMethodName, req, rsp)

	fmt.Printf("connect id: %d | ip: %s called | state: %s | message: %+v\n", conn.id, conn.addr, conn.GetState(), rsp)

	if err != nil {
		fmt.Printf("invoke fail. | id: %d | failCount: %d | err: %s\n", conn.id, failCount, err)
		failCount++
		return errors.New("invoke fail")
	}

	conn.Close()

	return nil
}

func (c *Client) GetGroupName(method string) (string, error) {
	item := strings.Split(method, "/")

	if len(item) < 3 {
		return "", errors.New("undefined grpc service")
	}

	return item[1], nil
}

func NewClient(conf *ClientConfig) (client *Client) {
	client = &Client{
		conf: conf,
	}
	client.pool = NewPool(conf.Capacity, time.Duration(conf.Idle)*time.Minute, time.Duration(conf.Ttl)*time.Minute)

	for _, addrs := range conf.AddrMap {
		for _, addr := range addrs {
			_, err := client.pool.InitClientConn(addr.GetAddress())

			if err != nil {
				fmt.Printf("InitClientConn fail.\n")
			}
		}
	}

	return
}
