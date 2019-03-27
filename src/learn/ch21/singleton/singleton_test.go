package singleton_test

// 之执行一次， 单例模式
// 核心 调用 sync.Once 这个包。通过  once.Do 方法来确保掉用一次。
import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

// 创建单例结构
type Singleton struct {
}

var singleInstance *Singleton

// 只调用一个的方法
var once sync.Once

// 单例的核心方法
func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj") // 这个只会输出一次
		singleInstance = new(Singleton)
	})
	return singleInstance
}

// 不使用单例 会调用多次

func GetNotSinglet() *Singleton {
	fmt.Println("Create Obj") // 这个会输出多次，这样子的话就证明调用了多次
	singleInstance = new(Singleton)
	return singleInstance
}
func TestGetSingletonObj(t *testing.T) {
	// 要等待所有协程都执行完毕 ,使用 sync.WaitGroup 方法包
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			// obj := GetNotSinglet()
			fmt.Println(unsafe.Pointer(obj)) // 输出对象的地址
			wg.Done()
		}()
	}
	wg.Wait()
}
