package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
}

func (p person) sayName() {
	fmt.Println(p.name)
}

func (p person) isAdult() bool {
	return p.age >= 18
}

func main() {
	var p1 = person{
		name:   "hcy",
		age:    19,
		gender: "男",
	}
	p1.sayName()              // hcy
	fmt.Println(p1.isAdult()) // true
}
