package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}

// 可以返回结构体，也可以返回结构体指针
// 如果返回的是结构体，那接收后的变量的值就是整个结构体，变量会变的很重
// 如果返回的是结构体指针，那么接收后的变量只存了一个地址值，每次使用时再从地址值取值
// 我个人比较倾向于使用结构体指针
func newPerson(name string, age int, gender string) *person {
	return &person{
		name:   name,
		age:    age,
		gender: gender,
	}
}

func main() {
	var p1 = newPerson("hcy", 19, "男") // 这时的 p1 是结构体指针
	fmt.Printf("%#v\n", *p1)           // main.person{name:"hcy", age:19, gender:"男"}
	fmt.Printf("%T,%p\n", p1, p1)      // *main.person,0xc000068330
}
