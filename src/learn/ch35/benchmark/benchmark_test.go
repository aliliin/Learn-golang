package benchmark

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 字符串相加，性能测试
func TestConcatStringByAdd(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5", "6", "7"}
	ret := ""
	for _, elem := range elems {
		ret += elem
	}
	assert.Equal("1234567", ret)
}

// 字符串切片连接在一起
func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert.New(t)
	elems := []string{"1", "2", "3", "4", "5", "6", "7"}
	var buf bytes.Buffer
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	// 这是断言
	assert.Equal("1234567", buf.String())
}

// benchmark 测试性能
// 通过命令行的方式运行 benchmark   go test -bench=. (全部运行)
// 想知道为啥运行的快和慢 go test -bench=. -benchmem 通过最后一个参数，才看运行速度
func BenchmarkConcatStringByAdd(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5", "6", "7"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
}

// 字符串切片连接在一起
func BenchmarkConcatStringByBytesBuffer(b *testing.B) {

	elems := []string{"1", "2", "3", "4", "5", "6", "7"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}
