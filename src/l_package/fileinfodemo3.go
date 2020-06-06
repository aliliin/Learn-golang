package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

/*
读取数据
	Reader 接口
		Read(p []byte)(n int,error)
*/
func main() {
	// 打开文件 abcdefghij
	fileName := "/Users/gaoyongli/sites/Learn-golang/file.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error:", err)
	}
	// 关闭文件
	defer file.Close()
	// 读取数据
	bs := make([]byte, 4, 4)
	n := -1
	// 从头开始读
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("读取结束")
			break
		}
		fmt.Println(string(bs[:n]))
	}
	// 指定位置开始读
	num, err := file.ReadAt(bs, 2)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(string(bs[:num]))

	/**
	// 第一次读取
	n, err := file.Read(bs)
	fmt.Println(err) // nil
	fmt.Println(n)   // 4
	fmt.Println(bs)  // [97 98 99 100]
	fmt.Println(string(bs))
	// 第二次读取
	n, err = file.Read(bs)
	fmt.Println(err)        // nil
	fmt.Println(n)          // 4
	fmt.Println(bs)         // [101 102 103 104]
	fmt.Println(string(bs)) //efgh

	// 第三次次读取
	n, err = file.Read(bs)
	fmt.Println(err)        // nil
	fmt.Println(n)          // 2
	fmt.Println(bs)         // [105 106 103 104]
	fmt.Println(string(bs)) //ijgh

	// 第四次读取
	n, err = file.Read(bs)
	fmt.Println(err)        // nil
	fmt.Println(n)          // 2
	fmt.Println(bs)         // [105 106 103 104]
	fmt.Println(string(bs)) //ijgh
	*/

	// 写数据
	// 打开文件
	wtiteFileName := "/Users/gaoyongli/sites/Learn-golang/files.txt"

	// os.O_CREATE 文件不存在创建文件
	//os.O_WRONLY  文件具有写的权限
	//os.O_APPEND  文件以追加的形式写入
	wtiteFile, err := os.OpenFile(wtiteFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
	}
	// 关闭文件
	defer wtiteFile.Close()
	// 写数据
	//ws := []byte{65, 66, 67, 68, 69} // ABCDE
	ws := []byte{97, 98, 99, 100} // abcd
	wn, err := wtiteFile.Write(ws)
	HandleErr(err)
	fmt.Println(wn)
	wtiteFile.WriteString("\n")
	// 转换格式写
	wn, err = wtiteFile.Write([]byte("golang"))
	HandleErr(err)
	fmt.Println(wn)
	wtiteFile.WriteString("\n")

	// 直接写出字符串
	wn, err = wtiteFile.WriteString("golang")
	HandleErr(err)
	fmt.Println(wn)

	// 拷贝文件 可以拷贝图片
	srcFile := "/Users/gaoyongli/sites/Learn-golang/file.txt"
	destFile := "abc3.txt"
	//total, err := CopyFile1(srcFile, destFile)
	//fmt.Println(total, err)
	//total, err := CopyFile2(srcFile, destFile)
	total, err := CopyFile3(srcFile, destFile)
	fmt.Println(total, err)
	/**
	copy(dst,src) 为复制 src 全部到 dst中
	copyN（dst，src，n(int64)）为复制src中n个字节到dst
	copyBuffer（dst，src，buf）为指定 一个buf缓存区，以这个大小完全复制
	*/

}

// 不适合大文件的读写，因为是一次性全部读取文件，然后写入新文件
func CopyFile3(srcFile, destFile string) (int, error) {
	bs, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return 0, err
	}
	err = ioutil.WriteFile(destFile, bs, 0777)
	if err != nil {
		return 0, err
	}
	return len(bs), nil
}
func CopyFile2(srcFile, destFile string) (int64, error) {
	file, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	defer file2.Close()
	return io.Copy(file2, file)
}

// 拷贝文件
func CopyFile1(srcFile, destFile string) (int, error) {
	file, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
	}
	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	defer file2.Close()

	// 读写 拷贝文件过大的话，就要修改这里的参数
	bs := make([]byte, 1020, 1024)
	n := -1 // 读取的数量
	total := 0
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("拷贝完毕")
			break
		} else if err != nil {
			fmt.Println("error", err)
			return total, err

		}
		total += n
		file2.Write(bs[:n])
	}
	return total, nil

}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
