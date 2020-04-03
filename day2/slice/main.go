package main

import "fmt"

func main() {
	var s1 []int           // 切片不需要指明长度，而且是引用数据类型
	fmt.Printf("%v\n", s1) // []
	// nil就是空，和别的语言null相同
	fmt.Printf("%v\n", s1 == nil) // true

	s2 := []int{1, 2, 3}
	fmt.Printf("切片s2的长度:%d\n", len(s2))                     // 切片s2的长度:3
	fmt.Printf("切片s2的值:%v,类型为:%T\n", s2, s2)                // 切片s2的值:[1 2 3],类型为:[]int
	fmt.Printf("len(s2):%d,cap(s2):%d\n", len(s2), cap(s2)) // len(s2):3,cap(s2):3
	// 这种声明后直接赋值的切片，len和cap相等

	// 切片可以以一个数组作为其依据的底层数组
	arr := [...]int{1, 3, 5, 7, 9, 11}
	// arr[start, end] 遵从左闭右开的规则
	s3 := arr[1:3]
	s4 := arr[:3]
	s5 := arr[3:]
	s6 := arr[:]
	fmt.Printf("%v为len(s3):%d,cap(s3):%d\n", s3, len(s3), cap(s3)) // [3 5]为len(s3):2,cap(s3):5
	fmt.Printf("%v为len(s4):%d,cap(s4):%d\n", s4, len(s4), cap(s4)) // [1 3 5]为len(s4):3,cap(s4):6
	fmt.Printf("%v为len(s5):%d,cap(s5):%d\n", s5, len(s5), cap(s5)) // [7 9 11]为len(s5):3,cap(s5):3
	fmt.Printf("%v为len(s6):%d,cap(s6):%d\n", s6, len(s6), cap(s6)) // [1 3 5 7 9 11]为len(s6):6,cap(s6):6
	// - s3的cap容量为5，s4的cap容量为6，而s5的cap容量为3
	// - 说明cap的算法是以其依赖的数组的长度的，算法为 start 到 len(arr) 之间有的数组成员的数量

	s7 := s4[1:]
	fmt.Printf("%v为len(s7):%d,cap(s7):%d\n", s7, len(s7), cap(s7)) // [3 5]为len(s7):2,cap(s7):5

	arr1 := [...]int{1, 2, 3, 4}
	sli1 := arr1[:2]
	sli2 := sli1[:1]
	fmt.Printf("修改前的,sli1:%v,sli2:%v\n", sli1, sli2) // 修改前的,sli1:[1 2],sli2:[1]
	sli1[0] = 100
	fmt.Printf("修改后的,sli1:%v,sli2:%v\n", sli1, sli2) // 修改后的,sli1:[100 2],sli2:[100]
}
