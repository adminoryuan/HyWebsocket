package HyWebsocket

type OnConnFunc func(IWsCli)
type OnCloseFunc func(addr string)

// 对外公布的websocket接口
type Websocket interface {
	StartServer(port string)

	OnConnect(OnConnFunc) //返回一个链接对象

	OnClose(OnCloseFunc)
	//收到数据时事件
	OnReadEvent(fun ReadEventFunc)
}
