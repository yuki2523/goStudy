package main

import (
	"fmt"
	"strings"
)

// 1、让闭包里引用的外层作用域的变量自增
func increment(x int) func() int {
	var count = 0
	return func() int {
		count++
		fmt.Printf("count:%d,", count)
		return count + x
	}
}

// 2、添加文件后缀
func makeSuffixFunc(suffix string) func(string) string {
	return func(fileName string) string {
		if !strings.HasSuffix(fileName, suffix) {
			return fileName + suffix
		}
		return fileName
	}
}

// 3、和第一个例子近似，一个公用的闭包外的变量 base 的变化过程
// base 这个变量，只要下面的 addBase, subBase != nil，也就是没有释放，它就永远存活着
func calc(base int) (func(int) int, func(int) int) {
	var add = func(num1 int) int {
		base += num1
		return base
	}
	var sub = func(num2 int) int {
		base -= num2
		return base
	}
	return add, sub
}

func main() {
	// 1、第一个例子的测试
	var add = increment(10)
	fmt.Println("count+10 =", add()) // count:1,count+10 = 11
	fmt.Println("count+10 =", add()) // count:2,count+10 = 12
	fmt.Println("count+10 =", add()) // count:3,count+10 = 13
	fmt.Println("count+10 =", add()) // count:4,count+10 = 14

	// 2、第二个例子的测试
	var jpgSuffix = makeSuffixFunc(".jpg")
	var txtSuffix = makeSuffixFunc(".txt")
	var fileName1 = "1.jpg"
	var fileName2 = "aaa"
	var fileName3 = "cc.txt"
	var fileName4 = "ddd"
	fmt.Println(jpgSuffix(fileName1)) // 1.jpg
	fmt.Println(jpgSuffix(fileName2)) // aaa.jpg
	fmt.Println(txtSuffix(fileName3)) // cc.txt
	fmt.Println(txtSuffix(fileName4)) // ddd.txt

	// 3、第三个例子的测试
	var addBase, subBase = calc(10) // 给base为10
	// base + 1 - 2 : 10 => 11 => 9
	fmt.Println(addBase(1), subBase(2)) // 11 9
	// base + 3 - 4 : 9 => 12 => 8
	fmt.Println(addBase(3), subBase(4)) // 12 8
	// base + 5 - 6 : 8 => 13 => 7
	fmt.Println(addBase(5), subBase(6)) // 13 7
}
