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
		fmt.Printf("链接成功 \n")
	})

	h.onReadEvent(func(rc request.RequestConn) {
		fmt.Printf("my Recive.. %s \n", string(rc.Bodys))

	})

	h.StartServer(":9091")
}
