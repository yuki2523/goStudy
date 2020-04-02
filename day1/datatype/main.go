package main

import "fmt"

func main() {
	// 这玩意不能直接指定二进制的数
	// 十进制
	var a int = 10
	fmt.Printf("%d\n", a) // 十进制在Printf中使用 %d 显示, 结果为：10
	fmt.Printf("%b\n", a) // 转换为二进制 %b, 结果为：1010
	fmt.Printf("%o\n", a) // 转换为八进制 %o, 结果为：12
	fmt.Printf("%x\n", a) // 转换为十六进制 %x, 结果为：a

	// 八进制的定义方式如下,0开头
	b := 021
	fmt.Printf("%d\n", b) // 转换为十进制的结果为：17
	// 十六进制定义方式,0x开头
	c := 0x1d
	fmt.Printf("%d\n", c) // 转换为十进制的结果为：29

	// 查看类型 %T ,还有go也是支持强制类型转换的 int8(10), int16(10)这样的
	d := int8(10)
	e := int16(10)
	f := 10
	fmt.Printf("%T\n", d) // int8
	fmt.Printf("%T\n", e) // int16
	fmt.Printf("%T\n", f) // int
}
