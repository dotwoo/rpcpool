package transfer

import (
	"context"
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"sync"
	"time"
)

var (
	// LatencyInit the latency init  is 3 seconds
	LatencyInit = time.Second * 3
)

// client the rpc client
// +gen * slice:"MinBy,SortBy,DistinctBy,Aggregate[int],Shuffle"
type client struct {
	client   *rpc.Client
	connTime time.Time
	latency  time.Duration
	lock     sync.RWMutex

	connTimeout time.Duration
	callTimeout time.Duration
	addr        string
}

// get get the rpc net conn
func (cn *client) get() (*rpc.Client, error) {
	cn.lock.Lock()
	defer cn.lock.Unlock()

	var err error
	//将对应latency 设置为最大避免,同时调用
	cn.latency = time.Minute * 3
	if cn.client == nil {
		if cn.client, err = cn.newConn(); err != nil {
			return nil, err
		}
		return cn.client, nil
	}
	if time.Since(cn.connTime).Hours() >= 2 {
		cn.client.Close()
		if cn.client, err = cn.newConn(); err != nil {
			return nil, err
		}
		return cn.client, nil
	}
	return cn.client, err
}

func (cn *client) newConn() (*rpc.Client, error) {
	conn, err := net.DialTimeout("tcp", cn.addr, cn.connTimeout)
	if err != nil {
		return nil, err
	}
	cn.connTime = time.Now()
	return jsonrpc.NewClient(conn), nil
}

// close close the conn
func (cn *client) close() {
	if cn.client != nil {
		cn.client.Close()
		cn.client = nil
	}
}

// call rpc method
func (cn *client) call(method string, args interface{}, resp interface{}) error {
	cl, err := cn.get()
	if err != nil {
		return err
	}

	ctx, cancelFn := context.WithTimeout(context.Background(), cn.callTimeout)
	defer cancelFn()

	done := make(chan error, 2)
	start := time.Now()
	go func() {
		done <- cl.Call(method, args, resp)
	}()

	select {
	case err := <-done:
		if err != nil {
			cn.close()
			cn.latency = time.Minute
			return fmt.Errorf("rpc call %s error:%s", cn.addr, err.Error())
		}

	case <-ctx.Done():
		cn.close()
		cn.latency = time.Minute * 2
		return fmt.Errorf("rpc call %s timeout", cn.addr)
	}
	cn.latency = time.Now().Sub(start)
	return nil
}

// NewClient create the rpc client
func NewClient(address string, connTW, callTW time.Duration) *client {
	c := &client{
		addr:        address,
		connTimeout: connTW,
		callTimeout: callTW,
		latency:     3 * time.Second,
	}
	return c
}

func quick(a, b *client) bool {
	return a.latency < b.latency
}

func aggregate(i int, c *client) int {
	c.latency = LatencyInit
	return i + 1
}

func addrLess(a, b *client) bool {
	return a.addr < b.addr
}

func distinctBy(a, b *client) bool {
	return a.addr == b.addr
}
