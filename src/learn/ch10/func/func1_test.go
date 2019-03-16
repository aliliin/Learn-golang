package func_test

import (
	"errors"
	"fmt"
	"testing"
)

// 函数定义
type operate func(x, y int) int
type calculator func(x, y int) (int, error)

func genCalculator(op operate) calculator {
	return func(x, y int) (int, error) {
		if op == nil {
			return 0, errors.New("Invalid operation")
		}
		return op(x, y), nil
	}

}
func TestFunc(t *testing.T) {
	// 匿名函数
	op := func(x, y int) int {
		return x + y
	}
	add := genCalculator(op)
	x, y := 56, 78
	result, err := add(x, y)
	fmt.Printf("The result: %d (error: %v)\n", result, err)
}

func TestComplexAreray(t *testing.T) {

	complexArray := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}
	ComplexAreray(complexArray)
	t.Log(complexArray)
}

// 更改指定数组下 切片的某个元素值, 原来的值也会更改
func ComplexAreray(array [3][]string) {
	for _, v := range array {
		v[0] = "aaa"
	}
	return
}
