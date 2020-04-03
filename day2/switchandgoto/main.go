package main

import "fmt"

func main() {
	// 1、break , continue
	// 直接跳出循环的方法，也就是使用 break
	// 跳过当次循环的方法，也就是使用 continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // 在i为5的时候跳出循环
			// continue // 在i为5时不执行这句循环里的代码直接跳过
		}
	}

	// 2、switch
	// 这用法不会有区别的，看看go怎么写的就好了
	n := 0
	switch n {
	case 3:
		fmt.Printf("n是3\n")
	case 4:
		fmt.Printf("n是4\n")
	case 5:
		fmt.Printf("n是5\n")
	default: // 设置默认值
		fmt.Printf("n未知\n")
	}

	// go里的switch在变量处可以赋值，case可以匹配多个值
	switch i := 3; i {
	case 1, 2, 3:
		fmt.Println("i是1,2,3里的一个")
	case 4, 5, 6:
		fmt.Println("i是4,5,6里的一个")
	}

	// go里的switch里的case还可以匹配表达式
	// 不过这种写法的话，switch 后面不要写变量
	j := 10
	switch { // 这个后面不要写变量
	case j < 5:
		fmt.Println("j在5以内")
	case j >= 5 && j <= 15:
		fmt.Println("j在5到15之间")
	case j > 15:
		fmt.Println("j大于15")
	}

	// go 里可以认为 case 匹配到一个就会马上结束，就不像js，c，c++一样还往下接着穿透
	// 但是如果就想它接着往下穿透，那么有一个关键字 fallthrough 可以解决这个问题
	x := 10
	switch {
	case x > 5:
		fmt.Println(">5")
		fallthrough // 没有这句的话只会打印 >5 ，有fallthrough后还会打印 >7
	case x > 7:
		fmt.Println(">7")
	case x > 10:
		fmt.Println(">10")
	}

	// goto
	// 直接把代码给引导到另外的位置
	// 下面展示一个循环里的用法
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				goto routerMark // 写指向的标签名
			}
			fmt.Println(i, j)
		}
	}

routerMark: // 直接写标签名: 后面接上想要写的代码
	fmt.Println("给我跳")
}
