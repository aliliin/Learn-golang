package client

// package 编写
import (
	"learn/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(5))
	t.Log(series.Fibonacci(3))
}
