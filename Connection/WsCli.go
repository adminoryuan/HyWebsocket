package connection

import (
	fream "Hywebsocket/Fream"
	webcontext "Hywebsocket/WebContext"
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
	freamObj fream.Fream
}

func NewWsCli() IWsCli {
	w := wsCli{}
	w.Locks = sync.RWMutex{}
	w.freamObj = fream.NewDataFreamCoding()
	return &w
}

func (c *wsCli) SetConn(cli net.Conn) {
	c.conn = cli
}
func (c *wsCli) SetReadFunc(Rfunc ReadEventFunc) {
	c.Rfunc = Rfunc
}

//当链接建立成功时 监听读
func (c *wsCli) OnRead() {
	go func() {
		var Bodys []byte = make([]byte, 512)
		for {
			n, err := c.conn.Read(Bodys)
			if err == io.EOF {
				break
			}

			f := c.freamObj.DecodeDataFream(Bodys[:n])
			c.Mask_key = f.Makeing_Key

			fmt.Printf("mask_key : = %s \n", c.Mask_key)

			ctx := webcontext.Context{
				Req: webcontext.RequestConn{
					LocalRemoter: net.IP(c.conn.RemoteAddr().Network()),
					Bodys:        f.PlayLoadData,
				},
				Resp: webcontext.NewWebsocketResp(c.conn, c.Mask_key),
			}
			c.Rfunc(ctx)

			//c.Write([]byte("qqwer"))

		}

	}()
}
func (c *wsCli) Write(meg []byte) error {

	//定义发送数据的接口
	frem := fream.DataFream{
		Fin:          0,
		Rsv:          true,
		OpCode:       0x01,
		Mask:         0,
		PayLoadLenth: byte(len(meg)),
		Makeing_Key:  c.Mask_key,
		PlayLoadData: meg,
	}
	go func() {
		fmt.Printf("make-send %s \n", c.Mask_key)
		c.conn.Write(c.freamObj.EnCodingDataFream(frem))
	}()

	return nil
}
