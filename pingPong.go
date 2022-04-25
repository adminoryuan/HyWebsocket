package main

import (
	fream "Hywebsocket/Fream"
	"io"
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
