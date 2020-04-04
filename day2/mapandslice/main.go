package main

import "fmt"

func main() {
	// 要注意一点，它们都是引用数据类型。因此都需要make

	// 1、在map里存入slice
	// 也就是说，map的 valueType 是 slice
	m1 := make(map[string][]int, 5)
	m1["first"] = make([]int, 5, 10)
	// 重点就在于这个 make 会用上两次，这种还好，下一种要格外注意
	fmt.Printf("%v, %d\n", m1, len(m1)) // map[first:[0 0 0 0 0]], 1
	m1["first"][1] = 100
	fmt.Println(m1["first"][1]) // 100

	// 2、在slice里存入map
	// slice 的类型是一个 map 类型、
	m2 := make([]map[string]int, 5, 10)
	m2[0] = make(map[string]int, 5)
	m2[0]["age"] = 20
	fmt.Println(m2)           // [map[age:20] map[] map[] map[] map[]]
	fmt.Println(m2[0] == nil) // false
	fmt.Println(m2[1] == nil) // true
	// 可以发现，没有 make 的，是 nil ，因为没有申请内存
}
