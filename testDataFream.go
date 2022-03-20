package main

import "fmt"

func TestDataFream() {
	//1
	DataFreamByteTest := []byte{129, 132, 41, 84, 68, 132, 24, 101, 119, 183}

	fmt.Println(DataFreamByteTest)
	//fmt.Println(string(d.PlayLoadData))
}
func EncodingFream() {

	obj := NewDataFreamCoding()

	d := DataFream{}
	d.Fin = 1
	d.Rsv = true
	d.OpCode = 0x01
	Btes := obj.EnCodingDataFream(d)
	fmt.Println(Btes)
}
