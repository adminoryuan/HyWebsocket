package main

import (
	connection "Hywebsocket/Connection"
)

func main() {
	h := NewWebsocket()

	h.StartServer(":9091")

	h.OnConnect(func(ic connection.IWsCli) {
		ic.Write([]byte("heelo"))
	})

}
