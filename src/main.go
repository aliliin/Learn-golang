package main

import (
	"fmt"
	"net/http"
	"time"
)

// http 服务
func main() {
	// 定义不同的url下的不同的处理逻辑
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})

	// 定义不同的url下的不同的处理逻辑
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\":\"%s\"}", t)
		w.Write([]byte(timeStr))
	})

	// 启动一个server 设置端口
	http.ListenAndServe(":8080", nil)
}


// import "learn"

// func main() {

	fmt.Print("hello world")
	// learn.functionCall()

	//learn.LearnSwitch()
	//learn.LearnBitwise() // 位运算符
	//learn.LearnOpearatorLogical() // 逻辑运算符
	//learn.LearnOpearatorRelation() // 关系运算符
	//learn.LearnOperatorArithmetic() // 算术运算符
	// 常量声明 iota
	//learn.LearnIota()
	// 变量声明学习
	//fmt.Print(learn.Cat) // 可以调用别的包中的大写字母的变量
	//learn.LearnVariable()
	// 数据类型学习
	//learn.LearnTypes()
	//hello.Hello()
	// hello.Learn()
	// mymath.Sqrt("Sqrt 字符串")
	// echo.Println("测试成功")
// }
