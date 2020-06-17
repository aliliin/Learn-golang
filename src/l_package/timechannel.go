package main

import (
	"fmt"
	"time"
)

func main() {
	/**
	创建一个定时器
	*/
	/**


	timer := time.NewTimer(3 * time.Second)
	fmt.Printf("%T\n", timer)
	fmt.Println(time.Now())

	// 这段代码 会阻塞 3秒
	ch2 := timer.C
	fmt.Println(<-ch2)
	*/
	timer2 := time.NewTicker(5 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2结束。。。开始")
	}()
	time.Sleep(2 * time.Second)
	timer2.Stop()
	fmt.Println("Timer 2 先停止了")

	/**
	 after
	返回一个通道 chan 存储的是 d 时间间隔之后的当前时间
	相当于  NewTimer(d).C
	*/
	ch := time.After(3 * time.Second)
	fmt.Println(time.Now())
	time2 := <-ch
	fmt.Println(time2)
}
