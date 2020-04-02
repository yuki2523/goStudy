package main

import (
	"fmt"
)

func main() {
	// 1、if / else if / else (学一下写法，套路都一样)
	age := 18
	if age < 18 {
		fmt.Println("小屁孩不要来上网")
	} else if age > 18 {
		fmt.Println("xxx网咖欢迎光临")
	} else {
		fmt.Println("刚满18就来上网？？？") // 刚满18就来上网？？？
	}
	fmt.Println(age) // 18

	// 还有一个类似js中let块作用域的东西
	// 写法是用一个分号; 隔开,然后在if这个控制语句内 sex := 3 变量有效
	if sex := 3; sex == 0 {
		fmt.Println(sex, "女")
	} else if sex == 1 {
		fmt.Println(sex, "男")
	} else {
		fmt.Println(sex, "性别未知") // 3 性别未知
	}
	// fmt.Println(sex) // 报错，undefined

	// 2、for 的基本写法
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i) // 0 1 2 3 4 5 6 7 8 9
	}
	// fmt.Println(i) // 报错undefined，同样是块作用域

	// 省略赋值语句，也就是写在外面
	j := 0
	for ; j < 10; j++ {
		fmt.Printf("%d ", j) // 0 1 2 3 4 5 6 7 8 9
	}
	fmt.Println(j) // 10,这个j就不是块作用域了

	// 就写个条件在外面
	x := 5
	for x < 10 {
		fmt.Printf("%d ", x) // 5 6 7 8 9
		x++
	}

	// 无限死循环
	// for {
	// 	fmt.Println("11")
	// }

	// for index, item := range 变量 {}
	// 这也是一种循环写法,下面拿一个字符串举例
	var str string = "学习Go"
	for index, item := range str {
		fmt.Printf("索引%d的字符为%c\n", index, item)
	}
	// 索引0的字符为学
	// 索引3的字符为习
	// 索引6的字符为G
	// 索引7的字符为o

	// 经典九九乘法表
	for a := 1; a < 10; a++ {
		for b := 1; b < a+1; b++ {
			fmt.Printf("%d * %d = %d\t", b, a, a*b)
		}
		fmt.Printf("\n")
	}
	// 1 * 1 = 1
	// 1 * 2 = 2       2 * 2 = 4
	// 1 * 3 = 3       2 * 3 = 6       3 * 3 = 9
	// 1 * 4 = 4       2 * 4 = 8       3 * 4 = 12      4 * 4 = 16
	// 1 * 5 = 5       2 * 5 = 10      3 * 5 = 15      4 * 5 = 20      5 * 5 = 25
	// 1 * 6 = 6       2 * 6 = 12      3 * 6 = 18      4 * 6 = 24      5 * 6 = 30      6 * 6 = 36
	// 1 * 7 = 7       2 * 7 = 14      3 * 7 = 21      4 * 7 = 28      5 * 7 = 35      6 * 7 = 42      7 * 7 = 49
	// 1 * 8 = 8       2 * 8 = 16      3 * 8 = 24      4 * 8 = 32      5 * 8 = 40      6 * 8 = 48      7 * 8 = 56      8 * 8 = 64
	// 1 * 9 = 9       2 * 9 = 18      3 * 9 = 27      4 * 9 = 36      5 * 9 = 45      6 * 9 = 54      7 * 9 = 63      8 * 9 = 72      9 * 9 = 81
}
