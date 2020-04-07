package main

import "fmt"

// multiplicable 这个接口类型,允许实现了 multiply 方法的数据类型使用
type multiplicable interface {
	multiply() // 方法标签
}

type number struct {
	value int
}

// number 类型实现的 multiply 方法
// 打印出结构体里 value 值二倍的结果
func (n number) multiply() {
	fmt.Println(n.value * 2)
}

type text struct {
	value string
}

// text 类型实现的 multiply 方法
// 打印出结构体里 value 值的两倍长度的字符串
func (t text) multiply() {
	fmt.Println(t.value + t.value)
}

func mul(structure multiplicable) {
	structure.multiply()
}

func main() {
	var num1 = number{10}
	var txt1 = text{"abc"}
	fmt.Println(num1, txt1) // {10} {abc}
	mul(num1)               // 20
	mul(txt1)               // abcabc

	var m1 multiplicable          // 这种是被允许的
	fmt.Printf("%v,%T\n", m1, m1) // <nil>,<nil>
	m1 = number{100}
	fmt.Printf("%v,%T\n", m1, m1) // {100},main.number
	m1 = text{"bbb"}
	fmt.Printf("%v,%T\n", m1, m1) // {bbb},main.text

	// var n2 = number{200}
	// fmt.Println(n2)
	// n2 = text{"ccc"} // 报错,类型问题
	// fmt.Println(n2)
}
