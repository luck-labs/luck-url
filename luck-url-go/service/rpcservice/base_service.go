package rpcservice

type Base struct {
}

var baseInstance = &Base{}

func NewBaseService() *Base {
	return baseInstance
}
