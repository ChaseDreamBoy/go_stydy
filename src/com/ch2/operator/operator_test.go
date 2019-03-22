package operator

import "testing"

func TestCompareArray(t *testing.T) {
	// 用 == 比较数组
	// 相同维数且含有相同个数元素的数组才可以比较
	// 每个元素都相同的才相等
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 2, 4}
	c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b, a == d)
	// 数组长度不等不可以作比较
	// t.Log(a == c)
	t.Log(c)
}

const (
	Readable   = 1 << iota // 0001
	Writable               // 0010
	Executable             // 0100
)

func TestBitClear(t *testing.T) {
	a := 7 // 0111  r 4 w 2 e 1
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	a = a &^ Readable
	t.Log(a)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	a = a &^ Writable
	t.Log(a)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	a = a &^ Executable
	t.Log(a)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
