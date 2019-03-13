package string_test

import (
	"strconv"
	"strings"
	"testing"
)

// 字符串函数的练习
func TestFuncString(t *testing.T) {
	s := "A,B,C"
	// 字符串分割
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}

	// 字符串链接
	t.Log(strings.Join(parts, "-"))

}

// 字符串 和其他类型转换
func TestStingConv(t *testing.T) {
	// 整数转换
	str := strconv.Itoa(10)
	t.Log("str" + str)

	// 字符串转整型 相加 必须这样子
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
