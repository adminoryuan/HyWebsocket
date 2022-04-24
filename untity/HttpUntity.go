package untity

import (
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

type HttpUntity struct{}

//解析http协议q
//reqByte 请求的http流
func (t *HttpUntity) AnalyHttp(reqByte []byte) map[string]string {
	https := strings.Split(string(reqByte), "\r\n")

	index := 1 //跳过第一行 http 1.0
	HtMaps := make(map[string]string, 20)
	var line string
	for index < len(https) {
		line = https[index]

		kvHttp := strings.Split(line, ":")
		///fmt.Println(line)
		if len(kvHttp) < 2 {
			index++

			continue
		}
		HtMaps[kvHttp[0]] = kvHttp[1]
		index++
	}

	return HtMaps
}

// 计算key
func (t *HttpUntity) EncodeSecWebsocketKey(req_scr_key string) string {
	//guid := tsgutils.GUID()

	req_scr_key = strings.ReplaceAll(req_scr_key, " ", "")
	req_scr_key = strings.ReplaceAll(req_scr_key, "0x00", "")
	req_scr_key = req_scr_key + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

	s := sha1.New()
	s.Write([]byte(req_scr_key))
	newSck := s.Sum(nil)

	m := base64.StdEncoding.EncodeToString([]byte(newSck))

	return m
}

//与客户端进行握手 返回握手包
func (t *HttpUntity) Handshake(https []byte) []byte {

	HttpMaps := t.AnalyHttp(https)

	//计算key
	Sce_Rpaly_Key := t.EncodeSecWebsocketKey(HttpMaps["Sec-WebSocket-Key"])

	//构造响应http协议
	ResponseString := "HTTP/1.1 101 Switching Protocols \r\n"
	ResponseString += "Upgrade: websocket \r\n"
	ResponseString += "Connection: Upgrade \r\n"
	ResponseString += "Sec-WebSocket-Accept: " + Sce_Rpaly_Key + "\r\n"
	//ResponseString += "Sec-WebSocket-Protocol: chat"
	ResponseString += "\r\n"

	return []byte(ResponseString)
}
