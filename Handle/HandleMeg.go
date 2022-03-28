package handle

import (
	fream "Hywebsocket/Fream"
	"context"
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
