package main

import (
	"fmt"
)

const (
	AddMeg    uint32 = 0x00 //附加消息
	TextMeg   uint32 = 0x01 //文本消息
	BinaryMeg uint32 = 0x2  //二进制消息
	PingMeg   uint32 = 0x9
	PongMeg   uint32 = 0xA
)

type DataFream struct {
	Fin    byte //是否为最后一位
	Rsv    bool //默认为0 不为0表示出错
	OpCode byte //消息类型

	Mask byte //是否掩码

	PayLoadLenth       byte  //消息长度
	ExtenDedPayLoadLen int64 //扩展长度

	Makeing_Key string //消息掩码

	PlayLoadData  []byte
	ExtensionData []byte
}

//将byte 数组转换为int
func ByteArrayToint(n []byte) uint32 {

	return TextMeg
}

//解析数据帧
func DecodeDataFream(meg []byte) DataFream {

	index := 0
	d := DataFream{}
	d.Fin = meg[index] >> 7
	d.Rsv = (meg[index]<<1)>>5 == 0
	d.OpCode = (meg[index] << 1) >> 1
	index += 1
	d.Mask = meg[index] >> 7
	d.PayLoadLenth = (meg[index] << 1)
	index += 1
	if d.PayLoadLenth == 126 {
		d.ExtenDedPayLoadLen = int64(ByteArrayToint(meg[index : index+2]))
		index += 2
	} else if d.PayLoadLenth == 128 {
		d.ExtenDedPayLoadLen = int64(ByteArrayToint(meg[index : index+4]))
		index += 4
	}
	if d.Mask == 1 {
		d.Makeing_Key = string(meg[index : index+4])
		index += 4
	}

	fmt.Println(d.PayLoadLenth)
	fmt.Println(d.ExtenDedPayLoadLen)
	//有效负载数据
	d.PlayLoadData = meg[index : index+(int(d.PayLoadLenth)+int(d.ExtenDedPayLoadLen))]

	fmt.Println(string(d.PlayLoadData))
	return d
}

//生产数据帧
func EnCodingDataFream() {

}
