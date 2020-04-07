package main

import "fmt"

type animal interface {
	move()
	eat(string) string
}

type chichen struct {
	feet int
}

func (c *chichen) move() {
	fmt.Println("鸡在飞")
}

func (c *chichen) eat(food string) string {
	fmt.Printf("鸡在吃%s\n", food)
	return "鸡蛋"
}

func main() {
	var an1 animal
	// var ch1 = chichen{2}
	var ch2 = &chichen{2}

	// 当方法为指针接收者时,interface类型的变量只能存地址值了
	// an1 = ch1
	// 报错,报错信息如下
	// cannot use ch1 (type chichen) as type animal in assignment:
	// chichen does not implement animal (eat method has pointer receiver)
	an1 = ch2
	fmt.Printf("%#v,%T\n", an1, an1)
}
