package main

import (
	connection "Hywebsocket/Connection"
	"fmt"
)

func main() {
	h := NewWebsocket()

	h.StartServer(":9091")

	h.OnConnect(func(ic connection.IWsCli) {
		ic.Write([]byte("heelo"))
	})
	h.onReadEvent(func(b RequestConn) {
		fmt.Println("收到了数据")
		fmt.Println(string(b))
	})
}
