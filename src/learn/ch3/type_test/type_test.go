package type_test

import "testing"

/**
 * 不允许隐式类型转换
 * 别名和原有类型也不能进行隐式类型转换
 * 可以支持支持指针类型，但是不能计算
 */
// 定义自己的类型别名
type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	var d int64
	b = int64(a) // 通过这种方式来实现类型的转换
	t.Log(d, b)
}

// 指针类型，不支持指针运算
// string 是值类型，其默认的初始化值为空字符串，而不是 nil null

func TestPoint(t *testing.T) {
	a := 1
	aPar := &a
	t.Log(a, aPar)
}

func TestSting(t *testing.T) {
	var s string // 定义空字符串
	t.Log(len(s))
	if s == "" {
		t.Log("*" + s + "*")
	}
}
func TestSliceInit(t *testing.T) {
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 2)
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}
