package main

import "fmt"

func deferTest() (ret int) {
	// 有 defer 时，前面加了 defer 的语句会放到 defer栈内
	// 返回值 ret 赋值 => defer栈里的语句依次出栈执行 => 真正的返回 ret
	fmt.Println("start")
	defer fmt.Println("deferFirst")
	defer fmt.Println("deferSecond")
	defer fmt.Println("deferThird")
	fmt.Println("end")
	return
	// defer 给我的感觉很像是js里的回调队列，但是它这个是一个栈，等主线程执行完毕，再把栈里的语句逐句出栈执行
}

// 下面开始，是一些展示 defer 特性的题目
func f1() (x int) { // 这个时候返回值是 x
	defer func() {
		x++
	}()
	return 5
	// 看到return
	// 先看看上面有没有声明返回值
	// 发现有，也就是这里的 x
	// 然后给 x 赋值为 5
	// 之后再找 defer,发现有 defer,执行defer里的语句
	// defer 里的语句是 x++，执行完后 x = 6
	// 真正的返回出返回值 x ,输出的结果就是 6
}

func f2() (x int) { // 这个时候返回值还是 x
	fmt.Println(x)      // 0
	defer func(x int) { // defer 虽然会暂时阻止其执行,但是参数会在这句决定好
		x++
		fmt.Println(x, "内部的") // 1 内部的
	}(x) // 在执行到这个函数时,x的值为0,因此里面 x++ 后,打印出来的 x 值为 1
	defer func() { // 这个就又不同了,因为在这里根本就没有参数传进来,那么自然是使用外层函数作用域的 x
		fmt.Println(x) // 5
	}()
	return 5
}

func f3() int { // 这里没有指定返回值的变量名
	// 没指明返回值名字时,可以认为它自己产生了一个返回值变量名为 ret
	x := 5
	defer func() {
		x++
		fmt.Println(x, "inner") // 6 inner
	}()
	return x
	// 看到return,先是给返回值 ret 赋值为 x = 5
	// 之后再执行 defer, x++ 之后 x = 6
	// 最后再返回返回值 ret = 5
}

func f4() (y int) { // 这个和没有指明返回值的变量名其实是可以归为一类的
	x := 5
	defer func() {
		x++
		fmt.Println(x, "inner")
	}()
	return y
	// 看到return,先是给返回值 y 赋值为 x = 5
	// 之后再执行 defer, x++ 之后 x = 6
	// 最后再返回返回值 y = 5
}

func main() {
	deferTest()
	// start
	// end
	// deferThird
	// deferSecond
	// deferFirst

	fmt.Println(f1()) // 6
	fmt.Println(f2())
	fmt.Println("f3()------")
	fmt.Println(f3())
}
