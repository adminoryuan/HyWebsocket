package main

// 对外公布的websocket接口
type Websocket interface {
	StartServer(port string)

	Read(data []byte) error

	Write(data []byte) error
}
