package main

import (
	"fmt"
	"gopractice/projects/testdemo/goconveydemo"
)

func main(){
	if goconveydemo.IsEqual(1,1){
		fmt.Printf("true")
	} else{
		fmt.Printf("false")
	}
}
