package learn

import "fmt"

// 变量练习

// 定义一个变量
//var variableName type
// 定义变量并初始化类型和值
var user string = "aliliin"
// 定义多个变量、
var newName,oldName string = "aliliin","oldAliliin"
// 也可以忽略变量的类型，编译器自动推导出相应的变量类型
//name1,name2,name3 := v1,v2,v3  本形式只能在函数体内部使用
// 简短声明 :=
func NewVariable() {
	name1,name2,name3 := "v1","v2","v3"
	// 可以使用下划线丢弃已经定义好的变量
	_,b := 3,5
	// 注意已经声明的变量如果下面没有用到的话，编译就会报错
	fmt.Print(user,newName,oldName,name1,name2,name3,b)
}

// 常量练习
// 只能定义类型为 数值，布尔值，字符串，
const Pi = 3.1415926
const Name = "Aliliin"

func testBool(){
	var name bool // 一般声明
	newName := false // 简短声明
	name = true  // 赋值声明
	fmt.Print(name,newName)
}

//iost 枚举 iost关键字
// Go 里面有一个关键字 iota ，这个关键字用来声明 enum 的时候采用，它的默认值是0 ，const 每增加一行，它就会加 1

const (
	x = iota // x == 0
	y = iota // y == 1
	z = iota // z == 2
	w  //常量声明省略值的话，默认和之前的值字面意思相同。这里 代表是 w = iota 因此 w的值为3
)
const v = iota // 每遇到一个const 关键字 iota 的值就会重置，这时候 v 的值又变成了 0

const (
	h,i,j = iota,iota,iota // h==0 i==0 j==0 在同一行中 iota的值是一样的
)

func PrintConst(){
	fmt.Print(Pi,i,Name,x,y,z,w,v,h,i,j,name,newName)
}

