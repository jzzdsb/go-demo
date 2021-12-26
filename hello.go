package main

import (
	"fmt"
	//引入自己的包
	"demo/constant"
)

func init(){
	fmt.Print("start func")
}

func main() {
	//fmt.Println("hello World")
	fmt.Println(constant.NAME)

	var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	balance = append(balance, 9)

	//循环
	for _,x:=range balance{
		//判断
		if x>8 {
			fmt.Println(x)
		}
	}

	//使用map
	var maps=make(map[string]string)

	maps["one"]=""
	maps["two"]="你"
	maps["three"]="妈"

	for v:=range maps{
		fmt.Print(v,maps[v])
	}
	
}
