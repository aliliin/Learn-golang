package main

import (
	"fmt"
	"time"
)

/**
 协程
一个打印数字，一个打印字母
// 当主的 Goroutine 结束了，其他的子 Goroutine 会被终止
*/
func main() {
	// 1. 先创建并启动子 goroutine 执行 printNum
	go printNum()
	for i := 1; i <= 100; i++ {
		fmt.Printf("\t 主 goroutine 中打印字母 ：A %d\n", i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("main over")
}

func printNum() {
	for i := 1; i <= 100; i++ {
		fmt.Printf("子 goroutine 中打印数字 ：%d\n", i)
	}
}
