# HyWebsocket
- # 介绍
  - 使用golang 开发的一个websoket 库
  - ![golang](https://img.shields.io/badge/golang-1.16.5-red) ![Mit](https://img.shields.io/badge/Mit-Passing-yellow)
- # 具体功能

- # 实例
```golang

func main() {
	h := NewWebsocket()

	h.OnConnect(func(ic connection.IWsCli) {
		//ic.Write([]byte("heelo"))
		fmt.Printf("链接成功 \n")
	})

	h.onReadEvent(func(c ctx.Context) {
		fmt.Printf("recv %s \n", string(c.Req.Bodys))

		c.Resp.Write([]byte("zhangsan"))

	}) 
	h.StartServer(":9091")
}
```
