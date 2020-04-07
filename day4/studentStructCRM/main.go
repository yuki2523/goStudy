package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int64
	name string
	age  int64
}

type crmSystem struct {
	students map[int64]student
}

var students = make(map[int64]student, 10)

func main() {
	// 初始化
	students[1] = student{
		id:   1,
		name: "hcy",
		age:  19,
	}
	students[2] = student{
		id:   2,
		name: "ying",
		age:  16,
	}
	var crm = crmSystem{students}

	// fmt.Println(crm)
	// 选项
	fmt.Printf("1.查看所有学生\n2.增加学生\n3.删除学生\n4.编辑学生信息\n5.退出\n")

	// 设置选项键入监听
	for {
		fmt.Print("请输入对应功能的数字:")
		var choise int64
		fmt.Scanf("%d\n", &choise)

		switch choise {
		case 1:
			crm.showAllStudents()
		case 2:
			crm.addStudent()
		case 3:
			crm.removeStudent()
		case 4:
			crm.editStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("键入的值有误,请重试")
		}
	}
}
