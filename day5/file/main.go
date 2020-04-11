package main

type fs struct{}

func main() {
	var f fs
	// result := f.readFile("./test.txt")
	// fmt.Println(result)
	// result := f.readFileByBufio("./test.txt")
	// fmt.Println(result)
	// f.writeFileByIoutil("./test1.txt", "会新建一个文件吗?")

	// result := f.readFileByIoutil("./test.txt")
	// fmt.Println(result)
	f.copyFile("./0.jpg", "./79.jpg")
}
