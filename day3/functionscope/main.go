package main

import "fmt"

var a = 100

func f1() {
	var c = 300
	fmt.Println(a) // 100
	// fmt.Println(b) // 还是找不到
	fmt.Println(c)

}

func main() {
	// var b = 200
	f1()

	// if 和 for 有块作用域的特性
	if i := 10; i < 18 {
		fmt.Println("快去上学")
	}
	// fmt.Println(i) // undefined

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
	// fmt.Println(j) // undefined

	if true {
		// var x = 10
	}
	// fmt.Println(x) // undefined
}
