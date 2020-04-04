package main

import (
	"fmt"
	"strings"
)

func main() {
	// // 引用数据类型: map
	// // 定义方式: map[keyType]valueType
	// var ageList map[string]int
	// fmt.Printf("type:%T, value:%v\n", ageList, ageList) // type:map[string]int, value:map[]
	// // 数据类型是 map[string]int, key的类型和value的类型都要写明确
	// fmt.Println(ageList == nil) // true
	// // var声明出来的 map 类型数据，由于是引用数据类型，一开始是没有指向任何内存空间的，那么就是 nil

	// 为了获取有内存空间的 map 数据类型的数据，可以使用 make 函数
	// 语法: make(map[keyType]valueType, 容量)
	ageList := make(map[string]int, 10)
	fmt.Printf("value:%v,isNil:%v\n", ageList, ageList == nil) // value:map[],isNil:false

	// 向 map 里存值的方法: map变量名[key] = value
	ageList["hcy"] = 19
	ageList["ying"] = 16

	// 在 map 里获取值的方式: map变量名[key],它有两个返回值
	// 如果 key 存在，value就是对应的值，ok返回布尔值 true
	// 如果 key 不存在，value就是对应类型的值的空或0，int就是0，string是空字符串，ok返回布尔值 false
	// 当然不要 ok 也是可以的，value 取值规则是相同的
	name := "hcy"
	value, ok := ageList[name]
	fmt.Printf("value:%v,ok:%v\n", value, ok) // value:19,ok:true

	name = "h"
	value, ok = ageList[name]
	fmt.Printf("value:%v,ok:%v\n", value, ok) // value:0,ok:false

	// 遍历 map
	// 使用 for range 就可以了
	// 语法: for key, value := range map类型变量名 {}
	for Name, age := range ageList {
		fmt.Println(Name, age)
	}
	// hcy 19
	// ying 16

	// 也可以只遍历 key
	// 语法: for key := range map类型变量名 {}
	for Name := range ageList {
		fmt.Println(Name)
	}
	// hcy
	// ying

	// 练习 : 统计单词数量
	str := "how do you do"
	wordCount := make(map[string]int, 10)
	var key string

	wordList := strings.Split(str, " ")
	for _, value := range wordList {
		key = value
		_, ok := wordCount[key]
		if !ok {
			wordCount[key] = 1
		} else {
			wordCount[key]++
		}
	}

	fmt.Println(wordCount) // map[do:2 how:1 you:1]
}
