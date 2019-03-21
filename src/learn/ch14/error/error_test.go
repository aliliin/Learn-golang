package error_test

// 错误机制
import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

// 定义错误变量
var LessThanTwoError = errors.New("n should be not less than 2")
var LessThanHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LessThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestFib(t *testing.T) {
	// t.Log(GetFibonacci(5))
	// 错误检查机制
	if v, err := GetFibonacci(5); err != nil {
		if err == LessThanTwoError {
			fmt.Println("小于 2")
		}
		if err == LessThanHundredError {
			fmt.Println("大于 100")
		}
	} else {
		t.Log(v)
	}
}

func GetFibonacci1(str string) {
	var (
		i    int
		err  error
		list []int
	)

	if i, err = strconv.Atoi(str); err == nil {
		if list, err = GetFibonacci(i); err == nil {
			fmt.Println(list)
		} else {
			fmt.Println("Error", err)
		}
	} else {
		fmt.Println("Error", err)
	}
}

func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)

	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("Error", err)
		return
	}
	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println(list)
}
