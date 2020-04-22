
# 条件与循环

# 循环

go 语言仅支持 for 循环，不支持 while 循环

简单的 for : 

```
func TestFor1(t *testing.T) {
	for i := 0; i < 6; i++ {
		t.Log("i == " + i)
	}
}
```

实现 `while (n < 5)`

```
func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}
```

实现无限循环 :

```
func TestInfiniteLoop(t *testing.T) {
    n := 0
    for {
        // do somthing
    }
}
```

# if

和其他语言最大的区别是，go 的 if 中可以写变量。

```
func TestMultiSec(t *testing.T) {
	// if 中写变量，改变量不能在 if 外面用
	// 常用的用法师处理方法的多返回值
	if a := 1 == 1; a {
		t.Log("1 == 1")
		t.Log("1 == 1", a)
	}
	// 普通 if
	b := 10 / 2
	if b > 3 {
		t.Log("b == ", b)
	}

}
```

# switch

与其他主要编程语⾔的差异

1. 条件表达式不限制为常量或者整数；
2. 单个 case 中，可以出现多个结果选项, 使⽤逗号分隔，命中其中一个就执行；
3. 与 C 语⾔等规则相反，Go 语⾔不需要⽤ break 来明确退出⼀个 case；
4. 可以不设定 switch 之后的条件表达式，在此种情况下，整个 switch 结构与多个 if…else… 的逻辑作⽤等同

```
func TestMultiSwitchTry(t *testing.T) {
	for i := 0; i < 6; i++ {
		switch i {
		case 0, 2:
			t.Log("event", i)
		case 1, 3:
			t.Log("Odd", i)
		default:
			t.Log("it is not 0-3")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 6; i++ {
		switch {
		case i%2 == 0:
			t.Log("event", i)
		case i%2 == 1:
			t.Log("Odd", i)
		default:
			t.Log("un know")
		}
	}
}
```