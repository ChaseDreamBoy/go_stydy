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
