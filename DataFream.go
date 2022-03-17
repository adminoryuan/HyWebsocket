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

	Makeing_Key []byte //消息掩码

	PlayLoadData  []byte //消息载体 消息体和扩展数据
	ExtensionData []byte //扩展数据

}

//解析数据帧
func DecodeDataFream(meg []byte) DataFream {
	fmt.Println(meg)
	index := 0
	d := DataFream{}
	d.Fin = meg[index] >> 7
	d.Rsv = (meg[index]<<1)>>5 == 0
	d.OpCode = (meg[index] << 1) >> 1
	index += 1
	d.Mask = meg[index] >> 7
	d.PayLoadLenth = (meg[index] << 1) >> 1
	index += 1
	fmt.Println("---")
	fmt.Print(d.PayLoadLenth)
	fmt.Println("--")

	if d.PayLoadLenth == 126 {
		d.ExtenDedPayLoadLen = int64(BytesToInt(meg[index : index+2]))
		index += 2
	} else if d.PayLoadLenth == 128 {
		d.ExtenDedPayLoadLen = int64(BytesToInt(meg[index : index+4]))
		index += 4
	}
	if d.Mask == 1 {
		fmt.Println("计算maskkey")
		d.Makeing_Key = meg[index : index+4]
		index += 4
	}
	fmt.Println(d.Makeing_Key)
	//有效负载数据
	d.PlayLoadData = meg[index : index+(int(d.PayLoadLenth)+int(d.ExtenDedPayLoadLen))]

	for i, _ := range d.PlayLoadData {
		d.PlayLoadData[i] ^= d.Makeing_Key[i%4]
	}

	return d
}

//生产数据帧
func EnCodingDataFream() {

}
