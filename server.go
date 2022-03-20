package main

import (
	"Hywebsocket/http"
	"context"
	"fmt"
	"net"
)

var HttpUntity http.HttpUntity = http.HttpUntity{}

type Hwebsocket struct{}

func (h *Hwebsocket) startServer(port string) {

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

//与客户端进行握手 返回握手包
func (h *Hwebsocket) Handshake(https []byte) []byte {

	HttpMaps := HttpUntity.AnalyHttp(https)

	//计算key
	Sce_Rpaly_Key := HttpUntity.EncodeSecWebsocketKey(HttpMaps["Sec-WebSocket-Key"])

	//构造响应http协议
	ResponseString := "HTTP/1.1 101 Switching Protocols \r\n"
	ResponseString += "Upgrade: websocket \r\n"
	ResponseString += "Connection: Upgrade \r\n"
	ResponseString += "Sec-WebSocket-Accept: " + Sce_Rpaly_Key + "\r\n"
	//ResponseString += "Sec-WebSocket-Protocol: chat"
	ResponseString += "\r\n"
	fmt.Println(ResponseString)
	return []byte(ResponseString)
}
func (h *Hwebsocket) dispServes(c net.Conn) {

	//defer c.Close()

	Hobj := NewDispMessage()

	Ctx, cal := context.WithCancel(context.Background())

	Hobj.Canle = cal

	go Hobj.OnWrite(c, Ctx)

	go Hobj.onRead(c, Ctx)

	Https := make([]byte, 1024)

	c.Read(Https)

	ShakeMeg := h.Handshake(Https)

	Hobj.Meg <- ShakeMeg

	Writeshake := DataFream{}
	Writeshake.Fin = 1
	Writeshake.Rsv = true
	Writeshake.OpCode = 0x01
	Writeshake.PlayLoadData = []byte("你好啊张三")

	Writeshake.PayLoadLenth = byte(len(Writeshake.PlayLoadData))

	a := NewDataFreamCoding()
	Hobj.Meg <- a.EnCodingDataFream(Writeshake)

	Hobj.Meg <- a.EnCodingDataFream(Writeshake)
}
