package main

import "fmt"

func calc(index string, a, b int) int {
	var result = a + b
	fmt.Println(index, a, b, result)
	return result
}

func main() {
	var a = 1
	var b = 2
	defer calc("1", a, calc("10", a, b)) // 算完参数后这步被先压入栈中
	// 10 1 2 3 // 参数要先被算出来，所以到这步就执行了calc("10", a, b) = 3

	a = 0
	defer calc("2", a, calc("20", a, b)) // 算完参数后这步被再压入栈中
	// 20 0 2 2 // 参数要先被算出来，所以到这步就执行了calc("20", a, b) = 2

	b = 1
	// 接下来开始执行栈里的代码
	// 先执行 calc("2", a, calc("20", a, b)) => calc("2", 0, 2)
	// 2 0 2 2
	// 再执行 calc("1", a, calc("10", a, b)) => calc("1", 1, 3)
	// 1 1 3 4
}

// main函数的结果,就如上面的分析一样
// 10 1 2 3
// 20 0 2 2
// 2 0 2 2
// 1 1 3 4
