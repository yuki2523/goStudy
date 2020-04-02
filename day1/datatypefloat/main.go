package main

import (
	"fmt"
)

func main() {
	a := 1.23456                       // 默认是 float64 类型
	fmt.Printf("a的类型为%T,值为%f\n", a, a) // a的类型为float64,值为1.234560

	b := float32(1.23456)              // 想要用float32，就要明确写出来
	fmt.Printf("b的类型为%T,值为%f\n", b, b) // b的类型为float32,值为1.234560

	// b = a // 这是不行的，类型不同 float32和float64是完全不相同的两种类型
}
