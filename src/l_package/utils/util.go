package utils

import "fmt"

func Count() {
	fmt.Println("utils 包下的 Count() 函数")
}

func init() {
	fmt.Println("utils 包下的 init() 函数先被执行")
}

func init() {
	fmt.Println("utils 包下另一个的 init() 函数先被执行")
}
