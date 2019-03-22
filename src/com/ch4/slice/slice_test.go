package slice

/*
	数组容量是不可以伸缩的
	切片容量是可以伸缩的

	对于相同维度，相同长度的数组是可以比较的，
	只要每一个位置的元素都相同，就相等

	切片只能与 nli 空 作比较   不能与哪怕是切片做比较
*/

import "testing"

func TestSliceInit(t *testing.T) {
	// 声明一个slice
	var s1 []int
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

func TestSliceGrowing(t *testing.T) {
	// cap 扩容
	s := []int{}
	for i := 0; i < 10; i++ {
		// 存储空间变化了  需要从新指向 新的存储空间
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	// cap 是 该 切片 剩下的存储空间
	summer[0] = "un know"
	t.Log(Q2)
}
