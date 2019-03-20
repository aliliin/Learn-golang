package emptyInterface_test

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	// 断言类型
	// if i, ok := p.(int); ok {
	// 	fmt.Println("Int ", i)
	// 	return
	// }

	// if i, ok := p.(string); ok {
	// 	fmt.Println("string ", i)
	// 	return
	// }

	// fmt.Println("Unknow Type")

	switch v := p.(type) {
	case int:
		fmt.Println("Int ", v)
	case string:
		fmt.Println("string ", v)
	case float64:
		fmt.Println("float", v)
	}
}

func TestEmpty(t *testing.T) {
	DoSomething(10)
	DoSomething("Hello")
	DoSomething(2.44)

}
