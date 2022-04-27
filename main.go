package main

import (
	"fmt"

	ctx "github.com/adminoryuan/HyWebsocket/WebContext"
)

func main() {

	h := NewWebsocket()

	h.OnConnect(func(ic IWsCli) {
		fmt.Printf("收到链接： %s \n", ic.GetRemoterAddr().String())
		ic.Write([]byte("heelo"))

	})

	h.onReadEvent(func(c ctx.Context) {
		fmt.Printf("recv %s \n", string(c.Req.Bodys))

		c.Resp.Write([]byte("zhangsan"))

	})
	h.StartServer(":9091")
}
