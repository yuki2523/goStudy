package main

import "fmt" // 导包

// 函数外面只能放一些变量常量的声明，逻辑必须放在代码块里

// 声明变量：声明的名字建议是驼峰命名法
var (
	name        string
	age         int
	isOK        bool
	studentName string // 这种小驼峰格式的，和js一样
)

func foo() (int, string) {
	return 10, "hahaha"
}

func main() { // 入口函数
	// fmt.Println("Hello World")
	name = "hcy"
	age = 19
	isOK = true
	studentName = "huachenyang"
	// go 里变量声明了就要用，不用不给过编译
	fmt.Printf("name:%s", name) // 加了f的是格式化输出，经典%s占位符
	fmt.Println(age)            // 这个是默认加换行的
	fmt.Print(isOK)             // 这个就是单纯的打印
	fmt.Println(studentName)

	// 声明的同时赋值
	var s1 string = "yuki"
	var a1 int = 19
	fmt.Printf("name:%s,age:%d", s1, a1)

	// 类型推导
	var s2 = "haru" // 它可以自己根据值分析自己的类型
	fmt.Println(s2)

	// 简短变量声明
	s3 := "Hello" // 这个就和上面的类型推导相同，不过只能在函数里面用
	fmt.Println(s3)

	// 匿名变量：_
	// 它不占内存空间，不占命名空间，所以也不存在重复命名这回事
	// 主要用来占位，表示忽略值
	x, _ := foo() // 这种来只要一个值的，不要的值就让 _ 接收
	fmt.Println(x)
}
