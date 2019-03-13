package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(s)
	t.Log(len(s)) // 是byte数

	// s[1] = '3' // string 是不可以变的 byte slice
	// s = "\xE4\xB8"  // 可以存储任何二进制的数据
	var a string
	a = "中"
	t.Log(len(a)) // 是byte数 3
	c := []rune(a)
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0]) // 4e2d
	t.Logf("中 utf8 %x", a)       // e4b8ad

}

func TestStingToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]d", c) // 1 是代表我都用 c
	}
}
