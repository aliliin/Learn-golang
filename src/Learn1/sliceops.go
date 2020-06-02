package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v,len=%d, cap=%d\n", s, len(s), cap(s))
}
func main() {
	var s []int // Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)
	s3 := make([]int, 8, 16)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s3, s1)
	printSlice(s3)

	fmt.Println("Deleting elements from slice")
	s3 = append(s3[:3], s3[4:]...)
	printSlice(s3)

	fmt.Println("Popping from front")
	front := s3[0]
	s3 = s3[1:]
	fmt.Println(front)
	printSlice(s3)
	fmt.Println("Popping from back")
	tail := s3[len(s3)-1]
	s3 = s3[:len(s3)-1]
	fmt.Println(tail)

	printSlice(s3)
}
