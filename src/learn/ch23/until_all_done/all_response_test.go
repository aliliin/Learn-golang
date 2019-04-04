package until_all_done

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Microsecond)
	return fmt.Sprintf("这个结果为 %d", id)
}

func AllResponse() string {
	numOfRuner := 10
	ch := make(chan string, numOfRuner)
	for i := 0; i < numOfRuner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < numOfRuner; j++ {
		// 字符串累加，这种方式不高效
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func FirstResponse() string {
	numOfRuner := 10
	ch := make(chan string, numOfRuner)
	for i := 0; i < numOfRuner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("之前", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("之后", runtime.NumGoroutine())
}
