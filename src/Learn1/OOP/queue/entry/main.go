package main

import (
	"Learn1/OOP/queue"
	"fmt"
)

func main() {
	q := queue.Queue{1}
	q.Push(4)
	q.Push(5)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
