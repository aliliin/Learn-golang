package main

import (
	"fmt"
)

func main() {

	// 十进制整数
	fmt.Printf("%d\n", 20) // 20
	// 十六进制
	fmt.Printf("%x\n", 20) // 14
	// 八进制
	fmt.Printf("%o\n", 20) // 24
	// 二进制
	fmt.Printf("%b\n", 20) // 10100

	// 浮点数 %f,$g,%e
	fmt.Printf("%f\n", 20.00) // 20.000000
	fmt.Printf("%g\n", 20.00) // 20
	fmt.Printf("%e\n", 20.00) // 2.000000e+01

	// 布尔型
	fmt.Printf("%t\n", true)  // true
	fmt.Printf("%t\n", false) // false

	// 字符 （%c Unicode 码点）
	fmt.Printf("%s\n", "string") // string

	// 内置格式的任何值
	fmt.Printf("%v\n", "string") // string

	// 任何值的类型
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("%T\n", arr) // [9]int

	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)
	fmt.Println(s) // Kim is 22 years old.
}
