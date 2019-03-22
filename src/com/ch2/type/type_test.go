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

// 测试指针
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

func TestString(t *testing.T) {
	var s string
	// string 是值类型，其默认的初始化值为空字符串，⽽不是 nil
	t.Log("*" + s + "*")
	t.Log(len(s))
}
