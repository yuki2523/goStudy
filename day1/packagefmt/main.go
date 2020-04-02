package main

import "fmt"

func main() {
	// fmt占位符总结
	var n int = 100
	var s string = "go python javascript,未完待续..."

	fmt.Printf("%T\n", n) // %T 数据类型,结果为：int
	fmt.Printf("%T\n", s) // 结果为：string

	fmt.Printf("%b\n", n) // %b 二进制数,结果为：1100100
	fmt.Printf("%d\n", n) // %d 十进制,结果为：100
	fmt.Printf("%o\n", n) // %o 八进制,结果为：144
	fmt.Printf("%x\n", n) // %x 十六进制,结果为：64

	fmt.Printf("%s\n", s)  // %s 字符串,结果为：go python javascript,未完待续...
	fmt.Printf("%v\n", s)  // %v 任何类型的值都可以匹配,结果为：go python javascript,未完待续...
	fmt.Printf("%#v\n", s) // %#v 对那个类型的变量的值做一些修饰(字符串会多出""),结果为："go python javascript,未完待续..."
	fmt.Printf("%#v\n", n) // (整型好像没什么变化)结果为：100
}
