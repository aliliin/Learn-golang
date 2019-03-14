package map_ext

import "testing"

func TestMap(t *testing.T) {
	aMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	k := "two"
	v, ok := aMap[k]

	if ok {
		t.Logf("The element of key %q : %d\n", k, v) // The element of key "two" : 2
	} else {
		t.Log("Not found!")
	}
}

// 字典的键类型不能是哪些类型？
// 它的典型回答是：Go 语言字典的键类型不可以是函数类型、字典类型和切片类型

func TestMap1(t *testing.T) {
	var badMap2 = map[interface{}]int{
		"1": 1,
		// []int(2): 2, // 这里会引发 panic,
		3: 3,
	}
	t.Log(badMap2)
}

// 应该优先考虑哪些类型作为字典的键类型？
// 求哈希和判等操作的速度越快，对应的类型就越适合作为键类型
// 在那些基本类型中应该优先选择哪一个？答案是，优先选用数值类型和指针类型，通常
//  情况下类型的宽度越小越好。如果非要选择字符串类型的话，最好对键值的长度进行额外的约束。
