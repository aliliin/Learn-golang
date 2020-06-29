package main

import "fmt"

func enemy() func() {
	key := 0
	return func() {
		key++
		fmt.Println(key)
	}
}

func curryAdder(a int) func(int) {
	return func(b int) {
		fmt.Println(a + b)
	}
}
func main() {
	play := enemy()

	play() //1

	addB := curryAdder(1)
	addB(3) // 4
	addB(5) // 6
}
