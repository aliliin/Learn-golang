package csp

// 异步返回结果的方式
import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

// 串行输出结果
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
	/**
	Done
	working on something else
	Task is done.
	*/
}

// 封装成了 异步的 Service
// 需要的时候再拿出来
func AsyncService() chan string {
	// 声明一个通道
	retCh := make(chan string, 1) // 这个 1  是容量的大小
	go func() {
		ret := service()
		fmt.Println("retunrned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

/**
输出结果为:
working on something else
retunrned result.
Task is done.
Done
service exited.
*/

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	// 最后一个输出
	time.Sleep(time.Second * 1)
}
