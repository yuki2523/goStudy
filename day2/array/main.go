package main

import "fmt"

func main() {
	var arr1 [3]int16
	var boolArr [3]bool
	// 数组定义时，要指明数组长度，还有类型，[长度]类型
	fmt.Printf("%v,type:%T\n", arr1, arr1)       // [0 0 0],type:[3]int16
	fmt.Printf("%v,type:%T\n", boolArr, boolArr) // [false false false],type:[3]bool
	// 数组的类型打印出来，会带上长度 [length]type, 这就是数组的类型

	// 下面是数组初始化的方法，可以直接这样写: [length]type{...value}
	arr2 := [4]int{1, 2, 3, 4}
	fmt.Printf("%v,type:%T\n", arr2, arr2) // [1 2 3 4],type:[4]int

	// 给数组指定的索引处响应的值
	// 写法：[length]type{index: value,...}
	arr3 := [5]int{0: 1, 2: 2, 4: 3}
	fmt.Printf("%v,type:%T\n", arr3, arr3) // [1 0 2 0 3],type:[5]int
	// 没有指定的就是 0

	// 如果太长，或者就是不想写length，可以直接写三个点： ...
	// 但是这种情况是必须已经明确把数组里有数据给写出来了
	arr4 := [...]int{1, 3, 5, 7, 9}
	fmt.Printf("%v,type:%T\n", arr4, arr4) // [1 3 5 7 9],type:[5]int

	arr5 := [...]int{1: 3, 5: 9}           // 这样写也可以，长度会自动计算出来，没指定的用0补全
	fmt.Printf("%v,type:%T\n", arr5, arr5) // [0 3 0 0 0 9],type:[6]int

	// 数组的遍历
	citys := [...]string{"芜湖", "六安", "合肥"}

	// 最基本的索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	// 芜湖
	// 六安
	// 合肥

	// range遍历
	for index, value := range citys {
		fmt.Printf("索引为%d的值为%v\n", index, value)
	}
	// 索引为0的值为芜湖
	// 索引为1的值为六安
	// 索引为2的值为合肥

	// 二维数组
	var items [3][2]int
	items = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6}, // 记得最后一行这里，一定要加上逗号`,`，不然是会报错的
	}
	for _, value1 := range items {
		for _, value2 := range value1 {
			fmt.Println(value2)
		}
	}
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6

	// - 数组是值类型
	// - 与js不同，这里的数组是值类型，是可以使用 ==，!= 这些比较运算符的
	// 由于是值类型，并非是引用类型，那么这里使用数组，根本不用考虑不同的变量指在同一个数组的地址值这些东西，根本就不存在
	var c1 [3]int = [3]int{4, 5, 6}
	c2 := c1
	c2[1] = 50
	fmt.Printf("c1:%v,c2:%v", c1, c2) // c1:[4 5 6],c2:[4 50 6]
	// c2改变，与c1无关，因为是值类型，两者是处在不同内存地址，互相独立的
}
