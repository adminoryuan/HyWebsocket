package Webcontext

import (
	fream "Hywebsocket/Fream"
	"fmt"
	"io"
)

type WebsocketResp struct {
	key     []byte
	ioWrite io.Writer
}

func NewWebsocketResp(write io.Writer, key []byte) WebsocketResp {
	wobj := WebsocketResp{}
	wobj.ioWrite = write
	wobj.key = key
	return wobj

}
func (w *WebsocketResp) Write(Body []byte) {
	fmt.Printf("send make_key=%s", w.key)
	frem := fream.DataFream{
		Fin:          0,
		Rsv:          true,
		OpCode:       0x01,
		Mask:         1,
		PayLoadLenth: byte(len(Body)),
		Makeing_Key:  w.key,
		PlayLoadData: Body,
	}
	untity := fream.NewDataFreamCoding()
	w.ioWrite.Write(untity.EnCodingDataFream(frem))
}
