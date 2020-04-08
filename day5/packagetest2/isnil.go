package packagetest2

import "fmt"

func init() {
	fmt.Println("packagetest2/isnil.go")
}

// Isnil 判断引用数据类型是否为 nil
func Isnil(d interface{}) bool {
	switch data := d.(type) {
	case []int:
		return data == nil
	case map[string]int:
		return data == nil
	}
	return false
}
