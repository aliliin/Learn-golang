package series

import "fmt"

// 初始化 init 函数  可以写多个 init
func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

// package 编写
// 包名必须大写
func GetFibonacci(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}

func Fibonacci(n int) []int {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}
