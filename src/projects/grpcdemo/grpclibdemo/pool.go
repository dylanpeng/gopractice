package grpclibdemo

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"sync"
	"time"
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

func (p *Pool) GetConnection(addr string) (conn *ClientConn, err error) {
	connChan, exist := p.conns[addr]

	if !exist {
		connChan = make(chan *ClientConn, p.capacity)
	}

	select {
	case conn = <-connChan:
		// all good
	default:
	}

	// 超时关闭连接
	if conn != nil && conn.ClientConn != nil && (conn.LastUseTime.Add(p.idle).Before(time.Now()) || conn.CreateTime.Add(p.ttl).Before(time.Now())) {
		_ = conn.ClientConn.Close()
		conn.ClientConn = nil
	}

	if conn == nil || conn.ClientConn == nil {
		conn, err = p.CreateConnection(addr)

		if err != nil {
			fmt.Printf("GetConnection CreateConnection fail. | addr: %s | err: %s", addr, err)
			return
		}

		return
	}

	return
}

func (p *Pool) CreateConnection(addr string) (conn *ClientConn, err error) {
	conn = &ClientConn{
		pool:        p,
		CreateTime:  time.Now(),
		LastUseTime: time.Now(),
		addr:        addr,
	}

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
		for conn := range connChan {
			_ = conn.ClientConn.Close()
		}
	}

	p.close = true

}
