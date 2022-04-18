package main

import (
	connection "Hywebsocket/Connection"
	request "Hywebsocket/context"
	"fmt"
)

func main() {
	h := NewWebsocket()

	h.OnConnect(func(ic connection.IWsCli) {
		ic.Write([]byte("heelo"))
		fmt.Printf("链接成功 \n")
	})

	h.onReadEvent(func(c request.Context) {
		

	})

	h.StartServer(":9091")
}
