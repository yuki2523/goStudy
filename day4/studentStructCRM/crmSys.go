package main

import "fmt"

// 下面为业务功能函数

// 查看所有学生
func (c *crmSystem) showAllStudents() {
	fmt.Printf("id\tname\tage\t\n")
	var allStudents = c.students
	for _, stu := range allStudents {
		fmt.Printf("%d\t%s\t%d\t\n", stu.id, stu.name, stu.age)
	}
}

// 增加学生
func (c *crmSystem) addStudent() {
	var id = createID(c.students)
	var (
		name string
		age  int64
	)
	fmt.Print("学生姓名:")
	fmt.Scanf("%s\n", &name)
	fmt.Print("学生年龄:")
	fmt.Scanf("%d\n", &age)
	c.students[id] = student{
		id,
		name,
		age,
	}
}

// 删除学生
func (c *crmSystem) removeStudent() {
	var id int64
	fmt.Print("学生id:")
	fmt.Scanf("%d\n", &id)
	if !studentExistByID(id, c.students) {
		fmt.Println("该学生不存在,请输入正确的id！")
		return
	}
	delete(c.students, id)
}

// 修改学生信息
func (c *crmSystem) editStudent() {
	var (
		id   int64
		name string
		age  int64
	)
	fmt.Print("学生id:")
	fmt.Scanf("%d\n", &id)
	if !studentExistByID(id, c.students) {
		fmt.Println("该学生不存在,请输入正确的id！")
		return
	}
	fmt.Print("学生姓名:")
	fmt.Scanf("%s\n", &name)
	fmt.Print("学生年龄:")
	fmt.Scanf("%d\n", &age)
	c.students[id] = student{
		id,
		name,
		age,
	}
}
