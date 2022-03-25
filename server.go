package main

import (
	handle "Hywebsocket/Handle"
	http "Hywebsocket/untity"
	"context"
	"fmt"
	"net"
)

var HttpUntity http.HttpUntity = http.HttpUntity{}

type hwebsocket struct {
	hobj handle.DispCliMessage
}

func NewWebsocket() Websocket {
	h := hwebsocket{}
	h.hobj = handle.NewDispMessage()
	return h
}
func (h hwebsocket) StartServer(port string) {

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
		go h.dispServes(cli)
	}
}

func (h hwebsocket) Read(body []byte) error {

	h.hobj.PlayLoadData <- body

	return nil
}

func (h hwebsocket) Write(meg []byte) error {

	h.hobj.Meg <- meg

	return nil
}

func (h hwebsocket) dispServes(c net.Conn) {

	//defer c.Close()

	Ctx, cal := context.WithCancel(context.Background())

	h.hobj.Canle = cal

	go h.hobj.OnWrite(c, Ctx)

	go h.hobj.OnRead(c, Ctx)

	Https := make([]byte, 1024)

	c.Read(Https)

	ShakeMeg := HttpUntity.Handshake(Https)

	h.hobj.Meg <- ShakeMeg

}
