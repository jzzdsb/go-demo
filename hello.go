package main

import (
	"fmt"
	//引入自己的包
	"demo/constant"
)

func main() {
	//fmt.Println("hello World")
	fmt.Println(constant.NAME)

	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for _,x:=range balance{
		fmt.Println(x)
	}
}
