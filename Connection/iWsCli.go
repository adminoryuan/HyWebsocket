package connection

import "net"

type ReadEventFunc func([]byte)

//握手成功后返回的对象
type IWsCli interface {
	Write([]byte) error
	SetReadFunc(ReadEventFunc)

	SetConn(cli net.Conn)
}
