package main

import "fmt"

// 1、这个接口，只能传一个无参数有一个 int 类型返回值的函数进去
func f1(function func() int) {
	var sumResult = function()
	fmt.Println(sumResult)
}

// 2、现在我手上有这样一个函数，它接收两个参数，返回一个值
func sum(x, y int) int {
	return x + y
}

// 3、我希望把 sum 函数传到 f1 里运行，但是函数类型不匹配
// 		这个时候，可以使用闭包封装一下
func closeSum(sum func(int, int) int, x, y int) func() int {
	var functionSum = func() int {
		return sum(x, y)
	}
	return functionSum
}

// 4、这样返回的函数体 functionSum 满足了一个 int 类型返回值无参数的 func() int 类型

func main() {
	// 5、这个时候，就可以愉快的调用接口 f1 了
	var functionSum = closeSum(sum, 10, 20)
	f1(functionSum) // 30
}
