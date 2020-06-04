package main

import (
	"fmt"
	"os"
)

/**
FileInfo :文件信息
interface
 Name() 文件名字
	Size() 文件大小
	Mode() 文件类型和权限 -rw-r--r--
	ModTime() time.Time 最后一次修改时间
	IsDir() bool     是否是目录
	Sys() interface{}


*/
func main() {
	fileInfo, err := os.Stat("/Users/gaoyongli/sites/Learn-golang/file.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", fileInfo)
	// 文件名
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Mode())

}
