package operator_test

import "testing"

// 数组比较相同的数组元素相同才能比较，且每个元素的值相同才相等
func TestOperator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{2, 1, 3, 5}
	c := [...]int{2, 3, 5, 6}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	t.Log(c != b)
	t.Log(d == a)
}

/**
 * 位运算符
 * &^ 按位置零
 * 1 &^ 0 ----1
 * 1 &^ 1 ----0
 * 0 &^ 1 ----0
 * 0 &^ 0 ----0
 */

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry1(t *testing.T) {
	a := 7 // 二进制位 0111
	// a := 1 // 二进制位  0001
	a = a &^ Readable
	a = a &^ Executable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}


