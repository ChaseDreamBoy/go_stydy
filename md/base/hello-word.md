
# hello word 

新建一个 HelloWord.go 文件 :

```go
// main方法的包必须用main   package main
// 文件名可以随便起
// 包，表明代码所在的模块（包）
package main

import (
	"fmt"
	"os"
)

// main 方法必须是 main
func main() {
	// 获取入参数组 ： os.Args
	fmt.Println("start run go main!!!")
	fmt.Println(os.Args)
	fmt.Println("hello word!!!")
	// 没有返回值   os.Exit 获取返回状态
	os.Exit(-9)
}
```

命令行运行方式 : 

```shell script
$ go run HelloWord.go
```

Go语言原生支持 Unicode，它可以处理全世界任何语言的文本。如果不只是一次性实验，可以编译这个程序，保存编译结果以备将来之用。可以用build子命令：

```shell script
$ go build HelloWord.go
```

这个命令生成一个名为 HelloWord 的可执行的二进制文件，Windows系统下生成的可执行文件是 HelloWord.exe，增加了 .exe 后缀名，之你可以随时运行它，不需任何处理，因为静态编译，所以不用担心在系统库更新的时候冲突。




