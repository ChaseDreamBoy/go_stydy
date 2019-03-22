package constant

import "testing"

// 声明常量
// 连续常量
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

// 连续的二进制常量
const (
	Readable   = 1 << iota // 0001
	Writable               // 0010
	Executable             // 0100
)

func TestConstTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstBinTry(t *testing.T) {
	a := 7 // binary a = 0111
	// & 运算  上下两个bit位都是 1 的情况下   才是1
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	a = 1
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	t.Log(Readable, Writable, Executable)
}
