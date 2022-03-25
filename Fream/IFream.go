package fream

//对外公布的接口
type Fream interface {
	DecodeDataFream(meg []byte) DataFream
	EnCodingDataFream(f DataFream) []byte
}
