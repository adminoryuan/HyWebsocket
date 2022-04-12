package connection

import (
	"Hywebsocket/context"
	"net"
)

type ReadEventFunc func(context.Context)

//握手成功后返回的对象
type IWsCli interface {
	Write([]byte) error
	SetReadFunc(ReadEventFunc)
	OnRead()
	SetConn(cli net.Conn)
}
