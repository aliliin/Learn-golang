package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

/**
返回值类型写在最后面
函数可以返回多个值，返回多个值的时候还可以定义名字
仅用于非常简单的函数
对于调用者而言没有区别
函数可以做为参数
没有默认参数，可选参数
*/
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation: " + op)
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, % d)", opName, a, b)
	return op(a, b)
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}
func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// 可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}
func swap(a, b int) (int, int) {
	return b, a
}
func main() {
	if result, err := eval(8, 4, "x"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}

	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(apply(pow, 12, 3))
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
	fmt.Println(sum(1, 3, 4, 5, 5, 5))
	a, b := 3, 4
	a, b = swap(a, b)
	fmt.Println(a, b)
}
