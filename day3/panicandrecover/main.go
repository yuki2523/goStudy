package main

import "fmt"

func f1() {
	fmt.Println("f1()")
}

func f2() {
	defer func() { // 在可能出现 panic 的前面加上 defer,写一个立即执行函数 "善后"
		// recover() 不是太推荐使用，因为既然都是崩溃级的错误了，那接着运行下去程序也不正常了
		// 这个时候就应该让它崩掉，只不过崩掉前会把 defer 里的代码执行完，这这里面把一些善后工作做完就OK了
		var err = recover() // recover() 是获取错误信息,而且程序就不崩,而是跳过这段错误代码,接着运行下去了
		fmt.Println(err)
	}() // 主要就是把打开的文件释放掉，连接的数据库，打开的网络连接给释放掉，不然会一直占用资源
	panic("出现严重错误,程序崩溃") // 直接崩溃，程序退出,f3()也不执行了
	fmt.Println("f2()")  // 下面这句是永远到不了了,panic出现时，这块代码就已经完蛋了
}

func f3() {
	fmt.Println("f3()")
}

func main() {
	f1()
	f2()
	f3()
}

// 程序执行结果:
// f1()
// 出现严重错误,程序崩溃
// f3()
