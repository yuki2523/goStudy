package main

import "fmt"

func main() {
	var age map[string]int
	age = make(map[string]int, 3)
	age["hcy"] = 19
	age["ying"] = 16
	age["yang"] = 20
	age["yuki"] = 14
	fmt.Println(age) // map[hcy:19 yang:20 ying:16 yuki:14]

	// delete 删除map里的键值对
	// 语法：delete(目标map, key)
	// key存在，就去删
	delete(age, "hcy")
	fmt.Println(age) // map[yang:20 ying:16 yuki:14]

	// key不存在，就什么都不做，不会报错的
	delete(age, "nameless")
	fmt.Println(age) // map[yang:20 ying:16 yuki:14]
}
