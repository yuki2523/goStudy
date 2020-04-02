package main

import "fmt"

func main() {
	// 下面是复数的定义法，就没见过用这种东西
	var c1 complex64 = 1 + 2i  // complex64的实部和虚部都是32位的
	var c2 complex128 = 2 + 3i // complex128的实部和虚部都是64位的

	fmt.Println(c1) // (1+2i)
	fmt.Println(c2) // (2+3i)
}
