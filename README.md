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
	"sync"

	"github.com/adminoryuan/HyWebsocket"
	Webcontext "github.com/adminoryuan/HyWebsocket/WebContext"
)

func main() {
	server := HyWebsocket.NewWebsocket()
	var GloabConn sync.Map = sync.Map{}

	server.OnConnect(func(ic HyWebsocket.IWsCli) {
		GloabConn.Store(ic.GetRemoterAddr().String(), ic)
		mes := ic.GetRemoterAddr().String()
		GloabConn.Range(func(key, value interface{}) bool {
			value.(HyWebsocket.IWsCli).Write([]byte(mes + "上线了"))
			return true
		})
		fmt.Printf("收到 ：%s 的链接 \n", ic.GetRemoterAddr().String())
	})
	server.OnReadEvent(func(ctx Webcontext.Context) {
		fmt.Printf("%s", string(ctx.Req.Bodys))
	})
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
    <div id="userList">
        

    </div>
</head>
<body>
    <button id="app">点击</button>
    <script>
      
        
            var socket=new WebSocket('ws://127.0.0.1:9999')
            socket.onopen=function(){
               // alert("链接成功") 
              
                socket.send("你好123123")
            }
            socket.close=function(){
                alert("断开了链接")
            }
            socket.onmessage=function(res){
               
                alert(res.data)
                //  socket.send("你好123123")
            }
            socket.onerror=function(){
               // alert("出错了")
            }
            socket.setConnectionLostTimeout(3000)
            
        var btn =document.getElementById("app")
        btn.onclick=function(){
            socket.close()
            
        }
            </script>
    
</body>
</html>

```

- # 具体功能
