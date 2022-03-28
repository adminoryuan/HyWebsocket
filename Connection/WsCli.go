package connection

import (
	fream "Hywebsocket/Fream"
	"io"
	"net"
	"sync"
)

type WsCli struct {
	conn     net.Conn
	Mask_key []byte //保存上一次读取到的Make_key
	Locks    sync.RWMutex
	Rfunc    ReadEventFunc
}

var FreamObj fream.Fream = fream.NewDataFreamCoding()

func (c WsCli) SetConn(cli net.Conn) {
	c.conn = cli
}
func (c WsCli) SetReadFunc(Rfunc ReadEventFunc) {
	c.Rfunc = Rfunc
}

//当链接建立成功时 监听读
func (c WsCli) Read() {

	go func() {
		var Bodys []byte = make([]byte, 512)
		for {
			n, err := c.conn.Read(Bodys)
			if err == io.EOF {
				break
			}
			f := FreamObj.DecodeDataFream(Bodys[:n])
			c.Locks.Lock()
			c.Mask_key = f.Makeing_Key
			c.Locks.Unlock()

			c.Rfunc(Bodys[:n])
		}

	}()
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
