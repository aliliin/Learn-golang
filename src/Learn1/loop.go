package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
for 的条件里面不需要括号
for 的条件里面可以省略初始条件，结束条件，递增表达式
*/
// 整数转二进制
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func forever() {
	for {
		fmt.Println("这是个死循环")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1011 -->1101
		convertToBin(0),  // 是空字符串
		convertToBin(3434323223),
	)
	printFile("abc.txt")
}
