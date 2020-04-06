package main

import "fmt"

// type关键字用来造类型
// 由内置 int 造了个 myInt 类型
type myInt int

// 类型别名
// 加上 = 号的为类型别名
type ageInt = int

func main() {
	var a myInt = 100
	fmt.Printf("a,value:%d,type:%T\n", a, a) // a,value:100,type:main.myInt
	// 类型为 main 包里的 myInt 类型

	var age ageInt = 19
	fmt.Printf("age,value:%d,type:%T\n", age, age) // age,value:19,type:int
	// 类型别名,在代码编写中是有效的,代码执行时会换成其真正的名字,这里的ageInt就换成了int
	// 内置别名 byte 就是 int8, rune 就是 int32
}
