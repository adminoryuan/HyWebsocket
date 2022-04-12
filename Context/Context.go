package context

//封装统一上下文
type Context struct {
	Req RequestConn

	Resp WebsocketResp
}
