package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Yes中文字符测试"
	fmt.Println(len(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}
	fmt.Println()
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println()
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
	fmt.Println()
	/**
		rune 相当于 go 的char
		使用 range 遍历 pos，rune 队
	   使用 utf8.RuneCountInString 获取字符的数量
	   使用 len 获取字节长度
	   使用[]byte 获得字节
	*/
}
