package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	a := 1
	b := 1
	size := 15
	t.Log(" ", a)
	for i := 0; i < size; i++ {
		t.Log(" ", b)
		temp := a
		a = b
		b = temp + a
	}
}

// 通常用于外部赋值
var num1 int

func TestVar(t *testing.T) {
	num1 = 3

	var num2 int = 2
	var num3 int = 2
	// 可以用另外一种写法定义变量
	var (
		num4 int = 4
		num5 int = 5
	)

	// 其实 go具有类型推断能力 可以不用写变量类型
	var num6 = 6
	num7 := 7

	fmt.Println(num1, num2, num3, num4, num5, num6, num7)
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a, b)
}
