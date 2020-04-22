
# 命令行参数

os 包以跨平台的方式，提供了一些与操作系统交互的函数和变量。程序的命令行参数可从 os 包的 Args 变量获取；os包外部使用 `os.Args` 访问该变量。

`os.Args` 变量是一个字符串 (string) 的切片 (slice), 先把切片 s 当作数组元素序列, 序列的长度动态变化, 用 `s[i]` 访问单个元素, 用 `s[m:n]` 获取子序列。序列的元素数目为 `len(s)`。

`os.Args` 的第一个元素 : `os.Args[0]`, 是命令本身的名字, 其它的元素(`os.Args[1:len(os.Args)]` 或者 `os.Args[1:]` 后面这个是前面的简写)则是程序启动时传给它的参数。

新建 CommandParam.go 文件 :

```go
// 命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("start run command param main!!!")
	args := os.Args
	// var args []string = os.Args
	fmt.Println("first param is : " + args[0])
	// var 声明定义了两个 string 类型的变量 s 和 sep。
	// 变量会在声明时直接初始化。
	// 如果变量没有显式初始化，则被隐式地赋予其类型的零值（zero value），数值类型是0，字符串类型是空字符串 ""
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("other command params is :" + s)
	fmt.Println("end run command param main!!!")
}
```

以命令 : `go run CommandParam.go param1 param2` 运行试试。