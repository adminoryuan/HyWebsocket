package handle

import (
	fream "Hywebsocket/Fream"
	"context"
	"io"
)

type DispCliMessage struct {
	codobj fream.Fream

	Canle        context.CancelFunc
	PlayLoadData chan []byte
}

//服务端的消息监听和发送
func NewDispMessage() DispCliMessage {

	d := DispCliMessage{}
	d.codobj = fream.NewDataFreamCoding()

	d.PlayLoadData = make(chan []byte, 10)
	return d

}

//监听读
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

		}
		//c.Write(nes)
	}
}
