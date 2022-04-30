package HyWebsocket

import (
	"io"

	fream "github.com/adminoryuan/HyWebsocket/Fream"
)

func Pong(w io.Writer) {
	f := fream.DataFream{
		Fin:          0,
		Rsv:          true,
		Mask:         0,
		PayLoadLenth: byte(0),
		PlayLoadData: []byte{},
		OpCode:       0x0A,
	}
	bodys := fream.NewDataFreamCoding().EnCodingDataFream(f)
	w.Write(bodys)
}
