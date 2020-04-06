package main

import "fmt"

func main() {
	var m1 = make(map[string]int, 5)
	m1["hcy"] = 19
	// %#v 打印时会带上 Go 的语法
	fmt.Printf("%#v\n", m1) // map[string]int{"hcy":19}
	fmt.Printf("%v\n", m1)  // map[hcy:19]

	// %% 表示一个 %
	var a = 90
	fmt.Printf("%d%%\n", a) // 90%

	var num = 65
	// %c 值对应的 Unicode 字符
	fmt.Printf("%c\n", num) // A
	// %q 该值对应的 '' 括起来的,打印字符串时也会带上 ""
	// %q 转 Unicode, %+q 转ASCII码
	fmt.Printf("%q\n", num)         // 'A'
	fmt.Printf("%q\n", "我想要我想要的未来") // "我想要我想要的未来"

	var num1 = 123456789.0
	var num2 = 1.23456789
	var num3 = 1.23
	// %e：科学计算法，%f：小数
	fmt.Printf("%e\n", num1) // 1.234568e+08
	fmt.Printf("%e\n", num2) // 1.234568e+00

	// %f 的格式化输出 %p.qf ,p为总宽度,q为精度
	fmt.Printf("%.3f\n", num2) // 1.235
	// 前面使用空格填充
	fmt.Printf("%020.2f\n", num2) // 00000000000000001.23
	fmt.Printf("%10.2f\n", num2)  //       1.23

	// %g 根据实际情况决定以 %e 还是 %f 格式输出
	// %g 自动根据数据决定输出方式，是一种比较精确的输出方式
	fmt.Printf("%g\n", num1) // 1.23456789e+08
	fmt.Printf("%g\n", num2) // 1.23456789
	fmt.Printf("%g\n", num3) // 1.23

	// 指针 %p
	var x = 10
	fmt.Printf("%p %x\n", &x, &x) // 0xc000010150 c000010150

	// 输入,Scan,Scanf,Scanln
	var str string
	var str1 string
	var str2 string
	fmt.Scan(&str)   // 可以输入字符串
	fmt.Println(str) // 可以输入字符串
	fmt.Scanf("%s\n", &str1)
	fmt.Println(str1)
	fmt.Scanln(&str2)
	fmt.Println(str2)
}
