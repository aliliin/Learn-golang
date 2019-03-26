package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

// 生产者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // 关闭通道 关闭的通道 继续发送数据,会报 Panic 的错误 前提是 接收端 有判断是不是结束了
		wg.Done()
	}()
}

// 消费者
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}

		}
		wg.Done()
	}()
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		data := <-ch
	// 		fmt.Println(data)
	// 	}
	// 	wg.Done()
	// }()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()

}

// channel 关闭
//  向关闭的 channel 发送数据 会导致 panic
//  v,OK <-ch;ok 为bool值 ,true 表示正常接收,false 表示通道关闭
// 所以的channel 接收者都会在 channel关闭时,
// 立刻从阻塞等待中返回且 上述 ok 值为false,.
// 这个广播机制常被利用,进行向多个订阅者同时发送信号.
// 如退出信号

// 开始发送
func startWriting(ch chan int) {
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // 发送完毕,关闭通道
	}()
}

// 开始接收
func startReading(ch chan int) {
	go func() {
		for {
			// 如果通道数据接收完毕,执行break
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}

	}()
}

// 测试关闭 channel
func TestCloseChannel1(t *testing.T) {
	ch := make(chan int)
	startWriting(ch)
	startReading(ch)
}

type Saiyan struct {
	Name  string
	Power int
}

func (s *Saiyan) Super() {
	s.Power += 10000
}
func TestBl(t *testing.T) {
	goku := &Saiyan{"Goku", 9001} //这个不加&也能打印出19001，是为什么呢
	goku.Super()
	fmt.Println(goku.Power) // 将会打印出 19001
	chane := goku.Power
	var ch *int
	ch = &chane
	*ch = 2
	fmt.Println(ch)         // 将会打印出 19001
	fmt.Println(goku.Power) // 将会打印出 19001

}
