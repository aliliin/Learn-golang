package run_test

import (
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	numbers := make(map[string]int)
	numbers["one"] = 1          // 赋值
	numbers["two"] = 2          // 赋值
	fmt.Println(numbers["one"]) //  输出1
	t.Log(numbers["two"])       //
}

func TestMap(t *testing.T) {
	// 初始化一个字典
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "PHP": 2}
	mapa := map[string]int{"one": 1, "two": 2}
	mapb := map[string]int{"yi": 1, "er": 2}
	if mapa["one"] == mapb["yi"] {
		fmt.Println("可以比较")
	}
	if mapa["one"] != mapb["er"] {
		fmt.Println("可以比较")
	}
	// 赋值
	csharpRating, ok := rating["C++"]
	if ok {
		fmt.Println("输出ok的值", csharpRating)
	} else {
		fmt.Println("没找到对应的OK的值")
	}

	// delete(rating, "C")      // 删除key为C的元素
	fmt.Println(rating["C"]) // 输出 0
	fmt.Println(len(rating))

	m := make(map[string]string)
	m["Hello"] = "你好"
	m1 := m
	m1["Hello"] = "不好"
	fmt.Println(m["Hello"])
	fmt.Println(m1["Hello"])
}
