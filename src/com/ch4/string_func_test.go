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
