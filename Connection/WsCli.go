package connection

import "net"

type WsCli struct {
	conn net.Conn
}

func (c WsCli) SetConn(cli net.Conn) {
	c.conn = cli
}
func (c WsCli) Write(meg []byte) error {
	//定义发送数据的接口

	return nil
}
