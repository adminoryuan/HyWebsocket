package main

import (
	connection "Hywebsocket/Connection"
	request "Hywebsocket/Request"
)

type ReadEventFunc func(request.RequestConn)
type OnConnFunc func(connection.IWsCli)
type OnCloseFunc func()

// 对外公布的websocket接口
type Websocket interface {
	StartServer(port string)

	OnConnect(OnConnFunc) //返回一个链接对象

	OnClose()
	//收到数据时事件
	onReadEvent(fun ReadEventFunc)
}
