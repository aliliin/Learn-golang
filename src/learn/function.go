package learn

import "fmt"

// Go语言函数练习
func init() {

}
func functionCall(){
	// 函数调用 已知 TwoSum 函数需要两个参数
	num1 := 20
	num2 := 30
	var res int
	res = TwoSum(num1,num2)
	fmt.Print("两数之和是：",res)
}
// 函数基本定义格式
//func function_name([parameter list ]) return_types{
	// 函数体
//}

// func：函数由 func 开始声明
// function_name 自己定义的函数名称
// paramter list 参数列表，参数就像一个占位符，当本函数被调用时，你可以将值传递给这个参数，这个值被称为实际参数
// 参数列表指定的是 参数类型、顺序、及参数个数、参数是可选的，函数定义可以不定义参数
// return_type 返回类型，函数返回一列值。return_types 是该列值的数据类型，有些功能不需要返回值，这种情况下 return_types 不是必须的。
// 函数体。函数定义的代码集合

// 计算两数之和
func TwoSum(num1,num2 int) int{
	// 声明局部 结果变量
	var res int
	res = num1 + num2
	return res
}

func LearnFunction() {

}