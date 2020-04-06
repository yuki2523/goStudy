package main

import "fmt"

// 现在这个场景为给每个 string 类型的变量添加方法
// 让它们都可以打招呼
// 但是 string 并非当前 main 包里的方法,无法直接使用 (s *string) 这样的函数去定义方法
// 那么 自定义类型 就可以帮我们达成这个需求

type myString string

func (m myString) hello() {
	fmt.Println("Hello I'm", m)
}

func main() {
	var str = myString("测试字符串")
	str.hello() // Hello I'm 测试字符串
}
