package main

import (
	"fmt"
	"strings"
)

func main() {
	// 字符串查找
	fmt.Println(strings.Contains("string", "ing"))      // true
	fmt.Println(strings.ContainsAny("string", "idngs")) // true

	// 字符串比较
	fmt.Println(strings.Compare("go", "go"))         // 0
	fmt.Println(strings.Compare("GOlang", "goLang")) // -1

	// 字符串比较不区分大小写
	fmt.Println(strings.EqualFold("GOlang", "goLang")) // true
	fmt.Println(strings.EqualFold("GOlang", "goLand")) // false

	// 字符串出现的次数
	fmt.Println(strings.Count("Google", "o")) // 2
	fmt.Println(strings.Count("four", ""))    // 5

	// 字符串分割
	fmt.Println(strings.Split("g,o,o,g,l,e", ",")) // [g o o g l e]

	// 字符串第一次出现的位置
	fmt.Println(strings.Index("Google", "l")) // 4
	fmt.Println(strings.Index("Google", "s")) // -1
	// 返回第一个字符出现的位置
	fmt.Println(strings.IndexAny("Google", "os")) // 1

	// 字符串最后一次出现的位置
	fmt.Println(strings.LastIndex("gopher", "p"))         // 2
	fmt.Println(strings.LastIndex("gopher gopher", "go")) // 7

	// 字符串数组转为单一字符串
	s := []string{"abc", "def", "xyz"}
	fmt.Println(strings.Join(s, ", ")) // abc, def, xyz

	// 字符串追加指定字符，可以指定追加几次
	fmt.Println("Go" + strings.Repeat("lang", 1)) // Golang

	// 字符串替换 原字符串、需要替换的字符、新的字符、
	fmt.Println(strings.Replace("accdefghi", "acc", "abc", 1)) // abcdefghi
	// 相同字符串多次替换 全部替换 写-1
	fmt.Println(strings.Replace("accaccghi", "acc", "abc", 2)) // abcdefghi

	// 字符串每个单词首字母大写
	fmt.Println(strings.Title("The golang language is google")) // The Golang Language Is Google
	// 字符串全部转换为大写
	fmt.Println(strings.ToTitle("google golang")) // GOOGLE GOLANG
	fmt.Println(strings.ToUpper("google golang")) // GOOGLE GOLANG

	// 字符串每个单词首字母小写转换
	fmt.Println(strings.ToLower("The Golang Language Is Google")) // the golang language is google

	// 字符串太多了，现用现查吧，如果你英语好的话。你基本也能猜到

}
