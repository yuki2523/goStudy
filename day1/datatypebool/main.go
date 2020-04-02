package main

import "fmt"

func main() {
	var b1 bool // 默认是 false
	b2 := true

	// %v 任何类型的值都可以匹配打印出来
	fmt.Printf("b1类型为：%T,b1的值为%v\n", b1, b1) // b1类型为：bool,b1的值为false
	fmt.Println(b2)                          // true
}
