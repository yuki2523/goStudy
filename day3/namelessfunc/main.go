package main

import "fmt"

// 1、匿名函数
// 这用法和js一样
var f1 = func() {
	fmt.Println("匿名函数,用变量f1接收")
}

// 3、匿名函数的作用域

var a = 100

func f2() {
	var b = 200
	func() {
		var c = 300
		fmt.Println(a) // 100
		fmt.Println(b) // 200
		fmt.Println(c) // 300
		// fmt.Println(d) // undefined
		// 反正不会去 main 里找变量 d
	}()
}

func main() {
	f1() // 匿名函数,用变量f1接收

	// 2、立即执行函数
	var x = 10
	func(x int) {
		fmt.Println(x * 2) // 20
	}(x)

	// var d = 400
	f2()
}
