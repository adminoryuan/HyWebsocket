package main

import (
	connection "Hywebsocket/Connection"
	request "Hywebsocket/Request"
	"fmt"
)

func main() {
	h := NewWebsocket()

	h.OnConnect(func(ic connection.IWsCli) {
		ic.Write([]byte("heelo"))
		fmt.Printf("链接成功")
	})
	h.onReadEvent(func(rc request.RequestConn) {
		fmt.Printf("my Recive.. %s",string(rc.Bodys))
	})
	h.StartServer(":9091")

}
