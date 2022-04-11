package main

import (
	connection "Hywebsocket/Connection"
	handle "Hywebsocket/Handle"
	http "Hywebsocket/untity"
	"context"
	"fmt"
	"net"
)

var HttpUntity http.HttpUntity = http.HttpUntity{}

type hwebsocket struct {
	ConnFunc OnConnFunc
	ReadFunc ReadEventFunc
}

func NewWebsocket() Websocket {
	h := hwebsocket{}
	return &h
}
func (h *hwebsocket) OnConnect(funs OnConnFunc) {
	h.ConnFunc = funs
}
func (h *hwebsocket) StartServer(port string) {

	fmt.Println("服务已经启动")
	conn, err := net.Listen("tcp", port)

	if err != nil {
		panic(err)
	}
	for {
		cli, err := conn.Accept()
		if err != nil {
			fmt.Println(err.Error())

		}
		fmt.Println("收到请求")
		h.ShakeCli(cli)

	}
}

//收到消息触发回调
func (h *hwebsocket) onReadEvent(fun ReadEventFunc) {

	h.ReadFunc = fun
}

//断开链接触发回调
func (h *hwebsocket) OnClose() {

}

//握手
func (h *hwebsocket) ShakeCli(c net.Conn) {

	//defer c.Close()

	Https := make([]byte, 1024)

	c.Read(Https)

	//握手
	ShakeMeg := HttpUntity.Handshake(Https)

	c.Write(ShakeMeg)

	//握手成功调用OnConnect回调

	Wscliobj := connection.NewWsCli()

	Wscliobj.SetConn(c)

	if h.ConnFunc != nil {

		h.ConnFunc(Wscliobj)
	}

	han := handle.NewDispMessage(h.ReadFunc)

	go han.OnRead(c, net.IP(c.LocalAddr().Network()), context.Background())
}
