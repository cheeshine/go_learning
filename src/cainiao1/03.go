package main

import "fmt"

// 这种不带声明格式的只能在函数体中出现
//g, h := 1, 2

func main() {
	var balance [10]float32
	balance[0] = 1
	var balance1 = [2]float32{0, 1}
	var balance2 = [...]float32{0, 1}
	fmt.Println(balance)
	fmt.Println(balance1)
	fmt.Println(balance2)
}
