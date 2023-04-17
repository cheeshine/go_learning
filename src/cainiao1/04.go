package main

import (
	"fmt"
)

func main() {
	var a byte = '1'
	var b rune = '一'
	fmt.Println(a, b)
	fmt.Printf("%T", b)
}
