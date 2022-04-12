package context

import "net"

//描述一个用户端发送数据
type RequestConn struct {
	LocalRemoter net.IP

	Bodys []byte
}


