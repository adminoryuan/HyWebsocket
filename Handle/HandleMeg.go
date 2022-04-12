package handle

import (
	fream "Hywebsocket/Fream"
	request "Hywebsocket/Request"
	"context"
	"errors"
	"fmt"
	"io"
	"net"
)

type DispCliMessage struct {
	codobj fream.Fream
	Canle  context.CancelFunc
	ReadEvent func(request.RequestConn)
}

//服务端的消息监听和发送
func NewDispMessage(f func(request.RequestConn)) DispCliMessage {

	d := DispCliMessage{}
	d.ReadEvent = f
	d.codobj = fream.NewDataFreamCoding()

	return d

}

//监听读
func (d *DispCliMessage) OnRead(c io.Reader, ip net.IP, ctx context.Context) {
	Mes := make([]byte, 128)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			c.Read(Mes)

			cliFream := d.codobj.DecodeDataFream(Mes)
			fmt.Println("recive cliform...")
			//执行一个回调函数

			if d.ReadEvent == nil {
				panic(errors.New("Notfould recive envent..."))
			}

			go func() {
				d.ReadEvent(
					request.RequestConn{
						LocalRemoter: ip,
						Bodys:        cliFream.PlayLoadData,
					})
			}()

		}
		//c.Write(nes)
	}
}
