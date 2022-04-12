package context

import "io"

type WebsocketResp struct {
	ioWrite io.Writer
}

func NewWebsocketResp(write io.Writer) WebsocketResp {
	wobj := WebsocketResp{}
	wobj.ioWrite = write

	return wobj

}
func (w *WebsocketResp) Write(Body []byte) {

}
