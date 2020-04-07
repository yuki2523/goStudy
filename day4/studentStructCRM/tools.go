package main

// 下面为业务辅助工具函数

// 生成一个没有出现过的id
func createID(stus map[int64]student) (newID int64) {
	if len(stus) == 0 {
		return 1
	}
	for _, stu := range stus {
		if stu.id >= newID {
			newID = stu.id
		}
	}
	newID++
	return
}

// 查看该 id 的学生是否存在
func studentExistByID(id int64, stus map[int64]student) (exist bool) {
	_, ok := stus[id]
	if ok {
		exist = true
	}
	return
}
