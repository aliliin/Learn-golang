package obj_cache

// sync.pool 对象缓存
// 适合用于通过复用，降低复杂对象的创建和GC 代价
// 协程安全，会有锁的开销
// 生命周期受 GC 的影响，不适合做连接池等，需要自己管理生命周期的资源的池化
import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/**
输出结果为
create a new Object.
100
3
*/
func TestSyncPool(t *testing.T) {
	syncPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new Object.")
			return 100
		},
	}

	v := syncPool.Get().(int) // 断言输出结果类型
	fmt.Println(v)
	syncPool.Put(3)
	v1, _ := syncPool.Get().(int)
	fmt.Println(v1)
}

/**
输出结果为
create a new Object.
100
create a new Object.
100
*/
func TestSyncPoolGC(t *testing.T) {
	syncPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new Object.")
			return 100
		},
	}
	v := syncPool.Get().(int) // 断言输出结果类型
	fmt.Println(v)
	syncPool.Put(3)
	// 生命周期
	runtime.GC() // GC 会清除 sync.pool 中缓存的对象
	v1, _ := syncPool.Get().(int)
	fmt.Println(v1)
}

func TestSyncPoolGoFunc(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("开始创建对象")
			return "初始化结果为 10"
		},
	}
	// 设置写入值
	pool.Put(11)
	pool.Put(11)
	pool.Put(11)
	pool.Put(11)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
