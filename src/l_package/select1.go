package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 100
	}()

	select {
	case <-ch1:
		fmt.Println("case 1 执行")
	case <-ch2:
		fmt.Println("case 2 执行")
	case <-time.After(3 * time.Second):
		fmt.Println("case 3 执行 ..timeout..")
	}
}
