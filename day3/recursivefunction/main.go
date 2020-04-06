package main

import "fmt"

// 算阶乘的递归函数
// 递归一定要有一个合适的退出条件
func factorial(num uint64) uint64 {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

// 楼梯，一次可以走 1 步，也可以走 2 步，n阶时多少走法
func floorWays(n uint64) uint64 {
	if n > 0 {
		if n == 2 { // 2 阶时
			// 可以走一步跨上去，也可以走两步，一次性上一阶，两种走法
			return 2
		} else if n == 1 { // 1 阶时
			// 只有一步跨上去的唯一一个选择
			return 1
		} else {
			// 有多阶(>2)时，可以把一次走两步的方法数 floorWays(n - 2)
			// 和一次走一步的方法数 floorWays(n - 1) 加起来
			return floorWays(n-1) + floorWays(n-2)
		}
	}
	return 0
}

func main() {
	fmt.Println(factorial(uint64(20)))
	fmt.Println(floorWays(10))
}
