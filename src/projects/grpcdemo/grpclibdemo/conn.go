package grpclibdemo

import (
	"google.golang.org/grpc"
	"time"
)

type ClientConn struct {
	*grpc.ClientConn
	addr        string
	pool        *Pool
	CreateTime  time.Time
	LastUseTime time.Time
}

func (c *ClientConn) Close() error {
	// 直接go出去防止锁影响性能
	go func() {
		// 加锁
		c.pool.Lock()
		defer c.pool.Unlock()

		// 判断pool是否关闭
		if c.pool.close {
			return
		}

		connChan := c.pool.conns[c.addr]
		// channel有容量加入channel,满了直接关闭连接，防止内存泄漏
		if len(connChan) < c.pool.capacity {
			connChan <- c
		} else {
			_ = c.ClientConn.Close()
		}

	}()

	return nil
}
