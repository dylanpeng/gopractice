package grpclibdemo

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"time"
)

var (
	// ErrClosed is the error when the client pool is closed
	ErrClosed = errors.New("grpc pool: client pool is closed")
	// ErrTimeout is the error when the client pool timed out
	ErrTimeout = errors.New("grpc pool: client pool timed out")
	// ErrAlreadyClosed is the error when the client conn was already closed
	ErrAlreadyClosed = errors.New("grpc pool: the connection was already closed")
	// ErrFullPool is the error when the pool is already full
	ErrFullPool = errors.New("grpc pool: closing a ClientConn into a full pool")

	index uint = 0
)

type Pool struct {
	sync.RWMutex
	conns    map[string]chan *ClientConn
	capacity int
	idle     time.Duration // 活跃时间
	ttl      time.Duration // 生命周期时间
	close    bool
}

func NewPool(capacity int, idle time.Duration, ttl time.Duration) *Pool {
	pool := &Pool{
		conns:    make(map[string]chan *ClientConn, 16),
		capacity: capacity,
		idle:     idle,
		ttl:      ttl,
	}

	if idle <= 0 {
		pool.idle = 20 * time.Minute
	}

	if ttl <= 0 {
		pool.ttl = 120 * time.Minute
	}

	return pool
}

func (p *Pool) InitClientConn(addr string) (chan *ClientConn, error) {
	p.Lock()
	defer p.Unlock()

	// 再次确认没有初始化过
	connChan, exist := p.conns[addr]
	if exist {
		return connChan, nil
	}

	connChan = make(chan *ClientConn, p.capacity)
	for i := 0; i < p.capacity; i++ {
		clientConn, err := p.CreateConnection(addr)

		if err != nil {
			fmt.Printf("InitClientConn fail. | addr: %s | err: %s\n", addr, err)
			return nil, err
		}

		connChan <- clientConn
	}

	p.conns[addr] = connChan

	return connChan, nil
}

func (p *Pool) GetConnection(ctx context.Context, addr string) (conn *ClientConn, err error) {
	// 没有过期时间，10秒过期
	if _, hasDeadline := ctx.Deadline(); !hasDeadline {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
	}

	connChan, exist := p.conns[addr]

	// 不存在，初始化channel
	if !exist {
		connChan, err = p.InitClientConn(addr)

		if err != nil {
			fmt.Printf("GetConnection InitClientConn. | addr: %s | err: %s", addr, err)
			return nil, err
		}
	}

	// 等待连接，超时返回error
	select {
	case conn = <-connChan:
		if conn.ClientConn != nil && (conn.LastUseTime.Add(p.idle).Before(time.Now()) || conn.CreateTime.Add(p.ttl).Before(time.Now())) {
			conn.ClientConn.Close()
			conn.ClientConn = nil
		}
	case <-ctx.Done():
		return nil, ErrTimeout
	}

	// 连接中断，重新创建连接
	if conn.ClientConn != nil {
		connState := conn.GetState()

		if connState == connectivity.TransientFailure || connState == connectivity.Shutdown {
			conn.ClientConn.Close()
			conn.ClientConn = nil
		}
	}

	if conn.ClientConn == nil {
		conn, err = p.CreateConnection(addr)

		if err != nil {
			fmt.Printf("GetConnection CreateConnection fail. | addr: %s | err: %s", addr, err)
			return
		}

		return
	}

	return
}

// CreateConnection 创建连接
func (p *Pool) CreateConnection(addr string) (conn *ClientConn, err error) {
	conn = &ClientConn{
		pool:        p,
		CreateTime:  time.Now(),
		LastUseTime: time.Now(),
		addr:        addr,
		id:          index,
	}

	index++

	conn.ClientConn, err = grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Printf("CreateConnection Dial fail. | err: %s", err)
		return
	}

	return
}

func (p *Pool) Close() {
	if p == nil {
		return
	}

	p.Lock()
	defer p.Unlock()

	for _, connChan := range p.conns {
		close(connChan)
		for conn := range connChan {
			_ = conn.ClientConn.Close()
		}
	}

	p.close = true

}
