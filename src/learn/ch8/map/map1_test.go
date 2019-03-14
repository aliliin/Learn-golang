package map_ext

import (
	"fmt"
	"testing"
	"time"
)

func TestMapmain(t *testing.T) {
	m := map[int]string{
		1: "haha",
	}

	read(m)
	time.Sleep(time.Second)
	write(m)
	time.Sleep(30 * time.Second)
	fmt.Println(m)
}

func read(m map[int]string) {
	for {
		_ = m[1]
		time.Sleep(1)
	}
}

func write(m map[int]string) {
	for {
		m[1] = "write"
		time.Sleep(1)
	}
}

func TestChannel(t *testing.T) {

	ch1 := make(chan int, 3)
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3
	elem1 := <-ch1
	elem2 := <-ch1
	elem3 := <-ch1
	fmt.Printf("The first element received from channel ch1: %v\n",
		elem1)
	fmt.Printf("The Second element received from channel ch2: %v\n",
		elem2)

	fmt.Printf("The Third element received from channel ch3: %v\n",
		elem3)

}
