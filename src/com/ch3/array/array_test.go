package array

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [5]int
	t.Log(arr[1], arr[2], arr[3])

	arr1 := [8]int{1, 2, 3, 4, 5}
	t.Log(arr1[1], arr1[2], arr1[3])

	// 数组任意长度 [...]
	arr2 := [...]int{1, 2, 3, 4, 5, 6}
	t.Log(arr2[1], arr2[2], arr2[5])

	t.Log(arr, arr1, arr2)
}

func TestTravelArray(t *testing.T) {
	arr := [8]int{1, 2, 3, 4, 5, 6}
	// 普通遍历
	for i := 0; i < len(arr); i++ {
		t.Log(i, arr[i])
	}
	// foreach
	for idx, e := range arr {
		t.Log(idx, e)
	}
	// 不用下表，就使用 _ 下划线站位
	for _, e := range arr {
		t.Log(e)
	}
}

func TestSectionArray(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	// 数组截取 arr[1:3]  颔首不含尾
	t.Log(arr[1:3])
	// 从开始到下标为 4
	t.Log(arr[:4])
	// 所有
	t.Log(arr[:])
	// 从 下标为 2 到结尾
	t.Log(arr[2:])
	t.Log(arr[2:len(arr)])
}

func TestMapWithFunValue(t *testing.T) {
	// map 的 key 是 int, map 的 value 是 一个方法
	// value 方法 的入参是 int 返回值是 int
	m := map[int]func(op int) int{}
	// 这里等于是定义不同的 key 对应的方法
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	// 这里对 value 的方法传参, 打印返回值
	t.Log(m[1](2), m[2](2), m[3](2))
}
