package learn

import "fmt"

// array 数组类型
// 定义方式 var arr [n]type
// 在[n]type 中， n表示数组的长度，type 代表存储元素的类型
// 由于长度也是数组类型的一部分，因此 [3]int 和 [4]int 是不同的类型，数组也就不能改变长度。。
// 数组之间的赋值是值的赋值，不是数组本身的值变了。 即当把一个数组作为参数传入函数的时候，传入的其实是该数组的副本，而不是它的指针
// 如果要使用指针，那么就需要用到后面介绍的slice类型了。
func NewTypes(){
	var arr [10]int // 声明一个int 类型的数组
	arr[0] = 43 // 数组都是从零下标开始的。
	arr[1] = 34 // 赋值操作
	newArr := [3]int {2,3,4}  // 声明了一个长度为3的int数组
	newStringArr := [3]string {"gao","yong","li"}  // 声明了一个长度为3的字符数组
	c := [...]int {3,4,5,6,7,7,} // 省略长度可以使用 ... 表示，go 会自动计算它的长度
	// 二维数组 声明方式
	easyArray := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8},{10,11,12,13}}
	fmt.Print("第一个数组值",arr[0])
	fmt.Print("第一个数组值",newArr[1])
	fmt.Print("第一个数组值",newStringArr[2])
	fmt.Print("第一个数组值",c[2])
	fmt.Print("第一个数组值",easyArray[0])
}