package connection

import (
	ctx "Hywebsocket/WebContext"
	"net"
)

type ReadEventFunc func(ctx.Context)

//握手成功后返回的对象
type IWsCli interface {
	Write([]byte) error
	SetReadFunc(ReadEventFunc)
	OnRead()
	SetConn(cli net.Conn)
	GetRemoterAddr() net.Addr
}
