package main

import (
	"fmt"
	"io"
	"os"
)

/**
seek( offset int64 ,whence int ) (int64,error) , 设置指针光标的位置
第一个参数：偏移量
第二个参数： 如何设置
	seekstart 表示对于文件开始
seekcurrent 表示相对于当前位置的偏移量
2 seekend 表示相对于文件末尾

*/
func main() {
	fileName := "/Users/gaoyongli/sites/Learn-golang/file.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer file.Close()
	bs := []byte{0}
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(4, io.SeekStart)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(2, 0) // seekStart
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(3, io.SeekCurrent)
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(0, io.SeekEnd)
	file.WriteString("ABC")

}
