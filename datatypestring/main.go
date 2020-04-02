package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1、转义(经典套路)是 \
	var s1 string = "\"就要打印一下双引号\""
	var s2 string = "打印反斜杠\\"
	fmt.Println(s1) // "就要打印一下双引号"
	fmt.Println(s2) // 打印反斜杠\

	// 2、模板字符串 ``
	s3 := `
		<h1>有点类型js里的模板字符串</h1>
		<p>看看打印出来的结果吧</p>
	`
	fmt.Println(s3) // 很像js里的模板字符串，可惜不支持${}

	// <h1>有点类型js里的模板字符串</h1>
	// <p>看看打印出来的结果吧</p>

	// 3、一些方法
	name := "hcy"
	age := "19"
	test := "Hello Go Python Javascript"
	// len():这个比较通用
	fmt.Println(len(name)) // 3

	// fmt.Sprintf():拼接字符串
	me1 := fmt.Sprintf("%s %s", name, age)
	fmt.Println(me1) // hcy 19

	// + 加号拼接字符串
	me := "名字：" + name + "年龄：" + age
	fmt.Println(me) // 名字：hcy年龄：19

	// strings.Split():分割字符串为数组
	// 第一个参数是指定分割的字符串。第二个是指定以什么去分割
	testList := strings.Split(test, " ") // [Hello Go Python Javascript]
	fmt.Println(testList)                // 用空格分割字符串test

	// strings.contains():是否包含，返回bool值
	fmt.Println(strings.Contains(test, "Go"))   // true
	fmt.Println(strings.Contains(test, "1223")) // false

	// strings.HasPrefix():判断开头，strings.HasSuffix():判断结尾
	fmt.Println(strings.HasPrefix(test, "Hell"))   // true
	fmt.Println(strings.HasSuffix(test, "script")) // true

	// strings.Index():求索引,strings.LastIndex():最后一次出现的索引
	fmt.Println(strings.Index(test, "Go"))    // 6
	fmt.Println(strings.Index(test, "o"))     // 4
	fmt.Println(strings.LastIndex(test, "o")) // 13

	// strings.Join()拼接
	fmt.Println(strings.Join(testList, "++")) // Hello++Go++Python++Javascript
}
