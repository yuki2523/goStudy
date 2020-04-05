package main

import "fmt"

func f1(x int, y string) (int, string) {
	return x, y
}

func f2() {
	fmt.Println("f2()")
}

func f3() int {
	return 10
}

func f4(x int) {
	fmt.Println(x)
}

// 这个函数接收一个函数类型的参数,类型为 func(int) int
// 返回一个函数类型的result
func f5(function func(int) int) (result func(int) int) {
	result = func(x int) int {
		x = function(x) * 10
		return x
	}
	return
}

func f6(x int) int {
	return x + 10
}

func main() {
	fmt.Printf("f1的类型为:%T\n", f1) // f1的类型为:func(int, string) (int, string)
	fmt.Printf("f2的类型为:%T\n", f2) // f2的类型为:func()
	fmt.Printf("f3的类型为:%T\n", f3) // f3的类型为:func() int
	fmt.Printf("f4的类型为:%T\n", f4) // f4的类型为:func(int)

	addAndMultiply := f5(f6)        // addAndMultiply这个函数 + 10 再 * 10
	fmt.Println(addAndMultiply(10)) // 200
}
