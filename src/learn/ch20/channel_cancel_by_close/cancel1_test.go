package channel_cancel_by_close_test

import (
	"fmt"
	"testing"
	"time"
)

func isCanceled(cancelChan chan struct{}) bool {
	// 判断是不是接收到了关闭的消息
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

// 内建的 关闭机制 广播机制
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan) // 关闭
}

func TestCancelClose(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCanceled(cancelCh) {
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i, "任务被取消")
		}(i, cancelChan)
	}
	cancel_2(cancelChan)
	// 输出为： i 的顺序可能不同
	/*
		4 任务被取消
		0 任务被取消
		1 任务被取消
		2 任务被取消
		3 任务被取消
	*/
	time.Sleep(time.Second * 1)
}
