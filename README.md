# HyWebsocket (websocket开发库)
- # 介绍
  - 使用golang 开发的一个websoket 库
  - ![golang](https://img.shields.io/badge/golang-1.16.5-red) ![Mit](https://img.shields.io/badge/Mit-Passing-yellow) 
- # 安装
```bash
 go mod init test
 go get github.com/adminoryuan/HyWebsocket v1.0.4
```
- # 快速入门
- 服务端
``` golang
package main

import (
	"fmt"

	"github.com/adminoryuan/HyWebsocket"
	Webcontext "github.com/adminoryuan/HyWebsocket/WebContext"
)

func main() {
	server := HyWebsocket.NewWebsocket()
	//链接时的回调
	server.OnConnect(func(ic HyWebsocket.IWsCli) {
		fmt.Printf("收到 ：%s 的链接 \n", ic.GetRemoterAddr().String())
	})
	//接收到消息的回调
	server.OnReadEvent(func(ctx Webcontext.Context) {
		fmt.Printf("%s", string(ctx.Req.Bodys))
	})
	//启动服务
	server.StartServer(":9999")
}

```
- 到此服务器已经启动成功
```客户端
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>

</head>
<body>
    <button id="app">点击</button>
    <script>
        console.log("asdsa")
        
            var socket=new WebSocket('ws://127.0.0.1:9999')
            socket.onopen=function(){
               alert("链接成功") 
              
                socket.send("你好")
            }
            socket.close=function(){
                alert("断开了链接")
            }
            socket.onmessage=function(res){
                console.log(res.data)
            //  socket.send("你好")
            }
            socket.onerror=function(){
               // alert("出错了")
            }
            socket.setConnectionLostTimeout(3000)
            
        var btn =document.getElementById("app")
        btn.onclick=function(){
            alert("..")
            socket.close()
            
        }
            </script>
    
</body>
</html>

```

- # 具体功能
 - #  OnConnect() 回调 返回一个链接对象
 - #  OnReadEvent(func(ctx Webcontext.Context))  //Context 描述了一个请求对象(客户端ip,信息消息体) 相应对象(json 等一些影响函数) 

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
