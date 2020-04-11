### 1、读取文件

#### 1.1、最原始的一种方式(`os.Open(path)`和`文件结构体.Read()`)

```go
package main

import (
	"fmt"
	"io"
	"os"
)

type fs struct{}

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

func main() {
	var f fs
	result := f.readFile("./test.txt")
	fmt.Println(result)
}

// 测试文件读写1
// 测试文件读写2
// 测试文件读333
```

#### 1.2、换一个好用的方式(`os.Open(path)`和`bufio`)

```go
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
```

#### 1.3、最方便的方式(`ioutil.ReadFile(path)`)

```go
func (f fs) readFileByIoutil(path string) (data string) {
	content, err := ioutil.ReadFile(path) // 这种最方便,直接全部读完,返回值是 []byte 类型的
	if err != nil {
		fmt.Printf("文件读取失败,错误信息:%v\n", err)
		return
	}
	fmt.Println("文件读取成功")
	return string(content)
}
```

###2、文件的写入

- 三个方法，全写在里面了

```go
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
```

### 3、文件的复制

```go
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
```

























