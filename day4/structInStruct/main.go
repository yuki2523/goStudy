package main

import "fmt"

type person struct {
	name   string
	age    int
	gender string
	detail personDetail
}

type personDetail struct {
	city     string
	school   string
	birthday string
	job      string
}

func main() {
	var (
		name   = "hcy"
		age    = 19
		gender = "男"
		detail = personDetail{
			city:     "芜湖市",
			school:   "安工程",
			birthday: "8月18",
			job:      "学生",
		}
		p1 = person{
			name,
			age,
			gender,
			detail,
		}
	)

	// 上面这就是最基本的嵌套方式,取嵌套结构体里的值时,一层层的点下去就行了,修改也是一样的
	fmt.Printf("%#v\n", p1)
	// main.person{name:"hcy", age:19, gender:"男", detail:main.personDetail{city:"芜湖市", school:"安工程", birthday:"8月18", job:"学生"}}
	fmt.Println(p1.detail.job) // 学生
	p1.detail.job = "程序员"
	fmt.Println(p1.detail.job) // 程序员
}
