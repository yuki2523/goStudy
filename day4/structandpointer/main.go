package main

import "fmt"

func personChange(p person) {
	p.name = "ying"
}

func personChangePointer(p *person) {
	// (*p).name = "huachenyang"
	p.name = "huachenyang" // 简写成这样也可以，Go会自动判断是否为指针并取值
}

type person struct {
	name   string
	age    int
	gender string
}

func main() {
	var p1 person
	p1.name = "hcy"
	p1.age = 19
	p1.gender = "男"
	fmt.Println(p1) // {hcy 19 男}
	personChange(p1)
	fmt.Println(p1) // {hcy 19 男}
	personChangePointer(&p1)
	fmt.Println(p1) // {huachenyang 19 男}
	// 直接传入 person 实例对象并没有成功修改,证明了结构体为值类型的
	// 传入指针的话，肯定是达成修改的，重点就是结构体为值类型

	// 获取结构体指针的方法
	// 1、new 一个，因为是值类型，所以用 new
	var p2 = new(person)             // p2 是一个指针,指向一个person类型变量(结构体)
	fmt.Printf("p2:type:%T\n", p2)   // p2:type:*main.person
	fmt.Printf("p2指向值的地址值:%p\n", p2) // p2指向值的地址值:0xc0000a63c0
	// 但是这种方式初始化起来会比较麻烦，还要一个个的写

	// 2、类似字面量直接定义并初始化的方式
	// 语法：&person{...}
	var p3 = &person{
		name:   "ying",
		age:    16,
		gender: "男",
	}
	fmt.Printf("%T\n", p3)   // *main.person
	fmt.Printf("%#v\n", *p3) // main.person{name:"ying", age:16, gender:"男"}

	// 3、这种简单的写法里还可以直接省略字段名，不过这样的话必须按照定义时字段顺序写下来
	var p4 = &person{
		"yuki",
		14,
		"女",
	}
	fmt.Printf("%T\n", p4)   // *main.person
	fmt.Printf("%#v\n", *p4) // main.person{name:"yuki", age:14, gender:"女"}

	// 4、结构体里的字段的地址值是紧挨着的
	var test struct {
		a int8
		b int8
		c int8
	}
	test.a = 1
	test.b = 2
	test.c = 3
	fmt.Println(&test.a) // 0xc0000100f8
	fmt.Println(&test.b) // 0xc0000100f9
	fmt.Println(&test.c) // 0xc0000100fa
}
