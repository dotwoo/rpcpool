package transfer

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestClientsPool_NewClient(t *testing.T) {
	type fields struct {
		cs          clientSlice
		nest        *sync.Pool
		lock        sync.RWMutex
		connTimeout time.Duration
		callTimeout time.Duration
		addrs       []string
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &ClientsPool{
				cs:          tt.fields.cs,
				nest:        tt.fields.nest,
				lock:        tt.fields.lock,
				connTimeout: tt.fields.connTimeout,
				callTimeout: tt.fields.callTimeout,
				addrs:       tt.fields.addrs,
			}
			if got := cp.NewClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientsPool.NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientsPool_Put(t *testing.T) {
	type fields struct {
		cs          clientSlice
		nest        *sync.Pool
		lock        sync.RWMutex
		connTimeout time.Duration
		callTimeout time.Duration
		addrs       []string
	}
	type args struct {
		c *client
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &ClientsPool{
				cs:          tt.fields.cs,
				nest:        tt.fields.nest,
				lock:        tt.fields.lock,
				connTimeout: tt.fields.connTimeout,
				callTimeout: tt.fields.callTimeout,
				addrs:       tt.fields.addrs,
			}
			cp.Put(tt.args.c)
		})
	}
}

func TestClientsPool_Get(t *testing.T) {
	type fields struct {
		cs          clientSlice
		nest        *sync.Pool
		lock        sync.RWMutex
		connTimeout time.Duration
		callTimeout time.Duration
		addrs       []string
	}
	tests := []struct {
		name   string
		fields fields
		want   *client
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &ClientsPool{
				cs:          tt.fields.cs,
				nest:        tt.fields.nest,
				lock:        tt.fields.lock,
				connTimeout: tt.fields.connTimeout,
				callTimeout: tt.fields.callTimeout,
				addrs:       tt.fields.addrs,
			}
			if got := cp.Get(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ClientsPool.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientsPool_ReListClients(t *testing.T) {
	type fields struct {
		cs          clientSlice
		nest        *sync.Pool
		lock        sync.RWMutex
		connTimeout time.Duration
		callTimeout time.Duration
		addrs       []string
	}
	type args struct {
		al []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &ClientsPool{
				cs:          tt.fields.cs,
				nest:        tt.fields.nest,
				lock:        tt.fields.lock,
				connTimeout: tt.fields.connTimeout,
				callTimeout: tt.fields.callTimeout,
				addrs:       tt.fields.addrs,
			}
			cp.ReListClients(tt.args.al)
		})
	}
}

func TestCreatePool(t *testing.T) {
	type args struct {
		al     []string
		connTW time.Duration
		callTW time.Duration
	}
	tests := []struct {
		name string
		args args
		want *ClientsPool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreatePool(tt.args.al, tt.args.connTW, tt.args.callTW); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClientsPool_Call(t *testing.T) {
	type fields struct {
		cs          clientSlice
		nest        *sync.Pool
		lock        sync.RWMutex
		connTimeout time.Duration
		callTimeout time.Duration
		addrs       []string
	}
	type args struct {
		method string
		args   interface{}
		reply  interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cp := &ClientsPool{
				cs:          tt.fields.cs,
				nest:        tt.fields.nest,
				lock:        tt.fields.lock,
				connTimeout: tt.fields.connTimeout,
				callTimeout: tt.fields.callTimeout,
				addrs:       tt.fields.addrs,
			}
			if err := cp.Call(tt.args.method, tt.args.args, tt.args.reply); (err != nil) != tt.wantErr {
				t.Errorf("ClientsPool.Call() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
