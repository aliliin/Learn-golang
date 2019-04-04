package unit

import (
	"fmt"
	"testing"
	// "github.com/stretchr/testify/assert" 单元测试
)

// go test -v - cover  全部测试命令
// 内置的单元测试框架的区别
// Fail, Error: 测试失败，该测试还会继续，其他测试继续执行
// FailNow ,Fatal：该测试失败，该测试中止，其他测试继续执行
func TestFunction(t *testing.T) {
	// 组合测试法
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d,the expected is %d, the actual %d", inputs[i], expected[i], ret)
		}
	}
	t.Log("正确的结果")
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("开始")
	t.Error("Error")
	fmt.Println("结束")
}

func TestFailInCode(t *testing.T) {
	fmt.Println("开始")
	t.Fatal("Fail")
	fmt.Println("结束")
}
