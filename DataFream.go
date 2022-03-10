package main

const (
	AddMeg    int = 0x00 //附加消息
	TextMeg   int = 0x01 //文本消息
	BinaryMeg int = 0x2  //二进制消息
	PingMeg   int = 0x9
	PongMeg   int = 0xA
)

type DataFream struct {
	Fin    bool    //是否为最后一位
	Rsv    [4]byte //默认为0
	OpCode uint32  //消息类型

	Make byte //是否掩码

	PayLoadLenth int64 //消息长度

	Makeing_Key int64 //消息掩码

}

//解析数据帧
func DecodeDataFream(meg []byte) DataFream {
	data := DataFream{}
	
	
	return data
}

//生产数据帧
func EnCodingDataFream() {

}
