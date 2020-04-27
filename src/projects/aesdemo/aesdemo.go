package main

import (
	"fmt"
	"gopractice/common"
)

func main() {
	order := &Order{
		Id:   111,
		Name: "aaa",
	}

	encode, e := common.DataAesEncrypt(order, "aCqonzvpalooWtVT", "oexpressgogo!@#$")

	if e != nil{
		fmt.Printf("error: %s \n", e)
		return
	}

	fmt.Printf("encode string: %s \n", encode)

	decode, e := common.DataAesDecrypt(encode, "aCqonzvpalooWtVT", "oexpressgogo!@#$")

	if e != nil{
		fmt.Printf("error: %s \n", e)
		return
	}

	fmt.Printf("decode string: %s \n", decode)

}

type Order struct {
	Id   int64
	Name string
}
