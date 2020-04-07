package main

import "fmt"

type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string) string
}

type chicken struct {
	feet int
}

func (c chicken) move() {
	fmt.Println("鸡在飞")
}

func (c chicken) eat(food string) string {
	fmt.Printf("鸡在吃%s\n", food)
	return "鸡蛋"
}

func main() {
	var (
		an1 animal
		// an2 animal
		// an3 animal
		m1 mover
		e1 eater
	)

	var c1 = chicken{2}
	m1 = c1
	e1 = c1
	fmt.Printf("%#v,%T\n", m1, m1) // main.chicken{feet:2},main.chicken
	fmt.Printf("%#v,%T\n", e1, e1) // main.chicken{feet:2},main.chicken
	// fmt.Println(m1 == e1) // 类型不同,无法比较

	an1 = c1
	// an2 = m1 // 报错,报错信息大意是 mover 类型缺少 eat 方法
	// an3 = e1 // 报错,报错信息大意是 eater 类型缺少 move 方法
	fmt.Printf("%#v,%T\n", an1, an1) // main.chicken{feet:2},main.chicken
	// fmt.Printf("%#v,%T\n", an2, an2)
	// fmt.Printf("%#v,%T\n", an3, an3)

	fmt.Println(an1 == m1) // true
	fmt.Println(an1 == e1) // true
	// 这两个 true 说明一点,被嵌套的和外层接口类型是兼容的
}
