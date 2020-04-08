package main

import (
	"fmt"

	// 直接在左边在加一个名字是给包取别名的意思
	p1 "yuki2523.goStudy.git/day5/packagetest1"
	"yuki2523.goStudy.git/day5/packagetest2"

	// 匿名导入包,只执行包里的 init 函数,无法使用包里暴露向外的数据
	_ "yuki2523.goStudy.git/day5/packagetest3"
)

func init() {
	fmt.Println("main.go")
}

func main() {
	var (
		a    = 1
		b    = 4
		sli  = []int{1, 2}
		map1 map[string]int
		bl1  = true
		s    = ""
	)

	fmt.Println(p1.Add(a, b))             // 5
	fmt.Println(p1.Sub(a, b))             // -3
	fmt.Println(packagetest2.Isnil(sli))  // false
	fmt.Println(packagetest2.Isnil(map1)) // true
	fmt.Println(packagetest2.Iszero(bl1)) // true
	fmt.Println(packagetest2.Iszero(s))   // true
}
