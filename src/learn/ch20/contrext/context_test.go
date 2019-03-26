package contrext_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// 接收 context消息
func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done(): // 通过接收到的消息来判断 来判断取消不取消
		return true
	default:
		return false
	}
}

func TestCancelByContext(t *testing.T) {
	// 根 context.Background() 创建
	// 子创建 context.WithCancel()
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i, "任务取消")
		}(i, ctx)
		cancel()
		time.Sleep(time.Second * 1)
	}
}
