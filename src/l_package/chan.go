package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 创建通道
	var channel chan int
	fmt.Printf("%T,%v\n", channel, channel) // chan int,<nil>
	channel = make(chan int)
	fmt.Printf("%T,%v\n", channel, channel) // chan int,0xc00008c060

	// 通道内写入数据
	go setData(channel)
	// 读取通道内的数据
	for v := range channel {
		fmt.Println("读取的数据: ", v) // 读取的数据：0，1，2，3，4
	}

	// 非缓冲通道
	var ch1 chan int
	ch1 = make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子的 goroutine 中，i", i) // 子的 goroutine 中，i 0，1，2，3，4
		}
		// 循环结束后向通道中写数据，表示结束了。
		ch1 <- 1
		fmt.Println("结束")
	}()

	data := <-ch1
	fmt.Println("main data --> ", data)

	/**
	缓冲通道（固定大小）
		接收：缓冲区数据空了会阻塞
		写入：缓冲区数据满了会阻塞
	*/
	ch2 := make(chan int, 3)
	fmt.Println(len(ch2), cap(ch2)) // 0 ,3
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	fmt.Println(len(ch2), cap(ch2)) // 3,3
	// 缓冲区满了，如果需要继续写入数据，需要有其他的 goroutine 进行读取（队列的结构来读取）
	// ch2 <- 4
	ch3 := make(chan string, 3)
	fmt.Println(len(ch3), cap(ch3)) // 0 ,3
	go setStringData(ch3, 3)
	for {
		v, ok := <-ch3
		if !ok {
			fmt.Println("读取完毕", ok) // 读取完毕 false
			break
		}
		fmt.Println("读取的数据是：", v) // 读取的数据是： 0,1,2
	}

	/**
	双向通道
		chan ch
			chan <- data 写入数据
			data <- chan 读取数据
	*/
	ch4 := make(chan string)
	isDone := make(chan bool)
	go SendString(ch4, isDone)
	data1 := <-ch4        // 读取
	fmt.Println(data1)    // golang
	ch4 <- "learn golang" // 发送

	<-isDone // 判断这个来看主程序结束

	/**
	单向通道：定向
		chan <- Write 只支持写
	    <- chan Read 只能读
	*/
	singleChanWrite := make(chan<- int) // 只能写，不能读
	singleChanRead := make(<-chan int)  // 只能读，不能写

	fmt.Println("main over")
}

func setData(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch) // 通知对方，通道关闭
}

func setStringData(ch chan string, number int) {
	for i := 0; i < number; i++ {
		ch <- strconv.Itoa(i)
	}
	close(ch) // 通知对方，通道关闭
}

func SendString(ch chan string, isDone chan bool) {
	ch <- "golang"                            // 发送
	data2 := <-ch                             // 读取数据
	fmt.Println("main goroutine 传来: ", data2) // main goroutine 传来:   learn golang

	isDone <- true
}

/**
	通道注意点
1.用于goroutine，传递消息
2.通道，每个都有相关联的数据类型
	nil chan 的不能使用
3.使用通道传递数据：<- （通道是在goroutine之间链接，数据接收和发送必须处于不同的goroutine中
	chan <- data 发送数据到通道，向通道中写数据
	data <- chan 获取通道中的数据，从通道中读取数据
4.阻塞
	发送数据 是阻塞的，直到另一条 goroutine 读取数据来解除阻塞
	读取数据 也是阻塞的 直到另一条 goroutine 写出数据解除阻塞
5.本身channel 就是同步的，意味着同一时间，只能有一条 goroutine 来操作

*/
