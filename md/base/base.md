
# 资源

1. [Go 语言的官网](https://golang.google.cn/)
2. [Go 语言规范 ](https://golang.google.cn/ref/spec)
3. [Go 命令教程](https://github.com/hyper0x/go_command_tutorial)
4. [Go语言圣经(中文版)](https://docs.hacknode.org/gopl-zh/ch1/ch1-01.html)

# 定义关键字

```
函数 --> func
变量 --> var
常量 --> const
类型 --> type
```

# 换行与代码分割

Go 语言不需要在语句或者声明的末尾添加分号，除非一行上有多条语句。

实际上，编译器会主动把特定符号后的换行符转换为分号, 因此换行符添加的位置会影响Go代码的正确解析（比如行末是标识符、整数、浮点数、虚数、字符或字符串文字、关键字 break、continue、fallthrough 或 return 中的一个、运算符和分隔符++、--、)、] 或 } 中的一个。

举个例子, 函数的左括号 { 必须和 func 函数声明在同一行上, 且位于末尾，不能独占一行，

在表达式 x + y中，可在 + 后换行，不能在 + 前换行（以 + 结尾的话不会被插入分号分隔符，但是以 x 结尾的话则会被分号分隔符，从而导致编译错误）。

# 区间形式

区间索引时，Go言里也采用左闭右开形式, 即，区间包括第一个索引元素，不包括最后一个, 因为这样可以简化逻辑。

比如 :
 
1. `a = [1, 2, 3, 4, 5], a[0:3] = [1, 2, 3]`，不包含最后一个元素。
2. `s[m:n]` 这个切片，`0 ≤ m ≤ n ≤ len(s)`，包含 `n - m` 个元素。

# 编写测试程序

1. 源码⽂件以 _test 结尾：`xxx_test.go`
2. 测试⽅法名以 Test 开头：`func TestXXX(t *testing.T) {…}`

例如:

```
package try_test

import "testing"

func TestFirstTry(t *testing.T) {
	t.Log("my first try!")
}
```

# common

```
a := 1
// a 的地址
aPoint := &a
t.Log(a, aPoint)

// 输出类型
t.Logf("%T %T", a, aPoint)
```