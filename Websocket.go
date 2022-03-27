package main

import (
	connection "Hywebsocket/Connection"
)

type ReadEventFunc func([]byte)
type OnConnFunc func(connection.IWsCli)

// 对外公布的websocket接口
type Websocket interface {
	StartServer(port string)

	OnConnect(OnConnFunc) //返回一个链接对象

	//收到数据时事件
	onReadEvent(fun ReadEventFunc)
}