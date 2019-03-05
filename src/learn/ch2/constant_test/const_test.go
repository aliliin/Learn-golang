package constant_test

import "testing"

// 常量定义

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
}

func TestConstantTry1(t *testing.T) {
	// a := 7 // 二进制位 0111
	a := 1 // 二进制位  0001

	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
