package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	// 这是一个多返回值的函数
	// 函数名后面的两个括号 第一个是参数 第二个是返回值
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	// 正常的多返回值
	a, b := returnMultiValues()
	t.Log(a)
	t.Log(b)
	// 忽略其中的返回值
	c, _ := returnMultiValues()
	t.Log(c)
}

func timeSpent(inner func(op int) int) func(op int) int {
	// 实现一个计算 方法调用时长的功能
	// 入参是一个函数 返回值也是一个函数
	return func(n int) int {
		// 开始时间
		start := time.Now()
		// 调用函数
		ret := inner(n)
		// 打印函数耗时
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestTimeSpent(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	// 可变长参数 -- 实际上都是被转换为数组
	ret := 0
	// 忽略 index
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}
