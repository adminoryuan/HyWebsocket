package main

func TestDataFream() {
	//1
	DataFreamByteTest := []byte{129, 132, 41, 84, 68, 132, 24, 101, 119, 183}

	DecodeDataFream(DataFreamByteTest)
	//fmt.Println(string(d.PlayLoadData))
}
