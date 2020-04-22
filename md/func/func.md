
# 函数

# go 的函数与其他主要编程语言的差异

1. 函数可以有多个返回值
2. 所有参数都是值传递：slice，map，channel 会有传引用的错觉，但是函数的所有参数是值传递，穿的是结构的复制，操作的是同一块空间
3. 函数可以作为变量的值
4. 函数可以作为参数和返回值

# 简单的函数 

多返回值的函数，忽略其中的返回值。

```
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
```

# 实现计算方法调用时长

```
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

func TestTimeSpent(t *testing.T){
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
```

# 可变长参数

```
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
```

# 延时执行函数 defer

类似于 java 的 `try...finally..`

```
func TestDefer(t *testing.T) {
    // 这是一个匿名函数，也可以是一个写好的函数
    // 不会立即执行这个函数
    defer func() {
        t.Log("Clear resources")
    }()
    // 限制向下面的语句
    t.Log("Started")
    // panic 代表异常中断, 不可以修复的异常
    panic("Fatal error”) //defer仍会执⾏
    // 在返回之前执行上面的 延时函数
}
```

例子 : 

```
func Clear() {
	fmt.Println("Clear resources.")
}

func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}
```