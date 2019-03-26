package select_test

// 多路选择 和 超时控制
import (
	"fmt"
	"testing"
	"time"
)

// 启动服务方法
func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

// 其他服务进程
func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func AsyncService() chan string {
	// 声明一个通道
	retCh := make(chan string, 1)
	go func() {
		ret := service() // 调用启动服务
		fmt.Println("returned result.")
		retCh <- ret                   // 把这个启动的服务发送到通道
		fmt.Println("service exited.") // // 把这个启动的服务结束
	}()
	return retCh // 返回通道中的数据
}

func TestSelect(t *testing.T) {
	// 超时机制
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
	// retCh := AsyncService()
	// otherTask()
	// t.Log(<-retCh)
	// t.Log(ret)
}
