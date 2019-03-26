package channel_cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

// 通知消息来关闭和开启
func isCancelled(cancelChan chan struct{}) bool {
	// 是否收到消息 如果收到就 返回 true 就要返回取消协程
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}
func cancel1(cancelChan chan struct{}) {
	// struct{} 表示空结构 ，第二个{} 表示实例化这个空结构
	cancelChan <- struct{}{}
}

func cancel2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			// 任务取消
			fmt.Println(i, "任务取消") //   输出为: 4 任务取消
		}(i, cancelChan)
	}
	cancel1(cancelChan)
	time.Sleep(time.Second * 1)
}
