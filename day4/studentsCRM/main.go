package main

import (
	"fmt"
	"os"
)

type student struct {
	id   int
	name string
}

var students = make([]student, 10, 20)

// 通过id查找学生在students切片里的索引
// 返回值一个是index索引,另一个是是否查询成功的bool值
func findStudentIndexByID(id int) (index int, success bool) {
	for studentIndex, stu := range students {
		if stu.id == id {
			index = studentIndex
			success = true
		}
	}
	return
}

// 查询展示所有学生信息
func showAllStudents() {
	fmt.Printf("id\tname\t\n")
	for _, stu := range students {
		fmt.Printf("%d\t%s\t\n", stu.id, stu.name)
	}
}

// 添加学生信息
func addStudent(name string) {
	var newID int
	var studentsLength = len(students)
	if studentsLength == 0 {
		newID = 1
	} else {
		newID = students[studentsLength-1].id + 1
	}
	students = append(students, student{
		id:   newID,
		name: name,
	})
}

// 删除学生信息
func removeStudent(id int) {
	var index, success = findStudentIndexByID(id)
	if success {
		students = append(students[:index], students[index+1:]...)
	} else {
		fmt.Println("输入的编号有误,请重试")
	}
}

// 编辑学生信息
func editStudent(id int, name string) {
	var index, success = findStudentIndexByID(id)
	if success {
		students[index] = student{
			id:   id,
			name: name,
		}
	} else {
		fmt.Println("输入的编号有误,请重试")
	}
}

func main() {
	students = []student{
		student{
			id:   1,
			name: "hcy",
		},
		student{
			id:   2,
			name: "ying",
		},
	}

	fmt.Println("系统已启动...")
	fmt.Printf("1.查看所有学生\n2.添加学生\n3.删除学生\n4.修改学生信息\n5.终止程序\n")
	fmt.Print("请输入想要执行的操作：")

	var num int
	for {
		fmt.Scanf("%d\n", &num)
		switch num {
		case 1:
			showAllStudents()
		case 2:
			var name string
			fmt.Print("请输入学生的姓名:")
			fmt.Scanf("%s\n", &name)
			addStudent(name)
		case 3:
			var removeStudentID int
			fmt.Print("请输入想要删除的学生的id:")
			fmt.Scanf("%d\n", &removeStudentID)
			removeStudent(removeStudentID)
		case 4:
			var (
				editStudentID int
				newName       string
			)
			fmt.Print("请输入想要编辑的学生的id:")
			fmt.Scanf("%d\n", &editStudentID)
			fmt.Print("请输入想要编辑的学生的name:")
			fmt.Scanf("%s\n", &newName)
			editStudent(editStudentID, newName)
		case 5:
			os.Exit(1)
		default:
			fmt.Println("编号错误,请重新输入")
		}
		fmt.Print("请再次输入想要进行的操作编号:")
	}
}
