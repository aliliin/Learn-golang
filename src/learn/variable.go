package learn

import (
	"fmt"
	"reflect"
)

//  变量学习
var a int // 单个变量声明 全局变量不能省略 var
var b int = 456 // 同一行进行变量声明及赋值
var Cat string  = "汽车" // 大写的变量在别的包中可以被调用
// 分组方式声明变量
var (
	c string = "分组声明赋值"
	d int    = 2345
)
//x,v := 9,8    只能用在函数体内
func init() {
}
func LearnVariable() {
	a = 123 // 变量赋值
	// 函数体内也支持分组声明
	var (
		f string = "分组声明赋值"
		g int    = 2345
	)
	var h,i,j int = 1,2,3 // 单个声明多个变量
	var l,k = 1,3 // 省略掉类型 (如果省略掉类型，系统会自动匹配类型)
	x,v := 9,8  // 如果省略掉 var  就要在等号前面加上 冒号 := 的方式
	// 下划线 是打印不出来值的
	var r,_,w = 1,3,5
	// 类型转换 不兼容的类型是无法转换的 bool不能转为 数字类型
	var a int = 3
	var b float32 = 3.01
	//c := float32(a) 	// 类型转换
	c := int32(b) // 精度转换会丢失小数点后面的数值
	//var d bool = true 不兼容的类型是无法转换的，

	fmt.Print(a)
	fmt.Print(b)
	fmt.Print(c)
	fmt.Print(d)
	fmt.Print(f)
	fmt.Print(g)
	fmt.Print(h,i,j)
	fmt.Print(l,k)
	fmt.Print(reflect.TypeOf(l)) // 查询未定义类型的 变量类型，
	fmt.Print(x,v)
	fmt.Print(r,w)
	fmt.Print("\n")
	fmt.Print(c)
	fmt.Print(reflect.TypeOf(c))
}

