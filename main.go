package main

import (
	connection "Hywebsocket/Connection"
	ctx "Hywebsocket/WebContext"
	"fmt"
)

func main() {
	h := NewWebsocket()

	h.OnConnect(func(ic connection.IWsCli) {
		ic.Write([]byte("heelo"))
		fmt.Printf("链接成功 \n")
	})

	h.onReadEvent(func(c ctx.Context) {
		fmt.Printf("recv %s \n", string(c.Req.Bodys))

		c.Resp.Write([]byte("zhangsan"))

	})
	h.StartServer(":9091")
}
