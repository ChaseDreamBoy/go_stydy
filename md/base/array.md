
# 数组与切片

# 声明

```
// 声明并初始化默认每个元素为零值
var a [3]int 
a[0] = 1
// 声明同时初始化
b := [3]int{1, 2, 3} 
// 多维数组初始化
c := [2][2]int{{1, 2}, {3, 4}} 
// 数组任意长度 [...]
arr2 := [...]int{1, 2, 3, 4, 5, 6}
```

# 遍历

```
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
```

# 切片

切片用起来是一个可变长(可变长有点像 java 里面的 ArrayList)的数组，切片内部的结构体包括三个基本元素:

1. 一是指针ptr，它指向一片连续的存储空间，也就是指向一个数组；
2. 第二个是 slice 里面的元素个数，也就是可以访问的元素个数，就是初始化或者设定的值；
3. 第三个是 cap 表示指针指向的空间的长度。

```
func TestSliceInit(t *testing.T) {
	// 声明一个slice, 和申明数组的差别是这里没有指定长度, 表示 slice 是可变长的
	var s1 []int
	// 初始的 len 和 cap 值
	t.Log(len(s1), cap(s1))
	// 填充容量
	s1 = append(s1, 0)
	t.Log(len(s1), cap(s1))
	s1 = append(s1, 1, 2, 3)
	t.Log(len(s1), cap(s1))

	// init
	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	t.Log(len(s2), cap(s2))

	// 创建某总类型的变量  make
	// make(type, len, cap)
	s3 := make([]int, 3, 5)
	t.Log(len(s3), cap(s3))
	t.Log(s3[0], s3[1], s3[2])
	//t.Log(s3[0], s3[1], s3[2], s3[3])

	s3 = append(s3, 5)
	t.Log(len(s3), cap(s3))
	t.Log(s3[0], s3[1], s3[2], s3[3])
}
```

cap 扩容(切片实现可变长) :

PS : 扩容类似于 java 里面的 ArrayList.

```
func TestSliceGrowing(t *testing.T) {
	// cap 扩容
	s := []int{}
	for i := 0; i < 10; i++ {
		// 存储空间变化了  需要从新指向 新的存储空间
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}
```

### 切片共享存储空间

cap 是 该 切片 剩下的存储空间.

```
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	// cap 是 该 切片 剩下的存储空间
    // 修改其中一个, 也会影响其他的存储空间
	summer[0] = "un know"
	t.Log(Q2)
}
```

# 数组与切片的比较

1. 数组的容量不可以伸缩，切片容量是可以伸缩的
2. 数组对于相同维度，相同长度的数组可以比较，只要每一个位置的元素都相同，就相等；切片只能与 nli(空) 作比较，不能与其他的(哪怕是切片)做比较。
