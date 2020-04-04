package main

import "fmt"

func main() {
	// append 为切片追加元素
	s1 := []string{"芜湖", "合肥", "六安"}
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1)) // len:3,cap:3
	// s1[3] = "上海" // panic: runtime error: index out of range [3] with length 3
	// 这种情况，编译时才会报错，索引越界了

	// append函数语法: 原切片变量 = append(原切片变量, 新的元素)
	// 注意一点，append函数的结果，必须使用原追加元素的切片变量去接收
	s1 = append(s1, "上海")
	fmt.Printf("len:%d,cap:%d\n", len(s1), cap(s1)) // len:4,cap:6
	// 从源码开来，这里应该是进入了 old.len < 1024 里,然后执行了 newcap = doublecap,因此cap是翻倍了
	fmt.Printf("s1:%v\n", s1) // s1:[芜湖 合肥 六安 上海]

	// 下面这段为go处理slice的源码里摘取的部分
	// newcap := old.cap
	// doublecap := newcap + newcap
	// if cap > doublecap { // 如果新申请的newcap*2小于原先的oldcap，那么cap不变
	// 	newcap = cap
	// } else {
	// 	if old.len < 1024 { // 如果旧切片的长度小于1024,那么新的newcap为原先oldcap的二倍
	// 		newcap = doublecap
	// 	} else {
	// 		// Check 0 < newcap to detect overflow
	// 		// and prevent an infinite loop.
	// 		for 0 < newcap && newcap < cap { // newcap大于0且小于原先cap，循环执行newcap += newcap / 4
	// 			newcap += newcap / 4
	// 		}
	// 		// Set newcap to the requested cap when
	// 		// the newcap calculation overflowed.
	// 		if newcap <= 0 { // 如果新的newcap小于等于0，那么cap不变
	// 			newcap = cap
	// 		}
	// 	}
	// }

	// 一次性 append 多个
	// 拆开切片的方式, `slice变量名...`
	s2 := []int{1, 2, 3}
	fmt.Printf("len:%d,cap:%d\n", len(s2), cap(s2)) // len:3,cap:3
	s3 := []int{4, 5, 6, 7, 8}
	s2 = append(s2, s3...)
	fmt.Printf("len:%d,cap:%d\n", len(s2), cap(s2)) // len:8,cap:8
	// 这种情况下，新申请的 cap 应该是 4
}
