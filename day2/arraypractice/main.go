package main

import "fmt"

func main() {
	arr1 := [...]int{1, 3, 5, 7, 8}

	// 1、求和 [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, value := range arr1 {
		sum += value
	}
	fmt.Printf("数组和为:%d\n", sum) // 数组和为:24

	// 2、在 [...]int{1, 3, 5, 7, 8} 里找出和为 8 的两个数的索引
	for i := 0; i < len(arr1); i++ {
		for j := i + 1; j < len(arr1); j++ {
			if arr1[i]+arr1[j] == 8 {
				fmt.Printf("索引为%d和%d的值之和为8\n", i, j)
				break // 如果以及算出一次8了，那么就可以跳出这次的内存for循环了
			}
		}
	}
	// 索引为0和3的值之和为8
	// 索引为1和2的值之和为8
}
