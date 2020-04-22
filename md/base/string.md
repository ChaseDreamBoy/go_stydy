
# 字符串

# 与其他主要编程语言的差异

1. string 是数据类型，不是引用或指针类型，所以它的 0 值是空字符串而不是空
2. string 是只读的 byte slice，len 函数可以它所包含的 byte 数
3. string 的 byte 数组可以存放任何数据

```
func TestString(t *testing.T) {
	var s string
	//初始化为默认零值 ""
	t.Log(s)
	s = "hello"
	t.Log(len(s))

	// string 是不可变的 byte slice
	// 下面的代码编译不过
	// s[1] = '3'

	// 指定二进制
	s = "\xE4\xB8\xA5"
	// s 的长度是 3，它是由 3 个 byte 组成的
	// 所以说 len() 求出来的是 byte 数， 不是字符数
	t.Log(len(s))

	// 可以存储任何二进制数据，下面就是任意的(随便写的)二进制
	// 也是可以存储的，只是显示为乱码而已
	s = "\xEb4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s))
}
```

# Unicode 与 UTF8

1. Unicode 是一种字符集（code point）是一种字符的编码
2. UTF8 是 unicode 的存储实现 （转换为字节序列的规则）

```
func TestUnicodeAndUtf8(t *testing.T) {
	var s = "中"
	t.Log(s)
	// 是 byte 数, 汉字 中 由 3 个 byte 组成, 所以这里是 3
	t.Log(len(s))

	// rune 取出字符串里面的 Unicode
	c := []rune(s)
	// Unicode 的长度是一
	t.Log(len(c))
	//	t.Log("rune size:", unsafe.Sizeof(c[0]))
	// Unicode 的值是 4e2d
	t.Logf("中 unicode %x", c[0])
	// utf8 存储是 3 个 byte 是 e4b8ad(分别是 e4、b8、ad)
	t.Logf("中 UTF8 %x", s)
}
```

编码与存储

| 字符 | "中" |
|-----------|------|
| Unicode | 0x4E2D |
| UTF-8 | 0xE4B8AD |
| string/[]byte | `[0xE4,0xB8,0xAD]` |

# 字符串遍历

```
func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	// range 字符串遍历的时候输出的是 rune 而不是 byte, 它里面是做了转换的
	for _, c := range s {
		// %[1] 都表示和第一个参数对应
		// 只是格式化方式不同
		// c表示汉字 x表示16进制的编码 d表示编码
		t.Logf("%[1]c %[1]x %[1]d", c)
	}
}
```

# 字符串函数

常用字符串函数

1. strings 包 : [https://golang.org/pkg/strings/](https://golang.org/pkg/strings/)
2. strconv 包 : [https://golang.org/pkg/strconv/](https://golang.org/pkg/strconv/)

常用方法 : 

```
package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	// 字符串分割 -- 它的返回是一个切片
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	// 字符串连接
	t.Log(strings.Join(parts, "-"))
}

func TestConvert(t *testing.T) {
	// 整数转为字符串
	s := strconv.Itoa(10)
	t.Log("str" + s)
	// 字符串转为整形
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
```