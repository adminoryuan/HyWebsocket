package connection

import (
	request "Hywebsocket/Request"
	"net"
)

type ReadEventFunc func(request.RequestConn)

//握手成功后返回的对象
type IWsCli interface {
	Write([]byte) error
	SetReadFunc(ReadEventFunc)
	OnRead()
	SetConn(cli net.Conn)
}
