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
	for i := 0; i < len(arr); i++ {
		t.Log(i, arr[i])
	}
	for idx, e := range arr {
		t.Log(idx, e)
	}
	// _ 下划线站位
	for _, e := range arr {
		t.Log(e)
	}
}

func TestSectionArray(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6}
	// 数组截取 arr[1:3]  颔首不含尾
	t.Log(arr[1:3])
	t.Log(arr[:4])
	t.Log(arr[:])
	t.Log(arr[2:])
	t.Log(arr[2:len(arr)])
}
