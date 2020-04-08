package packagetest1

import "fmt"

func init() {
	fmt.Println("packagetest1/sub.go")
}

// Sub 两数相减
func Sub(a int, b int) int {
	return a - b
}
