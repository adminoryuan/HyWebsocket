package main

import (
	"Hywebsocket/http"
	"fmt"
	"io"
	"net"
)

var HttpUntity http.HttpUntity = http.HttpUntity{}

func main() {

	conn, err := net.Listen("tcp", ":800")
	if err != nil {
		panic(err)
	}
	for {
		cli, err := conn.Accept()
		if err != nil {
			fmt.Println(err.Error())

		}
		fmt.Println("收到请求")
		go dispServes(cli, cli)
	}
}

func dispServes(c io.Reader, w io.Writer) {

	Https := make([]byte, 1024)
	c.Read(Https)

	HttpMaps := HttpUntity.AnalyHttp(Https)

	Sce_Rpaly_Key := HttpUntity.EncodeSecWebsocketKey(HttpMaps["Sec-WebSocket-Key"])

	ResponseString := "HTTP/1.1 101 Switching Protocols \r\n"
	ResponseString += "Upgrade: websocket \r\n"
	ResponseString += "Connection: Upgrade \r\n"
	ResponseString += "Sec-WebSocket-Accept:" + Sce_Rpaly_Key + "\r\n"
	ResponseString += "Sec-WebSocket-Protocol: chat \r\n"

	w.Write([]byte(ResponseString))


}
