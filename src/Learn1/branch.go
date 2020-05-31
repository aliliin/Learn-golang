package main

import (
	"fmt"
	"io/ioutil"
)

func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

/**
switch 后面可以没有表达式
switch 不需要break 也可以直接switch多个条件
*/
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "f"
	case score < 80:
		g = "c"
	case score < 90:
		g = "b"
	case score <= 100:
		g = "a"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	/**
	if 的条件里面可以赋值
	if 条件里赋值的变量作用域就只能在 if 语句里面使用
	*/
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println("cannot print file contents:", err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	a, b := 8, 4
	fmt.Println(eval(a, b, "/"))
	fmt.Println(
		grade(10),
		grade(100),
		grade(11),
		grade(101),
	)

}
