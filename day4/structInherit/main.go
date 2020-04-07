package main

import "fmt"

type object struct {
}

type person struct {
	name   string
	age    int
	gender string
	object
}

func (o object) hello() {
	fmt.Println(o) // {}
	fmt.Println("Hello 这是object里方法hello")
}

func main() {
	var p1 = person{
		name:   "hcy",
		age:    19,
		gender: "男",
	}

	p1.hello() // Hello 这是object里方法hello
	// 虽然 person 实例确实是点出了 object 类的方法,但是这和JS里的prototype有着本质的差别
	// 首先,无法让 object 类里的,也就是继承的类拿到传递进去的实例,也就是this是什么都点不出来的
	// 虽然是可以在 person 类的实例里使用 object 类的方法和属性,但是方法中无法获取当前对象,都是白搭,没用
}
