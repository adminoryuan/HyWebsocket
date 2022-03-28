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

var (
	ConnFunc OnConnFunc
	ReadFunc ReadEventFunc
)

type hwebsocket struct {
	hobj handle.DispCliMessage
}

func NewWebsocket() Websocket {
	h := hwebsocket{}
	h.hobj = handle.NewDispMessage()
	return h
}
func (h hwebsocket) OnConnect(funs OnConnFunc) {
	ConnFunc = funs
}
func (h hwebsocket) StartServer(port string) {

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

func (h hwebsocket) onReadEvent(fun ReadEventFunc) {
	ReadFunc = fun
}

//握手
func (h hwebsocket) ShakeCli(c net.Conn) {

	//defer c.Close()
	Ctx, cal := context.WithCancel(context.Background())

	h.hobj.Canle = cal

	go h.hobj.OnWrite(c, Ctx)

	go h.hobj.OnRead(c, Ctx)

	Https := make([]byte, 1024)

	c.Read(Https)

	//握手
	ShakeMeg := HttpUntity.Handshake(Https)

	h.hobj.Meg <- ShakeMeg

	//握手成功调用OnConnect回调

	Wscliobj := connection.WsCli{}

	Wscliobj.SetConn(c)

	ConnFunc(Wscliobj)

}
