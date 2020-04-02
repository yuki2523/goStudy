package main

import (
	"fmt"
	"unicode"
)

func main() {
	count := 0
	str := "hello沙河小王子"
	str1 := []rune(str)
	for i := 0; i < len(str1); i++ {
		if unicode.Is(unicode.Han, str1[i]) {
			fmt.Printf("%c\n", str1[i])
			count++
		}
	}
	fmt.Println(count)
}
