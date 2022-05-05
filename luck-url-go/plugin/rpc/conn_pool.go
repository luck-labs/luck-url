package rpc

import (
	"context"
	"fmt"
	"github.com/luck-labs/luck-url/plugin/rpc/transport"
	pool "github.com/jolestar/go-commons-pool/v2"
	"net"
	"sync/atomic"
)

type PoolConnObj struct {
	Transport *transport.Transport
}

type PoolConnFactory struct {
	serverAddr string
}

var count uint32 = 0

func (f *PoolConnFactory) MakeObject(ctx context.Context) (*pool.PooledObject, error) {
	conn, err := net.Dial("tcp", f.serverAddr)
	if err != nil {
		return nil, err
	}
	atomic.AddUint32(&count, 1)
	fmt.Printf("make obj count:%d \n", count)
	return pool.NewPooledObject(
			&PoolConnObj{
				Transport: transport.NewTransport(conn),
			}),
		nil
}

func (f *PoolConnFactory) DestroyObject(ctx context.Context, object *pool.PooledObject) error {
	connObj := object.Object.(*PoolConnObj)
	return connObj.Transport.Close()
}

func (f *PoolConnFactory) ValidateObject(ctx context.Context, object *pool.PooledObject) bool {
	connObj := object.Object.(*PoolConnObj)
	return connObj.Transport.IsConnected
}

func (f *PoolConnFactory) ActivateObject(ctx context.Context, object *pool.PooledObject) error {
	_ = object.Object.(*PoolConnObj)
	return nil
}

func (f *PoolConnFactory) PassivateObject(ctx context.Context, object *pool.PooledObject) error {
	_ = object.Object.(*PoolConnObj)
	return nil
}
