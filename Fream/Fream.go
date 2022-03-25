package fream

type Fream interface {
	DecodeDataFream(meg []byte) DataFream
	EnCodingDataFream(f DataFream) []byte
}
