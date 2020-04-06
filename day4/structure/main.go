package main

import "fmt"

type person struct {
	name    string
	age     int
	gender  string
	hobbies []string
}

func main() {
	// 1、属性值都写全的情况
	var p1 person
	p1.name = "hcy"
	p1.age = 19
	p1.gender = "男"
	p1.hobbies = []string{"code", "game", "sleep"}
	fmt.Printf("p1,type:%T,value:%#v\n", p1, p1)
	// p1,type:main.person,value:main.person{name:"hcy", age:19, gender:"男", hobbies:[]string{"code", "game", "sleep"}}

	// 2、只写一个属性值的情况
	var p2 person
	p2.name = "ying"
	fmt.Println(p2.age)            // 0
	fmt.Println(p2.gender == "")   // true
	fmt.Println(p2.hobbies == nil) // true
	// 可以发现,结构体使用 var 声明出来后
	// int 类型的属性，默认值为 0
	// string 类型的属性，默认值为 "" 空字符串
	// slice 引用数据类型的属性, 默认值为 nil

	// 3、匿名结构体，在只用一次的时候用
	var s struct {
		x string
		y int
	}
	s.x = "aaa"
	s.y = 10
	fmt.Printf("s,type:%T,value:%v\n", s, s) // s,type:struct { x string; y int },value:{aaa 10}

	// 4、结构体也是值类型的
	var p3 = &p1
	fmt.Println(p1)  // {hcy 19 男 [code game sleep]}
	fmt.Println(*p3) // {hcy 19 男 [code game sleep]}
	p1.name = "huachenyang"
	fmt.Println(p1)  // {huachenyang 19 男 [code game sleep]}
	fmt.Println(*p3) // {huachenyang 19 男 [code game sleep]}
}
