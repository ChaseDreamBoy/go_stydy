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
	fmt.Println(os.Args)
	fmt.Println("hello word!")
	// 没有返回值   os.Exit 获取返回状态
	os.Exit(-1)
}
