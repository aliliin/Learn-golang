package share_mem

// 共享内存 并发机制
import (
	"sync"
	"testing"
	"time"
)

// 不同的携程进行自增 但是 结果 不是 5000
// 不是线程安全的程序
func TestGroutine(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func(i int) {
			counter++
		}(i)
	}
	time.Sleep(time.Second * 1)
	t.Logf("counter = %d", counter)
}

func TestGroutineThreadSafe(t *testing.T) {
	// 使用锁
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func(i int) {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}(i)
	}
	// 加上这个等待是为了 更准确的看到携程结果的正确性,不然,它执行的太快结果就不一样了
	// 因为不知道要等多久结束,就使用了 一秒
	time.Sleep(time.Second * 1)
	t.Logf("counter = %d", counter)
}

// 这个方法,可以不用之前执行携程需要多久,
// 他会自动判断执行时间,重而来不丢失数据
func TestWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done() // 当这个都完成了之后,才能执行剩下的
		}()
	}
	wg.Wait() // 上面的所以方法执行成功了.才会执行这个下面的方法
	t.Logf("counter = %d", counter)
}

csp()并发机制