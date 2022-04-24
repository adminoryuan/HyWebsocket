package fream

//mask _key 解码
func Makeing_Key(make_key []byte, data []byte) []byte {
	for i, _ := range data {
		data[i] ^= make_key[i%4]
	}
	return data
	//
}
