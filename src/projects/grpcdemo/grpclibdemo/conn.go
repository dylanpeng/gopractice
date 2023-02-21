package grpclibdemo

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

type ClientConn struct {
	*grpc.ClientConn
	addr        string
	pool        *Pool
	CreateTime  time.Time
	LastUseTime time.Time
	id          uint
}

func (c *ClientConn) Close() {
	// 直接go出去防止锁影响性能
	go func() {
		// 判断pool是否关闭
		if c.pool.close {
			return
		}

		connChan, exist := c.pool.conns[c.addr]

		if !exist {
			c.ClientConn.Close()
			c.ClientConn = nil
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		// channel有容量加入channel,满了直接关闭连接，防止内存泄漏
		select {
		case connChan <- c:
			// All good
		case <-ctx.Done():
			fmt.Printf("channel drop. | id: %d\n", c.id)
			return
		}

	}()

	return
}
