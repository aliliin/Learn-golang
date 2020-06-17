package main

import (
	"fmt"
	"time"
)

func main() {
	/**
		分支语句 if switch select
		select 语句类型于 switch 语句
	 但是select语句会随机执行一个可运行的 case
	 如果没有 case 可以运行，要看是否有default 如果有就执行 default
	否则就进入阻塞，直到有 case 可以运行
	*/
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()
	go func() {
		ch2 <- 200
	}()

	select {
	case num1 := <-ch1:
		fmt.Println("ch1 中获取的数据 ", num1)
	case num2, ok := <-ch2:
		if ok {
			fmt.Println("ch2 数据", num2)
		} else {
			fmt.Println("ch2 通道关闭")
		}
	default: // 如果有这个，就直接运行这个
		fmt.Println("default 语句")
	}
	fmt.Println("main over")
}
