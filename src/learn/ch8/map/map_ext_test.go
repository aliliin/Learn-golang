package map_ext

import "testing"

// map  的 value可以是一个方法

func TestMapWithFunValue(t *testing.T) {

	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }

	t.Log(m[1](2), m[2](2), m[3](2))
}

// go 实现 Set  可以使用map[type]bool
// 元素的唯一性
// 基本操作, 1.添加元素,2判断元素是否存在,3,删除元素,4元素个数
func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	mySet[2] = true

	n := 3
	// 判断元素存在不存在
	if mySet[n] {
		t.Log(" is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 3)
	t.Log(len(mySet))
}

// 与go的 dock type 接口方式一起, 可以方便的实现单一方法的对象的工厂模式
