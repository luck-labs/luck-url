package rpc

import (
	"context"
	"errors"
	"github.com/luck-labs/luck-url/plugin/conf"
	"github.com/luck-labs/luck-url/plugin/log"
	"github.com/luck-labs/luck-url/plugin/rpc/dataserial"
	pool "github.com/jolestar/go-commons-pool/v2"
	"github.com/sirupsen/logrus"
	"net"
	"reflect"
)

var RPCClientPool *ClientPool

type ClientPool struct {
	ConnPool   *pool.ObjectPool
	ServerAddr string
	Ctx        context.Context
}

func NewClientPool() *ClientPool {
	ctx := context.Background()
	rpcConfig := conf.GlobalConfig.RpcServer
	config := pool.NewDefaultPoolConfig()
	config.MaxTotal = rpcConfig.MaxTotal
	config.MaxIdle = rpcConfig.MaxIdle
	config.MinIdle = rpcConfig.MinIdle
	connPool := pool.NewObjectPool(ctx, &PoolConnFactory{
		serverAddr: rpcConfig.Address,
	}, config)
	return &ClientPool{
		ConnPool:   connPool,
		ServerAddr: rpcConfig.Address,
		Ctx:        ctx,
	}
}

// Client struct
type Client struct {
	Conn net.Conn
}

// NewClient creates a new client
func NewClient(conn net.Conn) *Client {
	return &Client{Conn: conn}
}

// CallRPC Method
func (c *ClientPool) CallRPC(rpcName string, fPtr interface{}) {
	ctx := context.Background()
	container := reflect.ValueOf(fPtr).Elem()
	f := func(req []reflect.Value) []reflect.Value {
		poolConnObj, err := c.ConnPool.BorrowObject(c.Ctx)
		defer func() {
			if poolConnObj != nil {
				err = c.ConnPool.ReturnObject(c.Ctx, poolConnObj)
				if err != nil {
					log.Error(ctx, log.MduRpcFramework, log.IdxRpcFrameworkClient, logrus.Fields{})
				}
			}
		}()
		cReqTransport := poolConnObj.(*PoolConnObj).Transport
		//cReqTransport := transport.NewTransport(c.Conn)
		errorHandler := func(err error) []reflect.Value {
			outArgs := make([]reflect.Value, container.Type().NumOut())
			for i := 0; i < len(outArgs)-1; i++ {
				outArgs[i] = reflect.Zero(container.Type().Out(i))
			}
			outArgs[len(outArgs)-1] = reflect.ValueOf(&err).Elem()
			return outArgs
		}

		// Process input parameters
		inArgs := make([]interface{}, 0, len(req))
		for _, arg := range req {
			inArgs = append(inArgs, arg.Interface())
		}

		// ReqRPC
		reqRPC := dataserial.RPCdata{Name: rpcName, Args: inArgs}
		b, err := dataserial.Encode(reqRPC)
		if err != nil {
			panic(err)
		}
		err = cReqTransport.Send(b)
		if err != nil {
			return errorHandler(err)
		}
		// receive response from server
		rsp, err := cReqTransport.Read()
		if err != nil { // local network error or decode error
			return errorHandler(err)
		}
		rspDecode, _ := dataserial.Decode(rsp)
		if rspDecode.Err != "" { // remote server error
			return errorHandler(errors.New(rspDecode.Err))
		}

		if len(rspDecode.Args) == 0 {
			rspDecode.Args = make([]interface{}, container.Type().NumOut())
		}
		// unpack response arguments
		numOut := container.Type().NumOut()
		outArgs := make([]reflect.Value, numOut)
		for i := 0; i < numOut; i++ {
			if i != numOut-1 { // unpack arguments (except error)
				if rspDecode.Args[i] == nil { // if argument is nil (gob will ignore "Zero" in transmission), set "Zero" value
					outArgs[i] = reflect.Zero(container.Type().Out(i))
				} else {
					outArgs[i] = reflect.ValueOf(rspDecode.Args[i])
				}
			} else { // unpack error argument
				outArgs[i] = reflect.Zero(container.Type().Out(i))
			}
		}

		return outArgs
	}
	container.Set(reflect.MakeFunc(container.Type(), f))
}
