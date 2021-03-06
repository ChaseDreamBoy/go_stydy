package _map

import "testing"

func TestInitMap(t *testing.T) {
	t.Log("start init map")
	// 第一种方法声明 map
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m1)
	// 第二种方法声明 map
	m2 := map[string]int{}
	m2["one"] = 1
	t.Log(m2)
	// 第三种方法声明 map -- 在初始化指定大小时用到
	m3 := make(map[string]int, 10)
	// make 在创建切片的时候参数是  make(type, len, cap)
	// 这里使用 make 初始化 map 是 : make(type, Capacity)
	// 这里使用 make 初始化的时候为什么没有 len 呢？
	// 因为 len 对应的单元格都是默认初始化 0 值，但是在 map 是不能做到的
	m3["one"] = 13
	t.Log(m3)
	m4 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m4)
	// 获取 key
	t.Log(m4[2])
	// 获取长度
	t.Logf("len m1 = %d", len(m1))
	// make 初始化的 len 是 0
	t.Logf("len m3 = %d", len(m3))
	// 不能使用 cap
	t.Log("end init map")
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	// 访问一个不存在的 key, 输出 0
	t.Log(m1[1])
	m1[2] = 0
	// 访问值为 0 的 key 也返回 0
	// 好处是不会返回空指针, 坏处是不知道该 key 到底存在不
	t.Log(m1[2])
	// m1[3] = 0
	// v 表示返回值, ok 鄙视该 key 是否存在, true 表示存在, false 表示不存在
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

func TestLoop(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
	t.Log("loop sec...")
	m2 := map[string]string{"one": "one-value", "tow": "tow-value", "three": "three-value"}
	for k, v := range m2 {
		t.Log(k, v)
	}
}

func TestMapForSet(t *testing.T) {
	// go 实现 map
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
