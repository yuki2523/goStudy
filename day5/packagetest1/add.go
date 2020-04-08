package packagetest1

import "fmt"

func init() {
	fmt.Println("packagetest1/add.go")
}

// Add 两数相加
func Add(a int, b int) int {
	return a + b
}
