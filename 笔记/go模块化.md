### 1、包(package)

#### 1.1、首先看看Gopath下的`src`下文件目录

![包导入测试](image/包导入测试.png)

#### 1.2、三个自定义包里的内容

##### 1.2.1、`packagetest1`

**add.go**

```go
package packagetest1

import "fmt"

func init() {
	fmt.Println("packagetest1/add.go")
}

// Add 两数相加
func Add(a int, b int) int {
	return a + b
}

```

**sub.go**

```go
package packagetest1

import "fmt"

func init() {
	fmt.Println("packagetest1/sub.go")
}

// Sub 两数相减
func Sub(a int, b int) int {
	return a - b
}

```

##### 1.2.2、`packagetest2`

**isnil.go**

```go
package packagetest2

import "fmt"

func init() {
	fmt.Println("packagetest2/isnil.go")
}

// Isnil 判断引用数据类型是否为 nil
func Isnil(d interface{}) bool {
	switch data := d.(type) {
	case []int:
		return data == nil
	case map[string]int:
		return data == nil
	}
	return false
}

```

**iszero.go**

```go
package packagetest2

import "fmt"

func init() {
	fmt.Println("packagetest2/iszero.go")
}

// Iszero 判断数据是否为当前类型的 0 值
func Iszero(d interface{}) bool {
	switch data := d.(type) {
	case string:
		return data == ""
	case int:
		return data == 0
	case bool:
		return data
	}
	return false
}

```

##### 1.2.3、结合上面两个例子总结一下导出含时的注意事项

1. 包名和文件夹名一致，虽然是建议事项，但我建议这是强制事项
2. 向外暴露方法时，变量的首字母必须是大写
3. 向外暴露函数时，函数建议写注释，而且还有格式要求`// 变量名 说明`，空格都不能少

##### 1.2.4、`packagetest3`

**test.go**

```go
package packagetest3

import "fmt"

func init() {
	fmt.Println("packagetest3/test.go")
}

```

##### 1.2.5、结合上面三个包的里的内容再总结一下init函数

- 可以看到每个`.go`文件里都有一个`init`函数，这个函数没有参数没有返回值
- init函数每次包被导入时就会调用，在import语句执行时调用
- init执行时机(假设有main函数的情况下)：全局声明 => init() => main()

#### 1.3、使用这些包的主目录下的`main.go`

```go
package main

import (
	"fmt"

	// 直接在左边在加一个名字是给包取别名的意思
	p1 "yuki2523.goStudy.git/day5/packagetest1"
	"yuki2523.goStudy.git/day5/packagetest2"

	// 匿名导入包,只执行包里的 init 函数,无法使用包里暴露向外的数据
	_ "yuki2523.goStudy.git/day5/packagetest3"
)

func init() {
	fmt.Println("main.go")
}

func main() {
	var (
		a    = 1
		b    = 4
		sli  = []int{1, 2}
		map1 map[string]int
		bl1  = true
		s    = ""
	)

	fmt.Println(p1.Add(a, b))             // 5
	fmt.Println(p1.Sub(a, b))             // -3
	fmt.Println(packagetest2.Isnil(sli))  // false
	fmt.Println(packagetest2.Isnil(map1)) // true
	fmt.Println(packagetest2.Iszero(bl1)) // true
	fmt.Println(packagetest2.Iszero(s))   // true
}

```

**导入包的注意事项：**

1. 要写路径，路径就是gopath下的src下的完整路径
2. 如果包的名字很奇怪，那么可以给包起别名，`import 别名 "包的路径"`
3. 匿名导包，`import _ "包的路径"`，这种情况下是无法使用包里向外暴露的变量的，不过init函数是正常执行的，一般这种导包的适用条件就是仅仅需要那个包里的init函数的情况














