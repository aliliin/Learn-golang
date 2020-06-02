package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/**
变量类型写在变量名之后
编译器可推测变量类型
没有char， 只有 rune
原生支持复数类型
bool,string
int,int8,int16,int32,int64,uintptr
byte,rune
float32,float64,complex64,complex128
*/

var as = 3
var ss = "kkkk"

var (
	aaa = 3
	sss = "ssss"
)

func variableZeroValue() {
	var number int
	var s string
	fmt.Println("%d %q\n", number, s)
}
func variableInitialValue() {
	var number, a, s, c = 10, 11, "hello", true
	fmt.Println(number, a, s, c)
}
func variableShorter() {
	a, b, c, s := 3, 4, true, "a"
	b = 5
	fmt.Println(a, b, c, s)
}

// 欧拉公式
func euler() {
	fmt.Println(
		cmplx.Exp(1i*math.Pi)+1,
		cmplx.Pow(math.E, 1i*math.Pi)+1)
}

// 强制类型转换
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

// 常量定义
const (
	filename = "abc.ext"
	number   = 45
)

func consts() {
	const a, b = 3, 5
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(filename, c, number)

}

// 枚举类型
func enums() {
	// iota 会自增
	const (
		cpp = iota
		_
		java
		python
		golang
		php
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang, php)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func isUnique(astr string) bool {
	result := true
	m := []rune(astr)
	for k, v := range m {
		for key, value := range m {
			if v == value && k != key {
				result = false
			}
		}
	}
	return result
}
func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableShorter()
	fmt.Println(as, ss)
	fmt.Println(aaa, sss)
	euler()
	triangle()
	consts()
	enums()
	isUnique("leetcode")
}
