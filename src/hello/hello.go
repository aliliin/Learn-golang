package hello

import (
	"fmt"
	"mymath"
)

func init() {
	fmt.Printf("Hello World!\n")
}
func Hello() {
	mymath.Sqrt("测试导入package的包")
}
