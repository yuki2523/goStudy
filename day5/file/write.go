package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func (f fs) writeFile(path string, text string) {
	fileObj, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("文件打开失败,错误信息:%v\n", err)
		return
	}
	defer fileObj.Close() // os里Open系列方法,都不能忘了关闭

	// 写入,下面两种都是可以的,只不过直接Write需要的是二进制字节,WriteString把字符串装换这步已经封装好了
	// fileObj.Write([]byte(text))
	fileObj.WriteString(text)
}

func (f fs) writeFileByBufio(path string, text string) {
	fileObj, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("文件读取失败,错误信息:%v\n", err)
		return
	}
	defer fileObj.Close()

	writer := bufio.NewWriter(fileObj)
	writer.WriteString(text)
	writer.Flush() // 把内存里的写入硬盘
}

func (f fs) writeFileByIoutil(path string, text string) {
	// 这种只能新建,覆盖,没有追加功能
	err := ioutil.WriteFile(path, []byte(text), 0644) // 最方便的方式
	if err != nil {
		fmt.Printf("文件读取失败,错误信息:%v\n", err)
		return
	}
}
