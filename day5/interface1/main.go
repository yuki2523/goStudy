package main

import "fmt"

type animal interface {
	move()
	eat(string) string
}

type rabbit struct {
	name string
	feet int
}

type chicken struct {
	feet int
}

func (r rabbit) move() {
	fmt.Println("兔子在跳")
}

func (r rabbit) eat(food string) string {
	fmt.Printf("兔子%s在吃%s\n", r.name, food)
	return "兔子的排泄物"
}

func (c chicken) move() {
	fmt.Println("鸡在飞")
}

func (c chicken) eat(food string) string {
	fmt.Printf("鸡在吃%s\n", food)
	return "鸡蛋"
}

func main() {
	var r1 = rabbit{"小白兔", 4}
	var c1 = chicken{2}

	// 接口类型的变量声明出来后,分为两个部分,一个部分是存值,另一部分存数据类型
	// 存值实际是在存一个值的内存地址,根据内存地址取值
	// 因此接口类型的变量很灵活,它可以存储任何符合这个接口类型条件的值
	var an1 animal
	var an2 = an1
	fmt.Printf("%#v,%T\n", an1, an1) // <nil>,<nil>
	fmt.Printf("%#v,%T\n", an2, an2) // <nil>,<nil>

	an1 = r1
	fmt.Printf("%#v,%T\n", an1, an1) // main.rabbit{name:"小白兔", feet:4},main.rabbit
	fmt.Printf("%#v,%T\n", an2, an2) // <nil>,<nil>

	an1 = c1
	fmt.Printf("%#v,%T\n", an1, an1) // main.chicken{feet:2},main.chicken
	fmt.Printf("%#v,%T\n", an2, an2) // <nil>,<nil>
}
