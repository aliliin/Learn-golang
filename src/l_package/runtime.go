package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取 goroot 目录
	fmt.Println("GOROOT-->", runtime.GOROOT())
	// 查看操作系统
	fmt.Println("os/platforms-->", runtime.GOOS)
	// 获取逻辑cpu 的数量
	fmt.Println("CPU 数量", runtime.NumCPU())
	// 设置 go 程序执行最大的 cpu 数量(1- 256)
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)

	// gosched  让出时间片，先让别的 goroutine 执行
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine...")
		}
		fmt.Println("goroutine 开始。")
		// 调用 fun
		fun()
		fmt.Println("goroutine 结束。")
	}()
	for i := 0; i < 5; i++ {
		// 让出时间片，先让别的 goroutine 执行
		runtime.Gosched()
		fmt.Println("main...")
	}

}

func fun() {
	defer fmt.Println("defer....")
	// return // 终止函数执行 defer 会正常执行
	runtime.Goexit() // 终止当前的 goroutine
	fmt.Println("fun 函数")
}
