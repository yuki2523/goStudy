package main

import "fmt"

// 匿名字段的情况时,直接把类型名识别为了字段名
// 绝对不推荐使用这种东西!!!
type person struct {
	string
	int
	bool
}

type student struct {
	id   int
	name string
}

func main() {
	var p1 = person{
		"hcy",
		19,
		true,
	}
	fmt.Printf("p1:%#v\n", p1) // p1:main.person{string:"hcy", int:19, bool:true}
	fmt.Println(p1.string)     // hcy
	fmt.Println(p1.int)        // 19
	fmt.Println(p1.bool)       // true

	p1.int = 20
	fmt.Println(p1) // {hcy 20 true}

	// 字段名和变量名同名,直接写一个字段名就OK了,这点和JS一样
	var id = 1
	var name = "ying"
	var stu1 = student{
		id,
		name,
	}
	fmt.Println(stu1)
}
