package main

import "fmt"

const ( // 批量声明常量
	statusOK = 200
	notFound = 404
	forbiden = 403
)

const ( // 这种的，如果第一个写了值，后面三个没这样写，就全和第一个一样
	n1 = 100
	n2
	n3
)

// iota 常量计数器
// 最普通的
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
)

// 有插队的，还有匿名的
const (
	b1 = iota // 0
	_         // 1被丢弃了
	b2 = iota // 2
	b3 = 100  // 100, 3它不要了
	b4 = iota // 4
)

// 一行多个变量的
const (
	c1, c2 = iota + 1, iota + 2 // 1, 2
	c3, c4 = iota + 1, iota + 2 // 2, 3
)

// 可以把iota单独理解为一个会自增的变量
// 在const那个括号里,第一次出现的时候是 0
// 之后每多出有声明变量的一行,iota的值+1,必须是多出有声明变量的一行,不然iota的值是不会变化的,切记

// iota的一个适用场景
const (
	_  = iota             // 0不要
	KB = 1 << (10 * iota) // 1 << 10,就是左移十位， 10000000000二进制数，就是2^10 也就是1024
	MB = 1 << (10 * iota) // 那么这个就是 2^20， 下面就以此类推
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Printf("n1:%d,n2:%d,n3:%d", n1, n2, n3) // n1:100,n2:100,n3:100
	fmt.Println(a1, a2, a3)                     // 0 1 2
	fmt.Println(b1, b2, b3, b4)                 // 0 2 100 4
	fmt.Println(c1, c2, c3, c4)                 // 1 2 2 3
}
