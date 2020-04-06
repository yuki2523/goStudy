package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}

// 结构体的指针接收者
// 适用场景
// 1、需要修改调用该方法的实例里的值
// 2、调用该方法的实例较大,再拷贝一份会占用较大的内存空间
// 3、保证这个类型的结构体方法的一致性,只要有一个方法适用了指针接收者,别的也都用指针接收者
func (p *person) ageChange(newAge int) {
	// (*p).age = newAge // 可以简写为下面这种
	p.age = newAge
}

func main() {
	var p1 = person{
		name:   "hcy",
		age:    19,
		gender: "男",
	}
	fmt.Println(p1.age) // 19
	// 结构体的指针接收者, (p *person)，这样写的话,person实例在调用这个方法时,传入的p是它的指针
	p1.ageChange(20)
	fmt.Println(p1.age) // 20
}
