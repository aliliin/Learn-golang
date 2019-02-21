package main

// 多依赖包依赖文件
import (
	echo "fmt" // 起别名
	_ "hello" // 只会执行 依赖包中的 init方法
	"mymath"
)

func main() {
	//hello.Hello()
	//hello.Learn()
	mymath.Sqrt("Sqrt 字符串")
	echo.Println("测试成功")
}

