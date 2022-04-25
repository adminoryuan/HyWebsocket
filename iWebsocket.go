package main

type OnConnFunc func(IWsCli)
type OnCloseFunc func()

// 对外公布的websocket接口
type Websocket interface {
	StartServer(port string)

	OnConnect(OnConnFunc) //返回一个链接对象

	OnClose()
	//收到数据时事件
	onReadEvent(fun ReadEventFunc)
}
