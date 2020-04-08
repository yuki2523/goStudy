package packagetest2

import "fmt"

func init() {
	fmt.Println("packagetest2/iszero.go")
}

// Iszero 判断数据是否为当前类型的 0 值
func Iszero(d interface{}) bool {
	switch data := d.(type) {
	case string:
		return data == ""
	case int:
		return data == 0
	case bool:
		return data
	}
	return false
}
