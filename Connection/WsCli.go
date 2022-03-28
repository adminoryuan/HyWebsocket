package connection

import (
	fream "Hywebsocket/Fream"
	"net"
	"sync"
)

type WsCli struct {
	conn     net.Conn
	Mask_key []byte //保存上一次读取到的Make_key
	Locks    sync.RWMutex
}

var FreamObj fream.Fream = fream.NewDataFreamCoding()

func (c WsCli) SetConn(cli net.Conn) {
	c.conn = cli
}
func (c WsCli) Write(meg []byte) error {
	//定义发送数据的接口
	frem := fream.DataFream{
		Fin:          0,
		Rsv:          true,
		OpCode:       0x01,
		Mask:         1,
		PayLoadLenth: byte(len(meg)),
		PlayLoadData: meg,
	}
	c.Locks.RLocker()
	frem.Makeing_Key = c.Mask_key
	c.Locks.RLocker().Unlock()

	go func() {
		c.Write(FreamObj.EnCodingDataFream(frem))
	}()
	return nil
}
func (c WsCli) Read() {
	for {
		Body := make([]byte, 512)
		c.conn.Read(Body)
		c.Mask_key = FreamObj.DecodeDataFream(Body).Makeing_Key

	}
}
