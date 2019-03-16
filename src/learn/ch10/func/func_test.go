package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 多返回值的函数
func returnMultiValues() (int, int) {
	return rand.Intn(100), rand.Intn(20)
}

// 计算函数运行时长
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSf := timeSpent(slowFunc)
	t.Log(tsSf(10))
	// 忽略其中一个返回值可以用 下划线 _
	a, b := returnMultiValues()
	t.Log(a, b)
}

// 求和函数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

// 延迟执行函数
func TestDefer(t *testing.T) {
	defer func() {
		t.Log("lean Resources")
	}()
	t.Log("Started")
	panic("Fatal error") // defer 仍会执行
}

func Clear() {
	fmt.Println("Clear Resources.")
}
func TestDeferTwo(t *testing.T) {
	defer Clear()
	fmt.Println("Started Resources.")
	panic("Fatal error") // defer 仍会执行
	// 后面的代码不会执行
	fmt.Println("End Resources.")

}
func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 3, 3, 5))
	t.Log(Sum(10, 20, 30))
}

type Printer func(contents string) (n int, err error)

func printToStd(contents string) (bytesNum int, err error) {
	return fmt.Println(contents)
}

func TestPrintToStd(t *testing.T) {

	t.Log(printToStd("nihaoma"))

}
