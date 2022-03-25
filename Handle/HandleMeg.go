package handle

import (
	fream "Hywebsocket/Fream"
	"context"
	"fmt"
	"io"
)

type DispCliMessage struct {
	codobj       fream.Fream
	Meg          chan []byte
	Canle        context.CancelFunc
	PlayLoadData chan []byte
}

//服务端的消息监听和发送
func NewDispMessage() DispCliMessage {

	d := DispCliMessage{}
	d.codobj = fream.NewDataFreamCoding()
	d.Meg = make(chan []byte, 10)
	d.PlayLoadData = make(chan []byte, 10)
	return d

}
func (d DispCliMessage) OnRead(c io.Reader, ctx context.Context) {
	Mes := make([]byte, 128)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			c.Read(Mes)

			cliFream := d.codobj.DecodeDataFream(Mes)

			d.PlayLoadData <- cliFream.PlayLoadData

			fmt.Printf(string(cliFream.PlayLoadData))
		}
		//c.Write(nes)
	}
}
func (d DispCliMessage) OnWrite(w io.Writer, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关闭了")
			return
		case a := <-d.Meg:
			fmt.Println("接收到数据")
			w.Write(a)
		}
	}
}
