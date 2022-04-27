package Webcontext

import (
	"encoding/json"
	"io"

	fream "github.com/Hywebsocket/Fream"
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

	frem := fream.DataFream{
		Fin:                1,
		Rsv:                true,
		OpCode:             byte(0x01),
		Mask:               byte(0),
		PayLoadLenth:       byte(len(Body)),
		Makeing_Key:        w.key,
		PlayLoadData:       Body,
		ExtenDedPayLoadLen: []byte{},
	}
	if w.key == nil {
		frem.Mask = byte(0)
	}
	untity := fream.NewDataFreamCoding()

	w.ioWrite.Write(untity.EnCodingDataFream(frem))
}

func (w *WebsocketResp) WriteJson(obj interface{}) error {

	bodys, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	w.Write(bodys)
	return nil

}
