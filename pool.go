package transfer

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// ClientsPool is client pool for transfer servers
type ClientsPool struct {
	cs   clientSlice
	nest *sync.Pool
	lock sync.RWMutex

	connTimeout time.Duration
	callTimeout time.Duration
	addrs       []string
}

func (cp *ClientsPool) NewClient() interface{} {
	cp.lock.RLock()
	c, err := cp.cs.minBy(quick)
	cp.lock.RUnlock()
	if err == nil {
		return c
	}
	log.Println("NewClient error:", err.Error())
	return nil
}

// Put cache client
func (cp *ClientsPool) Put(c *client) {
	if time.Since(c.connTime).Hours() > 2 {
		c.close()
		return
	}
	cp.nest.Put(c)
}

// Get get client from cache or new
func (cp *ClientsPool) Get() *client {
	c, ok := (cp.nest.Get()).(*client)
	if ok {
		return c
	}
	cp.cs.aggregateInt(aggregate)
	return nil
}

// ReListClients reset the clients list
func (cp *ClientsPool) ReListClients(al []string) {
	newcs := make(clientSlice, 0)
	for _, a := range al {
		c := NewClient(a, cp.connTimeout, cp.callTimeout)
		newcs = append(newcs, c)
	}
	newcs.shuffle()
	cp.lock.Lock()
	cp.cs = newcs
	cp.addrs = al
	cp.lock.Unlock()
}

// CreatePool create client pool
func CreatePool(al []string, connTW, callTW time.Duration) *ClientsPool {
	cp := new(ClientsPool)
	cp.nest = &sync.Pool{
		New: cp.NewClient,
	}
	cp.connTimeout = connTW
	cp.callTimeout = callTW
	cp.ReListClients(al)
	return cp
}

// Call  the rpc call
func (cp *ClientsPool) Call(method string, args interface{}, reply interface{}) error {
	c := cp.Get()
	if c == nil {
		return fmt.Errorf("Cant get client")
	}

	err := c.call(method, args, reply)
	if err == nil {
		cp.Put(c)
		return nil
	}
	c.close()
	return err

}
