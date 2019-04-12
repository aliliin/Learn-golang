package unsafe_programming

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

// 不安全编程

type Customer struct {
}

// 	强制类型转化尽量别这样子使用这种方法
func TestUnsafe(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	t.Log(unsafe.Pointer(&i))
	t.Log(f)
}

type myInt int

// 相对合理的类型转换

func TestConvert(t *testing.T) {
	a := []int{1, 23, 4, 5, 8}
	b := *(*[]myInt)(unsafe.Pointer(&a))
	t.Log(b)
}

// 原子类型操作

func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println(data, *(*[]int)(data))
	}

	var wg sync.WaitGroup

	writeDataFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
	}
}
