package main

import "fmt"

type person struct {
	name         string
	age          int
	gender       string
	personDetail // 直接使用匿名字段的方式
}

type personDetail struct {
	city     string
	school   string
	birthday string
	job      string
}

func main() {
	var p1 = person{
		name:   "hcy",
		age:    19,
		gender: "男",
		personDetail: personDetail{
			city:     "芜湖市",
			school:   "安工程",
			birthday: "8月18",
			job:      "学生",
		},
	}

	fmt.Printf("%#v\n", p1)
	// main.person{name:"hcy", age:19, gender:"男", detail:main.personDetail{city:"芜湖市", school:"安工程", birthday:"8月18", job:"学生"}}

	// 匿名字段的方式和一般写法的区别就在于
	// 匿名字段不仅使用多个点可以点出来,而且直接使用一个点,也可以点出来

	fmt.Println(p1.personDetail.job) // 学生
	p1.personDetail.job = "程序员"
	fmt.Println(p1.personDetail.job) // 程序员

	// 这就是使用一个点把字段点出来的例子，同样可以进行取值和修改
	fmt.Println(p1.city) // 芜湖市
	p1.city = "六安市"
	fmt.Println(p1.city) // 芜湖市
}
