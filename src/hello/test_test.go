package hello_test

import (
	"fmt"
	"testing"
)

func TestRanges(t *testing.T) {
	m := make(map[string]float64)
	m["oneone"] = 1.1
	m["twotwo"] = 2.2
	for key, value := range m {
		fmt.Printf("[%s]: %.1f\n", key, value)
	}
	a := [3]int{3, 2, 1}
	for idx, value := range a {
		fmt.Printf("[%d]: %d\n", idx, value)
	}
	s := []int{30, 20, 10}
	for idx, value := range s {
		fmt.Printf("[%d]: %d\n", idx, value)
	}
	name := "Michał"
	for idx, code := range name {
		fmt.Printf("[%d]: %q\n", idx, code)
	}
}

func FibonacciProducer(ch chan int, count int) {
	n2, n1 := 0, 1
	for count >= 0 {
		ch <- n2
		count--
		n2, n1 = n1, n2+n1
	}
	close(ch)
}
func TestFibonacciProducer(t *testing.T) {
	ch := make(chan int)
	go FibonacciProducer(ch, 10)
	idx := 0
	for num := range ch {
		fmt.Printf("F(%d): \t%d\n", idx, num)
		idx++
	}
	// for idx, num := range ch {
	// 	fmt.Printf("idx: %d F(%d): \t%d\n", idx, num)
	// }
}
