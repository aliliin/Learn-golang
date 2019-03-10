package map_test

import "testing"

// map 初始化方式
func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))

}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	// 面对不存在的key ,默认会给 value 0值 这样子就会使一个key 存在它本地的值对应的也是 0 值
	t.Log(m1[1])
	// 这种情况,到底使 本身是 0 值,还是 key不存在给的 0 值
	m1[2] = 0
	t.Log(m1[2])

	// 在访问的key 不存在的时,仍然会返回零值,不能通过返回 nil 来判断元素是否存在
	m1[3] = 0
	// 判断 key的值是否真的存在
	if v, ok := m1[3]; ok {
		t.Logf("key 3 value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

// map 遍历

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	// forecah  返回的是两个参数 key 和 value
	for k, v := range m1 {
		t.Log(k, v)
	}
}
