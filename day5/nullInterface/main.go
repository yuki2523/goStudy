package main

import "fmt"

// 2、空接口适用场景之二, 可以接收任何数据类型的函数参数
func dataPrinter(data interface{}) {
	fmt.Println(data)
}

// 3、断言判断空接口传入的数据类型
func typeGetter(noType interface{}) {
	var data, ok = noType.(string)
	// 空接口类型可以使用 `变量.(datatype)`,type是一个具体的类型,它有两个返回值
	// 如果正确,第二个返回值 ok 为 true,错误则为 false
	// 正确时,第一个返回值 data 就是当前对应类型的值,也就是从空接口类型转回它实际对应的类型
	// 错误时,就是 datatype 类型对应的 0 值,int：0,string："",bool：false,引用数据类型: nil
	fmt.Println(data, ok)
}

// 4、断言判断类型在 switch case 版本里的特别使用方法
func typePrinter(noType interface{}) {
	switch data := noType.(type) {
	case string:
		fmt.Printf("string:%#v,%T\n", data, data)
	case bool:
		fmt.Printf("bool:%#v,%T\n", data, data)
	case int:
		fmt.Printf("int:%#v,%T\n", data, data)
	case uint:
		fmt.Printf("uint:%#v,%T\n", data, data)
	case []int:
		fmt.Printf("slice:[]int:%#v,%T\n", data, data)
	case map[string]int:
		fmt.Printf("map[string]int:%#v,%T\n", data, data)
	case func():
		fmt.Printf("func():%T\n", data)
	case [2]int:
		fmt.Printf("array:[2]int:%#v,%T\n", data, data)
	}
}

func main() {
	// 1、空接口的使用场景之一, map
	// 可以类似实现 python 里的 dict字典
	var self = make(map[string]interface{}, 10)
	self["name"] = "hcy"
	self["age"] = 19
	self["gender"] = "男"
	self["married"] = false
	self["hobbies"] = []string{"code", "game", "music"}

	fmt.Printf("%#v\n", self)
	// map[string]interface {}{"age":19, "gender":"男", "hobbies":[]string{"code", "game", "music"}, "married":false, "name":"hcy"}
	fmt.Printf("%T\n", self) // map[string]interface {}

	// 2、第二个适用场景的测试
	dataPrinter(100)   // 100
	dataPrinter("aaa") // aaa
	dataPrinter(true)  // true

	// 3、第三个适用场景的测试
	typeGetter(10)

	// 4、第四个使用场景的测试
	typePrinter("aaa")          // string:"aaa",string
	typePrinter(true)           // bool:true,bool
	typePrinter(10)             // int:10,int
	typePrinter(uint(10))       // uint:0xa,uint
	typePrinter([]int{1, 2, 3}) // slice:[]int:[]int{1, 2, 3},[]int
	typePrinter(map[string]int{ // map[string]int:map[string]int{"ying":16, "yuki":14},map[string]int
		"ying": 16,
		"yuki": 14,
	})
	typePrinter(func() {})    // func():func()
	typePrinter([2]int{1, 2}) // array:[2]int:[2]int{1, 2},[2]int
}
