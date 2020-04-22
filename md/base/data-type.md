
# 数据类型

# 基本数据类型

```
// 布尔
bool

// 字符
// string 的初始值为空字符串而不是 nil
// string 的判空方式 : if str == ""{}
string 

// 整形
// 向这里的 int 和下面的 uint 在 64 为系统就是 64 位，在 32 位 系统就是 32 位
int int8 int16 int32 int64

// 无符号整形
uint uint8 uint16 uint32 uint64 uintptr

// 字节，因为它是 8 为，所以 go 语言当中它也是 8 位无符号整形的别名
byte // alias for uint8

// 是 rune code 的编码值，根字符串相关
rune // alias for int32,represents a Unicode code point

// 浮点
float32 float64

// 复数类型
complex64 complex128
```

与其他主要编程语⾔的差异

1. Go 语⾔不允许隐式类型转换
2. 别名和原有类型也不能进⾏隐式类型转换

类型转换 : 

```go
package type_test

import (
	"math"
	"testing"
)

// 定义别名
type MyInt int64

func TestImplicit(t *testing.T) {
	// Go 语⾔不允许隐式类型转换
	var a int = 5
	var b int64
	// b = a
	b = int64(a)

	// 别名和原有类型也不能进⾏隐式类型转换
	var c MyInt
	// c = b
	c = MyInt(b)
	t.Log(a, b, c)

	// 类型的预定义值
	t.Log(math.MaxInt64, math.MaxFloat64, math.MaxUint32)

}
```

测试指针:

```
func TestPoint(t *testing.T) {
	a := 1
	// a 的地址
	aPoint := &a
	t.Log(a, aPoint)

	// 输出类型
	t.Logf("%T %T", a, aPoint)

	// 不⽀持指针运算
	// aPoint = aPoint + 1

}
```
