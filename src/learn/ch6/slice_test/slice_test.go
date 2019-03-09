package slice_test

import "testing"

func TestSliceInit(t *testing.T) {

	var s0 []int
	t.Log(len(s0), cap(s0))
	// 添加元素
	s0 = append(s0, 2)
	t.Log(len(s0), cap(s0))

	s1 := []int{2, 4, 5, 6, 9}
	t.Log(len(s1), cap(s1))

	// 长度为3，容量为5
	s2 := make([]int, 3, 9)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 3)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
		// 增长规律为，放不下了之后。前一次乘 2
		// 连续的存储空间不够了之后，会成倍的增加存储空间，
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) // [Apr May Jun] 3 9
	Q3 := year[5:8]
	t.Log(Q3, len(Q3), cap(Q3)) // [Jun Jul Aug] 3 7
	summer := year[5:8]
	summer[0] = "Unknow"
	t.Log(Q2) // [Apr May Unknow]
	t.Log(Q3) // [Unknow Jul Aug]
	// 后端共享的内存空间，无论是谁修改了里面的内容，底层的数组数据都会变

	// 数组容量不可以伸缩

}
func TestSliceCompare(t *testing.T) {
	// a := []int{1, 3, 4}
	// b := []int{1, 2, 3}
	// c := []int{}
	// 不能进行比较
	// if a == b {
	// 	t.Log(a, b)
	// }
}
