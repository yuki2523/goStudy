package main

import "fmt"

func main() {
	var num1 float64
	num2 := 100
	// num1 = num2 // 这明显是要报错的，因为类型不对，float64怎么接收int的值呢
	// 这时就可以用到强制类型转换
	num1 = float64(num2)
	fmt.Printf("num1的值为%f,类型为%T", num1, num1) // num1的值为100.000000,类型为float64
}
