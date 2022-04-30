package HyWebsocket

import (
	"fmt"
	"testing"

	ctx "github.com/adminoryuan/HyWebsocket/WebContext"
)

func TestServerTest(t *testing.T) {
	h := NewWebsocket()

	h.OnConnect(func(ic IWsCli) {
		fmt.Printf("收到链接： %s \n", ic.GetRemoterAddr().String())
		ic.Write([]byte("heelo"))
	})

	h.OnReadEvent(func(c ctx.Context) {
		fmt.Printf("recv %s \n", string(c.Req.Bodys))
		c.Resp.Write([]byte("zhangsan"))
	})
	h.StartServer(":9091")
}
