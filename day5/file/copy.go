package main

import (
	"fmt"
	"io"
	"os"
)

func (f fs) copyFile(newPath, oldPath string) {
	// 先打开原来的
	oldFile, err1 := os.Open(oldPath)
	if err1 != nil {
		fmt.Printf("源文件打开失败,错误信息:%v\n", err1)
	}
	defer oldFile.Close()

	// 创建一个新的文件
	newFile, err2 := os.OpenFile(newPath, os.O_CREATE|os.O_RDWR, 6)
	if err2 != nil {
		fmt.Printf("新文件创建失败,错误信息:%v\n", err2)
	}
	defer newFile.Close()

	// 复制
	written, err3 := io.Copy(newFile, oldFile)
	if err3 != nil {
		fmt.Printf("复制过程报错,错误信息:%v\n", err3)
	}
	fmt.Println(written) // 这个值是文件的字节数
}
