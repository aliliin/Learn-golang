package array_test

import "testing"

// 数组练习
func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{3, 3, 4, 65, 4}
	t.Log(arr[1], arr[2])
	t.Log(arr1[1], arr1[2])
	t.Log(arr2[1], arr2[2])
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 3, 4, 5, 6, 7, 8}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	// 遍历用法
	// idx 相当于 key ，e相当于 value
	// 如果不想要 key 可以 用下划线代替  写成  _,e:=range arr3
	for idx, e := range arr3 {
		t.Log(idx, e)
	}

}

// a[开始索引](包含) ， 结束索引[(不包含)]

func TestArraySection(t *testing.T) {
	arr4 := [...]int{1, 3, 4, 5, 6, 7, 8}
	arr3_sec := arr4[:3]
	t.Log(arr3_sec)
	arr4_sec := arr4[3:]
	t.Log(arr4_sec)
}
