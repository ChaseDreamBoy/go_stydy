package condition

import "testing"

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
