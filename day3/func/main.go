package main

import "fmt"

// Go 强强强强强强...数据类型语言,因此要把类型指明
// 语法
// 1、最完整的一种：func 函数名(形参 dataType) (返回值形参 dataType) {}
func sum(x int, y int) (result int) {
	result = x + y // 上面写明了返回值的变量名，函数里就可以直接用了
	return         // 而且直接 return 就行了，也不需要把 return 的值写在后面
}

// 2、不提前写返回值的名字
func resultNameless(x int, y int) int {
	ret := x + y
	return ret
}

// 3、多个返回值
func returnManyResult(x int, y int) (int, int, int) {
	sum := x + y
	subtract := x - y
	multiply := x * y
	return sum, subtract, multiply
}

// 4、无返回值
func printStringSum(firstName string, lastName string) {
	fmt.Println(firstName + lastName)
}

// 5、无参数无返回值
func onlyDoFunction() {
	fmt.Println("无参数无返回值函数")
}

// 6、有不定长的参数
// 不定长参数，必须放在参数里的最后一个，经典省略号... 接收多个参数 ,js也是这样
// 不定长参数在函数里的呈现形式是 slice ,可以不传值，也可以传入1个，多个
func unsetLengthParam(x int, y ...int) {
	fmt.Print(x)
	fmt.Printf("不定长参数值:%v,类型:%T\n", y, y)
}

// 7、同类型的参数写在一起，只要写一次类型就可以了
func datatypeOnceWrite(x, y, z int, i, j string, p, q bool) {
	fmt.Println(x, y, z, i, j, p, q)
}

func main() {
	fmt.Println(sum(1, 2))                              // 3
	fmt.Println(resultNameless(2, 3))                   // 5
	fmt.Println(returnManyResult(4, 5))                 // 9 -1 20
	printStringSum("h", "cy")                           // hcy
	onlyDoFunction()                                    // 无参数无返回值函数
	unsetLengthParam(100, 1, 2, 3, 4, 5, 6, 7)          // 不定长参数值:[1 2 3 4 5 6 7],类型:[]int
	datatypeOnceWrite(1, 2, 3, "aa", "bb", true, false) // 1 2 3 aa bb true false
}
