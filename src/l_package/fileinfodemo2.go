package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/**
文件操作
分相对路径和绝对路径

*/

func main() {
	// 路径
	filePath1 := "/Users/gaoyongli/sites/Learn-golang/file.txt"
	filePath2 := "file.txt"

	//是否是绝对路径
	fmt.Println(filepath.IsAbs(filePath1)) //true
	fmt.Println(filepath.IsAbs(filePath2)) // flse
	// 获取绝对路径
	fmt.Println(filepath.Abs(filePath1))
	fmt.Println(filepath.Abs(filePath2)) // /Users/gaoyongli/sites/Learn-golang/file.txt

	// 获取上一层目录
	fmt.Println(filepath.Join(filePath1, ".."))
	// 创建一层文件夹目录
	err := os.Mkdir("/Users/gaoyongli/sites/Learn-golang/src/l_package/newdir", os.ModePerm)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("文件夹创建成功")
	}
	//  创建多层文件夹  os.MkdirAll()

	// 创建文件
	// Create 采用模式是 0666（任何人可读写，不可执行）
	// 如果文件存在会截断它（清空文件原内容）
	fileName, err := os.Create("/Users/gaoyongli/sites/Learn-golang/src/l_package/newdir/new.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fileName)

	// 打开文件
	file3, err := os.Open(filePath1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file3)
	// 写文件需要执行打开文件的格式
	file4, err := os.OpenFile(filePath1, os.O_RDWR, os.ModePerm)
	/**
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	O_RDWR   int = syscall.O_RDWR   /z/ open the file read-write.
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	*/
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file4)
	// 关闭文件
	fmt.Println(file4.Close())

	// 删除文件和目录 os.Remove("aa.txt")
	// 递归删除目录 os.RemoveAll()
}
