package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// 读取文件中的所有数据 无需手动关闭打开的文件
	fileName := "/Users/gaoyongli/sites/Learn-golang/file.txt"
	data, err := ioutil.ReadFile(fileName)
	fmt.Println(err)
	fmt.Println(string(data)) // abcdefghijABC

	// 写入数据 （写入之前会先清空这个文件）
	fileName1 := "/Users/gaoyongli/sites/Learn-golang/files.txt"
	s1 := "你好 hello world "
	error := ioutil.WriteFile(fileName1, []byte(s1), os.ModePerm)
	fmt.Println(error) // nil

	// ReadAll()
	s2 := "hello world"
	r1 := strings.NewReader(s2)
	data1, err1 := ioutil.ReadAll(r1)
	fmt.Println(err1)
	fmt.Println(string(data1))

	// 读取当前目录下的文件（不保含子目录下目录）
	dirName := "/Users/gaoyongli/sites/Learn-golang"
	fileInfo, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(fileInfo)) // 我的目录下 14 项, 隐藏文件也能查到
	for i := 0; i < len(fileInfo); i++ {
		fmt.Printf("第 %d 个文件名称： %s,是不是目录：%t\n", i, fileInfo[i].Name(), fileInfo[i].IsDir())
		/**
		第 0 个文件名称： .DS_Store,是不是目录：false
		第 1 个文件名称： .git,是不是目录：true
		*/
	}

	//  创建临时目录(定义的名字和随机生成的字符串) 创建临时文件
	dir, err := ioutil.TempDir("/Users/gaoyongli/sites/Learn-golang", "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(dir) // 不用就删除
	fmt.Println(dir)     // /Users/gaoyongli/sites/Learn-golang/test441916090
	dirFile, err := ioutil.TempFile(dir, "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(dirFile.Name()) // 不用就删除
	fmt.Println(dirFile.Name())     // /Users/gaoyongli/sites/Learn-golang/test441916090/test359195345
}
