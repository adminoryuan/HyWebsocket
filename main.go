package main

import (
	connection "Hywebsocket/Connection"
	ctx "Hywebsocket/WebContext"
	"fmt"
)

func solve(arr []int, n, m int) int {
	// 请添加具体实现
	left := 0
	right := n - 1
	count := 0
	for left < right-count {
		if arr[left]+arr[right] == m {
			count += 1
			left += 1
		}
		if right == left {
			left += 1
			right = n - count - 1
		}
		right -= 1
	}
	return count
}

func main() {

	l := solve([]int{1, 3, 4, 5, 6}, 5, 5)
	fmt.Printf("%d", l)

	h := NewWebsocket()

	h.OnConnect(func(ic connection.IWsCli) {
		fmt.Printf("收到链接： %s \n", ic.GetRemoterAddr().String())
		ic.Write([]byte("heelo"))

	})

	h.onReadEvent(func(c ctx.Context) {
		fmt.Printf("recv %s \n", string(c.Req.Bodys))

		c.Resp.Write([]byte("zhangsan"))

	})
	h.StartServer(":9091")
}
