package main

import (
	"fmt"
	"testing"
)

func TestServerTest(t *testing.T) {

	h := NewWebsocket()

	h.StartServer(":9090")

	head := make([]byte, 128)

	h.Read(head)

	fmt.Println(string(head))

}
