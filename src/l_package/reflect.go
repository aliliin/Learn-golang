package main

import (
	"fmt"
	"reflect"
)

/**
go 的反射机制是要通过接口来进行的。
*/
func main() {
	// 反射操作，通过反射，可以获取一个接口类型变量的 类型 和 值

	var x float64 = 3.14

	// 获取反射类型
	fmt.Println("type", reflect.TypeOf(x))   // type float64
	fmt.Println("value", reflect.ValueOf(x)) // value 3.14

	// 根据反射的值获取对应的类型和数值
	v := reflect.ValueOf(x)
	fmt.Println("kind is float64: ", v.Kind() == reflect.Float64)
	fmt.Println("type:", v.Type())
	fmt.Println("value :", v.Float())
	// 查找参数类型
	var f int64 = 10
	CheckType(f)

	// 接口类型变量 -> 反射类型对象
	num := 1.23
	value := reflect.ValueOf(num)
	// 反射类型对象 --》 接口类型变量
	convertValue := value.Interface().(float64)
	fmt.Println(convertValue)

	/**
	反射类型的对象 -》 接口类型变量，理解为强制转换
	*float64 和 float64 是不一样的
	*/
	pointer := reflect.ValueOf(&num)
	convertPointer := pointer.Interface().(*float64)
	fmt.Println(convertPointer)

}

// 查找参数类型
func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int64, reflect.Int32:
		fmt.Println("Int")
	case reflect.String:
		fmt.Println("String")
	default:
		fmt.Println("Unknown", t)
	}
}
