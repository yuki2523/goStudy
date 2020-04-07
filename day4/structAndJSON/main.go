package main

import (
	"encoding/json"
	"fmt"
)

// 由于Go里把包里的变量或方法向外开放,需要首字母大写
// 那么想要指定这个变量名在相应格式的文件里以什么样的形式呈现
// 可以在后面接上 `格式:"呈现形式"`,如下面的 `json:"name"`,在json格式里字段名Name解析为name
type person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	Ok     bool   `json:"ok"`
}

var (
	jsonStr = `{"name":"hcy","age":19,"gender":"男","ok":true}`
	jsonArr = `[1,2,3,"abc","你好"]` // 问题:Go里可以解析这种JSON字符串吗?
)

func main() {
	// fmt.Println(jsonStr)
	fmt.Println(jsonArr)
	// json.Marshal 序列化：结构体 => json字符串
	// json.Unmarshal 反序列化：json字符串 => 结构体

	var p1 = person{
		Name:   "hcy",
		Age:    19,
		Gender: "男",
		Ok:     true,
	}

	var strP1, err = json.Marshal(p1)
	// json.Marshal(结构体实例)
	// 第一个返回值为 []byte 格式的字符串切片,可以使用 string 强制转换
	// 第二个返回值为错误信息对象,如果没错,也就是序列化成功,那么其值为 nil
	if err != nil {
		fmt.Printf("序列化失败:%v\n", err)
	} else {
		fmt.Printf("%v\n", string(strP1))     // {"name":"hcy","age":19,"gender":"男","ok":true}
		fmt.Println(string(strP1) == jsonStr) // true
	}

	var p2 person
	fmt.Printf("%v\n", p2) // { 0  false}
	// 使用反序列化时,必须先声明一个 person 格式结构体出来,也就是实例化一个出来
	json.Unmarshal([]byte(jsonStr), &p2)
	// json.Unmarshal([]byte(string类型json字符串), 格式匹配的结构体实例地址值)
	// 第一个参数是 []byte 格式字符串切片,把 string 用 []byte 强制转换就OK了
	// 第二个参数是提前声明好的格式相匹配的结构体实例的地址值
	// 之后那个结构体实例就获取了json字符串解析(反序列化)后的值
	fmt.Printf("%#v\n", p2) // main.person{Name:"hcy", Age:19, Gender:"男", Ok:true}
}
