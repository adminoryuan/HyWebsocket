package untity

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
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
	fmt.Println("target", req_scr_key)

	req_scr_key = strings.ReplaceAll(req_scr_key, " ", "")
	req_scr_key = strings.ReplaceAll(req_scr_key, "0x00", "")
	req_scr_key = req_scr_key + "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"

	s := sha1.New()
	s.Write([]byte(req_scr_key))
	newSck := s.Sum(nil)

	m := base64.StdEncoding.EncodeToString([]byte(newSck))

	return m
}
