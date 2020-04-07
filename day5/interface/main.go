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

}
