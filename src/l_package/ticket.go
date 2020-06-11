package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket int = 10
var mutex sync.Mutex     // 创建互斥锁
var wg sync.WaitGroup    // 创建同步等待组
var rwMutex sync.RWMutex // 创建读写锁

func main() {
	wg.Add(7)
	// 临界资源问题，会出现超卖的现象 （互斥锁）
	go getRich("脱贫机会1")
	go getRich("脱贫机会2")
	go getRich("脱贫机会3")

	// 读写锁
	go writeData(1)
	go readData(1)
	go writeData(2)
	go readData(2)

	wg.Wait() // main 等待
	fmt.Println("程序结束")
}

func getRich(task string) {
	rand.Seed(time.Now().UnixNano()) // 生成随机数量
	defer wg.Done()
	for {
		mutex.Lock() // 上锁之后，只能有一个 goroutine 执行
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
			fmt.Println("脱贫人口在减少：", ticket)
			ticket--
		} else {
			mutex.Unlock() // 解锁
			fmt.Println(task, "用完了,你注定是穷人了!")
			break
		}
		mutex.Unlock() // 使用互斥锁的时候操作结束之后，一定要解锁
	}
}

// 可以多个 goroutine 同时读
func readData(i int) {
	defer wg.Done()
	fmt.Println(i, "read start ...")
	rwMutex.RLock() // 读操作上锁
	fmt.Println(i, "reading...")
	time.Sleep(3 * time.Second)
	rwMutex.RUnlock() // 读操作解锁
	fmt.Println(i, "read over ...")
}

// 写的时候加锁，不能读也不能写
func writeData(i int) {
	defer wg.Done()
	fmt.Println(i, "write start ...")
	rwMutex.Lock()
	fmt.Println(i, "writing ...")
	time.Sleep(3 * time.Second)
	rwMutex.Unlock() // 写操作解锁
	fmt.Println(i, "write over ...")

}
