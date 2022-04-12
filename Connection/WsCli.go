package connection

import (
	fream "Hywebsocket/Fream"
	"fmt"
	"io"
	"net"
	"sync"
)

type wsCli struct {
	conn     net.Conn
	Mask_key []byte //保存上一次读取到的Make_key
	Locks    sync.RWMutex
	Rfunc    ReadEventFunc
}

func NewWsCli() IWsCli {
	w := wsCli{}
	w.Locks = sync.RWMutex{}
	return &w
}

var FreamObj fream.Fream = fream.NewDataFreamCoding()

func (c *wsCli) SetConn(cli net.Conn) {
	c.conn = cli
}
func (c *wsCli) SetReadFunc(Rfunc ReadEventFunc) {
	c.Rfunc = Rfunc
}

//当链接建立成功时 监听读
func (c *wsCli) Read() {
	go func() {
		var Bodys []byte = make([]byte, 512)
		for {
			n, err := c.conn.Read(Bodys)
			if err == io.EOF {
				break
			}
			f := FreamObj.DecodeDataFream(Bodys[:n])

			c.Mask_key = f.Makeing_Key 
			
			c.Rfunc(Bodys[:n])
		}

	}()
}
func (c *wsCli) Write(meg []byte) error {
	//定义发送数据的接口
	frem := fream.DataFream{
		Fin:          0,
		Rsv:          true,
		OpCode:       0x01,
		Mask:         1,
		PayLoadLenth: byte(len(meg)),
		PlayLoadData: meg,
	}

	frem.Makeing_Key = c.Mask_key

	fmt.Printf("mask_key ..%s",frem.Makeing_Key)


	go func() {
		c.Write(FreamObj.EnCodingDataFream(frem))
	}()
	return nil
}
