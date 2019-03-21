package panic

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

//  os.Exit 退出时不调用defer 指定的函数
// os.Exit 退出时不输出当前调用栈信息

func TestExit(t *testing.T) {
	defer func() {
		fmt.Println("这段代码不会执行!") // 这一部分的代码不会执行
	}()
	fmt.Println("开始执行")
	os.Exit(-1) // 直接退出 exit status 255
}

// panic  用于不可以恢复的错误
// Panic 退出前会执行 defer 指定的内容
func TestPanic(t *testing.T) {
	defer func() {
		fmt.Println("最后结果依旧执行!") // 这一部分的代码依旧执行
	}()
	fmt.Println("开始")
	panic(errors.New("错误信息!"))
}

// Recover  try ...
func TestRecover(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recoverd 开始\n", err)
		}
	}()
	fmt.Println("开始执行")
	panic(errors.New("错误信息!"))
}
