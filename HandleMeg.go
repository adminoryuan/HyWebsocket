package main

import (
	"context"
	"fmt"
	"io"
)

type dispCliMessage struct {
	Meg   chan []byte
	Canle context.CancelFunc
}

func NewDispMessage() dispCliMessage {
	d := dispCliMessage{}
	d.Meg = make(chan []byte)
	return d

}
func (d dispCliMessage) onRead(c io.Reader, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:

			Mes := make([]byte, 128)
			c.Read(Mes)

			cliFream := DecodeDataFream(Mes)

			fmt.Printf(string(cliFream.PlayLoadData))

		}
		//c.Write(nes)
	}
}
func (d dispCliMessage) OnWrite(w io.Writer, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("关闭了")
			return
		case a := <-d.Meg:
			fmt.Println("接收到数据")
			w.Write(a)
		}
	}
}
