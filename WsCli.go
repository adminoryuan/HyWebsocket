package HyWebsocket

import (
	"fmt"
	"io"
	"net"
	"sync"

	fream "github.com/adminoryuan/HyWebsocket/Fream"
	webcontext "github.com/adminoryuan/HyWebsocket/WebContext"
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
func (c *wsCli) GetRemoterAddr() net.Addr {
	return c.conn.RemoteAddr()
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

			fmt.Printf("%d", f.OpCode)
			switch f.OpCode {
			case 8:
				c.conn.Close()
				return
			case 9:
				Pong(c.conn)
				return
			case 1:

				c.Mask_key = f.Makeing_Key
				ctx := webcontext.Context{
					Req: webcontext.RequestConn{
						LocalRemoter: net.IP(c.conn.RemoteAddr().Network()),
						Bodys:        f.PlayLoadData,
					},
					Resp: webcontext.NewWebsocketResp(c.conn, c.Mask_key),
				}
				c.Rfunc(ctx)
			}

			//c.Write([]byte("qqwer"))

		}

	}()
}
func (c *wsCli) Write(Body []byte) error {

	frem := fream.DataFream{
		Fin:                1,
		Rsv:                true,
		OpCode:             byte(0x01),
		Mask:               byte(0),
		PayLoadLenth:       byte(len(Body)),
		Makeing_Key:        c.Mask_key,
		PlayLoadData:       Body,
		ExtenDedPayLoadLen: []byte{},
	}

	go func() {
		c.conn.Write(c.freamObj.EnCodingDataFream(frem))
	}()
	return nil
}
