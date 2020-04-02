package main

import "fmt"

func main() {
	str := "你好 world"
	fmt.Printf("len(str) = %d,类型为%T\n", len(str), str) // len(str) = 12,类型为string
	// 确实是符合UTF-8编码的，一个中文字符一般占3个字节，展示出来的就是3个字符长度

	str1 := []int32(str)
	str2 := []rune(str)
	fmt.Printf("len(str1) = %d,类型为%T,值为%c\n", len(str1), str1, str1) // len(str1) = 8,类型为[]int32,值为[你 好   w o r l d]
	fmt.Printf("len(str2) = %d,类型为%T,值为%c\n", len(str2), str2, str2) // len(str2) = 8,类型为[]int32,值为[你 好   w o r l d]
	// 既然一个UTF-8是3个字符长度，那么使用int32就能很好的去储存，那么现在的长度就是8了
	// 而且可以证明 []int32和[]rune 是一样的，我试了 []int16，报错了

	str3 := []uint8(str)
	str4 := []byte(str)
	fmt.Printf("len(str3) = %d,类型为%T,值为%c\n", len(str3), str3, str3) // len(str3) = 12,类型为[]uint8,值为[ä ½   å ¥ ½   w o r l d]
	fmt.Printf("len(str4) = %d,类型为%T,值为%c\n", len(str4), str4, str4) // len(str3) = 12,类型为[]uint8,值为[ä ½   å ¥ ½   w o r l d]
	// 这里可以看到 []uint8 和 []byte 是一样的,而且结果发生了乱码

	// 而且可以看到不管是 []rune 还是 []byte,它们都是把字符串分成了一个数组

	// 字符串的修改
	// str1[0] = '人' // 字符串默认不能修改，这个虽然没有报错，但是
	// fmt.Printf("%s\n", str) // 你好 world,没有发生改变

	// 字符串的修改方案(尤其是还有中文等各国语言时的情况)
	str5 := "いま、私はGoを学んでいる"
	str6 := []rune(str5) // 由于字符串不能修改，这就是在重新开辟一块内存空间储存这个 []rune 后的结果
	// 然后对结果进行修改，之后再转回 string 就能获取修改后的数组
	fmt.Printf("%c\n", str6) // [い ま 、 私 は G o を 学 ん で い る]
	str6[3] = '我'
	fmt.Printf("%s\n", string(str6)) // いま、我はGoを学んでいる
}
