package main

import (
	// "log"
	"fmt"
	"path"
	"runtime"
	"time"

	// "io"
	"bufio"
	"os"

	// "bytes"
	"yuki2523.goStudy.git/day6/logger"
)

type level uint16

const (
	b level = 1
	c level = 4
)

func test1() {
	// 下面为两种获取 Writer 的方式
	fileObj, _ := os.OpenFile("./logtest.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644) // fileObj 是原生 io.Writer
	writer := bufio.NewWriter(fileObj)                                                  // 这个是继承自 io.Writer 的 bufio.Writer
	defer fileObj.Close()

	// Fp 家族三个,第一个参数都是指定输出位置, 要 Writer 类型数据
	// os.OpenFile bufio.NewWriter os.Stdout 这些都是 Writer
	// 分别为直接输出到硬盘,输出到缓冲区,输出到 console
	fmt.Fprintln(fileObj, "aaaa")
	fmt.Fprintln(writer, "aaaa")
	fmt.Fprintln(os.Stdout, "aaaa")
	writer.Flush() // 缓存区(内存)写入硬盘

	var a level = 10
	fmt.Printf("%v,%T\n", a, a)
	fmt.Println(a == 10)
	fmt.Printf("%v,%T\n", b, b)
	fmt.Printf("%v,%T\n", c, c)
	fmt.Println(b < c)
}

func test2() {
	// fileObj, _ := os.OpenFile("./logtest.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	// writer := bufio.NewWriter(fileObj)
	// defer fileObj.Close()
	// defer writer.Flush()

	logger := logger.NewLogger("debug", "./", "log.log", false)
	for i := 0; i < 100; i++ {
		logger.Debug("aaa%s,%d\n", "aaa", 1223)
		logger.Info("bbb\n")
		logger.Warning("ccc\n")
		logger.Error("ddd\n")
		time.Sleep(time.Second)
	}

	logger.Close()
}

func test3() {
	test4()
}

func test4() {
	pc, file, line, ok := runtime.Caller(2)
	fmt.Printf("%v,%T\n", pc, pc)
	f := runtime.FuncForPC(pc)
	fmt.Printf("%#v,%v,%T\n", f, f, f)
	fmt.Println(f.Name())

	fmt.Printf("%v,%T\n", file, file)
	fmt.Println(path.Base(file))
	fmt.Printf("%v,%T\n", line, line)
	fmt.Printf("%v,%T\n", ok, ok)
}

func test5() {
	str1 := "aaa"
	str2 := time.Now().Format("06-01-02 15:04:05 ")
	str3 := "bbb"
	fmt.Println(path.Join(str1, str2+str3))
}

func test6() {
	fileObj, _ := os.OpenFile("./bbb.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	fileObj.Close()

	os.Rename("./bbb.txt", "./aaa.txt")
}

func main() {
	// test1()
	test2()

	// test3()

	// test5()
	// test6()
}
