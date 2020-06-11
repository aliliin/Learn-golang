package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	/**
	WattGroup :同步等待组
		add() 设置等待组中要执行的 子 goroutine 的数量
		Wait() 让主 goroutine 处于等待
		Done() 让等待组中的 counter 减 1
	*/
	wg.Add(2)
	go fun1()
	go fun2()
	fmt.Println("main 进入阻塞状态。等待wg中子goroutine结束")
	wg.Wait()
	fmt.Println("main 解除阻塞")
}

func fun1() {
	for i := 0; i < 5; i++ {
		fmt.Println("fun1() 函数中打印。。A", i)
	}
	wg.Done() // 给wg等待组中的counter 数量减 1 (同wg.Add(-1))
}

func fun2() {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Println("\tfun1() 函数中打印。。B", i)
	}
}
