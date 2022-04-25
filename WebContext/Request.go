package Webcontext

import (
	"encoding/json"
	"net"
)

//描述一个用户端发送数据
type RequestConn struct {
	LocalRemoter net.IP

	Bodys []byte
}

func (r *RequestConn) Bind(o interface{}) error {
	return json.Unmarshal(r.Bodys, o)
}
