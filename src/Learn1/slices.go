package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr[2:6] = ", arr[2:6]) // 2 3 4 5
	fmt.Println("arr[:6] = ", arr[:6])   //  [0 1 2 3 4 5]
	fmt.Println("arr[2:] = ", arr[2:])   //  [2 3 4 5 6 7 8]
	fmt.Println("arr[:] = ", arr[:])     // [0 1 2 3 4 5 6 7 8]

	s1 := arr[2:]
	fmt.Println("arr[2:] = ", s1) //  [2 3 4 5 6 7 8]
	s2 := arr[:]
	fmt.Println("arr[:] = ", s2) // [0 1 2 3 4 5 6 7 8]

	fmt.Println("after update slice (s1)")
	updateSlice(s1)
	fmt.Println(arr)
	updateSlice(s2)
	fmt.Println(arr)

	fmt.Println("Extending slice")
	arr[0] = 0
	arr[2] = 2
	fmt.Println(arr) // [0 1 2 3 4 5 6 7 8]
	s1 = arr[2:6]    // [2 3 4 5]
	s2 = s1[3:5]     // [4 5 6 7]
	fmt.Printf("s1=%v,len(s1)=%d,cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Printf("s2=%v,len(s2)=%d,cap(s2)=%d\n", s2, len(s2), cap(s2))
	fmt.Println(s1[3:7])

	s3 := append(s2, 10)
	fmt.Println(s3)

}
