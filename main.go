package main

import (
	connection "Hywebsocket/Connection"
	request "Hywebsocket/Request"
	"fmt"
)

func main() {
	h := NewWebsocket()

	h.StartServer(":9091")

	h.OnConnect(func(ic connection.IWsCli) {
		ic.Write([]byte("heelo"))
	})
	h.onReadEvent(func(rc request.RequestConn) {
		fmt.Println(string(rc.Bodys))
	})

}
