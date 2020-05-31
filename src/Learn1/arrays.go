package main

import "fmt"

func printArray(arr [5]int){
	for _, v := range arr {
		fmt.Println(v)
	}
}
func printArrays(arr *[5]int){
	arr[0] = 100
	for _, v := range arr {
		fmt.Println(v)
	}
}
func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 3, 5, 6,8}

	// 4 行 5 列
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3, grid)

	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	for _, v := range arr3 {
		fmt.Println(v)
	}
	printArray(arr3)
	fmt.Println(arr3)
	printArrays(&arr3)
	fmt.Println(arr3)

}
