
# 变量与常量

# 变量赋值方式


```
var a int = 1
var b int = 2
```

可以简写为 : 

```
var (
	a int = 1
	b int = 2
)
```

go 语言有类型推断能力，如果使用默认类型的话，可以:

```
var (
	a int = 1
	b     = 2
)
```

也可以不需要定义关键字，直接使用类型推断:

```
a := 1
b := 2
```

go 语言可以在一个赋值语句中可以对多个变量进行同时赋值:

```
a := 1
b := 2
// 交换
a,b = b,a
```

# 常量

go 语言的常量和其他语言的常量没有很大的区别，唯一的差异是 go 语言可以快速的设置连续的值。

连续常量

```
// 连续常量
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

// 打印
func TestConstTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

// 打印结果 : 1 2 3
```

连续的二进制常量

```
// 连续的二进制常量
const (
	Readable   = 1 << iota // 0001
	Writable               // 0010
	Executable             // 0100
)

func TestConstBinTry(t *testing.T) {
	a := 7 // binary a = 0111
	// & 运算  上下两个bit位都是 1 的情况下   才是1
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	a = 1
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	t.Log(Readable, Writable, Executable)
}
```