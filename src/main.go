package main

// 多依赖包依赖文件
import (
	// echo "fmt" // 起别名
	// _ "hello" // 只会执行 依赖包中的 init方法
	// "mymath"
	"learn"
	"fmt"
)

func main() {
	// 常量声明 iota
	learn.LearnIota()
	// 变量声明学习
	fmt.Print(learn.Cat) // 可以调用别的包中的大写字母的变量
	learn.LearnVariable()
	// 数据类型学习
	learn.LearnTypes()
	//hello.Hello()
	// hello.Learn()
	// mymath.Sqrt("Sqrt 字符串")
	// echo.Println("测试成功")
}

