package main

type ReadEventFunc func([]byte)

// 对外公布的websocket接口
type Websocket interface {
	StartServer(port string)

	//收到数据时事件
	onReadEvent(fun ReadEventFunc)
}
