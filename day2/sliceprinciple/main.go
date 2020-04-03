package main

import "fmt"

func main() {
	// make() 函数创建切片
	// 下面是初始化
	s1 := make([]int, 5, 10) // type, len, cap
	fmt.Printf("s1:%v, len(s1):%d, cap(s1):%d\n", s1, len(s1), cap(s1))
	// s1:[0 0 0 0 0], len(s1):5, cap(s1):10
	// 可以发现s1为容量10长度5的切片，值全部由0填充

	// make出来的切片，一定是有底层数组的，就算是空的，切片也不为 nil
	s2 := make([]int, 0, 0)
	fmt.Printf("s2:%v, len(s2):%d, cap(s2):%d\n", s2, len(s2), cap(s2))
	// s2:[], len(s2):0, cap(s2):0
	fmt.Println(s2 == nil) // false

	// 直接声明出来的切片，没有赋值时，没有底层数组，因此为 nil
	var s3 []int
	fmt.Printf("s3:%v, len(s3):%d, cap(s3):%d\n", s3, len(s3), cap(s3))
	// s3:[], len(s3):0, cap(s3):0
	fmt.Println(s3 == nil) // true

	// 下面还是通过切片的赋值拷贝继续证明切片为引用数据类型
	s4 := []int{1, 2, 3}
	s5 := s4 // s4和s5指向相同的内存地址，因此改s5时，s4也变
	// 其实本质是改掉了它们共同所指的那块内存里的值
	fmt.Println(s4, s5) // [1 2 3] [1 2 3]
	s5[0] = 100
	fmt.Println(s4, s5) // [100 2 3] [100 2 3]

	// 切片遍历
	s6 := []int{1, 2, 3}

	// 索引遍历
	for i := 0; i < len(s6); i++ {
		fmt.Println(s6[i])
	}
	// 1
	// 2
	// 3

	// range遍历
	for index, value := range s6 {
		fmt.Println(index, value)
	}
	// 0 1
	// 1 2
	// 2 3
}
