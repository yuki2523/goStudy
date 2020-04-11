package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func (f fs) readFile(path string) (data string) {
	// os.Open(路径), 第一个返回值的类型就是 *File 类型
	fileObj, err := os.Open(path)
	// 打开时,返回值第一个为 file 结构体指针,第二个为错误信息,如果没错时为 false
	defer fileObj.Close() // 打开后操作完毕后要关闭,不要忘了
	if err != nil {
		fmt.Printf("打开文件失败,错误信息为%v\n", err)
		return
	}
	var result = make([]byte, 0, 100)
	var current = make([]byte, 30, 30) // 要初始化一个 []byte 类型切片作为读取时的接收者

	for {
		length, readErr := fileObj.Read(current)
		// 读取时,返回值第一个为读取的内容的长度,第二个为错误信息,如果没错时为 false
		// Read 是结构体 File 的方法
		// Read(一个 []byte 类型切片),读取的长度和 len(切片) 是相同的,如果文件剩下的内容长度小于这个值,返回文件剩下内容长度
		if readErr != nil {
			if readErr == io.EOF { // 反正让它一直往后读,直到出现读完了,也就是这个错误时,返回结果
				return string(result) // 最后返回的时候不要忘了再用 string 把 []byte 类型强转回字符串
			}
			fmt.Printf("读取文件失败,错误信息为%v\n", readErr)
		}

		if length == 30 { // 相等说明读满了,直接追加就行了
			result = append(result, current...)
		} else { // 不相等就一定是小于,那么要切片一下,再追加,不然会多出一些不必要的内容
			result = append(result, current[:length]...)
		}
	}
}

func (f fs) readFileByBufio(path string) (data string) {
	fileObj, err := os.Open(path) // 同样要获取文件的结构体指针
	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj) // 需要这个方法的返回值,参数就是那个文件结构体指针
	for {
		line, err := reader.ReadString(' ') // 读取中可以指定以哪个字符去分隔,记得是字符 byte 类型的
		fmt.Printf("%#v,%T\n", line, line)
		data += line
		if err != nil {
			if err == io.EOF {
				fmt.Println("文件读取完毕")
				return
			}
			fmt.Println("文件读取失败")
		}
	}
}

func (f fs) readFileByIoutil(path string) (data string) {
	content, err := ioutil.ReadFile(path) // 这种最方便,直接全部读完,返回值是 []byte 类型的
	if err != nil {
		fmt.Printf("文件读取失败,错误信息:%v\n", err)
		return
	}
	fmt.Println("文件读取成功")
	return string(content)
}
